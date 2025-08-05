package engine

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Rigidbody2D struct {
	Parent     *GameObject
	Velocity   rl.Vector2
	Mass       float32
	Damping    float32 // 0 = pas de ralentissement, 1 = stop immédiat
	UseGravity bool
	Gravity    rl.Vector2
}

func NewRigidBody2D(mass float32, damping float32, useGravity bool)*Rigidbody2D{
	return &Rigidbody2D{
		Velocity: rl.NewVector2(0,0),
		Mass: mass,
		Damping: damping,
		UseGravity: useGravity,
		Gravity: rl.NewVector2(0,0),
	}
}

func (rb *Rigidbody2D) Start() {
	if rb.Mass == 0 {
		rb.Mass = 1
	}
	if rb.Damping < 0 {
		rb.Damping = 0
	}
	if rb.Damping > 1 {
		rb.Damping = 1
	}
	if rb.UseGravity && rb.Gravity.X == 0 && rb.Gravity.Y == 0 {
		rb.Gravity = rl.NewVector2(0, 9.81) // gravité par défaut
	}
}

func (rb *Rigidbody2D) Update() {
	if rb.Parent == nil || rb.Parent.FrameData == nil {
		return
	}

	dt := rb.Parent.FrameData.DeltaTime

	// Appliquer la gravité si activée
	if rb.UseGravity {
		rb.Velocity.X += rb.Gravity.X * dt
		rb.Velocity.Y += rb.Gravity.Y * dt
	}

	// Mise à jour position
	rb.Parent.Transform.Position.X += rb.Velocity.X * dt
	rb.Parent.Transform.Position.Y += rb.Velocity.Y * dt

	// Appliquer le damping
	rb.Velocity.X *= (1 - rb.Damping)
	rb.Velocity.Y *= (1 - rb.Damping)
}

func (rb *Rigidbody2D) Draw() {
	// Rien à dessiner, c'est juste de la physique
}

func (rb *Rigidbody2D) SetParent(o *GameObject) {
	rb.Parent = o
}

// Appliquer une force
func (rb *Rigidbody2D) AddForce(force rl.Vector2) {
	rb.Velocity.X += force.X / rb.Mass
	rb.Velocity.Y += force.Y / rb.Mass
}
