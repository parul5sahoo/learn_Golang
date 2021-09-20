package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type logWriter struct{}

func main() {
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	//fmt.Println(resp)

	// bs := make([]byte, 99999) //gives an empty byte slice with 99999 empty spaces
	// //bcoz the Read fucntion doesnt automatically shrinks or enlarges the slice
	// resp.Body.Read(bs)
	// fmt.Println(string(bs))

	//io.Copy(os.Stdout, resp.Body) // using the writer interface
	lw := logWriter{}
	io.Copy(lw, resp.Body)

}

func (logWriter) Write(bs []byte) (int, error) {
	fmt.Println(string(bs))
	fmt.Println("The no. of bytes written is:", len(bs))
	return len(bs), nil
}
