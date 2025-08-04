package time
import rl "github.com/gen2brain/raylib-go/raylib"


var DeltaTime float32

func Update() {
	DeltaTime = rl.GetFrameTime()
}
