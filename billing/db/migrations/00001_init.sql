-- +goose Up
-- +goose StatementBegin
CREATE TYPE subscription_status AS ENUM ('pending_payment', 'active', 'cancelled');
CREATE TYPE payment_status AS ENUM ('pending', 'completed', 'pay_cancelled', 'waiting_for_capture', 'waiting_for_cancel');

CREATE TABLE payment_provider
(
    id   uuid PRIMARY KEY,
    name varchar(255) NOT NULL,
    is_active boolean DEFAULT TRUE
);

COMMENT ON TABLE payment_provider IS 'Таблица для хранения информации о поставщиках платежных услуг (PSP)';

CREATE TABLE tariff_plan
(
    id          uuid PRIMARY KEY,
    name        varchar(255) NOT NULL,
    description text NOT NULL,
    price       numeric(10, 2) NOT NULL,
    duration    bigint NOT NULL --bigint to time.Duration in Go
);

COMMENT ON TABLE tariff_plan IS 'Таблица тарифных планов';
COMMENT ON COLUMN tariff_plan.duration IS 'Длительность подписки';

CREATE TABLE subscription
(
    id              uuid PRIMARY KEY,
    user_id         uuid NOT NULL,
    tariff_plan_id  uuid NOT NULL REFERENCES tariff_plan(id),
    start_date      timestamp NOT NULL,
    end_date        timestamp NOT NULL,
    status          subscription_status NOT NULL DEFAULT 'pending_payment'
);

COMMENT ON TABLE subscription IS 'Таблица подписок пользователей';
COMMENT ON COLUMN subscription.user_id IS 'Идентификатор пользователя из auth';
COMMENT ON COLUMN subscription.status IS 'Статус подписки: pending_payment, active, cancelled';

CREATE TABLE payment
(
    id              uuid PRIMARY KEY,
    user_id         uuid NOT NULL,
    subscription_id uuid NOT NULL REFERENCES subscription(id),
    amount          numeric(10, 2) NOT NULL,
    status          payment_status NOT NULL DEFAULT 'pending',
    payment_provider_id uuid NOT NULL REFERENCES payment_provider(id),
    psp_token       varchar(255),
    created_at      timestamp NOT NULL,
    updated_at      timestamp NOT NULL
);

COMMENT ON TABLE payment IS 'Таблица платежей';
COMMENT ON COLUMN payment.user_id IS 'Идентификатор пользователя из auth';
COMMENT ON COLUMN payment.payment_provider_id IS 'Идентификатор PSP из таблицы payment_provider';
COMMENT ON COLUMN payment.psp_token IS 'Токен платежа от PSP';

CREATE TABLE payment_notification
(
    id                  uuid PRIMARY KEY,
    provider_payment_id  varchar(255) NOT NULL,
    event               varchar(255) NOT NULL,
    payment_provider_id uuid NOT NULL REFERENCES payment_provider(id),
    received_at         timestamp NOT NULL,
    notification_data   json NOT NULL,
    is_valid boolean DEFAULT TRUE
);

COMMENT ON TABLE payment_notification IS 'Таблица для хранения истории уведомлений от PSP';
COMMENT ON COLUMN payment_notification.event IS 'Событие, о котором уведомляет PSP, например, payment.succeeded';
COMMENT ON COLUMN payment_notification.payment_provider_id IS 'Идентификатор PSP из таблицы payment_provider';
COMMENT ON COLUMN payment_notification.notification_data IS 'Полные данные уведомления в формате JSON';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS payment_notification CASCADE;
DROP TABLE IF EXISTS payment CASCADE;
DROP TABLE IF EXISTS subscription CASCADE;
DROP TABLE IF EXISTS tariff_plan CASCADE;
DROP TABLE IF EXISTS payment_provider CASCADE;
DROP TYPE IF EXISTS payment_status;
DROP TYPE IF EXISTS subscription_status;
-- +goose StatementEnd
