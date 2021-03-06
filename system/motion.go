package system

import (
	"github.com/javinc/go-space-shooter"
	"github.com/javinc/go-space-shooter/component"
)

// Motion system.
type Motion struct{}

// NewMotion Motion system constructor.
func NewMotion() *Motion {
	return &Motion{}
}

// Process Control system implements System interface.
func (s *Motion) Process(em *ecs.EntityManager) {
	for _, e := range em.All() {
		cm := e.ComponentManager()
		if !cm.Requires("rect", "position", "velocity", "projectile") {
			continue
		}

		rect := cm.Get("rect").(*component.Rect)
		if !rect.Active {
			continue
		}

		pos := cm.Get("position").(*component.Position)
		vel := cm.Get("velocity").(*component.Velocity)

		// Shoot up
		pos.Y -= vel.Speed

		// Out of bounds for reuse
		if pos.Y < float64(-rect.H) || pos.X < float64(-rect.W) {
			rect.Active = false
		}
	}
}
