package responses

const (
	SuccessHtmlPrefix    = "HTTP/1.1 200 OK\nContent-Type: text/html\nContent-Length: "
	NotFound             = "HTTP/1.1 404 Not Found\nContent-Type: text/plain\nContent-Length: 13\n\n404 Not Found"
	InvalidHttpMethod    = "HTTP/1.1 405 Method Not Allowed\nAllow: GET\nContent-Type: text/plain\nContent-Length: 22\n\n405 Method Not Allowed"
	UnsupportedMediaType = "HTTP/1.1 415 Unsupported Media Type\nContent-Type: text/plain\nContent-Length: 26\n\n415 Unsupported Media Type"
)
