package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Collider struct {
	Parent    *GameObject
	Bounds    rl.Rectangle
	IsTrigger bool
	Colliding bool
	Triggered bool
}

func NewCollider(isTrigger bool) *Collider {
	return &Collider{
		IsTrigger: isTrigger,
	}
}

func (c *Collider) Start() {
	c.UpdateBounds()
}

func (c *Collider) Update() {
	c.UpdateBounds()
}

func (c *Collider) Draw() {
}

func (t *Collider) SetParent(o *GameObject) {
	t.Parent = o
}

func (c *Collider) UpdateBounds() {
	c.Bounds.X = c.Parent.Transform.Position.X
	c.Bounds.Y = c.Parent.Transform.Position.Y
	c.Bounds.Width = c.Parent.Transform.Scale.X
	c.Bounds.Height = c.Parent.Transform.Scale.Y
}
