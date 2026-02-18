# go-crawl

[English](#english) | [Tiếng Việt](#tiếng-việt)

---

## English

A web crawler written in Go that scrapes JLPT N1 Japanese vocabulary from [Japanesetest4you.com](https://japanesetest4you.com/jlpt-n1-vocabulary-list/).

### Features

- 🚀 **Concurrent Processing**: Utilizes Go's goroutines for parallel vocabulary scraping
- 📚 **Structured Data Extraction**: Extracts Kanji, Kana, Romaji, Meaning, Type, and JLPT level
- 🏗️ **Clean Architecture**: Well-organized code with clear separation of concerns
- 💾 **JSON Output**: Saves collected vocabulary data in structured JSON format
- 🔄 **HTTP Client Wrapper**: Custom HTTP client utility for reliable web requests

### Prerequisites

- **Go 1.18** or higher
- Internet connection for scraping

### Installation

1. Clone the repository:
```bash
git clone https://github.com/morris102/go-crawl.git
cd go-crawl
```

2. Install dependencies:
```bash
go mod download
```

3. (Optional) Configure environment variables in `.env`:
```bash
DB_HOST=localhost
DB_PORT=5432
```

### Usage

#### Basic Usage

Run the crawler to scrape JLPT N1 vocabulary:

```bash
go run main.go
```

The scraped vocabulary will be saved to `output.json` in the project root directory.

#### Building the Binary

Build an executable binary:

```bash
go build -o go-crawl main.go
```

Run the binary:

```bash
./go-crawl
```

### Project Structure

```
go-crawl/
├── main.go                              # Application entry point
├── go.mod / go.sum                      # Go module dependencies
├── .env                                 # Environment variables
├── .gitignore                           # Git ignore rules
├── Jenkinsfile                          # CI/CD pipeline configuration
├── output.json                          # Scraped vocabulary output
├── common/
│   └── util/
│       └── http_util.go                 # HTTP client utility
└── internal/
    └── vocabulary/
        ├── model.go                     # Data models
        ├── vocabulary_repository.go     # Data access layer
        └── vocabulary_usecase.go        # Business logic layer
```

### Architecture

The project follows **Clean Architecture** principles with clear separation of concerns:

#### Layers

1. **Entry Point** (`main.go`)
   - Initializes the application
   - Sets up GOMAXPROCS for parallel processing
   - Starts the crawling process

2. **Use Case Layer** (`vocabulary_usecase.go`)
   - Contains business logic for crawling and parsing
   - Orchestrates parallel scraping using goroutines
   - Implements `VocalbularyUseCase` interface

3. **Repository Layer** (`vocabulary_repository.go`)
   - Data persistence interface
   - Implements `VocalbularyRepository` interface
   - Can be extended to save to databases

4. **Model Layer** (`model.go`)
   - Defines data structures
   - `Word`: Represents a complete vocabulary entry
   - `WordHtml`: Represents HTML content before parsing

5. **Utility Layer** (`common/util/`)
   - Shared utilities
   - HTTP client wrapper for making web requests

### Data Models

#### Word
Represents a complete Japanese vocabulary entry:

```go
type Word struct {
    Kanji     string `json:"kanji"`      // Japanese characters
    Kana      string `json:"kana"`       // Hiragana/Katakana reading
    Romaji    string `json:"romaji"`     // Romanized reading
    Type      string `json:"type"`       // Word type (noun, verb, etc.)
    Meaning   string `json:"meaning"`    // English meaning
    JLPTlevel string `json:"level"`      // JLPT difficulty level
    Url       string `json:"url"`        // Source URL
}
```

#### WordHtml
Represents raw HTML content before detailed parsing:

```go
type WordHtml struct {
    Content string `json:"content"`      // Link text
    Url     string `json:"url"`          // Link URL
}
```

### How It Works

1. **Initial Scraping**: The crawler fetches the main vocabulary list page
2. **Link Extraction**: Extracts all vocabulary detail page URLs
3. **Parallel Processing**: Uses goroutines and WaitGroups to scrape details concurrently
4. **Data Parsing**: Parses HTML using goquery to extract structured vocabulary data
5. **Data Mapping**: Uses mapstructure to map parsed data to Go structs
6. **Output**: Saves all vocabulary entries to `output.json`

### Code Example

```go
// Initialize repository and use case
repo, _ := vocabulary.NewVocalbularyRepository()
vc := vocabulary.NewVocalbularyUseCaseImpl(repo)

// Start crawling
words, err := vc.Crawl("https://japanesetest4you.com/jlpt-n1-vocabulary-list/")
if err != nil {
    log.Fatal(err)
}

// Save to repository
err = vc.Save(words)
```

### Dependencies

- [goquery](https://github.com/PuerkitoBio/goquery) - jQuery-like HTML parsing
- [mapstructure](https://github.com/mitchellh/mapstructure) - Decoding map values into Go structs

### Configuration

The project uses environment variables defined in `.env`:

- `DB_HOST`: Database host (default: localhost)
- `DB_PORT`: Database port (default: 5432)

*Note: Database integration is prepared but not currently active in the repository implementation.*

### Development

#### Running in Development Mode

```bash
# Run with verbose output
go run main.go
```

#### Adjusting Concurrency

Modify `GOMAXPROCS` in `main.go` to control parallel processing:

```go
runtime.GOMAXPROCS(4) // Use 4 CPU cores
```

### CI/CD

The project includes a Jenkinsfile for automated CI/CD:

- Environment preparation
- Git operations
- Docker Hub integration (commented out, ready for containerization)

### Troubleshooting

#### Issue: HTTP Request Timeouts

**Solution**: The site might be rate-limiting. Add delays between requests:

```go
time.Sleep(time.Second * 1)
```

#### Issue: Incomplete Data Extraction

**Solution**: Check if the website structure has changed. Inspect the HTML selectors in `vocabulary_usecase.go`.

#### Issue: Race Conditions with Concurrent Writes

**Solution**: Use mutex locks when appending to shared slices in goroutines:

```go
var mu sync.Mutex
mu.Lock()
wordItems = append(wordItems, *w)
mu.Unlock()
```

### Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

### License

MIT

---

## Tiếng Việt

Một web crawler được viết bằng Go, thu thập từ vựng JLPT N1 tiếng Nhật từ [Japanesetest4you.com](https://japanesetest4you.com/jlpt-n1-vocabulary-list/).

### Tính năng

- 🚀 **Xử lý đồng thời**: Sử dụng goroutines của Go để thu thập song song
- 📚 **Trích xuất dữ liệu có cấu trúc**: Trích xuất Kanji, Kana, Romaji, Meaning, Type và JLPT level
- 🏗️ **Clean Architecture**: Code được tổ chức tốt với sự phân tách rõ ràng
- 💾 **Xuất JSON**: Lưu dữ liệu từ vựng dạng JSON có cấu trúc
- 🔄 **HTTP Client Wrapper**: Tiện ích HTTP client tùy chỉnh cho các yêu cầu web đáng tin cậy

### Yêu cầu

- **Go 1.18** hoặc cao hơn
- Kết nối Internet để thu thập dữ liệu

### Cài đặt

1. Clone repository:
```bash
git clone https://github.com/morris102/go-crawl.git
cd go-crawl
```

2. Cài đặt dependencies:
```bash
go mod download
```

3. (Tùy chọn) Cấu hình biến môi trường trong `.env`:
```bash
DB_HOST=localhost
DB_PORT=5432
```

### Cách sử dụng

#### Sử dụng cơ bản

Chạy crawler để thu thập từ vựng JLPT N1:

```bash
go run main.go
```

Từ vựng đã thu thập sẽ được lưu vào `output.json` trong thư mục gốc của dự án.

#### Build Binary

Build file thực thi:

```bash
go build -o go-crawl main.go
```

Chạy binary:

```bash
./go-crawl
```

### Cấu trúc dự án

```
go-crawl/
├── main.go                              # Điểm bắt đầu ứng dụng
├── go.mod / go.sum                      # Dependencies của Go module
├── .env                                 # Biến môi trường
├── .gitignore                           # Quy tắc bỏ qua Git
├── Jenkinsfile                          # Cấu hình pipeline CI/CD
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

### Kiến trúc

Dự án tuân theo nguyên tắc **Clean Architecture** với sự phân tách rõ ràng:

#### Các layer

1. **Entry Point** (`main.go`)
   - Khởi tạo ứng dụng
   - Thiết lập GOMAXPROCS cho xử lý song song
   - Bắt đầu quá trình thu thập

2. **Use Case Layer** (`vocabulary_usecase.go`)
   - Chứa business logic cho việc thu thập và phân tích
   - Điều phối thu thập song song sử dụng goroutines
   - Triển khai interface `VocalbularyUseCase`

3. **Repository Layer** (`vocabulary_repository.go`)
   - Interface lưu trữ dữ liệu
   - Triển khai interface `VocalbularyRepository`
   - Có thể mở rộng để lưu vào database

4. **Model Layer** (`model.go`)
   - Định nghĩa cấu trúc dữ liệu
   - `Word`: Đại diện cho một mục từ vựng hoàn chỉnh
   - `WordHtml`: Đại diện cho nội dung HTML trước khi phân tích

5. **Utility Layer** (`common/util/`)
   - Các tiện ích dùng chung
   - HTTP client wrapper để thực hiện web requests

### Data Models

#### Word
Đại diện cho một mục từ vựng tiếng Nhật hoàn chỉnh:

```go
type Word struct {
    Kanji     string `json:"kanji"`      // Chữ Kanji
    Kana      string `json:"kana"`       // Cách đọc Hiragana/Katakana
    Romaji    string `json:"romaji"`     // Cách đọc La-tinh hóa
    Type      string `json:"type"`       // Loại từ (danh từ, động từ, v.v.)
    Meaning   string `json:"meaning"`    // Nghĩa tiếng Anh
    JLPTlevel string `json:"level"`      // Cấp độ JLPT
    Url       string `json:"url"`        // URL nguồn
}
```

#### WordHtml
Đại diện cho nội dung HTML thô trước khi phân tích chi tiết:

```go
type WordHtml struct {
    Content string `json:"content"`      // Văn bản liên kết
    Url     string `json:"url"`          // URL liên kết
}
```

### Cách hoạt động

1. **Thu thập ban đầu**: Crawler lấy trang danh sách từ vựng chính
2. **Trích xuất liên kết**: Trích xuất tất cả URL trang chi tiết từ vựng
3. **Xử lý song song**: Sử dụng goroutines và WaitGroups để thu thập chi tiết đồng thời
4. **Phân tích dữ liệu**: Phân tích HTML sử dụng goquery để trích xuất dữ liệu từ vựng có cấu trúc
5. **Ánh xạ dữ liệu**: Sử dụng mapstructure để ánh xạ dữ liệu đã phân tích vào Go structs
6. **Xuất dữ liệu**: Lưu tất cả các mục từ vựng vào `output.json`

### Ví dụ code

```go
// Khởi tạo repository và use case
repo, _ := vocabulary.NewVocalbularyRepository()
vc := vocabulary.NewVocalbularyUseCaseImpl(repo)

// Bắt đầu thu thập
words, err := vc.Crawl("https://japanesetest4you.com/jlpt-n1-vocabulary-list/")
if err != nil {
    log.Fatal(err)
}

// Lưu vào repository
err = vc.Save(words)
```

### Dependencies

- [goquery](https://github.com/PuerkitoBio/goquery) - Phân tích HTML giống jQuery
- [mapstructure](https://github.com/mitchellh/mapstructure) - Giải mã map values thành Go structs

### Cấu hình

Dự án sử dụng biến môi trường được định nghĩa trong `.env`:

- `DB_HOST`: Database host (mặc định: localhost)
- `DB_PORT`: Database port (mặc định: 5432)

*Lưu ý: Tích hợp database đã được chuẩn bị nhưng chưa hoạt động trong triển khai hiện tại.*

### Development

#### Chạy ở chế độ Development

```bash
# Chạy với output chi tiết
go run main.go
```

#### Điều chỉnh Concurrency

Chỉnh sửa `GOMAXPROCS` trong `main.go` để kiểm soát xử lý song song:

```go
runtime.GOMAXPROCS(4) // Sử dụng 4 CPU cores
```

### CI/CD

Dự án bao gồm Jenkinsfile cho CI/CD tự động:

- Chuẩn bị môi trường
- Các thao tác Git
- Tích hợp Docker Hub (đã comment, sẵn sàng cho containerization)

### Xử lý sự cố

#### Vấn đề: HTTP Request Timeouts

**Giải pháp**: Trang web có thể đang rate-limiting. Thêm delay giữa các requests:

```go
time.Sleep(time.Second * 1)
```

#### Vấn đề: Trích xuất dữ liệu không đầy đủ

**Giải pháp**: Kiểm tra xem cấu trúc website đã thay đổi chưa. Kiểm tra HTML selectors trong `vocabulary_usecase.go`.

#### Vấn đề: Race Conditions với Concurrent Writes

**Giải pháp**: Sử dụng mutex locks khi append vào shared slices trong goroutines:

```go
var mu sync.Mutex
mu.Lock()
wordItems = append(wordItems, *w)
mu.Unlock()
```

### Đóng góp

Chào đón các đóng góp! Vui lòng tạo Pull Request.

### License

MIT
