package models

import (
	"context"
	"github.com/Nelwhix/iCallOn/pkg/requests"
	"github.com/oklog/ulid/v2"
	"time"
)

type Game struct {
	ID          string
	UserID      string
	RoundLength int
	CreatedAt   time.Time
}

func (m *Model) InsertIntoGames(ctx context.Context, request requests.NewGame) (Game, error) {
	gameID := ulid.Make().String()

	sql := "insert into games(id, user_id, round_length) values ($1, $2, $3)"
	_, err := m.Conn.Exec(ctx, sql, gameID, request.UserID, request.RoundLength)

	if err != nil {
		return Game{}, err
	}

	game, err := m.GetGameById(ctx, gameID)
	if err != nil {
		return Game{}, err
	}

	return game, nil
}

func (m *Model) GetGameById(ctx context.Context, gameID string) (Game, error) {
	var game Game
	row := m.Conn.QueryRow(ctx, "select id, user_id, round_length, created_at FROM games WHERE id = $1", gameID)
	err := row.Scan(&game.ID, &game.UserID, &game.RoundLength, &game.CreatedAt)
	if err != nil {
		return Game{}, err
	}

	return game, nil
}
