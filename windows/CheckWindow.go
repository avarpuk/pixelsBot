package windows

type IntWindow interface {
	SetWindowPos(x int, y int)
	SetFirstIconPos(x int, y int)
	SetFirstFieldPos(x int, y int)
}
