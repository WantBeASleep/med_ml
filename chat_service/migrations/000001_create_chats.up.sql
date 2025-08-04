-- +goose Up
-- +goose StatementBegin

CREATE SCHEMA IF NOT EXISTS chat;

CREATE TABLE chat.chats (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    name VARCHAR(255) NOT NULL,
    description TEXT,
    patient_id UUID NOT NULL,
    last_activity TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

CREATE TABLE chat.chat_participants (
    chat_id UUID NOT NULL REFERENCES chat.chats(id) ON DELETE CASCADE,
    doctor_id UUID NOT NULL,
    PRIMARY KEY (chat_id, doctor_id)
);

-- +goose StatementEnd 