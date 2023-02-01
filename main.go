package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/atotto/clipboard"
)

const (
	RedText    = "\033[31m"
	NormalText = "\033[0m"
	NewLine    = "\n"
)

func ThrowError(err error) {
	fmt.Printf("%v[ERROR]%v : %v\n", RedText, NormalText, err.Error())
	os.Exit(1)
}

func Copy(file string) (string, error) {

	var content string

	readFile, err := os.Open(file)

	if err != nil {
		return "", err
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		content += (line + NewLine)
	}

	return content, nil
}

func ReadAll(argLength int, Args []string) {

	var text string
	var verbose string

	for i := 1; i <= argLength; i++ {
		content, err := Copy(Args[i])

		if err != nil {
			ThrowError(err)
		}

		text += (content + NewLine)
		verbose += "gclear: ✅ Copied " + Args[i] + " to Clipboard\n"

	}

	fmt.Printf("%v", verbose)

	clipboard.WriteAll(text)
}

func main() {
	argLength := len(os.Args[1:])

	if argLength < 1 {
		fmt.Printf("gclip : %vNot enough arguments%v\n", RedText, NormalText)
		os.Exit(1)
	}

	ReadAll(argLength, os.Args)

}
