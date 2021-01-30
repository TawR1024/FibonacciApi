# Сервис для вычисления чисел последовательнности Фибоначи от N до M

Сервис имеет 2 end-point url

## fibonacci_small
Для вычисления небольших чисел можно воспользоваться формулой Бине,
которая позволяет сразу вычислить N-e число в последовательности.
Но эта формула хороша только для маленьких N<100, при N>100 будет
значительно влиять ошибка округления, а в последствии возникнет выход
за пределы max.float64. Данный способ работает быстро даже без применения
кеширования.

## fibonacci_big
Для вычисления больших N реализована рекурсивная функция вычисления
с применением кеширования на основе Redis.
В реализации использован пакет math.Big, что позволяет выполнять
точные вычисления с большими числами которые не помещаются в стандартные типы
данных.
Для минимизации задержки при вычислении новых чисел уже вычисленные
значения помещаются в Redis. Такой подход позволяет использовать уже
вычисленные результаты начиная со 2 запроса.

Для примера:
```bash
# Первый запрос без попадания в кеш
2021/01/24 12:37:37 "GET http://127.0.0.1:3000/fibonacci_big HTTP/1.1" from 127.0.0.1:51953 - 200 42794B in 7.153471666s

# Повторный запрос, значения найдены в кеше
2021/01/24 12:38:34 "GET http://127.0.0.1:3000/fibonacci_big HTTP/1.1" from 127.0.0.1:56637 - 200 42794B in 1.621323189s

```

# How To Use

`Body`

```json
{
  "from": 1,
  "to": 1 
}
```
`Example`

```bash
curl -XGET http://127.0.0.1:3000/fibonacci_big -d '{"from":11,"to":14}' -H "Content-Type: application/json"
```

`Request`

```json
{
  "11": "89",
  "12": "144",
  "13": "233",
  "14": "377"
}
```

# Default config

Конфигурационный файл может быть расположен в проивзольном месте и
передан как аргумент при запуске приложения. В случае если аргумент не был
передан, то ожидается расположение конфигурационного файла в */etc/fibinacci/config.yaml*

`Config`

```yaml
service:
  host: 127.0.0.1
  port: 3000
redis:
  ip: 127.0.0.1
  port: 6380
  db: 0
  password:
```

# Service deployment

## Ручная сборка и запуск как systemd unit

1. ```bash git clone https://github.com/TawR1024/FibonacciApi/```
2. ```bash go build -o fibonacciApi .```
3. ```bash mv fibonacchiApi /usr/share/bin```
4. Создаём описание юнита  [example]( /etc/systemd/system/fibonacci-api.service)
5. Добавляем юнит ```bash systemct daemon-reload```
6. Создаём пользователя и группу ```bash useradd fibonacci```
7. Запускаем systemctl start fibonacci-api.service

## goreleaser

Для обеспечения постоянной сборки приложения под различные платформы при необходимости 
настроен  [pipeline](.github/workflows/releaseFibonacchi.yml)

Для деплоя приложения достаточно скачать необходимый архив.


## Docker image
Запуск приложения можно осуществить в docker контейнере.
Каждая новая версия собирается и загружается на docker.hub

[pipeline](.github/workflows/buid_image.yml)

```bash
docker pull tawr/fibonacciapi
```


```bash
 docker run -p 127.0.0.1:3000:3000 \
 -v /path/to/config/dir/:/etc/fibonacci/ \
 --name fiboApi tawe/fibonacciapi  
```

Для запуска приложения в docker в конфигурационном файле необходимо указывать *host:0.0.0.0* вместо *127.0.0.0*
```yaml
service:
  host: 0.0.0.0
```

## Запуск связки fibonacciapi + redis

Для простоты запуска всех необходимых компонентов, api + redis как сервис кеширования, подготовлена конфигурация
[docker-compose](docker-compose.yml)
Для корректной работы, в конфигурационном файле необходимо вместо ip адреса redis, подставить имя сервиса из
docker-compose:

```yaml
service:
  host: 0.0.0.0
  port: 3000
redis:
  ip: redis
  port: 6379
  db: 0
  password:
```





