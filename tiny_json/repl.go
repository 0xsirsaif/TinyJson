package tiny_json

import (
	"bufio"
	"fmt"
	"io"
)

const PROMPOT = ">> "

func Start(in io.Reader, writer io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPOT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}
		line := scanner.Text()
		l := NewLexer(line)

		for token := l.NextToken(); token.Type != EOF; token = l.NextToken() {
			fmt.Printf("%+v\n", token)
		}
	}
}
