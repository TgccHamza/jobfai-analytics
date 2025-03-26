package services

import (
	"errors"
	"fmt"

	"jobfai-analytics/internal/models"
	"jobfai-analytics/internal/repositories"
)

// GameService handles game-related business logic
type GameService struct {
	gameRepository       *repositories.GameRepository
	stageRepository      *repositories.StageRepository
	competenceRepository *repositories.CompetenceRepository
}

// NewGameService creates a new game service
func NewGameService(
	gameRepository *repositories.GameRepository,
	stageRepository *repositories.StageRepository,
	competenceRepository *repositories.CompetenceRepository,
) *GameService {
	return &GameService{
		gameRepository:       gameRepository,
		stageRepository:      stageRepository,
		competenceRepository: competenceRepository,
	}
}

// GetGameByID retrieves a game by its ID
func (s *GameService) GetGameByID(gameID string) (*models.Game, error) {
	game, err := s.gameRepository.FindByID(gameID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving game: %w", err)
	}

	if game == nil {
		return nil, errors.New("game not found")
	}

	return game, nil
}

// GetGameWithFullConfiguration retrieves a game with all its configuration
func (s *GameService) GetGameWithFullConfiguration(gameID string) (*models.Game, error) {
	game, err := s.gameRepository.FindWithFullConfiguration(gameID)
	if err != nil {
		return nil, fmt.Errorf("error retrieving game configuration: %w", err)
	}

	if game == nil {
		return nil, errors.New("game not found")
	}

	return game, nil
}

// CreateGame creates a new game
func (s *GameService) CreateGame(game *models.Game) error {
	if game.GameID == "" || game.GameName == "" {
		return errors.New("game ID and name are required")
	}

	return s.gameRepository.Create(game)
}

// UpdateGame updates an existing game
func (s *GameService) UpdateGame(game *models.Game) error {
	if game.GameID == "" {
		return errors.New("game ID is required")
	}

	existingGame, err := s.gameRepository.FindByID(game.GameID)
	if err != nil {
		return fmt.Errorf("error checking existing game: %w", err)
	}

	if existingGame == nil {
		return errors.New("game not found")
	}

	return s.gameRepository.Update(game)
}

// DeleteGame deletes a game
func (s *GameService) DeleteGame(gameID string) error {
	existingGame, err := s.gameRepository.FindByID(gameID)
	if err != nil {
		return fmt.Errorf("error checking existing game: %w", err)
	}

	if existingGame == nil {
		return errors.New("game not found")
	}

	return s.gameRepository.Delete(gameID)
}

// GetAllGames retrieves all games
func (s *GameService) GetAllGames() ([]models.Game, error) {
	games, err := s.gameRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("error retrieving games: %w", err)
	}

	return games, nil
}

// GetGameStages retrieves all stages for a game
func (s *GameService) GetGameStages(gameID string) ([]models.Stage, error) {
	return s.stageRepository.FindByGame(gameID)
}

// GetGameCompetencies retrieves all competencies for a game
func (s *GameService) GetGameCompetencies(gameID string) ([]models.Competence, error) {
	return s.competenceRepository.FindByGame(gameID)
}
