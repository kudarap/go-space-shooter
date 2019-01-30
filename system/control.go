package system

import (
	"github.com/javinc/ecs"
	"github.com/javinc/ecs/component"
	"github.com/veandco/go-sdl2/sdl"
)

// Control system.
type Control struct {
	W, H int32
}

// NewControl Control system constructor.
func NewControl(w, h int32) *Control {
	return &Control{w, h}
}

// Process Control system implements System interface.
func (s *Control) Process(ee []*ecs.Entity) {
	kk := sdl.GetKeyboardState()

	for _, e := range ee {
		if !e.Requires("input", "rect", "position") {
			continue
		}

		rect := e.Get("rect").(*component.Rect)
		pos := e.Get("position").(*component.Position)

		velocity := 0.4
		if kk[sdl.SCANCODE_LEFT] == 1 && pos.X > 0 {
			pos.X -= velocity
		} else if kk[sdl.SCANCODE_RIGHT] == 1 && pos.X < float64(s.W-rect.W) {
			pos.X += velocity
		}
	}
}