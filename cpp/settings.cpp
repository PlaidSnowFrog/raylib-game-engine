struct GameState {
	bool isGameOver;
	bool hasPlayerWon;
	bool isGamePaused;
};

GameState newState(const bool isGameOver, const bool hasPlayerWon, const bool isGamePaused) {
	GameState gs;
	gs.isGameOver = isGameOver;
	gs.isGamePaused = isGamePaused;
	gs.hasPlayerWon = hasPlayerWon;

	return gs;
}