package main

import (
	"math"

	rl "github.com/gen2brain/raylib-go/raylib"
)

type Rect struct {
	x     int32
	y     int32
	w     int32
	h     int32
	speed uint8

	src     rl.Rectangle
	texture rl.Texture2D
	tint    rl.Color
}

type gameState struct {
	isGameOver   bool
	hasPlayerWon bool
	isGamePaused bool
}

func main() {
	const screenWidth int32 = 900
	const screenHeight int32 = 500

	rl.InitWindow(screenWidth, screenHeight, "Game Engine")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)

	rect := newRect(0, 0, 50, 50, 10, rl.Blue)
	gs := newGameState(false, false, false)

	for !rl.WindowShouldClose() {
		if !gs.isGameOver {
			if !gs.isGamePaused { // The game is afoot!
				wasdMoveRect(&rect, screenWidth, screenHeight)

				if rl.IsKeyPressed(rl.KeyF) {
					gs.isGamePaused = true
				}
				if rl.IsKeyPressed(rl.KeyR) {
					gs.isGameOver = true
				}
			} else { // Game paused
				if rl.IsKeyPressed(rl.KeyF) {
					gs.isGamePaused = false
				}
			}
		} else if gs.hasPlayerWon { // Player won
			if rl.IsKeyPressed(rl.KeyF) {
				gs.hasPlayerWon = false
			}
		} else { // Game lost
			if rl.IsKeyPressed(rl.KeyR) {
				gs.isGameOver = false
			}
			if rl.IsKeyPressed(rl.KeyF) {
				gs.hasPlayerWon = true
			}
		}

		rl.BeginDrawing()

		rl.ClearBackground(rl.RayWhite)
		if !gs.isGameOver {
			if !gs.isGamePaused { // The game is afoot!
				drawRect(&rect)
			} else { // Game paused
				rl.DrawText("Game Paused", screenWidth/2, screenHeight/2, 50, rl.Black)
			}
		} else if gs.hasPlayerWon { // Player won
			rl.DrawText("Player Won", screenWidth/2, screenHeight/2, 50, rl.Black)

		} else { // Player lost lost
			rl.DrawText("Player Lost", screenWidth/2, screenHeight/2, 50, rl.Black)
		}

		rl.EndDrawing()
	}
}

func newRectTextured(x int32, y int32, w int32, h int32, speed uint8, tint rl.Color, texturePath string) Rect {
	var rect Rect
	rect.x = x
	rect.y = y
	rect.w = w
	rect.h = h
	rect.speed = speed

	rect.src.X = float32(x)
	rect.src.Y = float32(y)
	rect.src.Width = float32(w)
	rect.src.Height = float32(h)

	rect.tint = tint
	rect.texture = rl.LoadTexture(texturePath)

	return rect
}

func newRect(x int32, y int32, w int32, h int32, speed uint8, tint rl.Color) Rect {
	var rect Rect
	rect.x = x
	rect.y = y
	rect.w = w
	rect.h = h
	rect.speed = speed

	rect.src.X = float32(x)
	rect.src.Y = float32(y)
	rect.src.Width = float32(w)
	rect.src.Height = float32(h)

	rect.tint = tint

	return rect
}

func drawRect(rect *Rect) {
	rl.DrawRectangle(rect.x, rect.y, rect.w, rect.h, rect.tint)
}

func drawTexturedRect(rect *Rect) {
	rl.DrawTextureRec(rect.texture, rect.src, rl.NewVector2(float32(rect.x), float32(rect.y)), rl.White)
}

func wasdMoveRect(rect *Rect, screenWidth int32, screenHeight int32) {
	switch {
	case rl.IsKeyDown(rl.KeyW) && rect.y > 0:
		rect.y -= int32(rect.speed)
	case rl.IsKeyDown(rl.KeyS) && rect.y < (screenHeight-rect.h):
		rect.y += int32(rect.speed)
	case rl.IsKeyDown(rl.KeyA) && rect.x > 0:
		rect.x -= int32(rect.speed)
	case rl.IsKeyDown(rl.KeyD) && rect.x < (screenWidth-rect.w):
		rect.x += int32(rect.speed)
	}
}

func checkRectCollision(rect1, rect2 Rect) bool {
	if rect1.x+rect1.w < rect2.x {
		return false
	}

	if rect1.x > rect2.x+rect2.w {
		return false
	}

	if rect1.y+rect1.h < rect2.y {
		return false
	}

	if rect1.y > rect2.y+rect2.h {
		return false
	}

	return true
}

func areNear(rect1, rect2 Rect, distanceThreshold float64) bool {
	center1X := float64(rect1.x) + float64(rect1.w/2)
	center1Y := float64(rect1.y) + float64(rect1.h/2)
	center2X := float64(rect2.x) + float64(rect2.w/2)
	center2Y := float64(rect2.y) + float64(rect2.h/2)

	distance := math.Sqrt(math.Pow(center2X-center1X, 2) + math.Pow(center2Y-center1Y, 2))

	return distance <= distanceThreshold
}

func newGameState(isGameOver, hasPlayerWon, isGamePaused bool) gameState {
	var gs gameState
	gs.isGameOver = isGameOver
	gs.hasPlayerWon = hasPlayerWon
	gs.isGamePaused = isGamePaused

	return gs
}
