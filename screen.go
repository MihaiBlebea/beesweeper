package main

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

// Cell height and width
const (
	CellH = 20
	CellW = 20
)

func render() error {
	defer sdl.Quit()

	w, err := window()
	if err != nil {
		return err
	}
	defer w.Destroy()

	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			err = cell(w, i*(CellW+2), j*(CellH+2))
			if err != nil {
				return err
			}
		}
	}

	// Running event loop
	running := true
	for running {
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch eventT := event.(type) {
			case *sdl.QuitEvent:
				println("Quit")
				running = false
				break
			case *sdl.KeyboardEvent:
				switch eventT.Keysym.Sym {
				case sdl.K_UP:
					fmt.Println("Pressed key up")
				case sdl.K_DOWN:
					fmt.Println("Pressed key down")
				case sdl.K_LEFT:
					fmt.Println("Pressed key left")
				case sdl.K_RIGHT:
					fmt.Println("Pressed key right")
				}
			}
		}

		sdl.Delay(16)
	}

	return nil
}

func init() {
	if err := sdl.Init(sdl.INIT_EVERYTHING); err != nil {
		panic(err)
	}
}

func window() (*sdl.Window, error) {
	return sdl.CreateWindow(
		"test",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		800,
		600,
		sdl.WINDOW_SHOWN)
}

func renderer(window *sdl.Window) (*sdl.Renderer, error) {
	return sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
}

func cell(window *sdl.Window, x, y int) error {
	surface, err := window.GetSurface()
	if err != nil {
		return err
	}
	// surface.FillRect(nil, 0)

	rect := sdl.Rect{X: int32(x), Y: int32(y), W: CellW, H: CellH}
	surface.FillRect(&rect, 0xffff0000)
	window.UpdateSurface()

	return nil
}
