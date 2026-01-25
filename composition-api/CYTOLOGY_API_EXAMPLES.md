# Примеры вызовов API для Cytology

## Создание цитологического исследования

```bash
curl -X POST "http://localhost:8080/api/v1/cytology" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "external_id": "123e4567-e89b-12d3-a456-426614174000",
    "doctor_id": "123e4567-e89b-12d3-a456-426614174000",
    "patient_id": "123e4567-e89b-12d3-a456-426614174000",
    "diagnostic_number": 1,
    "diagnostic_marking": "П11",
    "material_type": "GS",
    "calcitonin": 10,
    "calcitonin_in_flush": 5,
    "thyroglobulin": 20
  }'
```

Ответ:
```json
{
  "id": "123e4567-e89b-12d3-a456-426614174000"
}
```

## Получение цитологического исследования

```bash
curl -X GET "http://localhost:8080/api/v1/cytology/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

Ответ:
```json
{
  "cytology_image": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "external_id": "123e4567-e89b-12d3-a456-426614174000",
    "doctor_id": "123e4567-e89b-12d3-a456-426614174000",
    "patient_id": "123e4567-e89b-12d3-a456-426614174000",
    "diagnostic_number": 1,
    "diagnostic_marking": "П11",
    "material_type": "GS",
    "diagnos_date": "2024-01-01T00:00:00Z",
    "is_last": true,
    "calcitonin": 10,
    "calcitonin_in_flush": 5,
    "thyroglobulin": 20,
    "create_at": "2024-01-01T00:00:00Z"
  },
  "original_image": {
    "id": "123e4567-e89b-12d3-a456-426614174000",
    "cytology_id": "123e4567-e89b-12d3-a456-426614174000",
    "image_path": "cytology/images/123e4567-e89b-12d3-a456-426614174000.png",
    "create_date": "2024-01-01T00:00:00Z",
    "viewed_flag": false
  }
}
```

## Обновление цитологического исследования

```bash
curl -X PATCH "http://localhost:8080/api/v1/cytology/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "diagnostic_marking": "Л23",
    "is_last": true
  }'
```

## Удаление цитологического исследования

```bash
curl -X DELETE "http://localhost:8080/api/v1/cytology/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Получение исследований по внешнему ID

```bash
curl -X GET "http://localhost:8080/api/v1/cytologies/external/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Получение исследований по ID врача и пациента

```bash
curl -X GET "http://localhost:8080/api/v1/cytologies/patient-card/123e4567-e89b-12d3-a456-426614174000/123e4567-e89b-12d3-a456-426614174001" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Создание оригинального изображения

```bash
curl -X POST "http://localhost:8080/api/v1/cytology/123e4567-e89b-12d3-a456-426614174000/original-image" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "image_path": "cytology/images/123e4567-e89b-12d3-a456-426614174000.png",
    "delay_time": 1.5
  }'
```

## Получение оригинальных изображений

```bash
curl -X GET "http://localhost:8080/api/v1/cytology/123e4567-e89b-12d3-a456-426614174000/original-image" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Обновление оригинального изображения

```bash
curl -X PATCH "http://localhost:8080/api/v1/cytology/original-image/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "delay_time": 2.0,
    "viewed_flag": true
  }'
```

## Создание группы сегментаций

```bash
curl -X POST "http://localhost:8080/api/v1/cytology/123e4567-e89b-12d3-a456-426614174000/segmentation-groups" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "seg_type": "NIL",
    "group_type": "CE",
    "is_ai": false
  }'
```

## Получение групп сегментаций

```bash
curl -X GET "http://localhost:8080/api/v1/cytology/123e4567-e89b-12d3-a456-426614174000/segmentation-groups?seg_type=NIL&group_type=CE" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Обновление группы сегментаций

```bash
curl -X PATCH "http://localhost:8080/api/v1/cytology/segmentation-group/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "seg_type": "NIR"
  }'
```

## Удаление группы сегментаций

```bash
curl -X DELETE "http://localhost:8080/api/v1/cytology/segmentation-group/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Создание сегментации

```bash
curl -X POST "http://localhost:8080/api/v1/cytology/segmentation-group/123e4567-e89b-12d3-a456-426614174000/segments" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "points": [
      {"x": 100, "y": 200},
      {"x": 200, "y": 300},
      {"x": 300, "y": 400}
    ]
  }'
```

## Получение сегментаций группы

```bash
curl -X GET "http://localhost:8080/api/v1/cytology/segmentation-group/123e4567-e89b-12d3-a456-426614174000/segments" \
  -H "Authorization: Bearer YOUR_TOKEN"
```

## Обновление сегментации

```bash
curl -X PATCH "http://localhost:8080/api/v1/cytology/segmentation/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_TOKEN" \
  -H "Content-Type: application/json" \
  -d '{
    "points": [
      {"x": 150, "y": 250},
      {"x": 250, "y": 350}
    ]
  }'
```

## Удаление сегментации

```bash
curl -X DELETE "http://localhost:8080/api/v1/cytology/segmentation/123e4567-e89b-12d3-a456-426614174000" \
  -H "Authorization: Bearer YOUR_TOKEN"
```
