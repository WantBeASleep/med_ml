-- +goose Down
-- +goose StatementBegin

DROP TABLE IF EXISTS chat.chat_participants;
DROP TABLE IF EXISTS chat.chats;

DROP SCHEMA IF EXISTS chat;

-- +goose StatementEnd 