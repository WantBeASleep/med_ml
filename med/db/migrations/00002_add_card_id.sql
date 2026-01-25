-- +goose Up
-- +goose StatementBegin
-- Добавляем поле id в таблицу card
-- Сначала удаляем составной первичный ключ
ALTER TABLE card
    DROP CONSTRAINT pk_card;

-- Добавляем поле id как serial (автоинкремент)
ALTER TABLE card
    ADD COLUMN id SERIAL;

-- Делаем id первичным ключом
ALTER TABLE card
    ADD CONSTRAINT pk_card PRIMARY KEY (id);

-- Добавляем уникальный индекс на (doctor_id, patient_id) для предотвращения дубликатов
CREATE UNIQUE INDEX idx_card_doctor_patient ON card(doctor_id, patient_id);

COMMENT ON COLUMN card.id IS 'Идентификатор карты пациента';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Удаляем индекс
DROP INDEX IF EXISTS idx_card_doctor_patient;

-- Удаляем поле id
ALTER TABLE card
    DROP CONSTRAINT pk_card;

ALTER TABLE card
    DROP COLUMN id;

-- Восстанавливаем составной первичный ключ
ALTER TABLE card
    ADD CONSTRAINT pk_card PRIMARY KEY (doctor_id, patient_id);
-- +goose StatementEnd
