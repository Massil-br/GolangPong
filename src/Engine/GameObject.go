package engine

import (
	"reflect"

	time "github.com/Massil-br/GolangPong/src/Engine/Time"
)

type Component interface {
	Start()
	Update()
	Draw()
	SetParent(*GameObject)
}

type GameObject struct {
	Name         string
	Active       bool
	Components   []Component
	componentMap map[reflect.Type]Component
	FrameData    *time.FrameData
	Transform    Transform2D
}

func (g *GameObject) AddComponent(c Component) {
	if g.componentMap == nil {
		g.componentMap = make(map[reflect.Type]Component)
	}
	g.Components = append(g.Components, c)
	g.componentMap[reflect.TypeOf(c)] = c
	c.SetParent(g)
	c.Start()
}

func GetComponent[T Component](g *GameObject) T {
	var zero T
	for _, comp := range g.Components {
		if c, ok := comp.(T); ok {
			return c
		}
	}
	return zero
}

func (g *GameObject) Update() {
	if !g.Active {
		return
	}
	for _, comp := range g.Components {
		comp.Update()
	}
}

func (g *GameObject) Draw() {
	if !g.Active {
		return
	}
	for _, comp := range g.Components {
		comp.Draw()
	}
}
