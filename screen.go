package main

import (
	"fmt"

	"github.com/MihaiBlebea/beesweeper/game"
	"github.com/veandco/go-sdl2/sdl"
)

// Cell height and width
const (
	CellH = 20
	CellW = 20
)

// Screen is a struct responsible for displaying the game
type Screen struct {
	cellH      int
	cellW      int
	cellCountH int
	cellCountW int
	spacer     int
	game       *game.Game
}

// NewScreen constructor for screen struct
func NewScreen(cellH, cellW, cellCountH, cellCountW, spacer int, gm *game.Game) *Screen {
	// b := gm.GetBoard()
	return &Screen{cellH, cellW, cellCountH, cellCountW, spacer, gm}
}

func (s *Screen) getSceenTotalWidth() int32 {
	screenW := s.cellW*s.cellCountW + s.spacer*(s.cellCountW-1)
	return int32(screenW)
}

func (s *Screen) getSceenTotalHeight() int32 {
	screenH := s.cellH*s.cellCountH + s.spacer*(s.cellCountH-1)
	return int32(screenH)
}

func (s *Screen) render() error {
	defer sdl.Quit()

	w, err := s.window()
	if err != nil {
		return err
	}
	defer w.Destroy()

	r, err := s.renderer(w)
	if err != nil {
		return err
	}
	defer r.Destroy()

	// Render everything in here
	s.drawScene(r)

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

func (s *Screen) window() (*sdl.Window, error) {
	return sdl.CreateWindow(
		"test",
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		s.getSceenTotalWidth(),
		s.getSceenTotalHeight(),
		sdl.WINDOW_SHOWN)
}

func (s *Screen) renderer(window *sdl.Window) (*sdl.Renderer, error) {
	return sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
}

func (s *Screen) cell(rect sdl.Rect, r *sdl.Renderer, isOpen bool) {
	r.SetDrawColor(255, 0, 0, 255)
	r.DrawRect(&rect)
	if isOpen == false {
		r.FillRect(&rect)
	}
}

func (s *Screen) drawScene(r *sdl.Renderer) {
	r.Clear()

	for i := 0; i < s.cellCountW; i++ {
		for j := 0; j < s.cellCountH; j++ {
			rect := sdl.Rect{
				X: int32(i * (CellW + s.spacer)),
				Y: int32(j * (CellH + s.spacer)),
				W: int32(s.cellW),
				H: int32(s.cellH),
			}
			s.cell(rect, r, false)
		}
	}

	r.Present()
}
