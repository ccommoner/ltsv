package ltsv

import (
	"bytes"
	"fmt"
	"log"
)

func ExampleDecoder() {
	data := `
Time:Mon, 02 Jan 2006 15:04:05 MST	Host:127.0.0.1	Method:POST
Time:Mon, 02 Jan 2006 15:05:05 MST	Host:127.0.0.1	Method:PUT
Time:Mon, 02 Jan 2006 15:06:05 MST	Host:127.0.0.1	Method:GET
`

	decoder := NewDecoder(bytes.NewBufferString(data))

	var logData TestSimpleLogData

	if err := decoder.Decode(&logData); err != nil {
		log.Fatal(err)
	}
	fmt.Println(logData)

	// Output:
	// {Mon, 02 Jan 2006 15:04:05 MST 127.0.0.1 POST}
}
