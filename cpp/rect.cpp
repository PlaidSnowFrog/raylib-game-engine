#include "raylib.h"

class Rect {
public:
	int x;
	int y;
	int w;
	int h;

	unsigned int speed;

	Color tint;
	Rectangle src;
	Texture2D texture;

	Rect(const int x, const int y, const int w, const int h, const unsigned int speed, const Color tint, const char* texturePath) {
		this->x = x;
		this->y = y;
		this->w = w;
		this->h = h;
		this->speed = speed;

		this->src.x = x;
		this->src.y = y;
		this->src.width = w;
		this->src.height = h;

		this->tint = tint;
		this->texture = LoadTexture(texturePath);
	}

	void Draw() {
		DrawRectangle(this->x, this->y, this->w, this->h, this->tint);
	}

	void DrawTextured() {
    Vector2 position = { static_cast<float>(this->x), static_cast<float>(this->y) };
    DrawTextureRec(this->texture, this->src, position, WHITE);
	}

	void WasdMove(const int screenWidth, const int screenHeight) {
		if (IsKeyDown(KEY_W) && this->y > 0) {
			this->y -=  this->speed;
		}
		if (IsKeyDown(KEY_S) && this->y < (screenHeight - this->h)) {
			this->y += this->speed;
		}
		if (IsKeyDown(KEY_A) && this->x > 0) {
			this->x -= this->speed;
		}
		if (IsKeyDown(KEY_D) && this->x < (screenWidth - this->w)) {
			this->x += this->speed;
		}
	}

	void UpdateSrc() {
		this->src.x = this->x;
		this->src.y = this->y;
		this->src.width = this->w;
		this->src.height = this->h;
	}
};