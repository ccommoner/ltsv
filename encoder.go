// encoder.go
package ltsv

import (
	"errors"
	"io"
	"reflect"
)

// During the experiment!!!!!!
type Encoder struct {
	recorder *Recorder
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{NewRecorder(w)}
}

func (enc *Encoder) Flush() {
	enc.recorder.Flush()
}

func (enc *Encoder) Encode(v interface{}) error {
	argType := reflect.TypeOf(v)
	if argType.Kind() != reflect.Struct {
		return errors.New("argument is not struct")
	}

	for i := 0; i < argType.NumField(); i++ {
		if argType.Field(i).Type.Kind() != reflect.String {
			return errors.New("struct field is not string")
		}
	}

	return enc.recorder.Record(v)
}
