package main

import (
	"fmt"
	"os"
	"os/exec"
	engine "snake-go/ui-engine"
	"time"
)

const (
	KeyCodeArrowUp    = "\x1b[A"
	KeyCodeArrowDown  = "\x1b[B"
	KeyCodeArrowRight = "\x1b[C"
	KeyCodeArrowLeft  = "\x1b[D"
)

func main() {

	exec.Command("stty", "-F", "/dev/tty", "cbreak", "min", "1").Run()
	var buf [3]byte

	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	engine.ClearScreen()
	engine.InitializeScreenBuffer()
	engine.DrawRectangle(engine.Coordinates{X: 0, Y: 0}, engine.Coordinates{X: engine.ScreenSize - 1, Y: engine.ScreenSize - 1})
	range_start := engine.Coordinates{X: 1, Y: 1}
	range_end := engine.Coordinates{X: engine.ScreenSize - 2, Y: engine.ScreenSize - 2}
	foodPosition := engine.RandomDotInsideBox(engine.Coordinates{X: 1, Y: 1}, engine.Coordinates{X: engine.ScreenSize - 2, Y: engine.ScreenSize - 2}, 'o')

	engine.PlaceElement(engine.Coordinates{X: 1, Y: 1}, '-')
	engine.RenderScreen()
	time.Sleep(500 * time.Millisecond)

	currentPosition := engine.Coordinates{X: 1, Y: 1}
	currentVelocity := engine.Coordinates{X: 0, Y: 1}

	go func() {
		for {
			_, err := os.Stdin.Read(buf[:])
			if err != nil {
				fmt.Println("Error reading from stdin:", err)
				return
			}
			input := string(buf[:])
			switch input {
			case KeyCodeArrowUp:
				currentVelocity = engine.Coordinates{X: -1, Y: 0}
			case KeyCodeArrowDown:
				currentVelocity = engine.Coordinates{X: 1, Y: 0}
			case KeyCodeArrowRight:
				currentVelocity = engine.Coordinates{X: 0, Y: 1}
			case KeyCodeArrowLeft:
				currentVelocity = engine.Coordinates{X: 0, Y: -1}
			case "q":
				fmt.Println("Exiting...")
				os.Exit(0)
			default:
				fmt.Println("Unknown key pressed:", input)
			}
		}
	}()
	for {
		engine.ClearScreen()
		currentPosition = engine.MoveForward(currentPosition, currentVelocity, '-', range_start, range_end)
		if currentPosition.X == foodPosition.X && currentPosition.Y == foodPosition.Y {
			foodPosition = engine.RandomDotInsideBox(engine.Coordinates{X: 1, Y: 1}, engine.Coordinates{X: engine.ScreenSize - 2, Y: engine.ScreenSize - 2}, 'o')

		}
		engine.RenderScreen()
		time.Sleep(500 * time.Millisecond)
	}

	/*for i := 0; i < 60; i++ {
		for j := 0; j < engine.ScreenSize; j++ {
			engine.ClearScreen()
			engine.InitializeScreenBuffer()
			engine.DrawLine(engine.Coordinates{X: 0, Y: j}, engine.Coordinates{X: engine.ScreenSize / 2, Y: engine.ScreenSize / 2})
			engine.RenderScreen()
			time.Sleep(time.Duration(1000*60/(4*engine.ScreenSize)) * time.Millisecond)
		}
		for j := 0; j < engine.ScreenSize; j++ {
			engine.ClearScreen()
			engine.InitializeScreenBuffer()
			engine.DrawLine(engine.Coordinates{X: j, Y: engine.ScreenSize - 1}, engine.Coordinates{X: engine.ScreenSize / 2, Y: engine.ScreenSize / 2})
			engine.RenderScreen()
			time.Sleep(time.Duration(1000*60/(4*engine.ScreenSize)) * time.Millisecond)
		}
		for j := 0; j < engine.ScreenSize; j++ {
			engine.ClearScreen()
			engine.InitializeScreenBuffer()
			engine.DrawLine(engine.Coordinates{X: engine.ScreenSize - 1, Y: engine.ScreenSize - 1 - j}, engine.Coordinates{X: engine.ScreenSize / 2, Y: engine.ScreenSize / 2})
			engine.RenderScreen()
			time.Sleep(time.Duration(1000*60/(4*engine.ScreenSize)) * time.Millisecond)
		}
		for j := 0; j < engine.ScreenSize; j++ {
			engine.ClearScreen()
			engine.InitializeScreenBuffer()
			engine.DrawLine(engine.Coordinates{X: engine.ScreenSize - 1 - j, Y: 0}, engine.Coordinates{X: engine.ScreenSize / 2, Y: engine.ScreenSize / 2})
			engine.RenderScreen()
			time.Sleep(time.Duration(1000*60/(4*engine.ScreenSize)) * time.Millisecond)
		}

	}*/
}
