package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"strings"
)

func func1() {
	reader := strings.NewReader("Talk is cheap ,Show me Code")
	p := make([]byte, 8)
	for {
		n, err := reader.Read(p)
		if err == io.EOF {
			break
		}
		fmt.Println(string(p[:n]))
	}
}

func func2() {
	meats := []string{
		"chicken",
		"|beef",
		"|pork",
		"|mutton",
	}

	var writer bytes.Buffer
	fmt.Println("11" + writer.String())
	for _, p := range meats {
		n, err := writer.Write([]byte(p))
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		if n != len(p) {
			fmt.Println("n !=  len(p)")
			os.Exit(1)
		}
	}

	fmt.Println("22" + writer.String())

}

func main() {

	//func1()
	func2()

}
