package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/mitchellh/mapstructure"
)

const (
	url            = "https://japanesetest4you.com/jlpt-n1-vocabulary-list/"
	pathOutputFile = "./output.json"
)

type Word struct {
	Japanese string `json:"jp_content"`
	Kanji    string `json:"kanji"`
	Latinh   string `json:"latinh"`
	English  string `json:"en_content"`
	Url      string `json:"url"`
}

type WordDetail struct {
	Kanji     string `json:"kanji"`
	Kana      string `json:"kana"`
	Romaji    string `json:"romaji"`
	Type      string `json:"type"`
	Meaning   string `json:"meaning"`
	JLPTlevel string `json:"level"`
	Url       string `json:"url"`
}

type Docucment struct {
	doc      *goquery.Document
	wordList WordDetailList
}

type WordDetailList []*WordDetail

var wordList WordDetailList

func NewDocument(url string) (*Docucment, error) {
	res, err := callHttpGet(url)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		return nil, err
	}

	return &Docucment{
		doc:      doc,
		wordList: wordList,
	}, nil
}

func (d *Docucment) getWordList() {
	var wg sync.WaitGroup

	d.doc.Find("#content p").Each(func(i int, s *goquery.Selection) {
		if s.Children().Is("a") {
			url, _ := s.Find("a").Attr("href")

			if len(url) > 1 {
				wg.Add(1)

				go func() {
					defer wg.Done()

					wordDetail := getWordDetail(url)
					if wordDetail != nil {
						d.wordList = append(d.wordList, wordDetail)
					}
				}()

			}
		}
	})

	wg.Wait()
}

func getWordDetail(url string) *WordDetail {

	res, _ := callHttpGet(url)
	defer res.Body.Close()

	// Load document
	doc, _ := goquery.NewDocumentFromReader(res.Body)

	mWord := map[string]string{
		"Kana":       "",
		"Kanji":      "",
		"Romaji":     "",
		"Meaning":    "",
		"JLPT level": "",
		"Type":       "",
	}

	word := &WordDetail{
		Url: url,
	}

	doc.Find("#content p").Each(func(i int, s *goquery.Selection) {
		var content = s.Text()
		lines := strings.Split(content, "\n")

		if len(lines) > 0 {
			for i := 0; i < len(lines); i++ {
				for k, v := range mWord {
					if strings.HasPrefix(strings.Trim(lines[i], " "), k) && v == "" {
						v = strings.TrimPrefix(lines[i], k+": ")
						mWord[k] = v
					}
				}
			}
		}

		fmt.Println(mWord)
	})

	mapstructure.Decode(mWord, &word)

	return word
}

func (d *Docucment) writeFile() error {
	var err error
	file, err := os.OpenFile(pathOutputFile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	b, err := json.Marshal(d.wordList)
	if err != nil {
		return err
	}
	_, err = file.Write(b)
	return err
}

func callHttpGet(url string) (*http.Response, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	// defer res.Body.Close()

	return res, err
}

func main() {

	doc, err := NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.getWordList()
	doc.writeFile()

}
