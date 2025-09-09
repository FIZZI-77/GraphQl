-- +goose Up
-- +goose StatementBegin

CREATE TABLE tasks(
    id SERIAL PRIMARY KEY,
    user_id UUID NOT NULL,
    title TEXT NOT NULL ,
    description TEXT,
    completed BOOLEAN DEFAULT FALSE,
    created_at TIMESTAMP DEFAULT now(),
    update_at TIMESTAMP DEFAULT now()
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE tasks CASCADE;
-- +goose StatementEnd