package tydaapi

import (
	"fmt"
	"log"
	"net/url"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// Languages is a map from short to long form of languages
var Languages = map[string]string{
	"sv": "Svenska",
	"en": "Engelska",
	"fr": "Franska",
	"de": "Tyska",
	"es": "Spanska",
	"la": "Latin",
	"nb": "Norska",
}

// Search makes a search on the tyda.se webpage and parses the result
// into the Response structure. Returns a response and nil, or nil and an error.
// Languages parameter uses the short form, eg. sv, en es.
func Search(query string, languages []string) (*Response, error) {
	u, err := BuildURL(query, languages)

	// Parse search reslult
	doc, err := goquery.NewDocument(u.String())
	if err != nil {
		log.Fatal(err)
	}
	// Limit selection on page
	res := doc.Find(".box-searchresult").First()

	ret := Parse(res)

	return ret, nil
}

// Parse takes a response and a goquery selection and parses data from the
// selection into the response structure.
func Parse(doc *goquery.Selection) *Response {
	response := &Response{}

	setSearchTermAndPronunciationURL(response, doc)
	setLanguage(response, doc)
	setConjugations(response, doc)
	setWordClass(response, doc)
	setTranslations(response, doc)
	setSynonyms(response, doc)

	return response
}

// BuildURL retuns a URL for tyda.se with correct format
func BuildURL(query string, languages []string) (*url.URL, error) {
	u, err := url.Parse("http://tyda.se/search/" + query)
	if err != nil {
		return nil, err
	}

	q := u.Query()
	for i, v := range languages {
		q.Set(fmt.Sprintf("lang[%d]", i), v)
	}
	u.RawQuery = q.Encode()
	return u, nil
}

func setLanguage(ret *Response, res *goquery.Selection) {
	id, _ := res.Find("h2").First().Attr("id")
	spl := strings.Split(id, "-")
	ret.Language = Languages[spl[0]]
}

func setSearchTermAndPronunciationURL(ret *Response, res *goquery.Selection) {
	h2 := res.Find("h2").First()
	s := h2.Find("b").First()
	ret.SearchTerm = strings.TrimSpace(s.Text())
	p, _ := h2.Find(".speaker").First().Attr("href")
	if p != "" {
		ret.PronunciationURL = "http://tyda.se" + p
	}
}

func setConjugations(ret *Response, res *goquery.Selection) {
	conjugations := res.Find(".conjugation")
	conjugations.Each(func(i int, s *goquery.Selection) {
		if !s.HasClass("missing") {
			ret.Conjugations = append(ret.Conjugations, strings.TrimSpace(s.Text()))
		}
	})
}

func setWordClass(ret *Response, res *goquery.Selection) {
	w := res.Find(".word-class").First()
	ret.WordClass = strings.TrimSpace(w.Text())
}

func setTranslations(ret *Response, res *goquery.Selection) {
	var translations []Translation
	var t Translation
	res.Find(".capsulated-content").Each(func(i int, c *goquery.Selection) {
		c.Find(".list-translations").Each(func(i int, tr *goquery.Selection) {
			tr.Find(".item").Each(func(i int, s *goquery.Selection) {
				if s.HasClass("item-title") {
					if t.Language != "" {
						d := c.Find(".description").First()
						t.Description = strings.TrimSpace(d.Text())
						translations = append(translations, t)
					}
					t = Translation{Language: strings.TrimSpace(s.Text()), Words: make([]Word, 0)}
				} else {
					w := Word{Value: s.Find("a").First().Text()}
					w.Context = strings.Trim(s.Find(".trans-desc").First().Text(), " \n\t[]")
					p, _ := s.Find(".speaker").First().Attr("href")
					if p != "" {
						w.PronunciationURL = "http://tyda.se" + p
					}
					w.DictionaryURL, _ = s.Find(".mm").First().Attr("href")
					t.Words = append(t.Words, w)
				}
			})
		})
	})
	if t.Language != "" {
		translations = append(translations, t)
	}
	ret.Translations = translations
}

func setSynonyms(ret *Response, res *goquery.Selection) {
	sy := res.Find(".list-synonyms").First()
	var synonyms []Word
	sy.Find(".item").Each(func(i int, s *goquery.Selection) {
		w := Word{Value: s.Find("a").First().Text()}
		w.Context = strings.Trim(s.Find(".syn-desc").First().Text(), " \n\t[]")
		synonyms = append(synonyms, w)
	})
	ret.Synonyms = synonyms
}
