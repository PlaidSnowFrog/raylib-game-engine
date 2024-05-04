#ifndef SETTINGS_HPP
#define SETTINGS_HPP

struct GameState {
    bool isGameOver;
    bool hasPlayerWon;
    bool isGamePaused;
};

GameState newState(const bool isGameOver, const bool hasPlayerWon, const bool isGamePaused);

#endif // SETTINGS_HPP