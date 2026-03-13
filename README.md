# Anime API Server

A high-performance backend API built with **Go** that aggregates anime metadata from multiple sources. The server exposes REST endpoints for searching anime, retrieving detailed information.

This project focuses on backend engineering concepts such as **API design, concurrency, caching, and provider aggregation**.

---

## Features

* RESTful API for anime data
* Aggregates results from multiple providers (Currently supported providers: **Anilist**)
* Concurrent provider requests using Go goroutines
* In-memory caching to reduce external API calls
* Modular service architecture
* Clean separation of handlers, services, and providers

---

## Tech Stack

* **Language:** Go
* **Architecture:** Layered backend architecture
* **Concurrency:** Goroutines & channels
* **Caching:** In-memory cache
* **API Style:** REST

---

## Project Structure

```
anime-api-server
│
├── handlers/       # HTTP request handlers
├── services/       # Business logic layer
├── providers/      # External anime data providers
├── models/         # Data models
├── internal/       # Internal utilities (cache, helpers)
│
├── main.go         # Server entry point
├── go.mod
├── go.sum
└── README.md
```

---

## Getting Started

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

## Example API Endpoints

### Search Anime

```
GET /anime/search?q=naruto
```

### Get Anime Details

```
GET /anime/{id}
```

## Goals of This Project

This project was built to practice:

* Designing scalable backend APIs
* Working with external data providers
* Implementing caching strategies
* Using Go concurrency effectively
* Structuring maintainable backend services
