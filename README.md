# 🚀 Parallel Web Scraper & Keyword Analyzer

## 📌 Project Overview
Parallel Web Scraper & Keyword Analyzer is a high-performance Go application that scrapes multiple websites concurrently, analyzes the most frequent words, and stores the results in a MongoDB database with Redis caching.

## ⚡ Features
- **Concurrent Web Scraping**: Uses Go's goroutines and channels for efficient parallel scraping.
- **Redis Caching**: Avoids unnecessary HTTP requests by caching previously scraped pages.
- **Keyword Analysis**: Extracts and counts the most frequently used words in the scraped content.
- **MongoDB Storage**: Persists results for future analysis and retrieval.
- **API Support**: Easily extendable with a REST API to fetch stored data.

## 📂 Project Structure
```
/ParallelScraper
│── main.go                    # Entry point of the application
│── /scraper
│   ├── scraper.go              # Handles web scraping logic
│── /analyzer
│   ├── analyzer.go             # Processes text and analyzes keywords
│── /database
│   ├── mongo.go                # MongoDB connection & operations
│   ├── redis.go                # Redis caching logic
│── go.mod                      # Go module dependencies
│── go.sum                      # Checksums for dependencies
│── README.md                   # Project documentation
```

## 🚀 Getting Started
### 1️⃣ Prerequisites
Ensure you have the following installed:
- Go 1.18+
- MongoDB
- Redis

### 2️⃣ Clone the Repository
```sh
git clone https://github.com/yourusername/ParallelScraper.git
cd ParallelScraper
```

### 3️⃣ Install Dependencies
```sh
go mod tidy
```

### 4️⃣ Configure Environment Variables
Create a `.env` file with:
```env
MONGO_URI=mongodb://localhost:27017
REDIS_ADDR=localhost:6379
```

### 5️⃣ Run the Application
```sh
go run main.go
```

## 📊 How It Works
1. User provides multiple URLs.
2. The scraper fetches web pages concurrently using goroutines and channels.
3. Redis checks for cached content before sending requests.
4. The analyzer processes and identifies frequently used words.
5. Data is stored in MongoDB for future retrieval.

## 🛠️ Technologies Used
- **Golang** (Concurrency, Goroutines, Channels)
- **Colly** (Web Scraping)
- **MongoDB** (Data Storage)
- **Redis** (Caching)
- **Gin** (Optional: API Framework)

## 🤝 Contributing
Feel free to fork this repository, submit issues, or create pull requests!

## 📜 License
This project is licensed under the MIT License.

---

Happy coding! 🚀
