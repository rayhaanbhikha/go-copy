package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	fmt.Println("hello world")

	// check if aws creds are provided.

	// TODO: scan input for aws access creds.
	// TODO: or path to file for aws creds.

	// install s3 util.

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(data))
}
