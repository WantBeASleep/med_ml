-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user"
(
    id          uuid         PRIMARY KEY,
    email       varchar(255) NOT NULL UNIQUE,
    "password"  varchar(255),
    refresh_token       text,
    "role"      varchar(255) NOT NULL
);

COMMENT ON TABLE "user" IS 'Таблица пользователей';
COMMENT ON COLUMN "user".refresh_token IS 'Refresh токен пользователя';
COMMENT ON COLUMN "user".role IS 'Роль пользователя';

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user" CASCADE;
-- +goose StatementEnd