package main

import (
	"fmt"
	"io/ioutil"
	"os"

	bfri "github.com/russross/blackfriday"
)

func main() {
	arguments := os.Args
	if len(arguments) < 2 {
		fmt.Println("Podaj nazwę pliku który chcesz wczytać")
		os.Exit(1)
	}
	filename := arguments[1]
	fcontents, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Błąd wczytywania pliku %v \n", filename)
		os.Exit(1)
	}

	os.Stdout.Write(bfri.MarkdownBasic(fcontents))
}
