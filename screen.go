package main

import (
	"fmt"
	"strconv"

	"github.com/MihaiBlebea/beesweeper/game"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

// Font and font size
const (
	fontPath = "./assets/test.ttf"
	fontSize = 14
)

// Screen is a struct responsible for displaying the game
type Screen struct {
	cellH  int
	cellW  int
	spacer int
	game   *game.Game
}

// NewScreen constructor for screen struct
func NewScreen(cellH, cellW, spacer int, gm *game.Game) *Screen {
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

	// init the ttf font
	if err = ttf.Init(); err != nil {
		return err
	}
	defer ttf.Quit()

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

				keyCode := eventT.Keysym.Sym

				switch keyCode {
				case sdl.K_UP:
					if eventT.State == sdl.PRESSED && selectedY > 0 {
						fmt.Println("Pressed key up")
						selectedY--
					}
				case sdl.K_DOWN:
					if eventT.State == sdl.PRESSED && selectedY < s.game.GetBoard().GetCellCountH()-1 {
						fmt.Println("Pressed key down")
						selectedY++
					}
				case sdl.K_LEFT:
					if eventT.State == sdl.PRESSED && selectedX > 0 {
						fmt.Println("Pressed key left")
						selectedX--
					}
				case sdl.K_RIGHT:
					if eventT.State == sdl.PRESSED && selectedX < s.game.GetBoard().GetCellCountW()-1 {
						fmt.Println("Pressed key right")
						selectedX++
					}
				case sdl.K_RETURN:
					if eventT.State == sdl.PRESSED {
						fmt.Println("Pressed key enter")
						// selectedX++
						s.game.GetBoard().UncoverCell(selectedX, selectedY)
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
	// Clear the screen to this color: black
	r.SetDrawColor(0, 0, 0, 255)

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
				X: int32(i * (s.cellW + s.spacer)),
				Y: int32(j * (s.cellH + s.spacer)),
				W: int32(s.cellW),
				H: int32(s.cellH),
			}
			s.cell(rect, r, cell.IsDiscovered(), cell.IsSelected())

			if cell.ShouldShowCount() == true {
				s.drawText(r, &rect, strconv.Itoa(cell.GetCount()))
			}
		}
	}

	r.Present()

	return nil
}

func (s *Screen) drawText(r *sdl.Renderer, rect *sdl.Rect, txt string) error {
	font, err := ttf.OpenFont(fontPath, fontSize)
	if err != nil {
		return err
	}
	defer font.Close()

	// Create a red text with the font
	surf, err := font.RenderUTF8Solid(
		txt,
		sdl.Color{R: 255, G: 255, B: 255, A: 255},
	)
	if err != nil {
		return err
	}
	defer surf.Free()

	t, err := r.CreateTextureFromSurface(surf)
	if err != nil {
		return err
	}
	defer t.Destroy()

	err = r.Copy(t, nil, rect)
	if err != nil {
		return err
	}

	return nil
}
