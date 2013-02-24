// decoder.go
package ltsv

import (
	"io"
	"reflect"
)

// During the experiment!!!!!!
type Decoder struct {
	reader *Reader
}

func NewDecoder(r io.Reader) *Decoder {
	return &Decoder{NewReader(r)}
}

func (dec *Decoder) Decode(v interface{}) error {
	record, err := dec.reader.Read()
	if err != nil {
		return err
	}
	retVal := reflect.Indirect(reflect.ValueOf(v))
	vType := retVal.Type()
	for i := 0; i < vType.NumField(); i++ {
		if val, exist := record[vType.Field(i).Name]; exist {
			retVal.Field(i).SetString(val)
		}
	}
	v = retVal
	return nil
}
