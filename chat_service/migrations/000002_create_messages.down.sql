-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS chat.messages;

-- +goose StatementEnd 