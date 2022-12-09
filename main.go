package main

import (
	"bufio"
	"fmt"
	"github.com/go-vgo/robotgo"
	"os"
	"pixels/windows"
	"strconv"
)

const colorClear = "\033[H\033[2J"

func main() {
	robotgo.Move(1, 2)

	scanner := bufio.NewScanner(os.Stdin)

	window := []windows.IntWindow{
		&windows.Info{},
	}

	for {

		var (
			inputStr, changeStr string
			number              = 0
			selectedWindow      int
		)

		fmt.Println(colorClear)
		fmt.Printf("[Selected window: %d]", selectedWindow)
		fmt.Println("Choose available command:\n" +
			"1. Set number of windows\n" +
			"2. Set window choose location\n" +
			"3. smth")

		if scanner.Scan() {
			inputStr = scanner.Text()
		}

		switch inputStr {
		case "1":
			fmt.Print("Number of windows (use back to switch command): ")
			if scanner.Scan() {
				changeStr = scanner.Text()
			}
			switch changeStr {
			case "back":
				break
			default:
				number = StoI(changeStr)
				// window := make([]windows.Info, number)
			}
		case "2":
			if number == 0 {
				fmt.Println("Error. Number of windows is not set.")
				break
			}
			for i := 3; i > 0; i-- {
				fmt.Printf("\nMouse position will be saved in %d seconds.", i)
			}
			robotgo.Sleep(3)
			// x, y := robotgo.GetMousePos()
		}
		break
	}

	/* window := []windows.Info{}
	for i := 0; i < 10; i++ {
		n := windows.Info{ID: i}
		window = append(window, n)
	}*/
	// fmt.Println(window[0:])
}

func StoI(a string) int {
	i, err := strconv.Atoi(a)
	if err != nil {
		// ... handle error
		panic(err)
	}
	return i
}
