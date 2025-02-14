# Typesense Go Demo

A minimal example of using Typesense search engine with Go.

## Quick Start
```bash
docker run \
  --name typesense-server \
  -p 8108:8108 \
  -v /tmp:/data \
  typesense/typesense:27.1 \
  --data-dir /data --api-key=xyz
```

## using v2 of the typesense-go client
```bash
go get "github.com/typesense/typesense-go/v2"
```