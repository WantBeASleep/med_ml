build:
	docker build -t ml_service .
run:
	docker run -d ml_service
up:
	make build && make run
gen:
	python3 -m grpc_tools.protoc -I .  --python_out=. --grpc_python_out=. ./ml_service/internal/events/kafka.proto