package windows

import "fmt"

type Info struct {
	ID int

	WindowSizeX int

	ChooseWindowX int
	ChooseWindowY int

	FirstIconX int
	FirstIconY int

	FirstFieldX int
	FirstFieldY int
}

func (c *Info) Set() {
	fmt.Println("hi")
}
