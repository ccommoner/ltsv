package ltsv

import (
	"bytes"
	"fmt"
	"log"
)

func ExampleWriter_format() {
	data := []map[string]string{
		{"time": "Mon, 02 Jan 2006 15:04:05 MST", "host": "127.0.0.1", "method": "POST"},
		{"time": "Mon, 02 Jan 2006 15:05:05 MST", "host": "127.0.0.1"},
		{"time": "Mon, 02 Jan 2006 15:06:05 MST", "host": "127.0.0.1", "method": "GET", "status": "404"},
	}

	fmtBuf := new(bytes.Buffer)
	fmtWriter := NewWriter(fmtBuf)
	fmtWriter.Format = []string{
		"time",
		"host",
		"method",
	}
	if err := fmtWriter.WriteAll(data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(fmtBuf)
	// Output:
	// time:Mon, 02 Jan 2006 15:04:05 MST	host:127.0.0.1	method:POST
	// time:Mon, 02 Jan 2006 15:05:05 MST	host:127.0.0.1	method:
	// time:Mon, 02 Jan 2006 15:06:05 MST	host:127.0.0.1	method:GET
}

func ExampleWriter_noFormat() {
	data := []map[string]string{
		{"time": "Mon, 02 Jan 2006 15:04:05 MST", "host": "127.0.0.1", "method": "POST"},
		{"time": "Mon, 02 Jan 2006 15:05:05 MST", "host": "127.0.0.1"},
		{"time": "Mon, 02 Jan 2006 15:06:05 MST", "host": "127.0.0.1", "method": "GET", "status": "404"},
	}

	buf := new(bytes.Buffer)
	writer := NewWriter(buf)
	if err := writer.WriteAll(data); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf)
	// Output:
	// method:POST	host:127.0.0.1	time:Mon, 02 Jan 2006 15:04:05 MST
	// host:127.0.0.1	time:Mon, 02 Jan 2006 15:05:05 MST
	// time:Mon, 02 Jan 2006 15:06:05 MST	host:127.0.0.1	status:404	method:GET
}
