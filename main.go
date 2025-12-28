package main

import (
	"fmt"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const WIDTH int32 = 800
const HEIGHT int32 = 600
const FPS int32 = 60

func drawBackground() {
	rl.ClearBackground(rl.RayWhite)
}

func main() {
	rl.InitWindow(WIDTH, HEIGHT, "CHIP8 Emulator")
	defer rl.CloseWindow()

	rl.SetTargetFPS(FPS)

	fmt.Println("Emulator starts")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		drawBackground()
		rl.DrawText("Congrats! You created your first window!", 190, 200, 20, rl.LightGray)
		rl.EndDrawing()
	}

}
