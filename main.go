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

func GetFileContent(readFile *os.File) string {
	var content string

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		content += (line + NewLine)
	}

	return content
}

func Copy(file string) (string, error) {

	var content string

	readFile, err := os.Open(file)

	if err != nil {
		return "", err
	}

	content = GetFileContent(readFile)

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

		if i != argLength {
			text += (content + NewLine)
		} else {
			text += content
		}

		verbose += "clipf: ✅ Copied " + Args[i] + " to Clipboard\n"

	}

	fmt.Printf("%v", verbose)

	clipboard.WriteAll(text)
}

func ReadStdin() {
	var content string

	content = GetFileContent(os.Stdin)

	clipboard.WriteAll(content)
}

func IsStdin() bool {
	stat, _ := os.Stdin.Stat()

	// We only treat piped input as stdin
	if (stat.Mode() & os.ModeCharDevice) == 0 {
		return true
	}

	return false
}

func WriteAll(argLength int, Args []string) {
	text, err := clipboard.ReadAll()
	if err != nil {
		ThrowError(err)
	}

	for i := 2; i <= argLength; i++ {
		file, err := os.Create(Args[i])
		if err != nil {
			ThrowError(err)
		}
		defer file.Close()

		_, err = file.WriteString(text)
		if err != nil {
			ThrowError(err)
		}
	}
}

func main() {
	var stdin bool
	write := flag.Bool("w", false, "Write content of clipboard to files")
	flag.Parse()

	argLength := len(os.Args[1:])
	stdin = IsStdin()

	if argLength < 1 && !stdin {
		fmt.Printf("clipf : %vNot enough arguments%v\n", RedText, NormalText)
		os.Exit(1)
	}

	if *write {
		WriteAll(argLength, os.Args)
		return
	}

	ReadAll(argLength, os.Args)
	if argLength < 1 && stdin {
		ReadStdin()
		fmt.Printf("clipf: ✅ Copied stdin to Clipboard\n")
	}
}
