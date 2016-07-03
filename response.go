package tydaapi

// Response contains all data parsed that should be delivered
type Response struct {
	SearchTerm       string        `json:"search_term"`
	Language         string        `json:"language"`
	PronunciationURL string        `json:"pronunciation_url"`
	Conjugations     []string      `json:"conjugations"`
	WordClass        string        `json:"word_class"`
	Translations     []Translation `json:"translations"`
	Synonyms         []Word        `json:"synonyms"`
}

// Translation represents one type of translation
type Translation struct {
	Language    string `json:"language"`
	Description string `json:"description"`
	Words       []Word `json:"words"`
}

// Word represents one specific word
type Word struct {
	Value            string `json:"value"`
	Context          string `json:"context"`
	PronunciationURL string `json:"pronunciation_url"`
	DictionaryURL    string `json:"dictionary_url"`
}
