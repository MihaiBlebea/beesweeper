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
	cellH int
	cellW int
	// cellCountH int
	// cellCountW int
	spacer int
	game   *game.Game
}

// NewScreen constructor for screen struct
func NewScreen(cellH, cellW, spacer int, gm *game.Game) *Screen {
	// b := gm.GetBoard()
	gm.GetBoard().SetSelected(0, 0)

	return &Screen{cellH, cellW, spacer, gm}
}

func (s *Screen) getSceenTotalWidth() int32 {
	cellCountW := s.game.GetBoard().GetCellCountW()
	screenW := s.cellW*cellCountW + s.spacer*(cellCountW-1)
	return int32(screenW)
}

func (s *Screen) getSceenTotalHeight() int32 {
	cellCountH := s.game.GetBoard().GetCellCountH()
	screenH := s.cellH*cellCountH + s.spacer*(cellCountH-1)
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
	err = s.drawScene(r)
	if err != nil {
		return err
	}

	// Running event loop
	selectedX := 0
	selectedY := 0

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
					if selectedY > 0 {
						selectedY--
					}
				case sdl.K_DOWN:
					fmt.Println("Pressed key down")
					if selectedY < s.game.GetBoard().GetCellCountH()-1 {
						selectedY++
					}
				case sdl.K_LEFT:
					fmt.Println("Pressed key left")
					if selectedX > 0 {
						selectedX--
					}
				case sdl.K_RIGHT:
					fmt.Println("Pressed key right")
					if selectedX < s.game.GetBoard().GetCellCountW()-1 {
						selectedX++
					}
				}

				s.game.GetBoard().UnselectAll()
				s.game.GetBoard().SetSelected(selectedX, selectedY)
				err = s.drawScene(r)
				if err != nil {
					return err
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

func (s *Screen) cell(rect sdl.Rect, r *sdl.Renderer, isOpen, isSelected bool) {
	// Set color
	r.SetDrawColor(s.getCellColor(isSelected))
	r.DrawRect(&rect)
	if isOpen == false {
		r.FillRect(&rect)
	}
}

func (s *Screen) getCellColor(isSelected bool) (r, g, b, a uint8) {
	if isSelected {
		// blue
		return 0, 0, 255, 255
	}

	// red
	return 255, 0, 0, 255
}

func (s *Screen) drawScene(r *sdl.Renderer) error {
	err := r.Clear()
	if err != nil {
		return err
	}

	board := s.game.GetBoard()
	cellCountW := board.GetCellCountW()
	cellCountH := board.GetCellCountH()

	for i := 0; i < cellCountW; i++ {
		for j := 0; j < cellCountH; j++ {
			cell := board.GetCell(i, j)

			rect := sdl.Rect{
				X: int32(i * (CellW + s.spacer)),
				Y: int32(j * (CellH + s.spacer)),
				W: int32(s.cellW),
				H: int32(s.cellH),
			}
			s.cell(rect, r, cell.HasBee(), cell.IsSelected())
		}
	}

	r.Present()

	return nil
}
