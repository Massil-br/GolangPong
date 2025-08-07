package game

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func SetupWalls(targetScene *engine.Scene) {
	topWall := engine.NewObj("topWall")
	topWall.AddComponent(engine.NewCollider(false))
	topWall.AddComponent(engine.NewRenderer(rl.White))
	topWall.Transform.Position.Y = -180
	topWall.Transform.Position.X -= 340
	topWall.Transform.Scale.Y += 20
	topWall.Transform.Scale.X += 680
	targetScene.AddObject(topWall)

	botWall := engine.NewObj("topWall")
	botWall.AddComponent(engine.NewCollider(false))
	botWall.AddComponent(engine.NewRenderer(rl.White))
	botWall.AddComponent(engine.NewCollider(false))
	botWall.Transform.Position.Y = 160
	botWall.Transform.Position.X -= 340
	botWall.Transform.Scale.Y += 20
	botWall.Transform.Scale.X += 680
	targetScene.AddObject(botWall)
}
