# TaskAPI

REST API для управления задачами, написанный на Go.

---

## Стек

- **Go 1.26**
- **PostgreSQL** — база данных
- **sqlx** — работа с БД
- **golang-migrate** — миграции
- **testify** — unit тесты
- **Docker + docker-compose** — контейнеризация

---

## Эндпоинты

| Метод | Путь | Описание |
|-------|------|----------|
| `GET` | `/tasks` | Получить все задачи |
| `GET` | `/tasks/{id}` | Получить задачу по ID |
| `POST` | `/tasks` | Создать задачу |
| `PUT` | `/tasks/{id}` | Обновить задачу |
| `DELETE` | `/tasks/{id}` | Удалить задачу |

---

## Структура проекта

```
TaskAPI/
├── cmd/
│   └── main.go          # Точка входа
├── Internal/
│   ├── config/          # Загрузка конфигурации
│   ├── db/              # Подключение к БД
│   ├── handlers/        # HTTP хендлеры
│   ├── models/          # Модели данных
│   ├── repository/      # Слой работы с БД
│   └── service/         # Бизнес логика
├── migrations/          # SQL миграции
├── config.yml           # Конфигурация (не в репо)
├── .env                 # Переменные окружения (не в репо)
├── Dockerfile
└── docker-compose.yml
```

---

## Запуск через Docker

```bash
docker-compose up --build
```

API будет доступен на `http://localhost:8080`

---

## Запуск локально

**1. Склонируй репозиторий**
```bash
git clone https://github.com/orionvega2343-cloud/TaskAPI.git
cd TaskAPI
```

**2. Создай конфиг файлы**

Скопируй пример конфига:
```bash
cp config.example.yml config.yml
```

Заполни `config.yml`:
```yaml
db:
  port: port
  host: host
  name: name
  user: user
  ssl_mode: mode
```

Создай `.env`:
```
DB_PASS=твой_пароль
```

**3. Создай базу данных**
```sql
CREATE DATABASE name;
```

**4. Запусти миграции**
```bash
migrate -path migrations -database "postgres://postgres:пароль@localhost:5432/taskapi?sslmode=disable" up
```

**5. Запусти приложение**
```bash
go run cmd/main.go
```

---

## Тесты

```bash
# Запустить все тесты
go test ./...

# С подробным выводом
go test -v ./...

# Покрытие кода
go test -cover ./...
```

---

## Пример запроса

**Создать задачу:**
```bash
curl -X POST http://localhost:8080/tasks \
  -H "Content-Type: application/json" \
  -d '{"title": "Новая задача", "desc": "Описание", "status": "pending"}'
```

**Получить все задачи:**
```bash
curl http://localhost:8080/tasks
```