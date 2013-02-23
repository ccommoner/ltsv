// encoder_test.go
package ltsv

import (
	"bytes"
	"fmt"
	"log"
	"time"
)

type ReqLogData struct {
	Size int
	Time time.Time
}

type TestLogData struct {
	Req    ReqLogData
	Host   string
	Method string
}

func ExampleRecorder_Record() {
	logTime, _ := time.Parse(time.UnixDate, time.UnixDate)
	logData := TestLogData{ReqLogData{100, logTime}, "127.0.0.1", "POST"}

	buf := new(bytes.Buffer)
	encoder := NewRecorder(buf)
	if err := encoder.Record(logData); err != nil {
		log.Fatal(err)
	}
	encoder.Flush()

	fmt.Println(buf)
	// Output:
	// Time:0001-01-01 00:00:00 +0000 UTC	Host:127.0.0.1	Method:POST
}
