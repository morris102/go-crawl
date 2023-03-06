package main

import (
	"log"
	"testing"
)

func TestMain(t *testing.T) {
	doc, err := NewDocument(url)
	if err != nil {
		log.Fatal(err)
	}

	doc.getWordList()
	if len(doc.wordList) == 0 {
		log.Fatal("word list is empty")
	}

	if err := doc.writeFile(); err != nil {
		log.Fatal(err)
	}
}
