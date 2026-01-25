FROM golang:alpine

# Устанавливаем зависимости для libvips (поддерживает SVS и compression value 7)
# vips-dev содержит заголовочные файлы для компиляции
# vips содержит runtime библиотеку
# gcc, musl-dev нужны для CGO компиляции
RUN apk add --no-cache \
	git \
	gcc \
	musl-dev \
	vips-dev \
	vips

WORKDIR /service

# Копируем go.mod и go.sum для кэширования слоев Docker
COPY go.mod go.sum ./

# Загружаем существующие зависимости (если go.sum есть)
RUN go mod download || true

# Копируем весь исходный код
COPY . .

# Получаем недостающие зависимости и обновляем go.mod/go.sum
RUN go mod tidy

# Собираем приложение с CGO (требуется для govips/libvips)
ENV CGO_ENABLED=1
RUN go build -o bin/tiler cmd/service/main.go

EXPOSE 50055

CMD ["./bin/tiler"]
