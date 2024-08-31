package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	printmenu()

	for {
		consoleReader := bufio.NewReaderSize(os.Stdin, 1)
		input, _ := consoleReader.ReadByte()
		var ascii = input

		// ESC = 27 and Ctrl-C = 3
		if ascii == 27 || ascii == 3 {
			fmt.Println("Exiting...")
			os.Exit(0)
		}
		fmt.Println("ASCII : ", ascii)
	}
}

func printmenu() {
	fmt.Println("1. D3")
	fmt.Println("2. D4")
	fmt.Println("3. D6")
	fmt.Println("4. D8")
	fmt.Println("5. D10")
	fmt.Println("6. D12")
	fmt.Println("7. D20")
	fmt.Println("8. D100")
	fmt.Println("A. Enter String")

}
