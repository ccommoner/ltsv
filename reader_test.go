package ltsv

import (
	"bytes"
	"fmt"
	"log"
)

func ExampleReader_format() {
	data := `
time:Mon, 02 Jan 2006 15:04:05 MST	host:127.0.0.1	method:POST
time:Mon, 02 Jan 2006 15:05:05 MST	host:127.0.0.1
time:Mon, 02 Jan 2006 15:06:05 MST	host:127.0.0.1	method:GET	status:404
`

	fmtReader := NewReader(bytes.NewBufferString(data))
	fmtReader.Format = []string{
		"time",
		"host",
		"method",
	}

	records, err := fmtReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		fmt.Println(record)
	}

	// Output Example:
	// map[method:POST time:Mon, 02 Jan 2006 15:04:05 MST host:127.0.0.1]
	// map[time:Mon, 02 Jan 2006 15:05:05 MST method: host:127.0.0.1]
	// map[method:GET host:127.0.0.1 time:Mon, 02 Jan 2006 15:06:05 MST]
}

func ExampleReader_noFormat() {
	data := `
time:Mon, 02 Jan 2006 15:04:05 MST	host:127.0.0.1	method:POST
time:Mon, 02 Jan 2006 15:05:05 MST	host:127.0.0.1
time:Mon, 02 Jan 2006 15:06:05 MST	host:127.0.0.1	method:GET	status:404
`

	reader := NewReader(bytes.NewBufferString(data))
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for _, record := range records {
		fmt.Println(record)
	}

	// Output Example:
	// map[host:127.0.0.1 method:POST time:Mon, 02 Jan 2006 15:04:05 MST]
	// map[time:Mon, 02 Jan 2006 15:05:05 MST host:127.0.0.1]
	// map[time:Mon, 02 Jan 2006 15:06:05 MST host:127.0.0.1 status:404 method:GET]
}
