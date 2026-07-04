package richerror

import "net/http"

// Kind represents the category of an error (useful for mapping to HTTP
// status codes or branching logic in calling code). Zero value means
// "not set" — do not assign it manually.
type Kind int

const (
	_ Kind = iota

	KindUnexpected

	KindInvalid
	KindUnauthorized
	KindForbidden
	KindNotFound
	KindConflict

	KindTooManyRequests
	KindTimeout
	KindUnavailable

	KindDatabase
	KindExternal

	KindNotImplemented
)

func (k Kind) String() string {
	switch k {
	case KindUnexpected:
		return "UNEXPECTED"

	case KindInvalid:
		return "INVALID"

	case KindUnauthorized:
		return "UNAUTHORIZED"

	case KindForbidden:
		return "FORBIDDEN"

	case KindNotFound:
		return "NOT_FOUND"

	case KindConflict:
		return "CONFLICT"

	case KindTooManyRequests:
		return "TOO_MANY_REQUESTS"

	case KindTimeout:
		return "TIMEOUT"

	case KindUnavailable:
		return "UNAVAILABLE"

	case KindDatabase:
		return "DATABASE"

	case KindExternal:
		return "EXTERNAL"

	case KindNotImplemented:
		return "NOT_IMPLEMENTED"

	default:
		return "UNKNOWN"
	}
}

func (k Kind) HTTPStatus() int {
	switch k {
	case KindInvalid:
		return http.StatusBadRequest // 400

	case KindUnauthorized:
		return http.StatusUnauthorized // 401

	case KindForbidden:
		return http.StatusForbidden // 403

	case KindNotFound:
		return http.StatusNotFound // 404

	case KindConflict:
		return http.StatusConflict // 409

	case KindTooManyRequests:
		return http.StatusTooManyRequests // 429

	case KindTimeout:
		return http.StatusRequestTimeout // 408

	case KindUnavailable:
		return http.StatusServiceUnavailable // 503

	case KindDatabase:
		return http.StatusInternalServerError // 500

	case KindExternal:
		return http.StatusBadGateway // 502

	case KindNotImplemented:
		return http.StatusNotImplemented // 501

	case KindUnexpected:
		fallthrough
	default:
		return http.StatusInternalServerError // 500
	}
}
