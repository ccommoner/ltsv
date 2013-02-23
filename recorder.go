// encoder.go
package ltsv

import (
	"errors"
	"io"
	"reflect"
)

// During the experiment!!!!!!
type Recorder struct {
	writer *Writer
}

func NewRecorder(w io.Writer) *Recorder {
	return &Recorder{NewWriter(w)}
}

func (enc *Recorder) Record(v interface{}) error {
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
		record[label] = argVal.FieldByName(label).String()
	}
	enc.writer.Format = format
	if err := enc.writer.Write(record); err != nil {
		return nil
	}
	enc.writer.Flush()
	return nil
}
