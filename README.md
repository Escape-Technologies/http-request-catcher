# http-request-catcher

[![Go Reference](https://pkg.go.dev/badge/github.com/Escape-Technologies/http-request-catcher.svg)](https://pkg.go.dev/github.com/Escape-Technologies/http-request-catcher)
[![CI](https://github.com/Escape-Technologies/http-request-catcher/actions/workflows/ci.yaml/badge.svg)](https://github.com/Escape-Technologies/http-request-catcher/actions/workflows/ci.yaml)
[![CD](https://github.com/Escape-Technologies/http-request-catcher/actions/workflows/cd.yaml/badge.svg)](https://github.com/Escape-Technologies/http-request-catcher/actions/workflows/cd.yaml)
[![Go Report Card](https://goreportcard.com/badge/github.com/Escape-Technologies/http-request-catcher)](https://goreportcard.com/report/github.com/Escape-Technologies/http-request-catcher)
[![codecov](https://codecov.io/gh/Escape-Technologies/http-request-catcher/branch/main/graph/badge.svg)](https://codecov.io/gh/Escape-Technologies/http-request-catcher)
![Docker Pulls](https://img.shields.io/docker/pulls/escapetech/http-request-catcher)
![Docker Image Size (latest by date)](https://img.shields.io/docker/image-size/escapetech/http-request-catcher)

## Options

```bash
CATCHER_PORT            = 8080
REDIS_HOST              = "localhost"
REDIS_PORT              = 6379
REDIS_PASSWORD          = ""
REDIS_DB                = 0
REQUEST_EXPIRATION_TIME = 120
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
```
