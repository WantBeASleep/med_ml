# Пакет server

Реализация RPC сервера и handler'ов protobuf.

### Договоренности по реализации

- 1 пакет = 1 handler
- Каждый handler регистрируется в server/server.go
- Handler не должен самостоятельно реализовывать логику, а должен делегировать ее в пакет service
- Вся валидация должна быть реализована средствами стандартной библиотеки go и реализована внутри пакета handler'а или
  в одном из сервисов пакета service, если валидация подразумевается не только в rpc.