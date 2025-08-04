package engine

import rl "github.com/gen2brain/raylib-go/raylib"

type Transform2D struct {
	Position rl.Vector2
	Rotation float32
	Scale    rl.Vector2
}

func (t *Transform2D) Start(){}
func (t *Transform2D) Update(){}
func (t *Transform2D) Draw(){}
