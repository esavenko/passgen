# Passgen

A simple cli util for generating passwords in a couple of clicks.

## Manual installation

1. Clone the repo:
``` bash
git clone https://github.com/esavenko/passgen.git
cd passgen
```

2. Install and check deps:
``` bash
go mod tidy
```

3. Build the application ( for macOS):
``` bash
make build
```

Or (binary for Linux/macOS/Windows):
``` bash
make build-all
```

4. Start app:
``` bash
./bin/{{ file u need }}
```

## Possibilities

- Password generation
  - with/without digits
  - with/without special symbols
  - arbitrary length

## Requirements

- For manual installation
  - Go (1.24+)