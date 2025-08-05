package src

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	time "github.com/Massil-br/GolangPong/src/Engine/Time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func AppLoop() {
	font := rl.LoadFont("assets/font/font.ttf")
	defer rl.UnloadFont(font)
	engine.Sm = &engine.SceneManager{}
	running := &engine.Running
	SceneMenu := &engine.Scene{FrameData: &time.Data}
	SceneGame := &engine.Scene{FrameData: &time.Data}
	engine.Sm.AddScene(SceneMenu)
	engine.Sm.AddScene(SceneGame)
	engine.Sm.SetScene(0)


	SetupMainMenu(SceneMenu, &font)
	SetupGameScene(SceneGame)

	for !rl.WindowShouldClose() && *running {

		time.Update()
		engine.Sm.Update()

		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		
		

		engine.Sm.Draw()

		rl.EndDrawing()
	}
}
