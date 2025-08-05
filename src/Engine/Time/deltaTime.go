package time
import rl "github.com/gen2brain/raylib-go/raylib"

type FrameData struct{
	DeltaTime float32
	Width float32
	Height float32
	VirtualWidth float32
	VirtualHeight float32
}
var Data FrameData

func Update() {
	Data.DeltaTime = rl.GetFrameTime()
	Data.Width = float32(rl.GetRenderWidth())
	Data.Height= float32(rl.GetRenderHeight())
	

}
