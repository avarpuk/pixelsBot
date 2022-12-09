package windows

type Info []struct {
	ID int

	ChooseWindowX int
	ChooseWindowY int

	FirstIconX int
	FirstIconY int

	FirstFieldX int
	FirstFieldY int
}

func (c *Info) SetWindowPos(x int, y int) {
	c.ChooseWindowX = x
	c.ChooseWindowY = y
}

func (c *Info) SetFirstIconPos(x int, y int) {
	c.FirstIconX = x
	c.FirstIconY = y
}

func (c *Info) SetFirstFieldPos(x int, y int) {
	c.FirstFieldX = x
	c.FirstFieldY = y
}
