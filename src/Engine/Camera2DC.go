package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera2DC struct {
	Parent *GameObject
	Camera rl.Camera2D
}

func NewCamera2DC() *Camera2DC {
	return &Camera2DC{
		Camera: rl.Camera2D{
			Offset:   rl.NewVector2(float32(rl.GetScreenWidth()/2), float32(rl.GetScreenHeight()/2)),
			Target:   rl.NewVector2(0, 0),
			Rotation: 0,
			Zoom:     1.0,
		},
	}
}

func (c *Camera2DC) Start() {
	if c.Parent.FrameData != nil {
		c.Camera.Offset = rl.NewVector2(
			float32(c.Parent.FrameData.Width)/2,
			float32(c.Parent.FrameData.Height)/2,
		)
	}
}

func (c *Camera2DC) Update() {

	c.Camera.Target = c.Parent.Transform.Position

	if c.Parent.FrameData != nil {
		c.Camera.Offset.X = c.Parent.FrameData.Width / 2
		c.Camera.Offset.Y = c.Parent.FrameData.Height / 2
	}
}
func (c *Camera2DC) Draw() {

}
func (c *Camera2DC) SetParent(p *GameObject) {
	c.Parent = p
}
