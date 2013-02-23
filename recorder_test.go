// encoder_test.go
package ltsv

import (
	"bytes"
	"fmt"
	"log"
)

type TestLogData struct {
	time   string
	host   string
	method string
}

func ExampleRecorder_Encode() {
	logData := TestLogData{"Mon, 02 Jan 2006 15:04:05 MST", "127.0.0.1", "POST"}

	buf := new(bytes.Buffer)
	encoder := NewRecorder(buf)
	if err := encoder.Record(logData); err != nil {
		log.Fatal(err)
	}
	fmt.Println(buf)
	// Output:
	// time:Mon, 02 Jan 2006 15:04:05 MST	host:127.0.0.1	method:POST
}
