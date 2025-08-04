package scripts

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Text struct {
	Parent      *engine.GameObject
	Font        *rl.Font
	Text        string
	Position    rl.Vector2
	FontSize    float32
	Spacing     float32
	Color       rl.Color
	TextSize    rl.Vector2
	DivFontSize float32
	DivSpacing  float32
	PXPourcent  float32
	PYPourcent  float32
}

func (t *Text) Start() {}

func (t *Text) Update() {
	t.FontSize = t.Parent.FrameData.Height / t.DivFontSize
	t.Spacing = t.FontSize / t.DivSpacing
	t.TextSize = rl.MeasureTextEx(*t.Font, t.Text, t.FontSize, t.Spacing)
	t.Position.X = (t.Parent.FrameData.Width - t.TextSize.X) * (t.PXPourcent / 100)
	t.Position.Y = (t.Parent.FrameData.Height - t.TextSize.Y) * (t.PYPourcent / 100)
}
func (t *Text) Draw() {
	rl.DrawTextEx(*t.Font, t.Text, t.Position, t.FontSize, t.Spacing, t.Color)
}
func (t *Text) SetText(text string) {
	t.Text = text
}
func (t *Text) SetColor(color rl.Color) {
	t.Color = color
}
func (t *Text) SetParent(o *engine.GameObject) {
	t.Parent = o
}
