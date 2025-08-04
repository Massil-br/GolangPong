package src

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	time "github.com/Massil-br/GolangPong/src/Engine/Time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

var Sm *engine.SceneManager

func AppLoop() {
	font := rl.LoadFont("assets/font/font.ttf")
	defer rl.UnloadFont(font)
	Sm = &engine.SceneManager{}

	SceneMenu := &engine.Scene{FrameData: &time.Data}

	Sm.AddScene(SceneMenu)

	Sm.SetScene(0)

	SetupMainMenu(SceneMenu, &font)

	for !rl.WindowShouldClose() {

		time.Update()
		Sm.Update()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		Sm.Draw()
		rl.EndDrawing()
	}
}
