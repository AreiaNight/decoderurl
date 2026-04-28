Aquí va tu README con el mismo estilo:

# ✦•┈๑⋅⋯ 𝐃𝐞𝐜𝐨𝐝𝐞𝐝𝐆𝐨 ⋯⋅๑┈•✦


<img src="https://github.com/AreiaNight/DecodedGo/blob/main/assets/kory.jpg?width=770&height=578&fit=crop&format=pjpg&auto=webp" alt="img" align="right" width="400px"> <br><br>
<br><br><br>
This project is a simple command-line tool built in Go for encoding and decoding Base64 files and URLs. It is designed for fast and clean output, allowing users to quickly transform data without heavy dependencies or complex setups.

The tool focuses on simplicity and speed, making it useful for developers, analysts, or anyone needing quick Base64 or URL transformations directly from the terminal.

<br><br>

## 【𝑭𝒆𝒂𝒕𝒖𝒓𝒆𝒔】
* **Base64 Decoding**
  Reads and decodes a Base64-encoded file directly from disk.
* **URL Decoding**
  Strips percent-encoding from URLs, returning clean readable strings.
* **URL Encoding**
  Converts strings into percent-encoded URL format.
* **Fast Execution**
  Built in Go for instant performance.
* **Minimal & Clean Output**
  Focused results without unnecessary noise.

---

## 【𝑰𝒏𝒔𝒕𝒂𝒍𝒍𝒂𝒕𝒊𝒐𝒏】

### Prerequisites
* Go (1.18+ recommended)

---

### 【𝑴𝒂𝒏𝒖𝒂𝒍 𝑰𝒏𝒔𝒕𝒂𝒍𝒍】

Linux and macOS:

```bash
git clone https://github.com/yourusername/DecodedGo.git
cd decoded64
go build -o decodedgo main.go
```

Move binary:
```bash
sudo mv decodedgo /usr/local/bin/
```

Or without sudo:
```bash
mkdir -p ~/.local/bin
mv decodedgo ~/.local/bin/
```

### 【𝑰𝒏𝒔𝒕𝒂𝒍𝒍】

```bash
git clone https://github.com/yourusername/DecodedGo.git
cd DecodedGo
chmod +x install.sh #If using mac, use chmod +x installMac.sh
./install.sh #Or ./installMac.sh for mac
```

---

## 【𝑼𝒔𝒂𝒈𝒆】

Decode a Base64 file:
```bash
decodedgo -b64 file.txt
```

Decode a URL:
```bash
decodedgo -u "https%3A%2F%2Fexample.com"
```

Encode a URL:
```bash
decodedgo -c "https://example.com"
```

---

### Example Output

```
# Base64 decode
Hello, World!

# URL decode
https://example.com/path?query=hello world

# URL encode
https%3A%2F%2Fexample.com%2Fpath%3Fquery%3Dhello%20world
```

---

### Help
```bash
decodedgo --help
```

---

## 【𝑪𝒐𝒎𝒎𝒂𝒏𝒅𝒔】

| Command                  | Description              |
| ------------------------ | ------------------------ |
| `decodedgo -b64 <file>`  | Decode a Base64 file     |
| `decodedgo -u <url>`     | Decode a URL             |
| `decodedgo -c <string>`  | Encode a string to URL   |
| `--help`                 | Show help                |

---

## 【𝑵𝒐𝒕𝒆𝒔】
* This tool is focused on **quick data transformation**
* Base64 input must be a valid encoded file
* URL encoding uses standard **path escaping**
* Intended for **fast terminal use**, not full data pipelines

---

## 【𝑫𝒊𝒔𝒄𝒍𝒂𝒊𝒎𝒆𝒓】
This project is a personal utility tool and will receive updates based on what I need. It is not intended as a production-grade solution.
```
