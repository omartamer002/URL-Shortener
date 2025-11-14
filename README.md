# URL Shortener

A simple URL shortening service built with Go that allows you to create short aliases for long URLs and redirect users to the original URL using the short alias.

---

## 🚀 Features
- Generate short URL aliases from long URLs  
- Redirect from short URLs to original URLs  
- Thread-safe storage using mutexes  
- RESTful HTTP API  

---

## 🛠 Getting Started

### **Prerequisites**
- Go **1.23.5** or higher

### **Installation**
1. **Clone the repository**
2. **Navigate to the project directory**
3. **Run the application**
   ```bash
   go run main.go

##The server will start at:
http://localhost:8080

## 🔗 API Endpoints
### ***Shorten URL***
Request
POST /shorten
Response
Returns a short URL string.

Redirect
Request
GET /{shortURL}
Response
Redirects the user to the original long URL.

## 📁 Project Structure

| File / Directory | Description |
|------------------|-------------|
| `main.go`        | Application entry point and HTTP server setup |
| `handler.go`     | HTTP handlers for URL shortening and redirection |
| `url.go`         | URL generation logic using random Base64 encoding |
| `storage.go`     | Thread-safe in-memory storage for URL mappings |
| `go.mod`         | Go module definition |


## 🧩 How It Works

1. User submits a long URL to the `/shorten` endpoint.  
2. A random **6-byte** string is generated and Base64-encoded to create a short URL.  
3. The application stores the short → long URL mapping in memory.  
4. The user visits `/{shortURL}`.  
5. The server redirects the user to the original URL.
