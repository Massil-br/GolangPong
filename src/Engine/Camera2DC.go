package engine

import (
	"math"
	time "github.com/Massil-br/GolangPong/src/Engine/Time"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Camera2DC struct {
	Parent           *GameObject
	Camera           rl.Camera2D
	ViewWidth        float32  // Taille de vue virtuelle (comme 640 dans votre SFML)
	ViewHeight       float32  // Taille de vue virtuelle (comme 360 dans votre SFML)
	FollowTarget     bool     // Si la caméra suit automatiquement le parent
	SmoothFollow     bool     // Suivi lissé
	SmoothSpeed      float32  // Vitesse du suivi lissé
}

func NewCamera2DC(viewWidth, viewHeight float32) *Camera2DC {
	// Sauvegarder les dimensions virtuelles globalement (comme dans votre SFML)
	time.Data.VirtualWidth = viewWidth
	time.Data.VirtualHeight = viewHeight
	
	return &Camera2DC{
		Camera: rl.Camera2D{
			Offset:   rl.NewVector2(0, 0), // Sera calculé dans Update()
			Target:   rl.NewVector2(0, 0), // Position du joueur
			Rotation: 0,
			Zoom:     1.0, // Sera calculé dans Update()
		},
		ViewWidth:    viewWidth,
		ViewHeight:   viewHeight,
		FollowTarget: true,
		SmoothFollow: false,
		SmoothSpeed:  5.0,
	}
}

func (c *Camera2DC) Start() {}

func (c *Camera2DC) Update() {
	if c.Parent == nil {
		return
	}

	// Calculer le zoom pour maintenir la taille de vue virtuelle
	// (équivalent à view.Size dans SFML)
	scaleX := c.Parent.FrameData.Width / c.ViewWidth
	scaleY := c.Parent.FrameData.Height / c.ViewHeight
	
	// Utiliser le plus petit scale pour garder les proportions
	// (comme votre méthode ResizeCamera qui garde le ratio)
	scale := float32(math.Min(float64(scaleX), float64(scaleY)))
	c.Camera.Zoom = scale

	// Centrer la caméra sur l'écran (équivalent à l'offset dans SFML)
	c.Camera.Offset = rl.NewVector2(
		c.Parent.FrameData.Width/2,
		c.Parent.FrameData.Height/2,
	)

	// Suivre la cible si activé (équivalent à view.Center = targetPosition)
	if c.FollowTarget {
		if c.SmoothFollow {
			// Suivi lissé
			deltaTime := rl.GetFrameTime()
			currentTarget := c.Camera.Target
			targetPos := c.Parent.Transform.Position
			
			lerpedX := currentTarget.X + (targetPos.X-currentTarget.X)*c.SmoothSpeed*deltaTime
			lerpedY := currentTarget.Y + (targetPos.Y-currentTarget.Y)*c.SmoothSpeed*deltaTime
			
			c.Camera.Target = rl.NewVector2(lerpedX, lerpedY)
		} else {
			// Suivi direct (comme dans votre SFML)
			c.Camera.Target = c.Parent.Transform.Position
		}
	}
}

func (c *Camera2DC) Draw() {}

func (c *Camera2DC) SetParent(p *GameObject) {
	c.Parent = p
}

// Méthodes équivalentes à votre classe Camera SFML

// Équivalent à camera.Update(player.GetPosition()) dans votre GameLoop
func (c *Camera2DC) UpdateTarget(targetPosition rl.Vector2) {
	if c.SmoothFollow {
		deltaTime := rl.GetFrameTime()
		currentTarget := c.Camera.Target
		
		lerpedX := currentTarget.X + (targetPosition.X-currentTarget.X)*c.SmoothSpeed*deltaTime
		lerpedY := currentTarget.Y + (targetPosition.Y-currentTarget.Y)*c.SmoothSpeed*deltaTime
		
		c.Camera.Target = rl.NewVector2(lerpedX, lerpedY)
	} else {
		c.Camera.Target = targetPosition
	}
}

// Équivalent à camera.Resize(640, 360)
func (c *Camera2DC) Resize(viewWidth, viewHeight float32) {
	c.ViewWidth = viewWidth
	c.ViewHeight = viewHeight
	time.Data.VirtualWidth = viewWidth
	time.Data.VirtualHeight = viewHeight
}

// Getters équivalents à votre SFML
func (c *Camera2DC) GetViewWidth() float32 {
	return c.ViewWidth
}

func (c *Camera2DC) GetViewHeight() float32 {
	return c.ViewHeight
}

func (c *Camera2DC) GetCameraPosition() rl.Vector2 {
	return c.Camera.Target
}

// Méthodes pour commencer/finir le mode caméra (équivalent à window.SetView(camera.GetView()))
func (c *Camera2DC) BeginMode() {
	rl.BeginMode2D(c.Camera)
}

func (c *Camera2DC) EndMode() {
	rl.EndMode2D()
}

// Configuration du comportement de suivi
func (c *Camera2DC) SetFollowTarget(follow bool) {
	c.FollowTarget = follow
}

func (c *Camera2DC) SetSmoothFollow(smooth bool, speed float32) {
	c.SmoothFollow = smooth
	c.SmoothSpeed = speed
}

// Méthodes utilitaires pour la conversion de coordonnées
// (équivalent à window.MapPixelToCoords dans SFML)
func (c *Camera2DC) ScreenToWorld(screenPos rl.Vector2) rl.Vector2 {
	return rl.GetScreenToWorld2D(screenPos, c.Camera)
}

func (c *Camera2DC) WorldToScreen(worldPos rl.Vector2) rl.Vector2 {
	return rl.GetWorldToScreen2D(worldPos, c.Camera)
}

// Exemple d'utilisation pour reproduire votre GameLoop SFML
/*
Utilisation dans votre GameLoop Go :

func (g *GameLoop) Update() {
	// Équivalent à camera.Update(player.GetPosition())
	g.camera.UpdateTarget(g.player.GetPosition())
	
	// Mettre à jour la caméra
	g.camera.Update()
}

func (g *GameLoop) Render() {
	rl.BeginDrawing()
	rl.ClearBackground(rl.Black)
	
	// Équivalent à window.SetView(camera.GetView())
	g.camera.BeginMode()
	
	// Dessiner le monde (map, joueurs, etc.)
	g.map.Draw()
	g.player.Render()
	
	// Fin du mode caméra
	g.camera.EndMode()
	
	// Dessiner l'UI (sans caméra, comme dans votre SFML)
	g.DrawUI()
	
	rl.EndDrawing()
}

// Configuration initiale (équivalent à ResizeCamera)
func InitCamera() *Camera2DC {
	camera := NewCamera2DC(640, 360) // Même taille que votre SFML
	camera.SetSmoothFollow(true, 5.0) // Optionnel : suivi lissé
	return camera
}
*/