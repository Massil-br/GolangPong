package mainmenu

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	time "github.com/Massil-br/GolangPong/src/Engine/Time"
	mainmenuscript "github.com/Massil-br/GolangPong/src/scripts/MainMenuScript"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func SetupQuitButton(targetScene *engine.Scene, font *rl.Font) {
	//gameObject
	quitbutton := &engine.GameObject{Name: "QuitButton",
		Active: true,
		Transform: engine.Transform2D{
			Position: rl.NewVector2(0, 0),
			Rotation: 0,
			Scale:    rl.NewVector2(1, 1),
		},
		FrameData: &time.Data,
	}

	text := engine.NewText(font, "Quit", rl.White, 15, 8, 50, 70)

	//add component to gameobject
	quitbutton.AddComponent(text)
	quitbutton.AddComponent(&mainmenuscript.QuitButton{})
	//add gamobject to scene
	targetScene.AddObject(quitbutton)
}
