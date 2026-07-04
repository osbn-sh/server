package richerror

type Translator interface {
	Translate(err error) string
}

type defaultTranslator struct{}

func (defaultTranslator) Translate(_ error) string {
	return "something went wrong"
}

var activeTranslator Translator = defaultTranslator{}

func SetTranslator(t Translator) {
	if t != nil {
		activeTranslator = t
	}
}
