package src

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	time "github.com/Massil-br/GolangPong/src/Engine/Time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func AppLoop() {

	sm := &engine.SceneManager{}

	sceneGame := &engine.Scene{DeltaTime: &time.DeltaTime}
	sceneMenu := &engine.Scene{DeltaTime: &time.DeltaTime}

	player := &engine.GameObject{Name: "Player", Active: true}

	player.AddComponent(&engine.Transform2D{
		Position: rl.NewVector2(0, 0),
		Scale:    rl.NewVector2(1, 1),
	})
	sceneGame.AddObject(player)

	sm.AddScene(sceneGame)
	sm.AddScene(sceneMenu)

	sm.SetScene(0)

	for !rl.WindowShouldClose() {
		time.Update()
		sm.Update()
		rl.BeginDrawing()
		rl.ClearBackground(rl.Black)
		sm.Draw()
		rl.EndDrawing()
	}
}
