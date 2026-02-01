# go-crawl

Một web crawler được viết bằng Go, thu thập từ vựng JLPT N1 tiếng Nhật từ [Japanesetest4you.com](https://japanesetest4you.com/jlpt-n1-vocabulary-list/).

## Tính năng

- Thu thập danh sách từ vựng JLPT N1 từ trang web kiểm tra tiếng Nhật
- Trích xuất dữ liệu có cấu trúc: Kanji, Kana, Romaji, Meaning, JLPT level
- Xử lý song song để thu thập hiệu quả
- Clean Architecture với sự phân tách rõ ràng

## Cấu trúc dự án

```
go-crawl/
├── main.go                              # Điểm bắt đầu ứng dụng
├── go.mod / go.sum                      # Dependencies của Go module
├── .env                                 # Biến môi trường
├── .gitignore                           # Quy tắc bỏ qua Git
├── Jenkinsfile                          # Pipeline CI/CD
├── output.json                          # Kết quả từ vựng đã thu thập
├── common/
│   └── util/
│       └── http_util.go                 # Tiện ích HTTP client
└── internal/
    └── vocabulary/
        ├── model.go                     # Data models
        ├── vocabulary_repository.go     # Data access layer
        └── vocabulary_usecase.go        # Business logic layer
```

## Kiến trúc

Dự án tuân theo nguyên tắc **Clean Architecture**:

### Các layer

1. **Entry Point** (`main.go`) - Khởi tạo ứng dụng và bắt đầu thu thập
2. **Use Case** (`vocabulary_usecase.go`) - Business logic cho việc thu thập và phân tích
3. **Repository** (`vocabulary_repository.go`) - Interface cho việc lưu trữ dữ liệu
4. **Model** (`model.go`) - Cấu trúc dữ liệu
5. **Utility** (`common/util/`) - Các tiện ích dùng chung

### Data Models

```go
type Word struct {
    Kanji     string `json:"kanji"`
    Kana      string `json:"kana"`
    Romaji    string `json:"romaji"`
    Type      string `json:"type"`
    Meaning   string `json:"meaning"`
    JLPTlevel string `json:"level"`
    Url       string `json:"url"`
}
```

## Cách sử dụng

```bash
# Chạy crawler
go run main.go
```

Từ vựng đã thu thập sẽ được lưu vào `output.json`.

## Dependencies

- [goquery](https://github.com/PuerkitoBio/goquery) - Phân tích HTML
- [mapstructure](https://github.com/mitchellh/mapstructure) - Giải mã map

## License

MIT
