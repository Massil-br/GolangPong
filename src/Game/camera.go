package game

import engine "github.com/Massil-br/GolangPong/src/Engine"

func SetupCamera(targetScene *engine.Scene) {
	camera := engine.NewObj("MainCamera")
	camera.AddComponent(engine.NewCamera2DC(640,360))
	
	targetScene.AddObject(camera)
	targetScene.SetCamera(camera)
}
