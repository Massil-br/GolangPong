package main

import (
	"github.com/Massil-br/GolangPong/src"
	rl "github.com/gen2brain/raylib-go/raylib"
)


const InitialWidth int32 = 1280
const InitialHeight int32 = 720

func main() {
	rl.SetConfigFlags(rl.FlagWindowResizable)
	rl.InitWindow(InitialWidth, InitialHeight, "Massil's Pong")
	src.AppLoop()
	defer rl.CloseWindow()
	
	
}
