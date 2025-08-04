package engine

import time "github.com/Massil-br/GolangPong/src/Engine/Time"

type Scene struct {
	GameObjects []*GameObject
	FrameData   *time.FrameData
}

func (s *Scene) AddObject(o *GameObject) {
	o.FrameData = s.FrameData
	s.GameObjects = append(s.GameObjects, o)
}

func (s *Scene) Update() {
	for _, o := range s.GameObjects {
		o.Update()
	}
}

func (s *Scene) Draw() {
	for _, o := range s.GameObjects {
		o.Draw()
	}
}