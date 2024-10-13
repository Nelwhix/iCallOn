CREATE TABLE IF NOT EXISTS games (
     id CHAR(26) PRIMARY KEY,
     user_id CHAR(26) NOT NULL,
     created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
     updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE INDEX idx_games_user_id ON games (user_id);