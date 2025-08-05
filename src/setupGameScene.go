package src

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	game "github.com/Massil-br/GolangPong/src/Game"
)

func SetupGameScene(targetScene *engine.Scene) {
	game.SetupPlayer(targetScene)
	game.SetupCamera(targetScene)
}


