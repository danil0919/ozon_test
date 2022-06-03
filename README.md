# Тестовое задание OZON <br><br>
# Описание сервиса

Сервис предоставляет API по созданию сокращенных ссылок. Принимает следующие запросы:
### GET /{token}
Делаем запрос по короткому URL, в ответ получаем длинный URL

token - строка, состоящая из 10 символов <br>

### Ответы:
```
//200 - запись найдена
{
  "long_url":"https://www.golangprograms.com/get-current-date-and-time-in-various-format-in-golang.html"
}

//404 - запись не найдена
{
  "error":"record not found"
}
``` 

### POST /
Генерируем короткий URL по длинному URL, который был передан

### Тело запроса

```
{
  "long_url":string
}
``` 
### Ответы
```
//201 - запись успешно добавлена
{
  "long_url":"https://github.com/danil0919/ozon_test",
  "short_url":"http://localhost:8080/Dt6QYmJYfB",
  "message":"Short link was successfully created"
}

//200 - запись с таким URL уже существует, возвращает уже существующую сокращенную ссылку
{
  "long_url":"https://github.com/danil0919/ozon_test",
  "short_url":"http://localhost:8080/Dt6QYmJYfB",
  "message":"Short url for this url already exists"
}

//400 - некорректные входные данные
{
  "error":"Incorrect input data"
}

``` 

### GET /info/{token}

Получаем подробную информацию о сокращенной ссылке: просмотры, дату создания

token - строка, состоящая из 10 символов <br>

### Ответы:
```
//200 - запись найдена
{
  "long_url":"https://github.com/danil0919/ozon_test/",
  "short_url":"http://localhost:8080/Dt6QYmJYfB",
  "views":4,
  "created_at":"2022-06-03T01:07:38.590275Z"
}

//404 - запись не найдена
{
  "error":"record not found"
}
``` 

# Настройка

Конфиг сервера можно найти в файле [config/apiserver.toml](config/apiserver.toml), либо указать путь к своему конфигу при запуске (подробнее в разделе "Запуск")

> host используется только для вывода short_url в http ответах, предшествуя токену
```
bind_addr = ":8080"
log_level = "debug"
database_url = "host=host.docker.internal dbname=ozon_test sslmode=disable user=postgres password=root"
store_type = "sql"
host = "http://localhost:8080"
```
# Запуск

### Вручную
```
//Миграция
$ migrate -path migrations -database "postgres://localhost:5432/ozon_test?sslmode=disable" up

//Сервер
$ make
$ ./apiserver -help
Usage of ./apiserver:
  -config-path string
        path to  config file (default "configs/apiserver.toml")
  -store-type string
        available stores: sql, internal (default "sql")
$ ./apiserver 
```
>Для работы с внутренним хранилищемм в памяти приложения
```
$ ./apiserver -store-type internal для работы с внутренним хранилищем
```

### Docker
```
$ docker-compose build
$ docker-compose up
```

# Тесты

```
$ make test
```

>Для успешного прохождения тестов пакета sqlstore, необходимо создать тестовую БД и запустить для нее миграцию. Путь к тестовой базе данных можно поменять в файле [internal/app/store/sqlstore/store_test.go](internal/app/store/sqlstore/store_test.go) либо передать через переменную окружения DATABASE_URL.
```
$ migrate -path migrations -database "postgres://localhost:5432/test_ozon_test?sslmode=disable" up
```

