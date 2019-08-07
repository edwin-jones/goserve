package status

// Status HTTP status code typedef
type Code int

// HTTP status codes
const (
	Success              Code = 200
	BadRequest           Code = 400
	NotFound             Code = 404
	InvalidHTTPMethod    Code = 405
	URITooLong           Code = 414
	UnsupportedMediaType Code = 415
)
