# Typesense Go Demo

A minimal example of using Typesense search engine with Go.

## Quick Start

```bash
docker run \
  --name typesense-server \
  -p 8108:8108 \
  -v /tmp:/data \
  typesense/typesense:27.1 \
  --data-dir /data --api-key=xyz \
  --enable-cors
```


For running a 3rd party ui locally, you can use the following command:
Note: Cors should be enabled in the typesense server for this to work.
```bash
docker run --name typesense-dashboard -p 80:80 bfritscher/typesense-dashboard
````



## using v2 of the typesense-go client
```bash
go get "github.com/typesense/typesense-go/v2"
```