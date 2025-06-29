-- +goose Up
-- +goose StatementBegin
SELECT 'up SQL query';
-- +goose StatementEnd

ALTER TABLE "user" DROP COLUMN refresh_token CASCADE;

CREATE TABLE refresh_token
(
    id              uuid,
    refresh_token   text,
    PRIMARY KEY (id, refresh_token)
);

COMMENT ON TABLE refresh_token IS 'Таблица refresh токенов для мульти авторизации';
COMMENT ON COLUMN refresh_token.id IS 'ID пользователя';
COMMENT ON COLUMN refresh_token.refresh_token IS 'Refresh токен пользователя';

-- +goose Down
-- +goose StatementBegin
SELECT 'down SQL query';
-- +goose StatementEnd

DROP TABLE refresh_token CASCADE;

ALTER TABLE "user" ADD COLUMN refresh_token text;
