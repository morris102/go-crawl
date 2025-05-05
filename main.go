package main

import (
	"go-crawl/internal/vocabulary"
	"runtime"
)

func main() {
	var source = "https://japanesetest4you.com/jlpt-n1-vocabulary-list/"

	runtime.GOMAXPROCS(4)

	repo, _ := vocabulary.NewVocalbularyRepository()
	vc := vocabulary.NewVocalbularyUseCaseImpl(repo)
	vc.Crawl(source)
}
