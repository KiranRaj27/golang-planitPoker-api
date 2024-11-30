-- +goose Up
CREATE TABLE cards (
    id SERIAL PRIMARY KEY,
    value VARCHAR(255) NOT NULL,
    set_type VARCHAR(50) NOT NULL, -- To identify the type of the card set (Scrum, Fibonacci, etc.)
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);


-- +goose Down
DROP TABLE cards;
