package engine

type Scene struct{
	GameObjects []*GameObject
	DeltaTime *float32
}

func (s *Scene) AddObject(o *GameObject){
	o.DeltaTime = s.DeltaTime
	s.GameObjects = append(s.GameObjects, o)
}

func (s *Scene) Update(){
	for _,o := range s.GameObjects{
		o.Update()
	}
}

func (s *Scene) Draw(){
	for _,o := range s.GameObjects{
		o.Draw()
	}
}