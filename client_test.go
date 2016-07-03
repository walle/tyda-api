package tydaapi

import (
	"bytes"
	"testing"

	"github.com/PuerkitoBio/goquery"
)

func TestBuildURL(t *testing.T) {
	u, err := BuildURL("foo", []string{})
	if err != nil {
		t.Errorf("Error building URL: %s\n", err)
	}
	if u.Host != "tyda.se" {
		t.Errorf("Expected %s got %s", "tyda.se", u.Host)
	}
	u, err = BuildURL("foo", []string{"bar"})
	if err != nil {
		t.Errorf("Error building URL: %s\n", err)
	}
	if u.RawQuery != "lang%5B0%5D=bar" {
		t.Errorf("Expected %s got %s", "lang%5B0%5D=bar", u.RawQuery)
	}
}

func TestParseHimmelAll(t *testing.T) {
	b := bytes.NewBufferString(HimmelAll)
	doc, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		t.Errorf("Error parsing fixture: %s\n", err)
	}

	r := Parse(doc.Selection)

	if r.Language != "Svenska" {
		t.Errorf("Expected %s got %s", "Svenska", r.Language)
	}
	if r.SearchTerm != "himmel" {
		t.Errorf("Expected %s got %s", "himmel", r.SearchTerm)
	}
	if r.PronunciationURL != "" {
		t.Errorf("Expected %s got %s", "", r.PronunciationURL)
	}
	if r.WordClass != "Substantiv" {
		t.Errorf("Expected %s got %s", "Substantiv", r.WordClass)
	}

	if len(r.Conjugations) != 2 {
		t.Errorf("Expected %d got %d", 2, len(r.Conjugations))
	}
	if r.Conjugations[0] != "himlen" {
		t.Errorf("Expected %s got %s", "himlen", r.Conjugations[0])
	}
	if r.Conjugations[1] != "himlarna" {
		t.Errorf("Expected %s got %s", "himlarna", r.Conjugations[1])
	}

	if len(r.Translations) != 8 {
		t.Errorf("Expected %d got %d", 8, len(r.Translations))
	}
	if r.Translations[0].Language != "Engelska" {
		t.Errorf("Expected %s got %s", "Engelska", r.Translations[0].Language)
	}
	if r.Translations[4].Language != "Latin" {
		t.Errorf("Expected %s got %s", "Latin", r.Translations[4].Language)
	}
	if r.Translations[0].Words[1].Value != "sky" {
		t.Errorf("Expected %s got %s", "sky", r.Translations[0].Words[1].Value)
	}
	if r.Translations[0].Words[1].PronunciationURL == "" {
		t.Errorf("Expected %s got %s", "nonempty", r.Translations[0].Words[1].PronunciationURL)
	}
	if r.Translations[0].Words[1].PronunciationURL == "" {
		t.Errorf("Expected %s got %s", "nonempty", r.Translations[0].Words[1].DictionaryURL)
	}
	if r.Translations[0].Words[3].Context == "" {
		t.Errorf("Expected %s got %s", "nonempty", r.Translations[0].Words[3].Context)
	}

	if len(r.Synonyms) != 1 {
		t.Errorf("Expected %d got %d", 1, len(r.Synonyms))
	}
	if r.Synonyms[0].Value != "sky" {
		t.Errorf("Expected %s got %s", "sky", r.Synonyms[0].Value)
	}
	if r.Synonyms[0].Context != "meteorologi" {
		t.Errorf("Expected %s got %s", "meteorologi", r.Synonyms[0].Context)
	}
}

func TestParseConjugationAll(t *testing.T) {
	b := bytes.NewBufferString(ConjugationAll)
	doc, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		t.Errorf("Error parsing fixture: %s\n", err)
	}

	r := Parse(doc.Selection)

	if r.Language != "Engelska" {
		t.Errorf("Expected %s got %s", "Engelska", r.Language)
	}
	if r.SearchTerm != "conjugation" {
		t.Errorf("Expected %s got %s", "conjugation", r.SearchTerm)
	}
	if r.PronunciationURL == "" {
		t.Errorf("Expected %s got %s", "nonempty", r.PronunciationURL)
	}
	if r.WordClass != "Substantiv" {
		t.Errorf("Expected %s got %s", "Substantiv", r.WordClass)
	}

	if len(r.Conjugations) != 1 {
		t.Errorf("Expected %d got %d", 1, len(r.Conjugations))
	}
	if r.Conjugations[0] != "conjugations" {
		t.Errorf("Expected %s got %s", "conjugations", r.Conjugations[0])
	}

	if len(r.Translations) != 4 {
		t.Errorf("Expected %d got %d", 4, len(r.Translations))
	}
	if r.Translations[0].Language != "Svenska" {
		t.Errorf("Expected %s got %s", "Svenska", r.Translations[0].Language)
	}
	if r.Translations[0].Description == "" {
		t.Errorf("Expected %s got %s", "nonempty", r.Translations[0].Description)
	}
	if r.Translations[3].Language != "Svenska" {
		t.Errorf("Expected %s got %s", "Svenska", r.Translations[3].Language)
	}
	if r.Translations[0].Words[0].Value != "böjning" {
		t.Errorf("Expected %s got %s", "böjning", r.Translations[0].Words[0].Value)
	}
	if r.Translations[0].Words[0].PronunciationURL != "" {
		t.Errorf("Expected %s got %s", "", r.Translations[0].Words[0].PronunciationURL)
	}
	if r.Translations[0].Words[0].PronunciationURL != "" {
		t.Errorf("Expected %s got %s", "", r.Translations[0].Words[0].DictionaryURL)
	}
	if r.Translations[0].Words[0].Context != "lingvistik" {
		t.Errorf("Expected %s got %s", "lingvistik", r.Translations[0].Words[0].Context)
	}

	if len(r.Synonyms) != 5 {
		t.Errorf("Expected %d got %d", 5, len(r.Synonyms))
	}
	if r.Synonyms[2].Value != "pairing" {
		t.Errorf("Expected %s got %s", "pairing", r.Synonyms[2].Value)
	}
	if r.Synonyms[2].Context != "politik" {
		t.Errorf("Expected %s got %s", "politik", r.Synonyms[2].Context)
	}
}

func TestParseHacerSv(t *testing.T) {
	b := bytes.NewBufferString(HacerSv)
	doc, err := goquery.NewDocumentFromReader(b)
	if err != nil {
		t.Errorf("Error parsing fixture: %s\n", err)
	}

	r := Parse(doc.Selection)

	if r.Language != "Spanska" {
		t.Errorf("Expected %s got %s", "Spanska", r.Language)
	}
	if r.SearchTerm != "hacer" {
		t.Errorf("Expected %s got %s", "hacer", r.SearchTerm)
	}
	if r.PronunciationURL != "" {
		t.Errorf("Expected %s got %s", "", r.PronunciationURL)
	}
	if r.WordClass != "Verb" {
		t.Errorf("Expected %s got %s", "Verb", r.WordClass)
	}

	if len(r.Conjugations) != 0 {
		t.Errorf("Expected %d got %d", 0, len(r.Conjugations))
	}

	if len(r.Translations) != 3 {
		t.Errorf("Expected %d got %d", 3, len(r.Translations))
	}
	if r.Translations[0].Language != "Svenska" {
		t.Errorf("Expected %s got %s", "Svenska", r.Translations[0].Language)
	}
	if r.Translations[2].Language != "Svenska" {
		t.Errorf("Expected %s got %s", "Svenska", r.Translations[3].Language)
	}
	if r.Translations[1].Words[0].Value != "göra" {
		t.Errorf("Expected %s got %s", "göra", r.Translations[1].Words[0].Value)
	}

	if len(r.Synonyms) != 0 {
		t.Errorf("Expected %d got %d", 0, len(r.Synonyms))
	}
}

const HimmelAll = `
<!doctype html>
<html>
<head>
  <script>preloaded_emediate_pageviewid = (new Date()).getTime()+"_"+Math.floor(Math.random()*100000);</script>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta charset="utf-8">
  <title>Tyda.se - Resultat för "himmel"</title>
  <meta name="description" content="Ett lexikon för översättning med över 1 miljon uppslagsord på svenska, engelska, tyska, franska och spanska">
  <meta name="keywords" content="lexikon engelska svenska tyska franska spanska översättning översätta ordlista ordbok betyder">
  <meta name="PUBLISHER" content="Tyda Sverige AB">
  <meta name="URL" content="http://tyda.se/">
  <link rel="search" href="http://tyda.se/search-engines" type="application/opensearchdescription+xml">
  <link rel="icon" sizes="32x32" href="/static/img/tyda-favicon-32" type="image/png"/>
  <link rel="icon" sizes="64x64" href="/static/img/tyda-favicon-64" type="image/png"/>
  <link rel="apple-touch-icon" sizes="64x64" href="/static/img/tyda-favicon-64" type="image/png"/>
  <meta name="google-site-verification" content="qPVeUQYR82XH_LmvDZHBEoscMhFEBHd730fj2AevaMM" />
    <meta name="mobile-web-app-capable" content="yes">

  <link href="http://fonts.googleapis.com/css?family=Open+Sans:300,400,500,600,700,800" rel="stylesheet" type="text/css">
  <link rel="stylesheet" type="text/css" href="/static/css/default.css?v=1453979674" media="screen" />
  <link rel="stylesheet" type="text/css" media="only screen and (max-width: 600px), only screen and (max-device-width: 600px)" href="/static/css/mobile.css?v=1454931518" />
  
      <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  
  <!--[if IE]>
    <style type="text/css">
      #social-bar li.item {
        background-color: #fffffd;
      }

      .box {
        background-color: #fffffd;
      }

      .input-text-search {
        border-color: #C0C0C0;
      }
    </style>
  <![endif]-->
  <meta name="twitter:card" content="summary_large_image">
  <meta name="twitter:domain" content="tyda.se">
            <meta name="twitter:title" content="Sökresultat för himmel">

                <meta name="twitter:description" content="Himlen är jordens atmosfär sedd från jordytan.

Himlen är blå (eller grå vid dåligt väder) på dagen och med nyanser av rött eller gult vid soluppgång och solnedgång, på grund av Rayleigh-spridning av solljus i atmosfären.[1] På natten är himlen mörk och när det inte finns moln kan man se månen (när den är över horisonten) och stjärnorna (stjärnhimlen). På månen och alla andra himlakroppar som saknar atmosfär är det ständigt svart och stjärnorna syns där även i dagsljus. Himlen kan ibland definieras som den tätare zonen av en planets atmosfär.

">

                <meta name="twitter:image" content="http://tyda.se/img/9314ced3483e4ae00aae2fd699efbfd1.jpg">

            <meta property="og:type" content="website">
      <meta property="og:url" content="http://tyda.se/search/himmel?lang%5B0%5D=en&amp;lang%5B1%5D=fr&amp;lang%5B2%5D=de&amp;lang%5B3%5D=es&amp;lang%5B4%5D=la&amp;lang%5B5%5D=nb&amp;lang%5B6%5D=sv">
      <meta property="og:site_name" content="Tyda.se">
      <meta property="og:description" content="Himlen är jordens atmosfär sedd från jordytan.

Himlen är blå (eller grå vid dåligt väder) på dagen och med nyanser av rött eller gult vid soluppgång och solnedgång, på grund av Rayleigh-spridning av solljus i atmosfären.[1] På natten är himlen mörk och när det inte finns moln kan man se månen (när den är över horisonten) och stjärnorna (stjärnhimlen). På månen och alla andra himlakroppar som saknar atmosfär är det ständigt svart och stjärnorna syns där även i dagsljus. Himlen kan ibland definieras som den tätare zonen av en planets atmosfär.

">
      <meta property="og:title" content="Sökresultat för himmel">
      <meta property="og:image" content="http://tyda.se/img/9314ced3483e4ae00aae2fd699efbfd1.jpg">
    <meta property="fb:admins" content="1563034500" />
    <script type="text/javascript">
        n24g = {"nuggad":"2212222222222221222222221229999999939994999499999949393922121212222222222222012241653232221229999999999999999919011000"};
    </script>
</head>
<body class="compactBoxMode veryCompactBoxMode">


<iframe id="splashy" src="/splash_iframe.html?cu=23528" name="emediate23528" style="width: 0px; height: 0px; border: 0px; display: none;"></iframe>

<div id="super_wrapper">

<script type="text/javascript" src="/static/js/pbt.js?v=1415624586"></script>

<div id="jplayer" class="tyda_jplayer"></div>
<div id="wrapper">
    <div id="accept-cookies-div">
        <style>
        #accept-cookies-div {
            width:100%;
            padding:10px 0;
            background: #EAEAEA;
            color: #000;
            font-size:14px;
            text-align: center;
            margin-top:3px;
        }
        </style>
        <script type="text/javascript">
            function n24_accept_cookies() {
                var d = new Date();
                d.setTime(d.getTime() + (365 * 24 * 60 * 60 * 1000));
                document.cookie = "accepted-cookies=1; path=/; expires=" + d.toGMTString();
                var el = document.getElementById('accept-cookies-div');
                el.parentNode.removeChild(el);
            }
        </script>
        Som besökare på Tyda samtycker du till användandet av s.k. <a target="_blank" href="http://nyheter24gruppen.se/integritetspolicy">cookies</a> för att förbättra din upplevelse hos oss.
        <a href="#" id="accept-cookies-button" onclick="n24_accept_cookies(); return false;" target="_blank">Jag förstår, ta bort denna ruta!</a>
    </div>
  <ul id="social-bar">
    <li class="item item-fb"><a title="Hitta oss på facebook" href="https://www.facebook.com/tyda.se"></a></li>
    <li class="item item-tw"><a title="Följ oss på Twitter" href="https://twitter.com/TydaSe"></a></li>
  </ul>
  <div id="page">
    <div class="advertise-skin" id="advertise-skin">
      <div class="wrap">
      <div class="topMargin"></div>
        <div id="header" class="search-box-size">
          <h1 class="header-large"><a href="/"><img src="/static/img/logo-large.png" alt="Tyda" title="Tyda" /></a></h1>
                  </div>
        <div id="search-box" class="box search-box-size">
  <form class="form-search" autocomplete="off" method="POST" action="/s">
    <div class="top">
      <div class="search-fields clearfix">
        <h1 class="header-small"><a href="/"><img src="/static/img/logo-small.png" alt="Tyda" title="Tyda" /></a></h1>
        <div class="fields clearfix">
          <div class="search-fields-item search-fields-input-holder">
            <a class="latin-button" href="#"><img src="/static/img/icon-keyboard.png" alt="" /></a>
            <div>
                              <input id="search-input" tabindex="1" type="text" placeholder="Slå upp ord i Sveriges största lexikon" name="word" class="input input-text input-text-search" value="himmel" autofocus>
                <script type="text/javascript">
                  document.getElementById('search-input').focus();
                  document.getElementById('search-input').select();
                </script>
                            <div class="autocomplete-wrapper">
                <div class="autocomplete"></div>
              </div>
            </div>
          </div>
          <div class="search-fields-item search-fields-button-holder">
            <input type="submit" value="" class="button-search" tabindex="2"/>
          </div>
        </div>
        <div class="latin clearfix">
          <div class="latin-holder">
          </div>
        </div>
      </div>
    </div>
    <div class="middle clearfix">
            <div class="language-selector clearfix">
        <select class="select-language" multiple name="lang[]" size="2" style="display:none;">
                          <option data-languageid="en" value="en"  selected="selected"  >Engelska</option>
                          <option data-languageid="fr" value="fr"  selected="selected"  >Franska</option>
                          <option data-languageid="de" value="de"  selected="selected"  >Tyska</option>
                          <option data-languageid="es" value="es"  selected="selected"  >Spanska</option>
                          <option data-languageid="la" value="la"  selected="selected"  >Latin</option>
                          <option data-languageid="nb" value="nb"  selected="selected"  >Norska</option>
                        <option data-languageid="sv" value="sv"  selected="selected"  >Svenska</option>
        </select>
        <ul class="list list-flags">
                                  <li title="Engelska" class="item selectable  choose  selected " data-languageid="en"><img src="/static/img/flags/en.jpg" alt="Engelska"></li>
                                  <li title="Franska" class="item selectable  choose  selected " data-languageid="fr"><img src="/static/img/flags/fr.jpg" alt="Franska"></li>
                                  <li title="Tyska" class="item selectable  choose  selected " data-languageid="de"><img src="/static/img/flags/de.jpg" alt="Tyska"></li>
                                  <li title="Spanska" class="item selectable  choose  selected " data-languageid="es"><img src="/static/img/flags/es.jpg" alt="Spanska"></li>
                                  <li title="Latin" class="item selectable  choose  selected " data-languageid="la"><img src="/static/img/flags/la.jpg" alt="Latin"></li>
                                  <li title="Norska" class="item selectable  choose  selected " data-languageid="nb"><img src="/static/img/flags/nb.jpg" alt="Norska"></li>
                  </ul>
      </div>
      <div class="language-selector clearfix" style="margin-right:10px; border:0; background:none;">
      <div class="arrowboth"></div>
      </div>
      <div class="language-selector clearfix" style="margin-right:10px;">
        <ul class="list list-flags">
          <li title="Svenska" class="item selected" data-languageid="sv"><img src="/static/img/flags/sv.jpg" alt="Svenska"></li>
        </ul>
      </div>
    </div>
    <div class="bottom">
      <ul class="nav nav-menu clearfix">
        <li class="item"><a href="/forum" title="Läs och diskutera i forumet">Forum</a></li>
        <li class="item"><a href="/tools" title="Verktyg för bättre tydande ...">Verktyg</a></li>
        <li class="item"><a href="/random" title="Visa ett ord slumpmässigt">Slumpa ord</a></li>
        <li class="item"><a href="/news" title="">Nyheter</a></li>
        <li class="item"><a href="/about" title="Kontakt, villkor, cookie-policy ...">Om Tyda</a></li>
        <li class="item"><a href="/panel">Panelen</a></li>
      </ul>
              <ul class="nav nav-help clearfix">
          <li class="item"><a href="/register" title="">Bli medlem</a></li>
          <li class="item"><a class="colored" href="/login" title="">Logga in</a></li>
        </ul>
      
    </div>
  </form>
</div>
      </div>
      <div class="ad ad-panorama">
  <center>
    <iframe id="emediate-19773" src="/static/html/ad.html?v=1413826447" name="emediate:19773:1467550798" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
  </center>
</div>
      
    </div>
    <div class="fiskpinne-wrap">
          </div>
    <div class="wrap wrap-content">
      <div class="divider"></div>
      <div id="content" class="clearfix">
        <div class="content-body">
<div class="page-searchresult">
  <div class="box box-searchresult"><h2 id="sv-nn"><b>himmel</b>
</h2>
<span class="conjugation nnd" title="Substantiv">himlen</span><span class="conjugation missing">(-)</span><span class="conjugation nnsd" title="Substantiv">himlarna</span><div class="word-class" id="899063">
  <span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt=""></span> Substantiv</div>
<div class="box-caption clearfix" id="sense-899063">
  <span class="displayer" title="Förekomst och ämnesområde">
      </span>
  <div class="social">

    <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=http%3A%2F%2Ftyda.se%2Fsearch%2Fhimmel%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-899063"><img src="/static/img/icon-fb.png" alt="" /></a>
    <a target="_blank" href="https://twitter.com/intent/tweet?text=himmel&url=http%3A%2F%2Ftyda.se%2Fsearch%2Fhimmel%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-899063&via=TydaSe"><img src="/static/img/icon-tw.png" alt="" /></a>
  </div>
</div>
<div class="capsulated-content">
    <div class="block-section clearfix">
    <div class="block block1-2">
                                <div class="image-wrapper">
            <img class="stretcher" src="/static/img/imagestretcher.png" alt="" /><img class="image" src="/img/9314ced3483e4ae00aae2fd699efbfd1.jpg" alt="" />
          </div>
                    <div class="photo-author">
                    <a href="/image/9314ced3483e4ae00aae2fd699efbfd1">Visa bild</a>
                                      Fotograf:                               WitekTHC                                  </div>
                                    <div class="description">
          Himlen är jordens atmosfär sedd från jordytan.<br />
<br />
Himlen är blå (eller grå vid dåligt väder) på dagen och med nyanser av rött eller gult vid soluppgång och solnedgång, på grund av Rayleigh-spridning av solljus i atmosfären.[1] På natten är himlen mörk och när det inte finns moln kan man se månen (när den är över horisonten) och stjärnorna (stjärnhimlen). På månen och alla andra himlakroppar som saknar atmosfär är det ständigt svart och stjärnorna syns där även i dagsljus. Himlen kan ibland definieras som den tätare zonen av en planets atmosfär.<br />
<br />
<a target="_blank" href="http://sv.wikipedia.org/wiki/Himmel"><img src="/static/img/icons/3rdparty/wikipedia.png?v=1413826447" height="24" width="24"></a>        </div>
            
      
    </div>
    <div class="block block1-2">
            <ul class="list list-translations">
                                                                                                                                              <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/en.png" alt="" /></span> Engelska</li>
                      <li class="item">
              <a href="/search/blue?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">blue</a>
                                                        <div class="extra">
                <a class="icon-item speaker" title="Läs upp uttal för blue" href="/audio/1459805"><img src="/static/img/icon-speaker.png" alt="" /></a>
                <a class="icon-item mm" href="http://www.macmillandictionary.com/search/british/?q=blue" target="_blank" title="Mer information på macmillan.com"><img src="/static/img/icon-mm.png" alt="" /></a>
              </div>
                          </li>
                      <li class="item">
              <a href="/search/sky?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">sky</a>
                                                        <div class="extra">
                <a class="icon-item speaker" title="Läs upp uttal för sky" href="/audio/1424413"><img src="/static/img/icon-speaker.png" alt="" /></a>
                <a class="icon-item mm" href="http://www.macmillandictionary.com/search/british/?q=sky" target="_blank" title="Mer information på macmillan.com"><img src="/static/img/icon-mm.png" alt="" /></a>
              </div>
                          </li>
                      <li class="item">
              <a href="/search/heaven?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">heaven</a>
                                                        <div class="extra">
                <a class="icon-item speaker" title="Läs upp uttal för heaven" href="/audio/1330677"><img src="/static/img/icon-speaker.png" alt="" /></a>
                <a class="icon-item mm" href="http://www.macmillandictionary.com/search/british/?q=heaven" target="_blank" title="Mer information på macmillan.com"><img src="/static/img/icon-mm.png" alt="" /></a>
              </div>
                          </li>
                      <li class="item">
              <a href="/search/cloud?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">cloud</a>
                                            <span class="trans-desc" title="Förekomst och ämnesområde">[ bildligt ]</span>
                                          <div class="extra">
                <a class="icon-item speaker" title="Läs upp uttal för cloud" href="/audio/1326017"><img src="/static/img/icon-speaker.png" alt="" /></a>
                <a class="icon-item mm" href="http://www.macmillandictionary.com/search/british/?q=cloud" target="_blank" title="Mer information på macmillan.com"><img src="/static/img/icon-mm.png" alt="" /></a>
              </div>
                          </li>
                    <li class="divider"></li>          <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/fr.png" alt="" /></span> Franska</li>
                      <li class="item">
              <a href="/search/ciel?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">ciel</a>
                                                      </li>
                    <li class="divider"></li>          <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/de.png" alt="" /></span> Tyska</li>
                      <li class="item">
              <a href="/search/Himmel?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">Himmel</a>
                                                      </li>
                    <li class="divider"></li>          <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/es.png" alt="" /></span> Spanska</li>
                      <li class="item">
              <a href="/search/cielo?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">cielo</a>
                                                      </li>
                    <li class="divider"></li>          <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/la.png" alt="" /></span> Latin</li>
                      <li class="item">
              <a href="/search/caelum?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">caelum</a>
                                                      </li>
                    <li class="divider"></li>          <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/nb.png" alt="" /></span> Norska</li>
                      <li class="item">
              <a href="/search/himmel?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">himmel</a>
                                                      </li>
                          </ul>
          </div>
  </div>
            <div class="divider"></div>
        <h5>Synonymer</h5>
  <ul class="list list-synonyms">
          <li class="item">
        <a href="/search/sky?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">sky</a>
                          <span class="syn-desc">
            [ meteorologi ]
          </span>
              </li>
      </ul>
      
</div>

            </div><div class="box box-searchresult"><h2 id="de-nn">Der&nbsp;<b>Himmel</b>
</h2>
<div class="word-class" id="1921033">
  <span class="flag-small"><img src="/static/img/flagssmall/de.png" alt=""></span> Substantiv</div>
<div class="box-caption clearfix" id="sense-1921033">
  <span class="displayer" title="Förekomst och ämnesområde">
      </span>
  <div class="social">

    <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=http%3A%2F%2Ftyda.se%2Fsearch%2Fhimmel%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-1921033"><img src="/static/img/icon-fb.png" alt="" /></a>
    <a target="_blank" href="https://twitter.com/intent/tweet?text=Himmel&url=http%3A%2F%2Ftyda.se%2Fsearch%2Fhimmel%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-1921033&via=TydaSe"><img src="/static/img/icon-tw.png" alt="" /></a>
  </div>
</div>
<div class="capsulated-content">
    <div class="block-section clearfix">
    <div class="block block1-2">
                                <div class="image-wrapper">
            <img class="stretcher" src="/static/img/imagestretcher.png" alt="" /><img class="image" src="/img/9314ced3483e4ae00aae2fd699efbfd1.jpg" alt="" />
          </div>
                    <div class="photo-author">
                    <a href="/image/9314ced3483e4ae00aae2fd699efbfd1">Visa bild</a>
                                      Fotograf:                               WitekTHC                                  </div>
                          
      
    </div>
    <div class="block block1-2">
            <ul class="list list-translations">
                                                    <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt="" /></span> Svenska</li>
                      <li class="item">
              <a href="/search/himmel?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">himmel</a>
                                                      </li>
                          </ul>
          </div>
  </div>
        
</div>

            </div><div class="box box-searchresult"><h2 id="nb-nn"><b>himmel</b>
</h2>
<div class="word-class" id="4039781">
  <span class="flag-small"><img src="/static/img/flagssmall/nb.png" alt=""></span> Substantiv</div>
<div class="box-caption clearfix" id="sense-4039781">
  <span class="displayer" title="Förekomst och ämnesområde">
      </span>
  <div class="social">

    <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=http%3A%2F%2Ftyda.se%2Fsearch%2Fhimmel%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-4039781"><img src="/static/img/icon-fb.png" alt="" /></a>
    <a target="_blank" href="https://twitter.com/intent/tweet?text=himmel&url=http%3A%2F%2Ftyda.se%2Fsearch%2Fhimmel%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-4039781&via=TydaSe"><img src="/static/img/icon-tw.png" alt="" /></a>
  </div>
</div>
<div class="capsulated-content">
    <div class="block-section clearfix">
    <div class="block block1-2">
                                <div class="image-wrapper">
            <img class="stretcher" src="/static/img/imagestretcher.png" alt="" /><img class="image" src="/img/9314ced3483e4ae00aae2fd699efbfd1.jpg" alt="" />
          </div>
                    <div class="photo-author">
                    <a href="/image/9314ced3483e4ae00aae2fd699efbfd1">Visa bild</a>
                                      Fotograf:                               WitekTHC                                  </div>
                          
      
    </div>
    <div class="block block1-2">
            <ul class="list list-translations">
                                                    <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt="" /></span> Svenska</li>
                      <li class="item">
              <a href="/search/himmel?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">himmel</a>
                                                      </li>
                          </ul>
          </div>
  </div>
        
</div>

              </div>
  </div>
        </div>
        <div class="content-sidebar">
            <div class="summary">
    <ul class="list-types box">
    
              <li class="item">
          <span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt=""></span>
                      <a href="#sv-nn">Substantiv</a> (1)
                  </li>
              <li class="item">
          <span class="flag-small"><img src="/static/img/flagssmall/de.png" alt=""></span>
                      <a href="#de-nn">Substantiv</a> (1)
                  </li>
              <li class="item">
          <span class="flag-small"><img src="/static/img/flagssmall/nb.png" alt=""></span>
                      <a href="#nb-nn">Substantiv</a> (1)
                  </li>
          </ul>
  </div>
                              <div class="box box-clean ads ads-sidebar">
            <div class="ad ad-sidebar ad-sidebar-1 ad-desktop">
  <textarea>
    <iframe id="emediate-16369" src="/static/html/ad2.html?v=1413826447" name="emediate:16369:1467550798" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
  </textarea>
</div>
            <div class="ad ad-sidebar ad-sidebar-1 ad-desktop">
  <textarea>
    <iframe id="emediate-16370" src="/static/html/ad2.html?v=1413826447" name="emediate:16370:1467550798" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
  </textarea>
</div>
            <div class="ad ad-sidebar ad-sidebar-3 ad-desktop ad-sticky">
  <div class="sticky-placeholder"></div>
  <div class="sticky-child">
    <textarea>
                  <iframe id="emediate-16371" src="/static/html/ad2.html?v=1413826447" name="emediate:16371:1467550798" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
    </textarea>
  </div>
</div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div id="footer">
    <div class="wrap">
      <div class="footer-content-wrapper">
        <div class="footer-margin">
                    <div class="block macmillan">
            <div class="left"><img src="/static/img/macmillan.jpg" alt="" /></div>
            <div class="desc">Uttalsljuden tillhandahålls av Macmillan Dictionary<br/>- Online english Dictionary and Thesaurus</div>
          </div>
                    <ul class="nav nav-group clearfix">
            <li class="item text">I tydagruppen:</li>
            <li class="item"><a href="http://tyda.se/" title="Ett lexikon för översättning med över 1 miljon uppslagsord på svenska, engelska, tyska, franska och spanska"><img src="/static/img/logo-tyda-small.png" alt="Tyda" /></a></li>
            <li class="item"><a href="http://korsord.cc/" title="nyttiga verktyg för dig som tycker om att lösa korsord"><img src="/static/img/logo-korsord-small.png" alt="Korsord" /></a></li>
            <li class="item"><a href="http://synonymer.cc/" title="Synonymlexikon med över 600.000 synonymer på svenska och engelska och integrerad tesaurus (associationsordbok)"><img src="/static/img/logo-synonymer-small.png" alt="Synonymer" /></a></li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/js/jquery.cookie.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/jquery.ddslick.min.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/jquery.jplayer.min.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/stickyMojo.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/default.js?v=1413826447"></script>

<!-- AT internet -->
<script type="text/javascript">
  xtnv = document;
  xtsd = "http://logc168";
  xtsite = "522730";
  xtn2 = "4";
</script>
<script type="text/javascript" src="/static/js/xtcore.js"></script>
<noscript><img width="1" height="1" alt="" src="http://logc168.xiti.com/hit.xiti?s=522730&amp;s2=4"/></noscript>
<!-- END AT internet -->
<!-- BEGIN adblock -->
<script type="text/javascript">
    check_ab = true;
</script>
<script type="text/javascript" src="/static/js/adsense.js"></script>
<script type="text/javascript">
if (check_ab) {
    adblock_state = 'Adblock Detected';
} else {
    adblock_state = "No Adblock";
}
</script>

<!-- END adblock -->

<!-- GA -->
<script type="text/javascript">
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
ga('create', 'UA-41030504-1', 'tyda.se');
ga('require', 'displayfeatures');
ga('send', 'pageview', { 'dimension1': adblock_state });
</script>
<!-- END GA -->

<!-- START UserEcho -->
<script type='text/javascript'>

var _ues = {
host:'tyda.userecho.com',
forum:'23258',
lang:'en',
tab_corner_radius:5,
tab_font_size:20,
tab_image_hash:'ZmVlZGJhY2s%3D',
tab_chat_hash:'Y2hhdA%3D%3D',
tab_alignment:'right',
tab_text_color:'#FFFFFF',
tab_text_shadow_color:'#00000055',
tab_bg_color:'#1EA56A',
tab_hover_color:'#188F6C'
};

(function() {
    var _ue = document.createElement('script'); _ue.type = 'text/javascript'; _ue.async = true;
    _ue.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + 'cdn.userecho.com/js/widget-1.4.gz.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(_ue, s);
  })();

</script>
<!-- END UserEcho -->
<iframe id="emediate-16747" src="/static/html/ad.html?v=1413826447" name="emediate:16747:1467550798" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
</div>
</body>
</html>
`

const ConjugationAll = `
<!doctype html>
<html>
<head>
  <script>preloaded_emediate_pageviewid = (new Date()).getTime()+"_"+Math.floor(Math.random()*100000);</script>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta charset="utf-8">
  <title>Tyda.se - Resultat för "conjugation"</title>
  <meta name="description" content="Ett lexikon för översättning med över 1 miljon uppslagsord på svenska, engelska, tyska, franska och spanska">
  <meta name="keywords" content="lexikon engelska svenska tyska franska spanska översättning översätta ordlista ordbok betyder">
  <meta name="PUBLISHER" content="Tyda Sverige AB">
  <meta name="URL" content="http://tyda.se/">
  <link rel="search" href="http://tyda.se/search-engines" type="application/opensearchdescription+xml">
  <link rel="icon" sizes="32x32" href="/static/img/tyda-favicon-32" type="image/png"/>
  <link rel="icon" sizes="64x64" href="/static/img/tyda-favicon-64" type="image/png"/>
  <link rel="apple-touch-icon" sizes="64x64" href="/static/img/tyda-favicon-64" type="image/png"/>
  <meta name="google-site-verification" content="qPVeUQYR82XH_LmvDZHBEoscMhFEBHd730fj2AevaMM" />
    <meta name="mobile-web-app-capable" content="yes">

  <link href="http://fonts.googleapis.com/css?family=Open+Sans:300,400,500,600,700,800" rel="stylesheet" type="text/css">
  <link rel="stylesheet" type="text/css" href="/static/css/default.css?v=1453979632" media="screen" />
  <link rel="stylesheet" type="text/css" media="only screen and (max-width: 600px), only screen and (max-device-width: 600px)" href="/static/css/mobile.css?v=1454931523" />
  
      <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  
  <!--[if IE]>
    <style type="text/css">
      #social-bar li.item {
        background-color: #fffffd;
      }

      .box {
        background-color: #fffffd;
      }

      .input-text-search {
        border-color: #C0C0C0;
      }
    </style>
  <![endif]-->
  <meta name="twitter:card" content="summary_large_image">
  <meta name="twitter:domain" content="tyda.se">
            <meta name="twitter:title" content="Sökresultat för conjugation">

                <meta name="twitter:description" content="the inflection of verbs">

                  <meta property="og:type" content="website">
      <meta property="og:url" content="http://tyda.se/search/conjugation?lang%5B0%5D=en&amp;lang%5B1%5D=fr&amp;lang%5B2%5D=de&amp;lang%5B3%5D=es&amp;lang%5B4%5D=la&amp;lang%5B5%5D=nb&amp;lang%5B6%5D=sv">
      <meta property="og:site_name" content="Tyda.se">
      <meta property="og:description" content="the inflection of verbs">
      <meta property="og:title" content="Sökresultat för conjugation">
    <meta property="fb:admins" content="1563034500" />
    <script type="text/javascript">
        n24g = {"nuggad":"2212222222222221222222221229999999939994999499999949393922121212222222222222012241653232221229999999999999999919011000"};
    </script>
</head>
<body class="compactBoxMode veryCompactBoxMode">


<iframe id="splashy" src="/splash_iframe.html?cu=23528" name="emediate23528" style="width: 0px; height: 0px; border: 0px; display: none;"></iframe>

<div id="super_wrapper">

<script type="text/javascript" src="/static/js/pbt.js?v=1415624586"></script>

<div id="jplayer" class="tyda_jplayer"></div>
<div id="wrapper">
    <div id="accept-cookies-div">
        <style>
        #accept-cookies-div {
            width:100%;
            padding:10px 0;
            background: #EAEAEA;
            color: #000;
            font-size:14px;
            text-align: center;
            margin-top:3px;
        }
        </style>
        <script type="text/javascript">
            function n24_accept_cookies() {
                var d = new Date();
                d.setTime(d.getTime() + (365 * 24 * 60 * 60 * 1000));
                document.cookie = "accepted-cookies=1; path=/; expires=" + d.toGMTString();
                var el = document.getElementById('accept-cookies-div');
                el.parentNode.removeChild(el);
            }
        </script>
        Som besökare på Tyda samtycker du till användandet av s.k. <a target="_blank" href="http://nyheter24gruppen.se/integritetspolicy">cookies</a> för att förbättra din upplevelse hos oss.
        <a href="#" id="accept-cookies-button" onclick="n24_accept_cookies(); return false;" target="_blank">Jag förstår, ta bort denna ruta!</a>
    </div>
  <ul id="social-bar">
    <li class="item item-fb"><a title="Hitta oss på facebook" href="https://www.facebook.com/tyda.se"></a></li>
    <li class="item item-tw"><a title="Följ oss på Twitter" href="https://twitter.com/TydaSe"></a></li>
  </ul>
  <div id="page">
    <div class="advertise-skin" id="advertise-skin">
      <div class="wrap">
      <div class="topMargin"></div>
        <div id="header" class="search-box-size">
          <h1 class="header-large"><a href="/"><img src="/static/img/logo-large.png" alt="Tyda" title="Tyda" /></a></h1>
                  </div>
        <div id="search-box" class="box search-box-size">
  <form class="form-search" autocomplete="off" method="POST" action="/s">
    <div class="top">
      <div class="search-fields clearfix">
        <h1 class="header-small"><a href="/"><img src="/static/img/logo-small.png" alt="Tyda" title="Tyda" /></a></h1>
        <div class="fields clearfix">
          <div class="search-fields-item search-fields-input-holder">
            <a class="latin-button" href="#"><img src="/static/img/icon-keyboard.png" alt="" /></a>
            <div>
                              <input id="search-input" tabindex="1" type="text" placeholder="Slå upp ord i Sveriges största lexikon" name="word" class="input input-text input-text-search" value="conjugation" autofocus>
                <script type="text/javascript">
                  document.getElementById('search-input').focus();
                  document.getElementById('search-input').select();
                </script>
                            <div class="autocomplete-wrapper">
                <div class="autocomplete"></div>
              </div>
            </div>
          </div>
          <div class="search-fields-item search-fields-button-holder">
            <input type="submit" value="" class="button-search" tabindex="2"/>
          </div>
        </div>
        <div class="latin clearfix">
          <div class="latin-holder">
          </div>
        </div>
      </div>
    </div>
    <div class="middle clearfix">
            <div class="language-selector clearfix">
        <select class="select-language" multiple name="lang[]" size="2" style="display:none;">
                          <option data-languageid="en" value="en"  selected="selected"  >Engelska</option>
                          <option data-languageid="fr" value="fr"  selected="selected"  >Franska</option>
                          <option data-languageid="de" value="de"  selected="selected"  >Tyska</option>
                          <option data-languageid="es" value="es"  selected="selected"  >Spanska</option>
                          <option data-languageid="la" value="la"  selected="selected"  >Latin</option>
                          <option data-languageid="nb" value="nb"  selected="selected"  >Norska</option>
                        <option data-languageid="sv" value="sv"  selected="selected"  >Svenska</option>
        </select>
        <ul class="list list-flags">
                                  <li title="Engelska" class="item selectable  choose  selected " data-languageid="en"><img src="/static/img/flags/en.jpg" alt="Engelska"></li>
                                  <li title="Franska" class="item selectable  choose  selected " data-languageid="fr"><img src="/static/img/flags/fr.jpg" alt="Franska"></li>
                                  <li title="Tyska" class="item selectable  choose  selected " data-languageid="de"><img src="/static/img/flags/de.jpg" alt="Tyska"></li>
                                  <li title="Spanska" class="item selectable  choose  selected " data-languageid="es"><img src="/static/img/flags/es.jpg" alt="Spanska"></li>
                                  <li title="Latin" class="item selectable  choose  selected " data-languageid="la"><img src="/static/img/flags/la.jpg" alt="Latin"></li>
                                  <li title="Norska" class="item selectable  choose  selected " data-languageid="nb"><img src="/static/img/flags/nb.jpg" alt="Norska"></li>
                  </ul>
      </div>
      <div class="language-selector clearfix" style="margin-right:10px; border:0; background:none;">
      <div class="arrowboth"></div>
      </div>
      <div class="language-selector clearfix" style="margin-right:10px;">
        <ul class="list list-flags">
          <li title="Svenska" class="item selected" data-languageid="sv"><img src="/static/img/flags/sv.jpg" alt="Svenska"></li>
        </ul>
      </div>
    </div>
    <div class="bottom">
      <ul class="nav nav-menu clearfix">
        <li class="item"><a href="/forum" title="Läs och diskutera i forumet">Forum</a></li>
        <li class="item"><a href="/tools" title="Verktyg för bättre tydande ...">Verktyg</a></li>
        <li class="item"><a href="/random" title="Visa ett ord slumpmässigt">Slumpa ord</a></li>
        <li class="item"><a href="/news" title="">Nyheter</a></li>
        <li class="item"><a href="/about" title="Kontakt, villkor, cookie-policy ...">Om Tyda</a></li>
        <li class="item"><a href="/panel">Panelen</a></li>
      </ul>
              <ul class="nav nav-help clearfix">
          <li class="item"><a href="/register" title="">Bli medlem</a></li>
          <li class="item"><a class="colored" href="/login" title="">Logga in</a></li>
        </ul>
      
    </div>
  </form>
</div>
      </div>
      <div class="ad ad-panorama">
  <center>
    <iframe id="emediate-19773" src="/static/html/ad.html?v=1413826447" name="emediate:19773:1467552809" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
  </center>
</div>
      
    </div>
    <div class="fiskpinne-wrap">
          </div>
    <div class="wrap wrap-content">
      <div class="divider"></div>
      <div id="content" class="clearfix">
        <div class="content-body">
<div class="page-searchresult">
  <div class="box box-searchresult"><h2 id="en-nn"><b>conjugation</b>
  <a class="icon-item speaker" title="Läs upp uttal för conjugation" href="/audio/1424785"><img src="/static/img/icon-speaker.png" alt="" /></a>
</h2>
<span class="conjugation nns" title="Substantiv">conjugations</span><div class="word-class" id="834753">
  <span class="flag-small"><img src="/static/img/flagssmall/en.png" alt=""></span> Substantiv</div>
<div class="box-caption clearfix" id="sense-834753">
  <span class="displayer" title="Förekomst och ämnesområde">
        lingvistik    </span>
  <div class="social">

    <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=http%3A%2F%2Ftyda.se%2Fsearch%2Fconjugation%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-834753"><img src="/static/img/icon-fb.png" alt="" /></a>
    <a target="_blank" href="https://twitter.com/intent/tweet?text=conjugation&url=http%3A%2F%2Ftyda.se%2Fsearch%2Fconjugation%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-834753&via=TydaSe"><img src="/static/img/icon-tw.png" alt="" /></a>
  </div>
</div>
<div class="capsulated-content">
    <div class="block-section clearfix">
    <div class="block block1-2">
                            <div class="description">
          the inflection of verbs        </div>
            
      
    </div>
    <div class="block block1-2">
            <ul class="list list-translations">
                                                    <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt="" /></span> Svenska</li>
                      <li class="item">
              <a href="/search/b%C3%B6jning?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">böjning</a>
                                            <span class="trans-desc" title="Förekomst och ämnesområde">[ lingvistik ]</span>
                                        </li>
                          </ul>
          </div>
  </div>
        
</div>

      <div class="box-caption clearfix" id="sense-1411645">
  <span class="displayer" title="Förekomst och ämnesområde">
        lingvistik    </span>
  <div class="social">

    <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=http%3A%2F%2Ftyda.se%2Fsearch%2Fconjugation%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-1411645"><img src="/static/img/icon-fb.png" alt="" /></a>
    <a target="_blank" href="https://twitter.com/intent/tweet?text=conjugation&url=http%3A%2F%2Ftyda.se%2Fsearch%2Fconjugation%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-1411645&via=TydaSe"><img src="/static/img/icon-tw.png" alt="" /></a>
  </div>
</div>
<div class="capsulated-content">
    <div class="block-section clearfix">
    <div class="block block1-2">
                            <div class="description">
          the complete set of inflected forms of a verb        </div>
            
      
    </div>
    <div class="block block1-2">
            <ul class="list list-translations">
                                                    <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt="" /></span> Svenska</li>
                      <li class="item">
              <a href="/search/konjugation?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">konjugation</a>
                                            <span class="trans-desc" title="Förekomst och ämnesområde">[ lingvistik ]</span>
                                        </li>
                          </ul>
          </div>
  </div>
        
</div>

      <div class="box-caption clearfix" id="sense-1411646">
  <span class="displayer" title="Förekomst och ämnesområde">
      </span>
  <div class="social">

    <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=http%3A%2F%2Ftyda.se%2Fsearch%2Fconjugation%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-1411646"><img src="/static/img/icon-fb.png" alt="" /></a>
    <a target="_blank" href="https://twitter.com/intent/tweet?text=conjugation&url=http%3A%2F%2Ftyda.se%2Fsearch%2Fconjugation%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-1411646&via=TydaSe"><img src="/static/img/icon-tw.png" alt="" /></a>
  </div>
</div>
<div class="capsulated-content">
    <div class="block-section clearfix">
    <div class="block block1-2">
                            <div class="description">
          a class of verbs having the same inflectional forms        </div>
            
      
    </div>
    <div class="block block1-2">
            <ul class="list list-translations">
                                                    <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt="" /></span> Svenska</li>
                      <li class="item">
              <a href="/search/b%C3%B6jningsklass?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">böjningsklass</a>
                                            <span class="trans-desc" title="Förekomst och ämnesområde">[ lingvistik ]</span>
                                        </li>
                          </ul>
          </div>
  </div>
        
</div>

      <div class="box-caption clearfix" id="sense-1411647">
  <span class="displayer" title="Förekomst och ämnesområde">
      </span>
  <div class="social">

    <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=http%3A%2F%2Ftyda.se%2Fsearch%2Fconjugation%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-1411647"><img src="/static/img/icon-fb.png" alt="" /></a>
    <a target="_blank" href="https://twitter.com/intent/tweet?text=conjugation&url=http%3A%2F%2Ftyda.se%2Fsearch%2Fconjugation%3Flang%255B0%255D%3Den%26lang%255B1%255D%3Dfr%26lang%255B2%255D%3Dde%26lang%255B3%255D%3Des%26lang%255B4%255D%3Dla%26lang%255B5%255D%3Dnb%26lang%255B6%255D%3Dsv%23sense-1411647&via=TydaSe"><img src="/static/img/icon-tw.png" alt="" /></a>
  </div>
</div>
<div class="capsulated-content">
    <div class="block-section clearfix">
    <div class="block block1-2">
                  
        <h5>Synonymer</h5>
  <ul class="list list-synonyms">
          <li class="item">
        <a href="/search/coupling?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">coupling</a>
                      </li>
          <li class="item">
        <a href="/search/mating?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">mating</a>
                      </li>
          <li class="item">
        <a href="/search/pairing?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">pairing</a>
                          <span class="syn-desc">
            [ politik ]
          </span>
              </li>
          <li class="item">
        <a href="/search/sexual+union?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">sexual union</a>
                      </li>
          <li class="item">
        <a href="/search/union?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">union</a>
                      </li>
      </ul>

    </div>
    <div class="block block1-2">
            <ul class="list list-translations">
                                                    <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt="" /></span> Svenska</li>
                      <li class="item">
              <a href="/search/sammansm%C3%A4ltning?lang%5B0%5D=en&lang%5B1%5D=fr&lang%5B2%5D=de&lang%5B3%5D=es&lang%5B4%5D=la&lang%5B5%5D=nb&lang%5B6%5D=sv">sammansmältning</a>
                                            <span class="trans-desc" title="Förekomst och ämnesområde">[ teknik ]</span>
                                        </li>
                          </ul>
          </div>
  </div>
  
</div>

              </div>
  </div>
        </div>
        <div class="content-sidebar">
                                        <div class="box box-clean ads ads-sidebar">
            <div class="ad ad-sidebar ad-sidebar-1 ad-desktop">
  <textarea>
    <iframe id="emediate-16369" src="/static/html/ad2.html?v=1413826447" name="emediate:16369:1467552809" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
  </textarea>
</div>
            <div class="ad ad-sidebar ad-sidebar-1 ad-desktop">
  <textarea>
    <iframe id="emediate-16370" src="/static/html/ad2.html?v=1413826447" name="emediate:16370:1467552809" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
  </textarea>
</div>
            <div class="ad ad-sidebar ad-sidebar-3 ad-desktop ad-sticky">
  <div class="sticky-placeholder"></div>
  <div class="sticky-child">
    <textarea>
                  <iframe id="emediate-16371" src="/static/html/ad2.html?v=1413826447" name="emediate:16371:1467552809" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
    </textarea>
  </div>
</div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div id="footer">
    <div class="wrap">
      <div class="footer-content-wrapper">
        <div class="footer-margin">
                    <div class="block macmillan">
            <div class="left"><img src="/static/img/macmillan.jpg" alt="" /></div>
            <div class="desc">Uttalsljuden tillhandahålls av Macmillan Dictionary<br/>- Online english Dictionary and Thesaurus</div>
          </div>
                    <ul class="nav nav-group clearfix">
            <li class="item text">I tydagruppen:</li>
            <li class="item"><a href="http://tyda.se/" title="Ett lexikon för översättning med över 1 miljon uppslagsord på svenska, engelska, tyska, franska och spanska"><img src="/static/img/logo-tyda-small.png" alt="Tyda" /></a></li>
            <li class="item"><a href="http://korsord.cc/" title="nyttiga verktyg för dig som tycker om att lösa korsord"><img src="/static/img/logo-korsord-small.png" alt="Korsord" /></a></li>
            <li class="item"><a href="http://synonymer.cc/" title="Synonymlexikon med över 600.000 synonymer på svenska och engelska och integrerad tesaurus (associationsordbok)"><img src="/static/img/logo-synonymer-small.png" alt="Synonymer" /></a></li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/js/jquery.cookie.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/jquery.ddslick.min.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/jquery.jplayer.min.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/stickyMojo.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/default.js?v=1413826447"></script>

<!-- AT internet -->
<script type="text/javascript">
  xtnv = document;
  xtsd = "http://logc168";
  xtsite = "522730";
  xtn2 = "4";
</script>
<script type="text/javascript" src="/static/js/xtcore.js"></script>
<noscript><img width="1" height="1" alt="" src="http://logc168.xiti.com/hit.xiti?s=522730&amp;s2=4"/></noscript>
<!-- END AT internet -->
<!-- BEGIN adblock -->
<script type="text/javascript">
    check_ab = true;
</script>
<script type="text/javascript" src="/static/js/adsense.js"></script>
<script type="text/javascript">
if (check_ab) {
    adblock_state = 'Adblock Detected';
} else {
    adblock_state = "No Adblock";
}
</script>

<!-- END adblock -->

<!-- GA -->
<script type="text/javascript">
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
ga('create', 'UA-41030504-1', 'tyda.se');
ga('require', 'displayfeatures');
ga('send', 'pageview', { 'dimension1': adblock_state });
</script>
<!-- END GA -->

<!-- START UserEcho -->
<script type='text/javascript'>

var _ues = {
host:'tyda.userecho.com',
forum:'23258',
lang:'en',
tab_corner_radius:5,
tab_font_size:20,
tab_image_hash:'ZmVlZGJhY2s%3D',
tab_chat_hash:'Y2hhdA%3D%3D',
tab_alignment:'right',
tab_text_color:'#FFFFFF',
tab_text_shadow_color:'#00000055',
tab_bg_color:'#1EA56A',
tab_hover_color:'#188F6C'
};

(function() {
    var _ue = document.createElement('script'); _ue.type = 'text/javascript'; _ue.async = true;
    _ue.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + 'cdn.userecho.com/js/widget-1.4.gz.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(_ue, s);
  })();

</script>
<!-- END UserEcho -->
<iframe id="emediate-16747" src="/static/html/ad.html?v=1413826447" name="emediate:16747:1467552809" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
</div>
</body>
</html>
`

const HacerSv = `
<!doctype html>
<html>
<head>
  <script>preloaded_emediate_pageviewid = (new Date()).getTime()+"_"+Math.floor(Math.random()*100000);</script>
  <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
  <meta charset="utf-8">
  <title>Tyda.se - Resultat för "hacer"</title>
  <meta name="description" content="Ett lexikon för översättning med över 1 miljon uppslagsord på svenska, engelska, tyska, franska och spanska">
  <meta name="keywords" content="lexikon engelska svenska tyska franska spanska översättning översätta ordlista ordbok betyder">
  <meta name="PUBLISHER" content="Tyda Sverige AB">
  <meta name="URL" content="http://tyda.se/">
  <link rel="search" href="http://tyda.se/search-engines" type="application/opensearchdescription+xml">
  <link rel="icon" sizes="32x32" href="/static/img/tyda-favicon-32" type="image/png"/>
  <link rel="icon" sizes="64x64" href="/static/img/tyda-favicon-64" type="image/png"/>
  <link rel="apple-touch-icon" sizes="64x64" href="/static/img/tyda-favicon-64" type="image/png"/>
  <meta name="google-site-verification" content="qPVeUQYR82XH_LmvDZHBEoscMhFEBHd730fj2AevaMM" />
    <meta name="mobile-web-app-capable" content="yes">

  <link href="http://fonts.googleapis.com/css?family=Open+Sans:300,400,500,600,700,800" rel="stylesheet" type="text/css">
  <link rel="stylesheet" type="text/css" href="/static/css/default.css?v=1453979674" media="screen" />
  <link rel="stylesheet" type="text/css" media="only screen and (max-width: 600px), only screen and (max-device-width: 600px)" href="/static/css/mobile.css?v=1454931518" />
  
      <meta name="viewport" content="width=device-width, initial-scale=1.0"/>
  
  <!--[if IE]>
    <style type="text/css">
      #social-bar li.item {
        background-color: #fffffd;
      }

      .box {
        background-color: #fffffd;
      }

      .input-text-search {
        border-color: #C0C0C0;
      }
    </style>
  <![endif]-->
  <meta name="twitter:card" content="summary_large_image">
  <meta name="twitter:domain" content="tyda.se">
            <meta name="twitter:title" content="Sökresultat för hacer">

                <meta name="twitter:description" content="Ett lexikon för översättning med över 1 miljon uppslagsord på svenska, engelska, tyska, franska och spanska">

                  <meta property="og:type" content="website">
      <meta property="og:url" content="http://tyda.se/search/hacer?lang%5B0%5D=es&amp;lang%5B1%5D=sv">
      <meta property="og:site_name" content="Tyda.se">
      <meta property="og:description" content="Ett lexikon för översättning med över 1 miljon uppslagsord på svenska, engelska, tyska, franska och spanska">
      <meta property="og:title" content="Sökresultat för hacer">
    <meta property="fb:admins" content="1563034500" />
    <script type="text/javascript">
        n24g = {"nuggad":"2212222222222221222222221229999999939994999499999949393922121212222222222222012241653232221229999999999999999919011000"};
    </script>
</head>
<body class="compactBoxMode veryCompactBoxMode">


<iframe id="splashy" src="/splash_iframe.html?cu=23528" name="emediate23528" style="width: 0px; height: 0px; border: 0px; display: none;"></iframe>

<div id="super_wrapper">

<script type="text/javascript" src="/static/js/pbt.js?v=1415624586"></script>

<div id="jplayer" class="tyda_jplayer"></div>
<div id="wrapper">
    <div id="accept-cookies-div">
        <style>
        #accept-cookies-div {
            width:100%;
            padding:10px 0;
            background: #EAEAEA;
            color: #000;
            font-size:14px;
            text-align: center;
            margin-top:3px;
        }
        </style>
        <script type="text/javascript">
            function n24_accept_cookies() {
                var d = new Date();
                d.setTime(d.getTime() + (365 * 24 * 60 * 60 * 1000));
                document.cookie = "accepted-cookies=1; path=/; expires=" + d.toGMTString();
                var el = document.getElementById('accept-cookies-div');
                el.parentNode.removeChild(el);
            }
        </script>
        Som besökare på Tyda samtycker du till användandet av s.k. <a target="_blank" href="http://nyheter24gruppen.se/integritetspolicy">cookies</a> för att förbättra din upplevelse hos oss.
        <a href="#" id="accept-cookies-button" onclick="n24_accept_cookies(); return false;" target="_blank">Jag förstår, ta bort denna ruta!</a>
    </div>
  <ul id="social-bar">
    <li class="item item-fb"><a title="Hitta oss på facebook" href="https://www.facebook.com/tyda.se"></a></li>
    <li class="item item-tw"><a title="Följ oss på Twitter" href="https://twitter.com/TydaSe"></a></li>
  </ul>
  <div id="page">
    <div class="advertise-skin" id="advertise-skin">
      <div class="wrap">
      <div class="topMargin"></div>
        <div id="header" class="search-box-size">
          <h1 class="header-large"><a href="/"><img src="/static/img/logo-large.png" alt="Tyda" title="Tyda" /></a></h1>
                  </div>
        <div id="search-box" class="box search-box-size">
  <form class="form-search" autocomplete="off" method="POST" action="/s">
    <div class="top">
      <div class="search-fields clearfix">
        <h1 class="header-small"><a href="/"><img src="/static/img/logo-small.png" alt="Tyda" title="Tyda" /></a></h1>
        <div class="fields clearfix">
          <div class="search-fields-item search-fields-input-holder">
            <a class="latin-button" href="#"><img src="/static/img/icon-keyboard.png" alt="" /></a>
            <div>
                              <input id="search-input" tabindex="1" type="text" placeholder="Slå upp ord i Sveriges största lexikon" name="word" class="input input-text input-text-search" value="hacer" autofocus>
                <script type="text/javascript">
                  document.getElementById('search-input').focus();
                  document.getElementById('search-input').select();
                </script>
                            <div class="autocomplete-wrapper">
                <div class="autocomplete"></div>
              </div>
            </div>
          </div>
          <div class="search-fields-item search-fields-button-holder">
            <input type="submit" value="" class="button-search" tabindex="2"/>
          </div>
        </div>
        <div class="latin clearfix">
          <div class="latin-holder">
          </div>
        </div>
      </div>
    </div>
    <div class="middle clearfix">
            <div class="language-selector clearfix">
        <select class="select-language" multiple name="lang[]" size="2" style="display:none;">
                          <option data-languageid="en" value="en"  >Engelska</option>
                          <option data-languageid="fr" value="fr"  >Franska</option>
                          <option data-languageid="de" value="de"  >Tyska</option>
                          <option data-languageid="es" value="es"  selected="selected"  >Spanska</option>
                          <option data-languageid="la" value="la"  >Latin</option>
                          <option data-languageid="nb" value="nb"  >Norska</option>
                        <option data-languageid="sv" value="sv"  selected="selected"  >Svenska</option>
        </select>
        <ul class="list list-flags">
                                  <li title="Engelska" class="item selectable  choose " data-languageid="en"><img src="/static/img/flags/en.jpg" alt="Engelska"></li>
                                  <li title="Franska" class="item selectable  choose " data-languageid="fr"><img src="/static/img/flags/fr.jpg" alt="Franska"></li>
                                  <li title="Tyska" class="item selectable  choose " data-languageid="de"><img src="/static/img/flags/de.jpg" alt="Tyska"></li>
                                  <li title="Spanska" class="item selectable  choose  selected " data-languageid="es"><img src="/static/img/flags/es.jpg" alt="Spanska"></li>
                                  <li title="Latin" class="item selectable  choose " data-languageid="la"><img src="/static/img/flags/la.jpg" alt="Latin"></li>
                                  <li title="Norska" class="item selectable  choose " data-languageid="nb"><img src="/static/img/flags/nb.jpg" alt="Norska"></li>
                  </ul>
      </div>
      <div class="language-selector clearfix" style="margin-right:10px; border:0; background:none;">
      <div class="arrowboth"></div>
      </div>
      <div class="language-selector clearfix" style="margin-right:10px;">
        <ul class="list list-flags">
          <li title="Svenska" class="item selected" data-languageid="sv"><img src="/static/img/flags/sv.jpg" alt="Svenska"></li>
        </ul>
      </div>
    </div>
    <div class="bottom">
      <ul class="nav nav-menu clearfix">
        <li class="item"><a href="/forum" title="Läs och diskutera i forumet">Forum</a></li>
        <li class="item"><a href="/tools" title="Verktyg för bättre tydande ...">Verktyg</a></li>
        <li class="item"><a href="/random" title="Visa ett ord slumpmässigt">Slumpa ord</a></li>
        <li class="item"><a href="/news" title="">Nyheter</a></li>
        <li class="item"><a href="/about" title="Kontakt, villkor, cookie-policy ...">Om Tyda</a></li>
        <li class="item"><a href="/panel">Panelen</a></li>
      </ul>
              <ul class="nav nav-help clearfix">
          <li class="item"><a href="/register" title="">Bli medlem</a></li>
          <li class="item"><a class="colored" href="/login" title="">Logga in</a></li>
        </ul>
      
    </div>
  </form>
</div>
      </div>
      <div class="ad ad-panorama">
  <center>
    <iframe id="emediate-19773" src="/static/html/ad.html?v=1413826447" name="emediate:19773:1467557255" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
  </center>
</div>
      
    </div>
    <div class="fiskpinne-wrap">
          </div>
    <div class="wrap wrap-content">
      <div class="divider"></div>
      <div id="content" class="clearfix">
        <div class="content-body">
<div class="page-searchresult">
  <div class="box box-searchresult"><h2 id="es-vb"><b>hacer</b>
</h2>
<div class="word-class" id="2125281">
  <span class="flag-small"><img src="/static/img/flagssmall/es.png" alt=""></span> Verb</div>
<div class="box-caption clearfix" id="sense-2125281">
  <span class="displayer" title="Förekomst och ämnesområde">
      </span>
  <div class="social">

    <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=http%3A%2F%2Ftyda.se%2Fsearch%2Fhacer%3Flang%255B0%255D%3Des%26lang%255B1%255D%3Dsv%23sense-2125281"><img src="/static/img/icon-fb.png" alt="" /></a>
    <a target="_blank" href="https://twitter.com/intent/tweet?text=hacer&url=http%3A%2F%2Ftyda.se%2Fsearch%2Fhacer%3Flang%255B0%255D%3Des%26lang%255B1%255D%3Dsv%23sense-2125281&via=TydaSe"><img src="/static/img/icon-tw.png" alt="" /></a>
  </div>
</div>
<div class="capsulated-content">
    <div class="block-section clearfix">
    <div class="block block1-2">
                  
      
    </div>
    <div class="block block1-2">
            <ul class="list list-translations">
                                                    <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt="" /></span> Svenska</li>
                      <li class="item">
              <a href="/search/handla?lang%5B0%5D=es&lang%5B1%5D=sv">handla</a>
                                                      </li>
                          </ul>
          </div>
  </div>
  
</div>

      <div class="box-caption clearfix" id="sense-2125285">
  <span class="displayer" title="Förekomst och ämnesområde">
      </span>
  <div class="social">

    <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=http%3A%2F%2Ftyda.se%2Fsearch%2Fhacer%3Flang%255B0%255D%3Des%26lang%255B1%255D%3Dsv%23sense-2125285"><img src="/static/img/icon-fb.png" alt="" /></a>
    <a target="_blank" href="https://twitter.com/intent/tweet?text=hacer&url=http%3A%2F%2Ftyda.se%2Fsearch%2Fhacer%3Flang%255B0%255D%3Des%26lang%255B1%255D%3Dsv%23sense-2125285&via=TydaSe"><img src="/static/img/icon-tw.png" alt="" /></a>
  </div>
</div>
<div class="capsulated-content">
    <div class="block-section clearfix">
    <div class="block block1-2">
                  
      
    </div>
    <div class="block block1-2">
            <ul class="list list-translations">
                                                    <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt="" /></span> Svenska</li>
                      <li class="item">
              <a href="/search/g%C3%B6ra?lang%5B0%5D=es&lang%5B1%5D=sv">göra</a>
                                                      </li>
                          </ul>
          </div>
  </div>
  
</div>

      <div class="box-caption clearfix" id="sense-2125289">
  <span class="displayer" title="Förekomst och ämnesområde">
      </span>
  <div class="social">

    <a target="_blank" href="https://www.facebook.com/sharer/sharer.php?u=http%3A%2F%2Ftyda.se%2Fsearch%2Fhacer%3Flang%255B0%255D%3Des%26lang%255B1%255D%3Dsv%23sense-2125289"><img src="/static/img/icon-fb.png" alt="" /></a>
    <a target="_blank" href="https://twitter.com/intent/tweet?text=hacer&url=http%3A%2F%2Ftyda.se%2Fsearch%2Fhacer%3Flang%255B0%255D%3Des%26lang%255B1%255D%3Dsv%23sense-2125289&via=TydaSe"><img src="/static/img/icon-tw.png" alt="" /></a>
  </div>
</div>
<div class="capsulated-content">
    <div class="block-section clearfix">
    <div class="block block1-2">
                  
      
    </div>
    <div class="block block1-2">
            <ul class="list list-translations">
                                                    <li class="item item-title"><span class="flag-small"><img src="/static/img/flagssmall/sv.png" alt="" /></span> Svenska</li>
                      <li class="item">
              <a href="/search/utf%C3%B6ra?lang%5B0%5D=es&lang%5B1%5D=sv">utföra</a>
                                                      </li>
                      <li class="item">
              <a href="/search/f%C3%A5?lang%5B0%5D=es&lang%5B1%5D=sv">få</a>
                                                        <div class="extra">
                <a class="icon-item speaker" title="Läs upp uttal för få" href="/audio/1387685"><img src="/static/img/icon-speaker.png" alt="" /></a>
                <a class="icon-item mm" href="http://www.macmillandictionary.com/search/british/?q=få" target="_blank" title="Mer information på macmillan.com"><img src="/static/img/icon-mm.png" alt="" /></a>
              </div>
                          </li>
                          </ul>
          </div>
  </div>
  
</div>

              </div>
  </div>
        </div>
        <div class="content-sidebar">
                                        <div class="box box-clean ads ads-sidebar">
            <div class="ad ad-sidebar ad-sidebar-1 ad-desktop">
  <textarea>
    <iframe id="emediate-16369" src="/static/html/ad2.html?v=1413826447" name="emediate:16369:1467557255" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
  </textarea>
</div>
            <div class="ad ad-sidebar ad-sidebar-1 ad-desktop">
  <textarea>
    <iframe id="emediate-16370" src="/static/html/ad2.html?v=1413826447" name="emediate:16370:1467557255" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
  </textarea>
</div>
            <div class="ad ad-sidebar ad-sidebar-3 ad-desktop ad-sticky">
  <div class="sticky-placeholder"></div>
  <div class="sticky-child">
    <textarea>
                  <iframe id="emediate-16371" src="/static/html/ad2.html?v=1413826447" name="emediate:16371:1467557255" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
    </textarea>
  </div>
</div>
          </div>
        </div>
      </div>
    </div>
  </div>
  <div id="footer">
    <div class="wrap">
      <div class="footer-content-wrapper">
        <div class="footer-margin">
                    <div class="block macmillan">
            <div class="left"><img src="/static/img/macmillan.jpg" alt="" /></div>
            <div class="desc">Uttalsljuden tillhandahålls av Macmillan Dictionary<br/>- Online english Dictionary and Thesaurus</div>
          </div>
                    <ul class="nav nav-group clearfix">
            <li class="item text">I tydagruppen:</li>
            <li class="item"><a href="http://tyda.se/" title="Ett lexikon för översättning med över 1 miljon uppslagsord på svenska, engelska, tyska, franska och spanska"><img src="/static/img/logo-tyda-small.png" alt="Tyda" /></a></li>
            <li class="item"><a href="http://korsord.cc/" title="nyttiga verktyg för dig som tycker om att lösa korsord"><img src="/static/img/logo-korsord-small.png" alt="Korsord" /></a></li>
            <li class="item"><a href="http://synonymer.cc/" title="Synonymlexikon med över 600.000 synonymer på svenska och engelska och integrerad tesaurus (associationsordbok)"><img src="/static/img/logo-synonymer-small.png" alt="Synonymer" /></a></li>
          </ul>
        </div>
      </div>
    </div>
  </div>
</div>
<script type="text/javascript" src="//ajax.googleapis.com/ajax/libs/jquery/1.10.1/jquery.min.js"></script>
<script type="text/javascript" src="/static/js/jquery.cookie.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/jquery.ddslick.min.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/jquery.jplayer.min.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/stickyMojo.js?v=1413826447"></script>
<script type="text/javascript" src="/static/js/default.js?v=1413826447"></script>

<!-- AT internet -->
<script type="text/javascript">
  xtnv = document;
  xtsd = "http://logc168";
  xtsite = "522730";
  xtn2 = "4";
</script>
<script type="text/javascript" src="/static/js/xtcore.js"></script>
<noscript><img width="1" height="1" alt="" src="http://logc168.xiti.com/hit.xiti?s=522730&amp;s2=4"/></noscript>
<!-- END AT internet -->
<!-- BEGIN adblock -->
<script type="text/javascript">
    check_ab = true;
</script>
<script type="text/javascript" src="/static/js/adsense.js"></script>
<script type="text/javascript">
if (check_ab) {
    adblock_state = 'Adblock Detected';
} else {
    adblock_state = "No Adblock";
}
</script>

<!-- END adblock -->

<!-- GA -->
<script type="text/javascript">
(function(i,s,o,g,r,a,m){i['GoogleAnalyticsObject']=r;i[r]=i[r]||function(){
(i[r].q=i[r].q||[]).push(arguments)},i[r].l=1*new Date();a=s.createElement(o),
m=s.getElementsByTagName(o)[0];a.async=1;a.src=g;m.parentNode.insertBefore(a,m)
})(window,document,'script','//www.google-analytics.com/analytics.js','ga');
ga('create', 'UA-41030504-1', 'tyda.se');
ga('require', 'displayfeatures');
ga('send', 'pageview', { 'dimension1': adblock_state });
</script>
<!-- END GA -->

<!-- START UserEcho -->
<script type='text/javascript'>

var _ues = {
host:'tyda.userecho.com',
forum:'23258',
lang:'en',
tab_corner_radius:5,
tab_font_size:20,
tab_image_hash:'ZmVlZGJhY2s%3D',
tab_chat_hash:'Y2hhdA%3D%3D',
tab_alignment:'right',
tab_text_color:'#FFFFFF',
tab_text_shadow_color:'#00000055',
tab_bg_color:'#1EA56A',
tab_hover_color:'#188F6C'
};

(function() {
    var _ue = document.createElement('script'); _ue.type = 'text/javascript'; _ue.async = true;
    _ue.src = ('https:' == document.location.protocol ? 'https://' : 'http://') + 'cdn.userecho.com/js/widget-1.4.gz.js';
    var s = document.getElementsByTagName('script')[0]; s.parentNode.insertBefore(_ue, s);
  })();

</script>
<!-- END UserEcho -->
<iframe id="emediate-16747" src="/static/html/ad.html?v=1413826447" name="emediate:16747:1467557255" scrolling="no" scrollBorder="0" frameborder="0" width="100%" height="0" seamless="seamless"></iframe>
</div>
</body>
</html>
`
