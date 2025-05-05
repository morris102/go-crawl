package vocabulary

type WordHtml struct {
	Content string `json:"content"`
	Url     string `json:"url"`
}

type Word struct {
	Kanji     string `json:"kanji"`
	Kana      string `json:"kana"`
	Romaji    string `json:"romaji"`
	Type      string `json:"type"`
	Meaning   string `json:"meaning"`
	JLPTlevel string `json:"level"`
	Url       string `json:"url"`
}
