[![Build Status](https://img.shields.io/travis/walle/tyda-api.svg?style=flat)](https://travis-ci.org/walle/tyda-api)
[![Coverage](https://img.shields.io/codecov/c/github/walle/tyda-api.svg?style=flat)](https://codecov.io/github/walle/tyda-api)
[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/walle/tyda-api)
[![license](http://img.shields.io/badge/license-MIT-red.svg?style=flat)](https://raw.githubusercontent.com/walle/tyda-api/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/walle/tyda-api)](http:/goreportcard.com/report/walle/tyda-api)

# tyda-api

Local binary for exposing http://tyda.se as an API.
The binary returns JSON on stdout as default, optionaly indented.

## Installation

```shell
$ go get github.com/walle/tyda-api/...
```

## Usage

```shell
usage: tyda-api [--indented] [--languages LANGUAGES] QUERY

positional arguments:
  query

options:
  --indented, -i         If the output should be indented
  --languages LANGUAGES, -l LANGUAGES
                         Languages to translate to (en fr de es la nb sv) [default: [en]]
  --help, -h             display this help and exit
```

## Testing

Use the `go test` tool.

```shell
$ go test -cover
```

## Contributing

All contributions are welcome! See [CONTRIBUTING](CONTRIBUTING.md) for more
info.

## License

The code is under the MIT license. See [LICENSE](LICENSE) for more
information.
