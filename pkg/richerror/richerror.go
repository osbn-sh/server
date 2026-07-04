package richerror

import (
	"encoding/json"
	"errors"
	"fmt"
	"runtime"
	"strings"
)

// Op identifies the operation/layer producing the error, e.g.
// "userRepository.RegisterUserByEmailAndPassword".
type Op string

// frame is a lightweight, JSON-friendly stack frame.
type frame struct {
	Func string `json:"func"`
	File string `json:"file"`
	Line int    `json:"line"`
}

type RichError struct {
	op      Op
	err     error
	kind    Kind
	message string
	meta    map[string]interface{}
	stack   []frame
}

// New starts a RichError for the given operation. Call it at every layer
// you want represented in the operation chain (handler -> service ->
// repository, ...). Only the layer that first wraps a plain (non-RichError)
// error actually captures a stack trace, so the trace always points at the
// true origin.
func New(op Op) RichError {
	return RichError{op: op}
}

func (r RichError) WithOp(op Op) RichError {
	r.op = op
	return r
}

// WithErr attaches the underlying error. If it's the first "raw" error in
// the chain (not itself a RichError), the current stack is captured here.
func (r RichError) WithErr(err error) RichError {
	r.err = err
	if err != nil {
		if _, ok := err.(RichError); !ok {
			r.stack = captureStack()
		}
	}
	return r
}

func (r RichError) WithKind(kind Kind) RichError {
	r.kind = kind
	return r
}

// WithMessage sets the internal/debug message (logs only).
func (r RichError) WithMessage(message string) RichError {
	r.message = message
	return r
}

func (r RichError) WithMeta(meta map[string]interface{}) RichError {
	newMeta := make(map[string]interface{}, len(r.meta)+len(meta))
	for k, v := range r.meta {
		newMeta[k] = v
	}
	for k, v := range meta {
		newMeta[k] = v
	}
	r.meta = newMeta
	return r
}

// WithRequestID / WithTraceID are small convenience wrappers around
// WithMeta for carrying request-scoped context (e.g. from an HTTP
// middleware) through the whole error chain.
func (r RichError) WithRequestID(id string) RichError {
	return r.WithMeta(map[string]interface{}{"request_id": id})
}

func (r RichError) WithTraceID(id string) RichError {
	return r.WithMeta(map[string]interface{}{"trace_id": id})
}

// ---------------------------------------------------------------------
// error interface + errors.Is / errors.As support
// ---------------------------------------------------------------------

func (r RichError) Error() string {
	if r.message != "" {
		return r.message
	}
	if r.err != nil {
		var richError RichError
		if errors.As(r.err, &richError) {
			return r.err.Error()
		}
		return activeTranslator.Translate(r.err)
	}
	return "something went wrong"
}

// Unwrap is what makes errors.Is / errors.As work through the whole chain.
func (r RichError) Unwrap() error {
	return r.err
}

// Kind returns the first explicitly-set Kind found walking down the chain.
func (r RichError) Kind() Kind {
	if r.kind != 0 {
		return r.kind
	}
	var re RichError
	if errors.As(r.err, &re) {
		return re.Kind()
	}
	return KindInvalid
}

// Message returns the deepest non-empty internal message in the chain.
func (r RichError) Message() string {
	if r.message != "" {
		return r.message
	}
	var re RichError
	if errors.As(r.err, &re) {
		return re.Message()
	}
	if r.err != nil {
		return r.err.Error()
	}
	return ""
}

func (r RichError) Operation() Op {
	return r.op
}

// OperationChain returns every operation from outermost to innermost, e.g.
// ["handler.Register", "service.Register", "userRepository.Insert"].
func (r RichError) OperationChain() []Op {
	chain := []Op{r.op}
	var re RichError
	if errors.As(r.err, &re) {
		chain = append(chain, re.OperationChain()...)
	}
	return chain
}

// RootCause walks the whole chain and returns the deepest underlying error.
func (r RichError) RootCause() error {
	var re RichError
	if errors.As(r.err, &re) {
		return re.RootCause()
	}
	if r.err != nil {
		return r.err
	}
	return r
}

func (r RichError) Meta() map[string]interface{} {
	return r.meta
}

// Stack returns the captured stack trace as readable lines. It's only
// present at the frame where a raw (non-RichError) error was first wrapped.
func (r RichError) Stack() []string {
	if len(r.stack) > 0 {
		out := make([]string, 0, len(r.stack))
		for _, f := range r.stack {
			out = append(out, fmt.Sprintf("%s\n\t%s:%d", f.Func, f.File, f.Line))
		}
		return out
	}
	var re RichError
	if errors.As(r.err, &re) {
		return re.Stack()
	}
	return nil
}

func captureStack() []frame {
	pcs := make([]uintptr, 32)
	// skip=3: runtime.Callers, captureStack, WithErr
	n := runtime.Callers(3, pcs)
	frames := runtime.CallersFrames(pcs[:n])
	var out []frame
	for {
		f, more := frames.Next()
		if strings.Contains(f.File, "runtime/") {
			if !more {
				break
			}
			continue
		}
		out = append(out, frame{Func: f.Function, File: f.File, Line: f.Line})
		if !more {
			break
		}
	}
	return out
}

// ---------------------------------------------------------------------
// Formatters
// ---------------------------------------------------------------------

// DebugReport is the full internal view — safe for logs, never for clients.
type DebugReport struct {
	Operation string                 `json:"operation"`
	Chain     []Op                   `json:"operation_chain"`
	Kind      string                 `json:"kind"`
	Message   string                 `json:"message"`
	RootCause string                 `json:"root_cause"`
	Meta      map[string]interface{} `json:"meta,omitempty"`
	Stack     []string               `json:"stack,omitempty"`
}

// UserReport is the minimal, safe-for-clients view.
type UserReport struct {
	Message string `json:"message"`
}

func (r RichError) ToDebug() DebugReport {
	return DebugReport{
		Operation: string(r.op),
		Chain:     r.OperationChain(),
		Kind:      r.Kind().String(),
		Message:   r.Message(),
		RootCause: r.RootCause().Error(),
		Meta:      r.Meta(),
		Stack:     r.Stack(),
	}
}

func (r RichError) ToUser() UserReport {
	return UserReport{
		Message: r.Message(),
	}
}

func (r RichError) ToJSON() ([]byte, error) {
	return json.Marshal(r.ToDebug())
}

// LogFields returns a flat map ready for any structured logger:
//
//	slog.Error("request failed", "error", richErr.LogFields())
//	logrus.WithFields(richErr.LogFields()).Error("request failed")
func (r RichError) LogFields() map[string]interface{} {
	fields := map[string]interface{}{
		"operation":       string(r.op),
		"operation_chain": r.OperationChain(),
		"kind":            r.Kind().String(),
		"message":         r.Message(),
		"root_cause":      r.RootCause().Error(),
	}
	for k, v := range r.meta {
		fields[k] = v
	}
	return fields
}

// ---------------------------------------------------------------------
// Panic recovery
// ---------------------------------------------------------------------

// Recover converts a panic into a RichError with a stack trace, assigning
// it to *errOut. Use it via defer at the top of a handler or goroutine:
//
//	func Handler() (err error) {
//	    defer richerror.Recover("handler.Register", &err)
//	    ...
//	}
func Recover(op Op, errOut *error) {
	rec := recover()
	if rec == nil {
		return
	}
	var err error
	switch v := rec.(type) {
	case error:
		err = v
	default:
		err = fmt.Errorf("panic: %v", v)
	}
	re := New(op).
		WithErr(err).
		WithKind(KindUnexpected).
		WithMessage(fmt.Sprintf("recovered from panic: %v", rec))
	re.stack = captureStack()
	*errOut = re
}
