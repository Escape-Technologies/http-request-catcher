# http-request-catcher

[![Go Reference](https://pkg.go.dev/badge/github.com/Escape-Technologies/http-request-catcher.svg)](https://pkg.go.dev/github.com/Escape-Technologies/http-request-catcher)
[![CI](https://github.com/Escape-Technologies/http-request-catcher/actions/workflows/ci.yaml/badge.svg)](https://github.com/Escape-Technologies/http-request-catcher/actions/workflows/ci.yaml)
[![CD](https://github.com/Escape-Technologies/http-request-catcher/actions/workflows/cd.yaml/badge.svg)](https://github.com/Escape-Technologies/http-request-catcher/actions/workflows/cd.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Escape-Technologies/http-request-catcher)](https://goreportcard.com/report/github.com/Escape-Technologies/http-request-catcher)
[![codecov](https://codecov.io/gh/Escape-Technologies/http-request-catcher/branch/main/graph/badge.svg)](https://codecov.io/gh/Escape-Technologies/http-request-catcher)
![Docker Pulls](https://img.shields.io/docker/pulls/escapetech/http-request-catcher)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/escapetech/http-request-catcher)

## Purpose

- Example of caught http request data.

```text
{
  "data": [
    {
      "id": "a",
      "bucket_id": "aaa",
      "method": "GET",
      "path": "/aaa/a",
      "ip": "172.20.0.1",
      "time": "2022-05-18T08:34:21Z",
      "headers": {
        "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,*/*;q=0.8",
        "Accept-Encoding": "gzip, deflate, br",
        "Accept-Language": "en-US,en;q=0.5",
        "Connection": "keep-alive",
        "Dnt": "1",
        "Sec-Fetch-Dest": "document",
        "Sec-Fetch-Mode": "navigate",
        "Sec-Fetch-Site": "none",
        "Sec-Fetch-User": "?1",
        "Upgrade-Insecure-Requests": "1",
        "User-Agent": "Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:100.0) Gecko/20100101 Firefox/100.0"
      },
      "data": ""
    }
  ],
  "message": "Ok"
}
```

## Options

```bash
CATCHER_PORT            = 8080
REDIS_HOST              = "localhost"
REDIS_PORT              = 6379
REDIS_PASSWORD          = ""
REDIS_DB                = 0
ENTRY_EXPIRATION_TIME   = 120
```

## Deployement

### Using docker-compose

```bash
# No .env file
CATCHER_PORT=8080 REDIS_PASSWORD=any docker-compose up --build
```

## Installation

- Only supports redis for the moment.

```bash
docker run --name redis -p 6379:6379 -e ALLOW_EMPTY_PASSWORD=yes bitnami/redis:latest
make run
```
