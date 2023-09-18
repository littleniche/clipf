package main

import (
	"bufio"
	"flag"
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
		verbose += "clipf: ✅ Copied " + Args[i] + " to Clipboard\n"

	}

	fmt.Printf("%v", verbose)

	clipboard.WriteAll(text)
}

func WriteAll(argLength int, Args []string) {
	text, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Unable to read clipboard contents")
	}

	for i := 2; i <= argLength; i++ {
		file, err := os.Create(Args[i])
		if err != nil {
			fmt.Println("Failed to create file:", err)
			continue
		}
		defer file.Close()

		_, err = file.WriteString(text)
		if err != nil {
			fmt.Println("Failed to write to file:", err)
			continue
		}
	}
}

func main() {
	write := flag.Bool("w", false, "Write content of clipboard to files")
	flag.Parse()

	argLength := len(os.Args[1:])

	if *write {
		WriteAll(argLength, os.Args)
		return
	}

	if argLength < 1 {
		fmt.Printf("gclip : %vNot enough arguments%v\n", RedText, NormalText)
		os.Exit(1)
	}

	ReadAll(argLength, os.Args)

}
