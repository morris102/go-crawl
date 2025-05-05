package vocabulary

import (
	"errors"
	"fmt"
	"go-crawl/common/util"
	"log"
	"strings"
	"sync"

	"github.com/PuerkitoBio/goquery"
	"github.com/mitchellh/mapstructure"
)

type VocalbularyUseCase interface {
	Crawl(url string) ([]WordHtml, error)
	Save(items []Word) error
}

type VocalbularyUseCaseImpl struct {
	VocabularyRepository VocalbularyRepository
}

func NewVocalbularyUseCaseImpl(vocabularyRepository VocalbularyRepository) *VocalbularyUseCaseImpl {
	return &VocalbularyUseCaseImpl{
		VocabularyRepository: vocabularyRepository,
	}
}

func (inst *VocalbularyUseCaseImpl) Crawl(url string) ([]Word, error) {
	var items []WordHtml
	var wordItems []Word
	// call API Resource to get data
	httpClient := util.NewHttpClient()
	resp, err := httpClient.Do(url, "GET", nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		return nil, err
	}

	doc.Find("#content p").Each(func(i int, s *goquery.Selection) {
		if s.Children().Is("a") {
			element := s.Find("a")
			url, _ := element.Attr("href")
			content := element.Text()
			if len(url) > 1 {
				items = append(items, WordHtml{
					Url:     url,
					Content: content,
				})
			}
		}
	})

	var wg sync.WaitGroup
	for i := 0; i < len(items); i++ {
		wg.Add(1)

		var url = items[i].Url
		go func(url string) {
			w, err := crawlDetail(url)
			if err == nil {
				wordItems = append(wordItems, *w)
			}
			wg.Done()
		}(url)
	}
	wg.Wait()

	return wordItems, nil
}

func crawlDetail(url string) (*Word, error) {
	// call API Resource to get data
	httpClient := util.NewHttpClient()
	res, err := httpClient.Do(url, "GET", nil)
	if err != nil {
		return nil, err
	}

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

	word := &Word{
		Url: url,
	}

	doc.Find("#content p").Each(func(i int, s *goquery.Selection) {
		var content = s.Text()
		lines := strings.Split(content, "\n")

		for i := 0; i < len(lines); i++ {
			for k, v := range mWord {
				if strings.HasPrefix(strings.Trim(lines[i], " "), k) && v == "" {
					v = strings.TrimPrefix(lines[i], k+": ")
					mWord[k] = v
				}
			}
		}

		fmt.Println(mWord)
	})

	mapstructure.Decode(mWord, &word)

	return word, nil
}

func (inst *VocalbularyUseCaseImpl) Save(items []Word) error {
	if len(items) == 0 {
		return errors.New("cannot found items")
	}

	err := inst.VocabularyRepository.CreateMany(items)
	if err != nil {
		log.Printf("VocabularyRepository.CreateMany err: %v", err.Error())
	}
	return err
}
