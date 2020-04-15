package repl

import (
	"bufio"
	"fmt"
	"github.com/uchiiii/gibbon/lexer"
	"github.com/uchiiii/gibbon/token"
	"io"
)

const (
	PROMPT = "gibbon> "
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf("\x1b[36m%s\x1b[0m", PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}
