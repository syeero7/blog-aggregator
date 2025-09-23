-- +goose Up
CREATE TABLE feed_follows (
  id INTEGER PRIMARY KEY GENERATED ALWAYS AS IDENTITY,
  feed_id INTEGER NOT NULL,
  user_id uuid NOT NULL,
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  UNIQUE(feed_id,user_id),
  FOREIGN KEY (feed_id) REFERENCES feeds(id) ON DELETE CASCADE,
  FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feed_follows;
