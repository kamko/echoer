# echoer
[![Actions Status](https://github.com/kamko/echoer/workflows/Docker%20build/badge.svg)](https://github.com/kamko/echoer/actions "docker build status badge")
[![image metadata](https://images.microbadger.com/badges/image/kamko/echoer.svg)](https://microbadger.com/images/kamko/echoer "kamko/echoer image metadata")

Simple http server which echoes HTTP requests.

## Usage
Automatically built image.

`docker run -p 80:80 kamko/echoer`

## Configuration
via env vars (shown values are defaults)

```
PORT=80
BODY_AS_STRING=TRUE
```

## Example
```json
{
  "body": "{\n\t\"value1\": \"ABC\"\n}",
  "headers": {
    "Accept": "*/*",
    "Content-Length": "20",
    "User-Agent": "insomnia/7.0.6"
  },
  "uri": "/?param1=1&param2=2"
}
```

When `BODY_AS_STRING` is set to false the body is returned as [base64 string](https://golang.org/pkg/encoding/json/).
```json
{
  "body": "ewoJInZhbHVlMSI6ICJBQkMiCn0=",
  "headers": {
    "Accept": "*/*",
    "Content-Length": "20",
    "User-Agent": "insomnia/7.0.6"
  },
  "uri": "/?param1=1&param2=2"
}
```

## Build with
- go 1.13.5 with [Echo](https://echo.labstack.com/)

## License
[MIT](LICENSE)
