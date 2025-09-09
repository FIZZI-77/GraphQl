-- +goose Up
-- +goose StatementBegin

CREATE  TABLE users(
    id UUID PRIMARY KEY ,
    name varchar(50) NOT NULL ,
    email varchar(500) NOT NULL ,
    created_at timestamp DEFAULT now(),
    update_at timestamp DEFAULT now()
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE users CASCADE;
-- +goose StatementEnd