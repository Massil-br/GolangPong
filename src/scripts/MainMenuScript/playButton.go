package mainmenuscript

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"

	rl "github.com/gen2brain/raylib-go/raylib"
	//rl "github.com/gen2brain/raylib-go/raylib"
)

type PlayButton struct {
	Parent        *engine.GameObject
	Collider      rl.Rectangle
	MousePosition rl.Vector2
	Text          *engine.Text
	BaseColor     rl.Color
	HoverColor    rl.Color
}

func (p *PlayButton) Start() {
	p.MousePosition = rl.GetMousePosition()
	p.Text = engine.GetComponent[*engine.Text](p.Parent)
	p.Collider = rl.NewRectangle(p.Text.Position.X, p.Text.Position.Y, p.Text.TextSize.X, p.Text.TextSize.Y)
	p.BaseColor = p.Text.Color
	p.HoverColor = rl.Red
}
func (p *PlayButton) Update() {
	p.Collider.X = p.Text.Position.X
	p.Collider.Y = p.Text.Position.Y
	p.Collider.Width = p.Text.TextSize.X
	p.Collider.Height = p.Text.TextSize.Y
	p.MousePosition = rl.GetMousePosition()
	if rl.CheckCollisionPointRec(p.MousePosition, p.Collider) {
		p.Text.Color = p.HoverColor
		if rl.IsMouseButtonPressed(rl.MouseLeftButton) {
			engine.Sm.SetScene(1)
		}
	} else {
		p.Text.Color = p.BaseColor
	}
}
func (p *PlayButton) Draw() {

}

func (p *PlayButton) SetParent(o *engine.GameObject) {
	p.Parent = o
}
