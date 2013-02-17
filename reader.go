package ltsv

import (
	"encoding/csv"
	"io"
	"strings"
)

type Reader struct {
	svReader *csv.Reader
	Format   []string
}

func NewReader(r io.Reader) *Reader {
	base := csv.NewReader(r)
	base.Comma = '\t'
	base.FieldsPerRecord = -1
	return &Reader{base, nil}
}

func (r *Reader) Read() (map[string]string, error) {
	kvRecord, err := r.svReader.Read()
	if err != nil {
		return nil, err
	}
	record := map[string]string{}
	for _, key := range r.Format {
		record[key] = ""
	}
	for _, kv := range kvRecord {
		pair := strings.SplitN(kv, ":", 2)
		if _, exist := record[pair[0]]; exist || r.Format == nil {
			record[pair[0]] = pair[1]
		}
	}
	return record, nil
}

func (r *Reader) ReadAll() ([]map[string]string, error) {
	records := make([]map[string]string, 0)
	for {
		record, err := r.Read()
		if err == io.EOF {
			return records, nil
		}
		if err != nil {
			return nil, err
		}
		records = append(records, record)
	}
	panic("unreachable")
}
