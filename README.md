# URL Shortener Microservice / Микросервис для сокращения URL

### Основные требования

- Длина сокращенного URL-адреса должна быть *как можно короче*.
- Сокращенный URL может содержать цифры (`0-9`) и буквы (`a-z`, `A-Z`).

### Эндпоинты

#### 1. Создание сокращенного URL

- **Метод**: POST
- **URL**: [http://localhost:8080/](http://localhost:8080/)
- **Request (body)**: 
http://cjdr17afeihmk.biz/123/kdni9/z9d112423421

- **Response**: 
http://localhost:8080/qtj5opu

#### 2. Получение полного URL по сокращенному URL

- **Метод**: GET
- **URL**: [http://localhost:8080/qtj5opu](http://localhost:8080/qtj5opu)
- **Request (url query)**: 
http://localhost:8080/qtj5opu

- **Response (body)**: 
http://cjdr17afeihmk.biz/123/kdni9/z9d112423421

### Хранение данных

Микросервис имеет возможность хранить информацию в *памяти* или в базе данных *PostgreSQL* в зависимости от флага запуска.

### Язык программирования

Этот микросервис написан на языке программирования *Go* (Golang), обеспечивая *быстродействие* и *эффективное использование ресурсов*.

### Зависимости

Для работы микросервиса требуется наличие *Go* версии X.X.X и установленных зависимостей, указанных в файле `go.mod`.

### Использование

Чтобы запустить микросервис, выполните следующие шаги:

1. Клонируйте репозиторий на локальную машину.
2. Убедитесь, что все зависимости установлены с помощью `go mod tidy`.
3. Используйте `go build` для сборки проекта.
4. Запустите исполняемый файл, указав необходимые параметры (например, флаг хранения данных).