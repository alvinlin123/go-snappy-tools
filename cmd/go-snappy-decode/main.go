package main

import (
	"io"
	"os"

	"fmt"

	"github.com/golang/snappy"
)

func main() {
	_, err := io.Copy(os.Stdout, snappy.NewReader(os.Stdin))

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
