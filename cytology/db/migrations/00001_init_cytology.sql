-- +goose Up
-- +goose StatementBegin
CREATE TABLE cytology_image
(
    id                  uuid            PRIMARY KEY,
    external_id         uuid            NOT NULL,
    doctor_id           uuid            NOT NULL,
    patient_id          uuid            NOT NULL,
    diagnostic_number   integer         NOT NULL,
    diagnostic_marking  varchar(10),
    material_type       varchar(10),
    diagnos_date        timestamp       NOT NULL,
    is_last             boolean         NOT NULL DEFAULT true,
    calcitonin          integer,
    calcitonin_in_flush integer,
    thyroglobulin       integer,
    details             jsonb,
    prev_id             uuid            REFERENCES cytology_image (id),
    parent_prev_id      uuid            REFERENCES cytology_image (id),
    create_at           timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE cytology_image IS 'Хранилище цитологических исследований';
COMMENT ON COLUMN cytology_image.external_id IS 'Внешний идентификатор исследования';
COMMENT ON COLUMN cytology_image.doctor_id IS 'ID врача';
COMMENT ON COLUMN cytology_image.patient_id IS 'ID пациента';
COMMENT ON COLUMN cytology_image.diagnostic_number IS 'Номер диагностики';
COMMENT ON COLUMN cytology_image.diagnostic_marking IS 'Маркировка диагностики (П11, Л23)';
COMMENT ON COLUMN cytology_image.material_type IS 'Тип материала (GS, BP, TP, PTP, LNP)';
COMMENT ON COLUMN cytology_image.is_last IS 'Является ли данная версия последней';
COMMENT ON COLUMN cytology_image.details IS 'Детали диагностики (JSON)';

CREATE TABLE original_image
(
    id          uuid            PRIMARY KEY,
    cytology_id uuid            NOT NULL REFERENCES cytology_image (id) ON DELETE CASCADE,
    image_path  varchar(512)    NOT NULL,
    create_date timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP,
    delay_time  real,
    viewed_flag boolean         NOT NULL DEFAULT false
);

COMMENT ON TABLE original_image IS 'Хранилище оригинальных изображений';
COMMENT ON COLUMN original_image.cytology_id IS 'ID цитологического исследования';
COMMENT ON COLUMN original_image.image_path IS 'Путь к изображению в S3';
COMMENT ON COLUMN original_image.viewed_flag IS 'Флаг просмотра изображения';

CREATE TABLE segmentation_group
(
    id          SERIAL          PRIMARY KEY,
    cytology_id uuid            NOT NULL REFERENCES cytology_image (id) ON DELETE CASCADE,
    seg_type    varchar(10)     NOT NULL,
    group_type  varchar(10)     NOT NULL,
    is_ai       boolean         NOT NULL DEFAULT false,
    details     jsonb,
    create_at   timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE segmentation_group IS 'Хранилище групп сегментаций';
COMMENT ON COLUMN segmentation_group.cytology_id IS 'ID цитологического исследования';
COMMENT ON COLUMN segmentation_group.seg_type IS 'Тип сегментации (NIL, NIR, NIM, CNO, CGE, C2N, CPS, CFC, CLY, SOS, SDS, SMS, STS, SPS, SNM, STM)';
COMMENT ON COLUMN segmentation_group.group_type IS 'Тип группы (CE, CL, ME)';
COMMENT ON COLUMN segmentation_group.is_ai IS 'Создана ли группа AI';

CREATE TABLE segmentation
(
    id                  SERIAL          PRIMARY KEY,
    segmentation_group_id INTEGER       NOT NULL REFERENCES segmentation_group (id) ON DELETE CASCADE,
    create_at           timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE segmentation IS 'Хранилище сегментаций';

CREATE TABLE segmentation_point
(
    id              SERIAL          PRIMARY KEY,
    segmentation_id INTEGER         NOT NULL REFERENCES segmentation (id) ON DELETE CASCADE,
    x               integer         NOT NULL,
    y               integer         NOT NULL,
    uid             integer         NOT NULL,
    create_at       timestamp       NOT NULL DEFAULT CURRENT_TIMESTAMP
);

COMMENT ON TABLE segmentation_point IS 'Хранилище точек сегментации';
COMMENT ON COLUMN segmentation_point.x IS 'Координата X';
COMMENT ON COLUMN segmentation_point.y IS 'Координата Y';
COMMENT ON COLUMN segmentation_point.uid IS 'Уникальный идентификатор точки';

CREATE INDEX idx_cytology_image_external_id ON cytology_image(external_id);
CREATE INDEX idx_cytology_image_doctor_id_patient_id ON cytology_image(doctor_id, patient_id);
CREATE INDEX idx_original_image_cytology_id ON original_image(cytology_id);
CREATE INDEX idx_segmentation_group_cytology_id ON segmentation_group(cytology_id);
CREATE INDEX idx_segmentation_segmentation_group_id ON segmentation(segmentation_group_id);
CREATE INDEX idx_segmentation_point_segmentation_id ON segmentation_point(segmentation_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS segmentation_point CASCADE;
DROP TABLE IF EXISTS segmentation CASCADE;
DROP TABLE IF EXISTS segmentation_group CASCADE;
DROP TABLE IF EXISTS original_image CASCADE;
DROP TABLE IF EXISTS cytology_image CASCADE;
-- +goose StatementEnd
