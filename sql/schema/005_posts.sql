-- +goose Up
CREATE TABLE posts (
    id UUID PRIMARY KEY,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    title VARCHAR(100) NOT NULL,
    url VARCHAR(100) NOT NULL UNIQUE,
    description VARCHAR(255),
    published_at TIMESTAMP,
    feed_id UUID NOT NULL,
    CONSTRAINT feeds_id
    FOREIGN KEY (feed_id)
    REFERENCES feeds(id) ON DELETE CASCADE
);

-- +goose Down
DROP TABLE posts;

