package ltsv

import (
	"encoding/csv"
	"io"
)

type Writer struct {
	svWriter *csv.Writer
	Format   []string
}

func NewWriter(w io.Writer) *Writer {
	base := csv.NewWriter(w)
	base.Comma = '\t'
	return &Writer{base, nil}
}

func (w *Writer) Flush() {
	w.svWriter.Flush()
}

func (w *Writer) Write(record map[string]string) error {
	svRecord := make([]string, 0)
	if w.Format == nil {
		for k, v := range record {
			svRecord = append(svRecord, k+":"+v)
		}
		return w.svWriter.Write(svRecord)
	}

	for _, key := range w.Format {
		svRecord = append(svRecord, key+":"+record[key])
	}
	return w.svWriter.Write(svRecord)
}

func (w *Writer) WriteAll(records []map[string]string) error {
	defer w.Flush()
	for _, record := range records {
		if err := w.Write(record); err != nil {
			return err
		}
	}
	return nil
}
