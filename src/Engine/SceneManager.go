package engine

type SceneManager struct{
	Scenes []*Scene
	CurrentIdx int
}

func (sm *SceneManager) AddScene(s *Scene){
	sm.Scenes = append(sm.Scenes, s)
}

func(sm *SceneManager) SetScene(idx int){
	if idx >= 0 && idx < len(sm.Scenes){
		sm.CurrentIdx = idx
	}
}

func(sm *SceneManager) Update(){
	if len(sm.Scenes)==0{
		return
	}
	sm.Scenes[sm.CurrentIdx].Update()
}

func (sm *SceneManager) Draw(){
	if len(sm.Scenes)==0{
		return
	}
	sm.Scenes[sm.CurrentIdx].Draw()
}