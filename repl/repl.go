package repl

import (
	"fmt"
	"io"
	"os"

	"github.com/peterh/liner"
	"github.com/syedazeez337/monkey-interpreter/lexer"
	"github.com/syedazeez337/monkey-interpreter/token"
)

const PROMT = ">>"

func Start(in io.Reader, out io.Writer) {

	line := liner.NewLiner()
	defer line.Close()

	line.SetCtrlCAborts(true)
	line.SetCompleter(
		func(line string) (c []string) {
			if line == "pri" {
				c = append(c, "print", "println")
			}
			return
		})

	// scanner := bufio.NewScanner(in)

	for {
		input, err := line.Prompt("mnky>>")
		if err != nil {
			if err == liner.ErrPromptAborted {
				fmt.Println("\nAborted...")
				break
			}

			fmt.Fprintln(os.Stderr, "Error reading line: ", err)
		}

		if input == "exit" {
			fmt.Println("Goodbye...")
			break
		}

		/*
			fmt.Printf(PROMT)
			scanned := scanner.Scan()
			if !scanned {
				return
			}
		*/

		// line := scanner.Text()

		l := lexer.New(input)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}

		line.AppendHistory(input)
	}
}
