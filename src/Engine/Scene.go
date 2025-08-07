package engine

import (
	time "github.com/Massil-br/GolangPong/src/Engine/Time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	GameObjects []*GameObject
	FrameData   *time.FrameData
	Camera      *Camera2DC
}

func (s *Scene) AddObject(o *GameObject) {
	o.FrameData = s.FrameData
	s.GameObjects = append(s.GameObjects, o)
}

func (s *Scene) Update() {
	for _, o := range s.GameObjects {
		o.Update()
	}
}

func (s *Scene) Draw() {
	if s.Camera != nil {
		rl.BeginMode2D(s.Camera.Camera)
	}
	for _, o := range s.GameObjects {
		o.Draw()
	}
	if s.Camera != nil {
		rl.EndMode2D()
	}
}

func (s *Scene) SetCamera(obj *GameObject) {
	camera := GetComponent[*Camera2DC](obj)
	if camera == nil {
		return
	}
	s.Camera = camera
}

func (s *Scene) PhysicsUpdate() {
	for i, objA := range s.GameObjects {
		colA := GetComponent[*Collider](objA)
		if colA == nil {
			continue
		}
		for j, objB := range s.GameObjects {
			if i == j {
				continue
			}
			colB := GetComponent[*Collider](objB)
			if colB == nil {
				continue
			}
			ResolveCollision(colA, colB)
		}
	}
}

func ResolveCollision(a, b *Collider) {
	if !rl.CheckCollisionRecs(a.Bounds, b.Bounds) {
		return
	}

	if a.IsTrigger || b.IsTrigger {
		a.Triggered = true
		b.Triggered = true
		return
	}

	a.Colliding = true
	b.Colliding = true

	// On identifie qui peut bouger
	aRb := GetComponent[*Rigidbody2D](a.Parent)
	bRb := GetComponent[*Rigidbody2D](b.Parent)

	overlapX := min(a.Bounds.X+a.Bounds.Width, b.Bounds.X+b.Bounds.Width) - max(a.Bounds.X, b.Bounds.X)
	overlapY := min(a.Bounds.Y+a.Bounds.Height, b.Bounds.Y+b.Bounds.Height) - max(a.Bounds.Y, b.Bounds.Y)

	

	// Appliquer uniquement si l'objet a un Rigidbody
	if overlapX < overlapY {
		if a.Bounds.X < b.Bounds.X {
			if aRb != nil {
				a.Parent.Transform.Position.X -= overlapX
			} else if bRb != nil {
				b.Parent.Transform.Position.X += overlapX
			}
		} else {
			if aRb != nil {
				a.Parent.Transform.Position.X += overlapX
			} else if bRb != nil {
				b.Parent.Transform.Position.X -= overlapX
			}
		}
	} else {
		if a.Bounds.Y < b.Bounds.Y {
			if aRb != nil {
				a.Parent.Transform.Position.Y -= overlapY
			} else if bRb != nil {
				b.Parent.Transform.Position.Y += overlapY
			}
		} else {
			if aRb != nil {
				a.Parent.Transform.Position.Y += overlapY
			} else if bRb != nil {
				b.Parent.Transform.Position.Y -= overlapY
			}
		}
	}
}

func min(a, b float32) float32 {
	if a < b {
		return a
	}
	return b
}

func max(a, b float32) float32 {
	if a > b {
		return a
	}
	return b
}
