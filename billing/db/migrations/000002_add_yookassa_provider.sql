-- +goose Up
-- +goose StatementBegin

-- Установка расширения uuid-ossp, если оно еще не установлено
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

-- Вставка данных в таблицу payment_provider
INSERT INTO payment_provider (id, name, is_active) VALUES
    (uuid_generate_v4(), 'Yookassa', TRUE);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

-- Удаление данных из таблицы payment_provider
DELETE FROM payment_provider WHERE name = 'Yookassa';

-- +goose StatementEnd