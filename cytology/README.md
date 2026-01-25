# Cytology Service

Микросервис для работы с цитологическими исследованиями и снимками.

## Структура

Сервис следует архитектуре, аналогичной сервису `uzi`:
- `cmd/service/main.go` - точка входа
- `db/migrations/` - миграции БД
- `proto/grpc/service.proto` - gRPC протобуфы
- `internal/domain/` - доменные модели
- `internal/repository/` - репозитории для работы с БД
- `internal/services/` - бизнес-логика
- `internal/server/` - gRPC handlers

## Основные сущности

1. **CytologyImage** - цитологическое исследование
2. **OriginalImage** - оригинальное изображение
3. **SegmentationGroup** - группа сегментаций
4. **Segmentation** - сегментация с точками

## База данных

Миграции находятся в `db/migrations/`. Для применения миграций используется goose.

База данных: `cytologydb`

## Запуск

Сервис запускается через docker-compose:

```bash
docker-compose up cytology_service
```

Порт: `50070:50055`

## Генерация proto файлов

Для генерации Go кода из proto файлов:

```bash
protoc --go_out=. --go_opt=paths=source_relative \
  --go-grpc_out=. --go-grpc_opt=paths=source_relative \
  proto/grpc/service.proto
```

## Переменные окружения

См. `.env-docker` для примера конфигурации.

## TODO

- [ ] Добавить полные маппинги для segmentation_group и segmentation
- [ ] Добавить валидацию входных данных
- [ ] Добавить тесты
- [ ] Добавить обработку ошибок для всех handlers
