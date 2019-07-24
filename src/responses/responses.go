package responses

const (
	Success           = "HTTP/1.1 200 OK\nContent-Type: text/plain\nContent-Length: 12\n\nHello world!"
	InvalidHttpMethod = "HTTP/1.1 405 Method Not Allowed\nAllow: GET\nContent-Type: text/plain\nContent-Length: 19\n\nMethod Not Allowed!"
)
