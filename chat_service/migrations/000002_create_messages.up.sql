-- +goose Up
-- +goose StatementBegin

CREATE TABLE chat.messages (
    id UUID PRIMARY KEY,
    chat_id UUID NOT NULL REFERENCES chat.chats(id) ON DELETE CASCADE,
    sender_id UUID NOT NULL,
    content TEXT NOT NULL,
    type VARCHAR(3) NOT NULL CHECK (type IN ('in', 'out')),
    content_type VARCHAR(16) NOT NULL CHECK (content_type IN ('chat_message', 'notification', 'system_message')),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- +goose StatementEnd 