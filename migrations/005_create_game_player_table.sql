CREATE TABLE IF NOT EXISTS game_player (
   game_id CHAR(26) NOT NULL,
   player_id CHAR(26) NOT NULL,
   created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
   PRIMARY KEY (game_id, player_id)
);

