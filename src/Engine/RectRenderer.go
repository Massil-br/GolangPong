package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type RectRenderer struct {
	Parent *GameObject
	Color  rl.Color
}

func NewRenderer(color rl.Color) *RectRenderer {
	return &RectRenderer{Color: color}
}

func (r *RectRenderer) Start() {
}

func (r *RectRenderer) Update() {
}

func (r *RectRenderer) Draw() {
	pos := r.Parent.Transform.Position
	scale := r.Parent.Transform.Scale

	rl.DrawRectangle(int32(pos.X), int32(pos.Y), int32(scale.X), int32(scale.Y), r.Color)
}

func (r *RectRenderer) SetParent(o *GameObject) {
	r.Parent = o
}
