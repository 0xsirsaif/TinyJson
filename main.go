package main

import (
	"flag"
	"fmt"
	"github.com/0xsirsaif/TinyJson/tiny_json"
	"io"
	"os"
)

func checkErr(e error, msg string) {
	if e != nil {
		fmt.Println(msg, e)
		os.Exit(1)
	}
}

func readJsonFile(filePath string) string {
	var fileBytes []byte
	var err error

	if filePath == "-" {
		fileBytes, err = io.ReadAll(os.Stdin)

	} else {
		fileBytes, err = os.ReadFile("./" + filePath)
	}

	checkErr(err, "Error loading File:")

	return string(fileBytes)
}

func main() {
	fileFlag := flag.String("f", "-", "file path")

	flag.Parse()

	jsonAsStr := readJsonFile(*fileFlag)

	lexer := tiny_json.NewLexer(jsonAsStr)

	for {
		token := lexer.NextToken()
		fmt.Printf("%+v\n", token)
		if token.Type == tiny_json.EOF {
			break
		}
	}
}
