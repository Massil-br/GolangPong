package playerscripts

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayerController struct {
	Parent *engine.GameObject
	Speed  float32
	rb     *engine.Rigidbody2D
}

func NewPlayerController(speed float32) *PlayerController {
	return &PlayerController{
		Speed: speed,
	}
}

func (pc *PlayerController) Start() {
	pc.rb = engine.GetComponent[*engine.Rigidbody2D](pc.Parent)
}
func (pc *PlayerController) Update() {
	if pc.rb == nil {
		return
	}

	if rl.IsKeyDown(rl.KeyW) {
		pc.rb.Velocity.Y = -pc.Speed
	}
	if rl.IsKeyDown(rl.KeyS) {
		pc.rb.Velocity.Y = pc.Speed
	}

}
func (pc *PlayerController) Draw() {

}
func (pc *PlayerController) SetParent(o *engine.GameObject) {
	pc.Parent = o
}

func (pc *PlayerController) CheckBounds() {
	
}
