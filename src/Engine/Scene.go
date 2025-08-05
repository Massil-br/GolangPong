package engine

import (
	time "github.com/Massil-br/GolangPong/src/Engine/Time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Scene struct {
	GameObjects []*GameObject
	FrameData   *time.FrameData
	Camera		*Camera2DC
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
	if s.Camera != nil{
		rl.BeginMode2D(s.Camera.Camera)
	}
	for _, o := range s.GameObjects {
		o.Draw()
	}
	if s.Camera != nil{
		rl.EndMode2D()
	}
}

func (s *Scene)SetCamera(obj *GameObject){
	camera:= GetComponent[*Camera2DC](obj)
	if camera == nil{
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
		// Mode Trigger → pas de blocage, juste état activé
		a.Triggered = true
		b.Triggered = true
		return
	}

	// --- Mode collision physique ---
	a.Colliding = true
	b.Colliding = true

	// Trouver le déplacement minimum pour séparer les objets
	overlapX := min(a.Bounds.X+a.Bounds.Width, b.Bounds.X+b.Bounds.Width) - max(a.Bounds.X, b.Bounds.X)
	overlapY := min(a.Bounds.Y+a.Bounds.Height, b.Bounds.Y+b.Bounds.Height) - max(a.Bounds.Y, b.Bounds.Y)

	if overlapX < overlapY {
		// Séparation horizontale
		if a.Bounds.X < b.Bounds.X {
			a.Parent.Transform.Position.X -= overlapX
		} else {
			a.Parent.Transform.Position.X += overlapX
		}
	} else {
		// Séparation verticale
		if a.Bounds.Y < b.Bounds.Y {
			a.Parent.Transform.Position.Y -= overlapY
		} else {
			a.Parent.Transform.Position.Y += overlapY
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
