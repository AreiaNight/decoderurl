# UnCodeURL

A simple and fast command-line tool to decode (and encode) URLs. No more copying URLs to online tools or struggling to read percent-encoded mess.

## Features

- **Decode URLs**: Remove percent-encoding (`%20` → space)
- **Encode URLs**: Add percent-encoding (space → `%20`)
- **Fast**: Written in Go, runs instantly
- **Simple**: Just one command, no configuration needed
- **Clean output**: No clutter, just the result

## Installation

### Linux 

Run the installation script:
```bash
chmod +x install.sh
./install.sh
```

The script will:
- Update your system
- Install Go (if needed)
- Compile and install `uncodeurl` to `/usr/local/bin`

### Manual Installation

1. Make sure you have Go installed:
```bash
go version
```

2. Clone and build:
```bash
git clone https://github.com/yourusername/uncodeurl.git
cd uncodeurl
go build -o uncodeurl main.go
sudo mv uncodeurl /usr/local/bin/
```

## Usage\

![](/showg.png)

### Decode a URL (remove %)
```bash
uncodeurl u "https://example.com/hello%20world"
# Output: https://example.com/hello world
```

### Encode a URL (add %)
```bash
uncodeurl c "https://example.com/hello world"
# Output: https://example.com/hello%20world
```

### Show help
```bash
uncodeurl h
```
## Commands

| Command | Alias | Description |
|---------|-------|-------------|
| `u` | `uncode` | Decode URL (remove percent-encoding) |
| `c` | `code` | Encode URL (add percent-encoding) |
| `h` | `--help` | Show help message |

---

**Note:** This tool uses Go's `net/url` package for encoding/decoding, which follows RFC 3986 standards.

