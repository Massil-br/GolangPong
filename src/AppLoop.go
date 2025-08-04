package src

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	time "github.com/Massil-br/GolangPong/src/Engine/Time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func AppLoop() {
	font := rl.LoadFont("assets/font/font.ttf")
	defer rl.UnloadFont(font)
	sm := &engine.SceneManager{}

	SceneMenu := &engine.Scene{FrameData: &time.Data}

	sm.AddScene(SceneMenu)
	
	sm.SetScene(0)

	SetupMainMenu(SceneMenu,&font)

	for !rl.WindowShouldClose() {

		time.Update()
		sm.Update()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		sm.Draw()
		rl.EndDrawing()
	}
}
