# Shortly
Сервис по обрезке ссылок + аналитика переходов


архитектура проекта: 
```
shortly/
├── cmd/
│   └── api/          # Точка входа main.go
├── internal/
│   ├── config/       # Конфигурация (viper или env)
│   ├── entity/       # Сущности (Url, Click)
│   ├── repository/   # Интерфейсы и реализация под PostgreSQL
│   │   └── postgres/
│   ├── service/      # Бизнес-логика
│   ├── transport/    # HTTP-обработчики
│   │   └── rest/
│   └── pkg/          # Вспомогательные пакеты (hasher, etc.)
├── migrations/       # SQL-миграции (golang-migrate)
├── proto/            # .proto файлы (позже)
├── docker-compose.yml
├── Dockerfile.api
└── Makefile
```

# План выполнения
- [x] Спроектировать архитектуру проека.
- [x] Создать рабочий docker-compose.yaml с PostgreSQL и Redis 
- [x] Настроить переменные окружения
- [x] Создана базовая структура проекта
- [x] Создать HTTP server
- [x] Создать entity URL
- [x] Разработать сервисный слой для создания коротких ссылок
- [x] Реализовать генератор коротких кодов на основе snowflake + base62
- [x] Написать handler для эндпоинта /v1/shorten
- [ ] OpenAPI
- [ ] Подключить линтеры
- [ ] Graceful shutdown
- [ ] Создать интерфейсы хранищ (storage interfaces)
- [ ] Подключение к PostgreSQL
- [ ] Миграции к PostgreSQL
- [ ] CRUD операции с PostgreSQL
- [ ] Подключение к Redis
- [ ] Реализация CashStorage
- [ ] Создать entity Click
- [ ] Написать эндпоинт редиректа
- [ ] Асинхронная запись кликов
- [ ] Написать эндроинт для статистики
- [ ] Логирование ошибок
- [ ] Метрики
- [ ] Trace запросов
- [ ] Интеграционные тесты
- [ ] Dockerfile
- [ ] Github actions (ci/cd)
- [ ] deploy
