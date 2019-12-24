# echoer
[![Actions Status](https://github.com/kamko/echoer/workflows/Docker%20build/badge.svg)](https://github.com/kamko/echoer/actions "docker build status badge")
[![image metadata](https://images.microbadger.com/badges/image/kamko/echoer.svg)](https://microbadger.com/images/kamko/echoer "kamko/echoer image metadata")

Simple http server which echoes headers.

## Usage
Automatically built image.

`docker run -p 80:80 kamko/echoer`

## Configuration
set port via environment variable `PORT` (defaults to 80)

## Example
```json
{
  "Accept": "text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8",
  "Accept-Encoding": "gzip, deflate",
  "Accept-Language": "sk,en-US;q=0.7,en;q=0.3",
  "Cache-Control": "max-age=0",
  "Connection": "keep-alive",
  "Cookie": "JSESSIONID=26CB92A40E1C5D145B54D3FD3C145112",
  "Dnt": "1",
  "Upgrade-Insecure-Requests": "1",
  "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:71.0) Gecko/20100101 Firefox/71.0"
}
```

## Build with
- go 1.13.5 with [Echo](https://echo.labstack.com/)

## License
[MIT](LICENSE)
