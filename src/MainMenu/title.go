package mainmenu

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	time "github.com/Massil-br/GolangPong/src/Engine/Time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func SetupTitle(targetScene *engine.Scene, font *rl.Font) {
	//gameObject
	title := &engine.GameObject{Name: "Title",
		Active: true,
		Transform: engine.Transform2D{
			Position: rl.NewVector2(0, 0),
			Rotation: 0,
			Scale:    rl.NewVector2(1, 1),
		},
		FrameData: &time.Data,
	}
	//component
	text := &engine.Text{
		Parent:      title,
		Font:        font,
		Text:        "Massil's Pong",
		Position:    title.Transform.Position,
		FontSize:    title.FrameData.Height / 15,
		Spacing:     title.FrameData.Height / 15 / 8,
		Color:       rl.White,
		TextSize:    rl.NewVector2(0, 0),
		DivFontSize: 15,
		DivSpacing:  8,
		PXPourcent:  50,
		PYPourcent:  20,
	}
	//add component to gameobject
	title.AddComponent(text)
	//add gamobject to scene
	targetScene.AddObject(title)
}
