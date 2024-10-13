package models

import (
	"context"
	"github.com/Nelwhix/iCallOn/pkg/requests"
	"github.com/oklog/ulid/v2"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type User struct {
	ID       string `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (m *Model) GetUserByEmail(ctx context.Context, email string) (User, error) {
	var user User
	row := m.Conn.QueryRow(ctx, "select id, username, email, password FROM users WHERE email = $1", email)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (m *Model) GetUserByToken(ctx context.Context, token string) (User, error) {
	cToken, err := m.FindToken(ctx, token)
	if err != nil {
		return User{}, err
	}

	lastUsedAt := time.Now()
	cToken.LastUsedAt = &lastUsedAt
	err = m.UpdateToken(ctx, cToken)
	if err != nil {
		return User{}, err
	}

	user, err := m.GetUserById(ctx, cToken.UserID)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (m *Model) GetUserById(ctx context.Context, userID string) (User, error) {
	var user User
	row := m.Conn.QueryRow(ctx, "select id, username, email, password FROM users WHERE id = $1", userID)
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.Password)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (m *Model) InsertIntoUsers(ctx context.Context, request requests.SignUp) (User, error) {
	userID := ulid.Make().String()
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(request.Password), 12)
	if err != nil {
		return User{}, err
	}

	sql := "insert into users(id, username, email, password) values ($1, $2, $3, $4)"
	_, err = m.Conn.Exec(ctx, sql, userID, request.Username, request.Email, string(passwordHash))

	if err != nil {
		return User{}, err
	}

	user, err := m.GetUserById(ctx, userID)
	if err != nil {
		return User{}, err
	}

	return user, nil
}

func (m *Model) UpdateUser(ctx context.Context, user User) error {
	sql := "update users set username = $1, updated_at = $5 where id = $6"
	_, err := m.Conn.Exec(ctx, sql, user.Username, time.Now(), user.ID)
	if err != nil {
		return err
	}

	return nil
}
