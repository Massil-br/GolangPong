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
	player.AddComponent(playerscripts.NewPlayerController(400))
	player.Transform.Scale.X = 16
	player.Transform.Scale.Y = 64
	player.Transform.Position.X -= 300
	targetScene.AddObject(player)

	player2 := engine.NewObj("player2")
	player2.AddComponent(engine.NewCollider(false))
	player2.AddComponent(engine.NewRenderer(rl.White))
	player2.AddComponent(engine.NewRigidBody2D(0, 1, false))
	player2.Transform.Scale.X = 16
	player2.Transform.Scale.Y = 64
	player2.Transform.Position.X += 300
	targetScene.AddObject(player2)

	
	
}
