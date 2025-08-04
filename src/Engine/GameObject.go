package engine

type Component interface{
	Start()
	Update()
	Draw()
}

type GameObject struct{
	Name string
	Active bool
	Components []Component
	DeltaTime *float32
}

func(g *GameObject) AddComponent(c Component){
	g.Components = append(g.Components, c)
	c.Start()
}

func (g *GameObject) Update(){
	if !g.Active{
		return
	}
	for _,comp := range g.Components{
		comp.Update()
	}
}

func (g *GameObject) Draw(){
	if !g.Active{
		return
	}
	for _,comp := range g.Components{
		comp.Draw()
	}
}