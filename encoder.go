// encoder.go
package ltsv

import (
	"errors"
	"io"
	"reflect"
)

// During the experiment!!!!!!
type Encoder struct {
	writer *Writer
}

func NewEncoder(w io.Writer) *Encoder {
	return &Encoder{NewWriter(w)}
}

func (enc *Encoder) Encode(v interface{}) error {
	argType := reflect.TypeOf(v)
	if argType.Kind() != reflect.Struct {
		return errors.New("argument is other than struct")
	}

	format := make([]string, 0)
	for i := 0; i < argType.NumField(); i++ {
		// TODO: Support "ltsv" StructTag
		format = append(format, argType.Field(i).Name)
	}

	argVal := reflect.ValueOf(v)
	record := map[string]string{}
	for _, label := range format {
		// TODO?: If the field is not a string, should occur an error?
		// Or Marshal JSON?
		record[label] = argVal.FieldByName(label).String()
	}
	enc.writer.Format = format
	if err := enc.writer.Write(record); err != nil {
		return nil
	}
	enc.writer.Flush()
	return nil
}
