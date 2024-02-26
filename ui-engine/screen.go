package uiengine

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const ScreenSize int = 40

var Screen2DBuffer [ScreenSize][ScreenSize]byte

func ClearScreen() {
	cmd := exec.Command("clear")
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func InitializeScreenBuffer() {
	for i := 0; i < ScreenSize; i++ {
		for j := 0; j < ScreenSize; j++ {
			Screen2DBuffer[i][j] = ' '
		}
	}
}

type Coordinates struct {
	X int
	Y int
}

func isValidCoordinates(axis Coordinates) bool {
	if axis.X < 0 || axis.Y < 0 || axis.X >= ScreenSize || axis.Y >= ScreenSize {
		return false
	}
	return true
}

func PlaceElement(position Coordinates, element byte) {
	if isValidCoordinates(position) {
		Screen2DBuffer[position.X][position.Y] = element
	}
}

func MoveForward(currentPosition Coordinates, velocity Coordinates, element byte, range_start Coordinates, range_end Coordinates) Coordinates {
	if isValidCoordinates(currentPosition) {
		Screen2DBuffer[currentPosition.X][currentPosition.Y] = ' '
	}
	nextPosition := Coordinates{X: (((currentPosition.X + velocity.X - range_start.X) % (range_end.X - range_start.X + 1)) + range_start.X), Y: ((currentPosition.Y+velocity.Y-range_start.Y)%(range_end.Y-range_start.Y+1) + range_start.Y)}
	if isValidCoordinates(nextPosition) {
		Screen2DBuffer[nextPosition.X][nextPosition.Y] = element
	}
	return nextPosition
}

func RandomDotInsideBox(start Coordinates, end Coordinates, element byte) Coordinates {
	rand.Seed(time.Now().UnixNano())
	x_min := start.X
	y_min := start.Y
	x_max := end.X
	y_max := end.Y
	x_rand := rand.Intn(x_max-x_min+1) + x_min
	y_rand := rand.Intn(y_max-y_min+1) + y_min
	if isValidCoordinates(Coordinates{X: x_rand, Y: y_rand}) {
		Screen2DBuffer[x_rand][y_rand] = element
	}
	return Coordinates{X: x_rand, Y: y_rand}
}
func DrawLine(start Coordinates, end Coordinates) {
	for x_coordinate := start.X; x_coordinate < end.X && x_coordinate >= 0 && x_coordinate < ScreenSize; x_coordinate++ {
		y_coordinate := (end.Y-start.Y)*(x_coordinate-start.X)/(end.X-start.X) + start.Y
		if isValidCoordinates(Coordinates{X: x_coordinate, Y: y_coordinate}) {
			Screen2DBuffer[x_coordinate][y_coordinate] = '.'
		}
	}
	for y_coordinate := start.Y; y_coordinate < end.Y && y_coordinate >= 0 && y_coordinate < ScreenSize; y_coordinate++ {
		x_coordinate := (end.X-start.X)*(y_coordinate-start.Y)/(end.Y-start.Y) + start.X
		if isValidCoordinates(Coordinates{X: x_coordinate, Y: y_coordinate}) {
			Screen2DBuffer[x_coordinate][y_coordinate] = '.'
		}
	}
	for x_coordinate := start.X; x_coordinate > end.X && x_coordinate >= 0 && x_coordinate < ScreenSize; x_coordinate-- {
		y_coordinate := (end.Y-start.Y)*(x_coordinate-start.X)/(end.X-start.X) + start.Y
		if isValidCoordinates(Coordinates{X: x_coordinate, Y: y_coordinate}) {
			Screen2DBuffer[x_coordinate][y_coordinate] = '.'
		}
	}
	for y_coordinate := start.Y; y_coordinate > end.Y && y_coordinate >= 0 && y_coordinate < ScreenSize; y_coordinate-- {
		x_coordinate := (end.X-start.X)*(y_coordinate-start.Y)/(end.Y-start.Y) + start.X
		if isValidCoordinates(Coordinates{X: x_coordinate, Y: y_coordinate}) {
			Screen2DBuffer[x_coordinate][y_coordinate] = '.'
		}
	}
}

func DrawRectangle(start Coordinates, end Coordinates) {
	DrawLine(Coordinates{X: start.X, Y: start.Y}, Coordinates{X: start.X, Y: end.Y})
	DrawLine(Coordinates{X: start.X, Y: start.Y}, Coordinates{X: end.X, Y: start.Y})
	DrawLine(Coordinates{X: end.X, Y: start.Y}, Coordinates{X: end.X, Y: end.Y})
	DrawLine(Coordinates{X: start.X, Y: end.Y}, Coordinates{X: end.X, Y: end.Y})
}

func RenderScreen() {
	for i := 0; i < ScreenSize; i++ {
		for j := 0; j < ScreenSize; j++ {
			fmt.Print(string(Screen2DBuffer[i][j]))
		}
		fmt.Println()
	}
}
