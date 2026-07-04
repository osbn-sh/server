package richerror

// Translator converts a raw/underlying error into a user-safe message.
// Provide your own implementation (e.g. wrapping your existing errMsgs
// package, reading Postgres error codes, etc.) via SetTranslator.
type Translator interface {
	Translate(err error) string
}

type defaultTranslator struct{}

func (defaultTranslator) Translate(_ error) string {
	return "something went wrong"
}

var activeTranslator Translator = defaultTranslator{}

// SetTranslator plugs a custom Translator in. Call this once at startup,
// e.g.:
//
//	richerror.SetTranslator(myErrMsgsAdapter{})
func SetTranslator(t Translator) {
	if t != nil {
		activeTranslator = t
	}
}
