package src

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func SetupPlayer(targetScene *engine.Scene) {
	player := &engine.GameObject{Name: "Player",
	Active: true,
	Transform :engine.Transform2D{
		Position:rl.NewVector2(0,0),
		Rotation: 0,
		Scale: rl.NewVector2(1,1)},
	}

	targetScene.AddObject(player)
}


