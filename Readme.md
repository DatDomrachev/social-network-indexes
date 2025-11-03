# Social Network

Монолит и сервис диалогов в отдельном микросервисе.

## Структура проекта

```
├── monolith/                # Директория монолита
│   ├── main.go              # Основной код приложения
│   ├── go.mod               # Зависимости Go
│   ├── analyze.go           # Генерация графиков
│   └── Dockerfile           # Сборка образа монолита
├── dialog-service/          # Директория микросервиса диалогов
│   ├── main.go              # Код Dialog Service
│   ├── go.mod               # Зависимости Go
│   └── Dockerfile           # Сборка образа микросервиса
├── docker-compose.yml       # Оркестрация сервисов
└── postman/                 # Коллекции для тестирования
    └── microservices.json   # Тесты для микросервисов
```

## Запуск

### Часть A: Запуск

```bash
# 1. Запуск всех сервисов через docker-compose
git clone https://github.com/DatDomrachev/social-network-indexes
cd social-network-indexes
docker-compose up -d --build

# 2. Проверка работы сервисов
curl http://localhost:8080/health  # Монолит
curl http://localhost:8081/health  # Dialog Service

# 3. Просмотр логов
docker-compose logs -f
```

## Архитектура

```
Client → Monolith (8080) ──────────→ In-Memory Storage
         │                          (Users, Posts, Friends)
         │
         └── /dialog/* requests
             │
             ↓
         Dialog Service (8081) ─────→ In-Memory Storage
                                      (Messages only)
```

## Cтек

- **Язык**: Go 1.21
- **Web Framework**: Gin
- **Хранилище**: PostgreSQL + In-memory (для демонстрации)
- **Контейнеризация**: Docker & Docker Compose
- **Аутентификация**: JWT токены (упрощенная)

## API Endpoints

### Монолит (порт 8080)
- `GET /health` - Проверка работоспособности
- `POST /login` - Аутентификация
- `POST /user/register` - Регистрация
- `GET /user/get/{id}` - Получение профиля
- `GET /user/search` - Поиск пользователей
- `PUT /friend/set/{user_id}` - Добавить друга
- `POST /post/create` - Создать пост
- `GET /post/feed` - Лента новостей
- `POST /dialog/{user_id}/send` - Отправка сообщения *(проксируется)*
- `GET /dialog/{user_id}/list` - История диалога *(проксируется)*

### Dialog Service (порт 8081)
- `GET /health` - Проверка работоспособности
- `POST /dialog/{user_id}/send` - Отправка сообщения
- `GET /dialog/{user_id}/list` - История диалога
- `GET /dialogs` - Все диалоги пользователя

## 🧪 Тестирование

### 1. Импорт Postman коллекций
- **Микросервисы**: `postman/microservices.json`

### 2. Ключевые тест-кейсы

#### Регистрация и аутентификация:
```bash
# Регистрация пользователя
curl -X POST http://localhost:8080/user/register \\
  -H "Content-Type: application/json" \\
  -d '{"first_name":"Ivan","second_name":"Ivanov","password":"ivan123"}'

# Логин
curl -X POST http://localhost:8080/login \\
  -H "Content-Type: application/json" \\
  -d '{"id":"<user-id>","password":"ivan123"}'
```

#### Диалоги (главная функция):
```bash
# Отправка сообщения через монолит
curl -X POST http://localhost:8080/dialog/<friend-id>/send \\
  -H "Authorization: Bearer <token>" \\
  -H "Content-Type: application/json" \\
  -d '{"text":"Привет! Это сообщение обрабатывается микросервисом!"}'

# Получение истории диалога
curl -X GET http://localhost:8080/dialog/<friend-id>/list \\
  -H "Authorization: Bearer <token>"
```

## 🔍 Проверка миграции

### 1. Обратная совместимость
Клиенты продолжают использовать те же эндпоинты `/dialog/*`, но теперь они обрабатываются микросервисом.

### 2. Изоляция сервисов
```bash
# Прямой вызов Dialog Service (должен провалиться без заголовка)
curl -X POST http://localhost:8081/dialog/<user-id>/send \\
  -H "Content-Type: application/json" \\
  -d '{"text":"Test"}'
# 401 Unauthorized

# Прямой вызов с заголовком X-User-ID (успех)
curl -X POST http://localhost:8081/dialog/<user-id>/send \\
  -H "X-User-ID: <current-user-id>" \\
  -H "Content-Type: application/json" \\
  -d '{"text":"Direct call to microservice"}'
# 200 OK
```

### 3. Мониторинг состояния
```bash
# Health check показывает статус обоих сервисов
curl http://localhost:8080/health
# Ответ: {"status":"ok","service":"monolith","dialog_service_status":"ok"}

curl http://localhost:8081/health
# Ответ: {"status":"ok","service":"dialog-service","stats":{...}}
```

## Docker образы

### Сборка и публикация
```bash
# Сборка образов
docker build -t DatDomrachev/social-network-monolith ./monolith
docker build -t DatDomrachev/dialog-service ./dialog-service

# Публикация в Docker Hub
docker push DatDomrachev/social-network-monolith
docker push DatDomrachev/dialog-service
```

### Использование готовых образов
```yaml
# docker-compose.yml  
services:
  monolith:
    image: DatDomrachev/social-network-monolith:v2.0
    # ...
  
  dialog-service:
    image: DatDomrachev/dialog-service:v1.0
    # ...
```





