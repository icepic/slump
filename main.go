package main

import (
	"fmt"
	"time"
	"math/rand/v2"
	"github.com/pkg/term"
)

func main() {

	for {
		printmenu()
		ascii, keycode, err := getChar()
		if err != nil {
			continue
		}
		if ascii == 0 {
			ascii = keycode
		}
		if ascii >= 48 && ascii < 58 {
			fmt.Print("Result: ")
			fmt.Println(do_rnd(byte(ascii - 48)))
			fmt.Println()
			fmt.Println("Press a key")
			time.Sleep(10)
		}
	}
}

func do_rnd(tmp byte) int {
	newtmp := values(int(tmp))
	if newtmp > 0 {
		return rand.N(newtmp) + 1
	} else {
		return 0
	}
}

func getChar() (ascii int, keyCode int, err error) {
	t, _ := term.Open("/dev/tty")
	term.RawMode(t)
	bytes := make([]byte, 3)

	defer t.Restore()
	defer t.Close()

	var numRead int
	numRead, err = t.Read(bytes)
	if err != nil {
		return
	}
	if numRead == 3 && bytes[0] == 27 && bytes[1] == 91 {
		// Three-character control sequence, beginning with "ESC-[".

		// Since there are no ASCII codes for arrow keys, we use
		// Javascript key codes.
		if bytes[2] == 65 {
			// Up
			keyCode = 38
		} else if bytes[2] == 66 {
			// Down
			keyCode = 40
		} else if bytes[2] == 67 {
			// Right
			keyCode = 39
		} else if bytes[2] == 68 {
			// Left
			keyCode = 37
		}
	} else if numRead == 1 {
		ascii = int(bytes[0])
	} else {
		// Two characters read??
	}
	return
}

func printmenu() {

	// clear screen
	fmt.Print("\033[H\033[2J")

	for i := 1; i < 10; i++ {

		fmt.Print(i)
		fmt.Print(". D")
		fmt.Println(values(i))
	}
	fmt.Println("A. Enter String")

}

func values(tmp int) int {
	if tmp == 1 {
		return 2
	}
	if tmp == 2 {
		return 3
	}
	if tmp == 3 {
		return 4
	}
	if tmp == 4 {
		return 6
	}
	if tmp == 5 {
		return 8
	}
	if tmp == 6 {
		return 10
	}
	if tmp == 7 {
		return 12
	}
	if tmp == 8 {
		return 20
	}
	if tmp == 9 {
		return 100
	}
	return 0
}
