package richerror

import (
	"ostadbun/pkg/errMsgs"
)

type Kind int

const (
	KindInvalid Kind = iota + 1
	KindForbidden
	KindNotFound
	KindUnexpected
)

type Op string

type RichError struct {
	operation    Op
	wrappedError error
	message      string
	kind         Kind
	meta         map[string]interface{}
}

func New(op Op) RichError {
	return RichError{operation: op}
}

func (r RichError) WithOp(op Op) RichError {
	r.operation = op
	return r
}

func (r RichError) WithErr(err error) RichError {
	r.wrappedError = err
	return r
}

func (r RichError) WithMessage(message string) RichError {
	r.message = message
	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind
	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	newMeta := make(map[string]interface{})
	for k, v := range meta {
		newMeta[k] = v
	}
	r.meta = newMeta
	return r
}

//func (r RichError) Error() string {
//	fmt.Println(r.message)
//
//	if r.message != "" && r.wrappedError != nil {
//		return fmt.Sprintf("%s: %s", r.message, r.wrappedError.Error())
//	}
//	if r.message != "" {
//		return r.message
//	}
//	if r.wrappedError != nil {
//		return r.wrappedError.Error()
//	}
//	return "something went wrong"
//}

func (r RichError) Error() string {
	if r.message == "" && r.wrappedError != nil {
		return errMsgs.TranslateErrorMessage(r.wrappedError.Error())
	}

	return r.message
}

func (r RichError) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}

	re, ok := r.wrappedError.(RichError)
	if !ok {
		return 0
	}

	return re.Kind()
}

func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}

	re, ok := r.wrappedError.(RichError)
	if ok {
		return re.Message()
	}

	if r.wrappedError != nil {
		return r.wrappedError.Error()
	}

	return ""
}

func (r RichError) Operation() Op {
	return r.operation
}

func (r RichError) Unwrap() error {
	return r.wrappedError
}

func (r RichError) Meta() map[string]interface{} {
	return r.meta
}
