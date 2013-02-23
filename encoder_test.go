// encoder_test.go
package ltsv

import (
	"bytes"
	"fmt"
	"log"
)

type TestSimpleLogData struct {
	Time   string
	Host   string
	Method string
}

func ExampleEncoder_Encode() {
	logData := TestSimpleLogData{"Mon, 02 Jan 2006 15:04:05 MST", "127.0.0.1", "POST"}

	buf := new(bytes.Buffer)
	encoder := NewEncoder(buf)
	if err := encoder.Encode(logData); err != nil {
		log.Fatal(err)
	}
	encoder.Flush()
	fmt.Println(buf)
	// Output:
	// Time:Mon, 02 Jan 2006 15:04:05 MST	Host:127.0.0.1	Method:POST
}
