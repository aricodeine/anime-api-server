# Anime API Server

A high-performance backend API built with **Go** that aggregates anime metadata from multiple sources. The server exposes REST endpoints for searching anime and retrieving detailed information across providers.

This project demonstrates strong backend engineering fundamentals including **API design, concurrency, caching, rate limiting, graceful shutdown, and provider-based architecture**.

---

## 🚀 Features

* RESTful API for anime data
* Multi-provider support:

  * AniList
  * Jikan (MyAnimeList)
* Provider-based architecture for easy extensibility
* Concurrent request handling using Go goroutines
* In-memory caching to reduce external API calls
* Per-IP rate limiting for API protection
* Graceful shutdown for reliability
* Clean separation of concerns (handlers, services, providers)

---

## 🛠 Tech Stack

* **Language:** Go
* **Architecture:** Layered backend architecture
* **Concurrency:** Goroutines & channels
* **Caching:** In-memory cache
* **API Style:** REST

---

## 📁 Project Structure

```
anime-api-server
│
├── handlers/       # HTTP request handlers
├── services/       # Business logic layer
├── providers/      # External anime data providers (AniList, Jikan)
├── models/         # Data models
├── internal/       # Internal utilities (cache, middleware)
│
├── main.go         # Server entry point
├── go.mod
├── go.sum
└── README.md
```

---

## ⚙️ Getting Started

### 1. Clone the repository

```
git clone https://github.com/yourusername/anime-api-server.git
cd anime-api-server
```

### 2. Install dependencies

```
go mod tidy
```

### 3. Run the server

```
go run main.go
```

Server will start on:

```
http://localhost:8080
```

---

## 📡 API Endpoints

### 🔍 Search Anime

```
GET /{provider}/anime/search?q={query}
```

**Examples:**

```
GET /anilist/anime/search?q=naruto
GET /jikan/anime/search?q=naruto
```

---

### 📖 Get Anime Details

```
GET /{provider}/anime/{id}
```

**Examples:**

```
GET /anilist/anime/20
GET /jikan/anime/1
```

---

## 🔌 Supported Providers

* AniList
* Jikan (MyAnimeList via Jikan API)

The provider-based architecture allows seamless integration of additional data sources without modifying core business logic.

---

## 🎯 Goals of This Project

This project was built to practice and demonstrate:

* Designing scalable backend APIs
* Implementing provider-based architecture
* Integrating multiple external APIs
* Handling concurrency in Go
* Implementing caching and rate limiting
* Structuring maintainable backend systems

---

## ⚠️ Disclaimer

This project is for educational purposes only.

It aggregates publicly available metadata from third-party APIs and does not host or distribute any copyrighted media.
