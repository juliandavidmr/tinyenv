# tinyenv

A small tool to manage environment variables system written in Go

## Usage
```bash
git clone https://github.com/juliandavidmr/tinyenv.git
cd tinyenv
go run main.go -list keys
```

## Flags

| Flag |   Subflags   | Description |Example|
|------|--------------|-------------|-------|
| list | keys, values | list environment variables | `-list keys` |
| version | | show version        | `-version` |
| cpus    | | number cpus         | `-cpus`    |
| inter   | | network interfaces  | `-inter`   |

### Example

```bash
go run main.go -list keys
```

_License MIT - By [juliandavidmr](https://github.com/juliandavidmr)_
