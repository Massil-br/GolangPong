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
	player.Transform.Scale.X = 32
	player.Transform.Scale.Y = 128
	player.Transform.Position.X -= 200
	targetScene.AddObject(player)

	player2 := engine.NewObj("player2")
	player2.AddComponent(engine.NewCollider(false))
	player2.AddComponent(engine.NewRenderer(rl.White))
	player2.AddComponent(engine.NewRigidBody2D(0, 1, false))
	player2.AddComponent(playerscripts.NewPlayerController(400))
	player2.Transform.Scale.X = 32
	player2.Transform.Scale.Y = 128
	player2.Transform.Position.X += 200
	targetScene.AddObject(player2)

	player3 := engine.NewObj("player2")
	player3.AddComponent(engine.NewCollider(false))
	player3.AddComponent(engine.NewRenderer(rl.White))
	player3.AddComponent(engine.NewRigidBody2D(0, 1, false))
	player3.AddComponent(playerscripts.NewPlayerController(400))
	player3.Transform.Scale.X = 32
	player3.Transform.Scale.Y = 128
	player3.Transform.Position.X -= 300
	targetScene.AddObject(player3)

	
}
