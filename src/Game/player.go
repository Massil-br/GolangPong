package game

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	playerscripts "github.com/Massil-br/GolangPong/src/scripts/PlayerScripts"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func SetupPlayer(targetScene *engine.Scene) {
	player := engine.NewObj("player")
	player.AddComponent(engine.NewCollider(false))
	player.AddComponent(engine.NewRenderer(rl.White))
	player.AddComponent(engine.NewRigidBody2D(0, 1, false))
	player.AddComponent(playerscripts.NewPlayerController(500))
	player.Transform.Scale.X = 50
	player.Transform.Scale.Y = 200
	player.Transform.Position.X -= 600
	targetScene.AddObject(player)
}
