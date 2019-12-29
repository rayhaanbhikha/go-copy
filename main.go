package main

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/rayhaanbhikha/go-copy/s3"
)

func main() {
	// check if aws creds are provided.

	// TODO: scan input for aws access creds.
	// TODO: or path to file for aws creds.

	// install s3 util.

	data, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		log.Fatal(err)
	}

	s3.Upload(string(data))

}
