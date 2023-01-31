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
	NewLine = "\n"
)


func ThrowError(err error) {
	fmt.Printf("%v[ERROR]%v : %v\n", RedText, NormalText, err.Error())
	os.Exit(1)
}


func Copy(file string)string {

	var content string

	readFile, err := os.Open(file)

	if err != nil {
		ThrowError(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	for fileScanner.Scan() {
		line := fileScanner.Text()

		content += ( line + NewLine )
	}

	return content
}


func ReadAll(argLength int, Args []string){

	var content string;

	for i:=1 ; i<=argLength; i++ {
		content += Copy(Args[i]) + NewLine
		fmt.Printf("gclear: ✅ Copied %v to Clipboard\n", Args[i]);
	}

	clipboard.WriteAll(content)
}


func main() {
	argLength := len(os.Args[1:])

	if argLength < 1 {
		fmt.Printf("gclip : %varguments%v\n", RedText, NormalText)
		os.Exit(1)
	}

	ReadAll(argLength, os.Args)

}