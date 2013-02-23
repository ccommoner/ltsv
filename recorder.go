// recorder.go
package ltsv

import (
	"errors"
	"fmt"
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

func (rec *Recorder) Flush() {
	rec.writer.Flush()
}

func (rec *Recorder) Record(v interface{}) error {
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
	for i := 0; i < argType.NumField(); i++ {
		if argVal.Type().Field(i).PkgPath == "" {
			record[format[i]] = fmt.Sprint(argVal.Field(i).Interface())
		}
	}
	rec.writer.Format = format
	if err := rec.writer.Write(record); err != nil {
		return nil
	}
	return nil
}
