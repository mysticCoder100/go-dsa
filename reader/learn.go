package reader

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func TestScan() {
	var name string
	fmt.Print("Entre your name: ")
	scanned, err := fmt.Scanln(&name)
	if err != nil {
		return
	}
	fmt.Println("I scanned ", scanned, " characters")
	fmt.Printf("Hello, %s!\n", name)
}

func TestBufio() {
	r := bufio.NewScanner(os.Stdin)
	fmt.Print("Entre your name: ")

	if r.Scan() {
		line := r.Text()
		fmt.Printf("Hello, %s!\n", line)
	}
}

func TestReadingFile() {
	file, err := os.Open("resume.pdf")

	if err != nil {
		fmt.Println(err)
		return
	}

	defer file.Close()

	content, err := io.ReadAll(file)
	fmt.Println(content)
}
