#include "raylib.h"
#include "rect.cpp"
#include "settings.hpp"

int main(void) {
  // Initialization
  const int screenWidth = 900;
  const int screenHeight = 500;

  InitWindow(screenWidth, screenHeight, "Game Engine");
  InitAudioDevice();

  SetTargetFPS(60);

  // Variables definition
  Rect rect(0, 0, 50, 50, 10, BLUE, "assets/car.png");
  GameState gameState = newState(false, false, false);

  Music song = LoadMusicStream("assets/song.wav");
  PlayMusicStream(song);

  // Main game loop
  while (!WindowShouldClose()) // Detect window close button or ESC key
  {
    UpdateMusicStream(song);

    BeginDrawing();

    ClearBackground(RAYWHITE);

    if (!gameState.isGameOver) {
      if (!gameState.isGamePaused) { // The game is afoot!
        rect.WasdMove(screenWidth, screenHeight);
        rect.DrawTextured();

        if (IsKeyPressed(KEY_F)) {
          gameState.isGamePaused = true;
        }
        if (IsKeyPressed(KEY_R)) {
          gameState.isGameOver = true;
        }
      } else { // Game paused
        if (IsKeyPressed(KEY_F)) {
          gameState.isGamePaused = false;
        }
        DrawText("game paused", screenWidth / 2, screenHeight / 2, 50, BLACK);
      }
    } else if (gameState.hasPlayerWon) { // Player won
      if (IsKeyPressed(KEY_R)) {
        gameState.isGameOver = false;
      }
      DrawText("player won", screenWidth / 2, screenHeight / 2, 50, BLACK);
    } else { // Player lost
      if (IsKeyPressed(KEY_R)) {
        gameState.isGameOver = false;
      }
      if (IsKeyPressed(KEY_F)) {
        gameState.hasPlayerWon = true;
      }
      DrawText("player lost", screenWidth / 2, screenHeight / 2, 50, BLACK);
    }

    EndDrawing();
  }

  // De-Initialization
  UnloadMusicStream(song);
  CloseAudioDevice();
  CloseWindow();

  return 0;
}