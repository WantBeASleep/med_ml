-- +goose Up
-- +goose StatementBegin

-- Вставка данных в таблицу tariff_plan
INSERT INTO tariff_plan (id, name, description, price, duration) VALUES
 (uuid_generate_v4(), 'Test Plan 1 Minute', 'Test subscription plan for 1 minute', 10.00, 60000000000),
 (uuid_generate_v4(), 'Test Plan 1 Day', 'Test subscription plan for 1 day', 50.50, 86400000000000),
 (uuid_generate_v4(), 'Normal Plan 1 Month', 'Normal subscription plan for 1 month', 500.00, 2592000000000000);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin


-- Удаление связанных записей из таблицы payment
DELETE FROM payment WHERE subscription_id IN (
    SELECT id FROM subscription WHERE tariff_plan_id IN (
        SELECT id FROM tariff_plan WHERE name IN ('Test Plan 1 Minute', 'Test Plan 1 Day', 'Normal Plan 1 Month')
    )
);
-- Удаление связанных данных
DELETE FROM subscription WHERE tariff_plan_id IN (
    SELECT id FROM tariff_plan WHERE name IN ('Test Plan 1 Minute', 'Test Plan 1 Day', 'Normal Plan 1 Month')
);
-- Удаление данных из таблицы tariff_plan
DELETE FROM tariff_plan WHERE name IN ('Test Plan 1 Minute', 'Test Plan 1 Day', 'Normal Plan 1 Month');

-- +goose StatementEnd