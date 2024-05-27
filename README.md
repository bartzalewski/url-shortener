# URL Shortener API

Welcome to the URL Shortener API! This is a simple service to shorten URLs, track clicks, and provide basic analytics.

## Features

- URL shortening
- Redirection to original URLs
- Click tracking and analytics

## Getting Started

Follow these instructions to get a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

Ensure you have the following installed on your system:

- [Go](https://golang.org/doc/install) (version 1.16+)
- [Git](https://git-scm.com/)

### Installation

1. **Clone the repository**:

   ```sh
   git clone https://github.com/bartzalewski/url-shortener.git
   cd url-shortener
   ```

2. **Initialize Go modules**:

   ```sh
   go mod tidy
   ```

3. **Run the application**:

   ```sh
   go run main.go
   ```

The server will start on `http://localhost:8080`.

## API Endpoints

### URL Shortening

- **Shorten URL**

  `POST /shorten`

  Request:

  ```json
  {
    "url": "https://www.example.com"
  }
  ```

  Response:

  ```json
  {
    "short_url": "http://localhost:8080/abc123"
  }
  ```

### URL Redirection

- **Redirect to Original URL**

  `GET /{shortCode}`

  Example:

  `GET /abc123`

  Redirects to `https://www.example.com`

### Analytics

- **Get Click Analytics**

  `GET /analytics/{shortCode}`

  Example:

  `GET /analytics/abc123`

  Response:

  ```json
  {
    "clicks": 10
  }
  ```

## Built With

- [Go](https://golang.org/) - The Go programming language
- [Gorilla Mux](https://github.com/gorilla/mux) - A powerful URL router and dispatcher for Golang

## Contributing

Feel free to submit issues or pull requests. For major changes, please open an issue first to discuss what you would like to change.

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for their invaluable resources and support.
- [Gorilla Mux](https://github.com/gorilla/mux) for making routing simple and efficient.

---

Happy coding! 🚀
