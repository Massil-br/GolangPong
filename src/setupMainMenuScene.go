package src

import (
	engine "github.com/Massil-br/GolangPong/src/Engine"
	mainmenu "github.com/Massil-br/GolangPong/src/MainMenu"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func SetupMainMenu(targetScene *engine.Scene, font *rl.Font) {

	mainmenu.SetupTitle(targetScene, font)
	mainmenu.SetupPlayButton(targetScene, font)
}
