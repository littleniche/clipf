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
)

func ThrowError(err error) {
	fmt.Printf("%v[ERROR]%v : %v\n", RedText, NormalText, err.Error())
	os.Exit(1)
}

func Copy(file string) {

	var content string

	readFile, err := os.Open(file)

	if err != nil {
		ThrowError(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		content += line + "\n"
	}

	clipboard.WriteAll(content)
}

func main() {
	argLength := len(os.Args[1:])

	if argLength < 1 {
		fmt.Printf("gclip : %vEnter name of the file to be copied to clipboard%v\n", RedText, NormalText)
		os.Exit(1)
	}

	filename := os.Args[1]
	Copy(filename)

}
