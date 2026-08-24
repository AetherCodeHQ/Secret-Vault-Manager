# 🔐 Secret Vault Manager

![Go](https://img.shields.io/badge/Go-1.21%2B-00ADD8?style=flat-square&logo=go&logoColor=white)
![Version](https://img.shields.io/badge/Version-v3.0.0-00ADD8?style=flat-square)
![License](https://img.shields.io/badge/License-MIT-green?style=flat-square)
![PRs](https://img.shields.io/badge/PRs-Welcome-brightgreen?style=flat-square)

> Security tool by [AetherCodeHQ](https://github.com/AetherCodeHQ)

`security` `cryptography` `cli` `golang` `json` `crypto`

---

## What is Secret-Vault-Manager?

**Secret-Vault-Manager** is a security-focused tool that analyzes and validates code, configurations, or data for vulnerabilities.

## Features

- ✅ JSON data handling
- ✅ Cryptographic operations
- 🚀 **Zero dependencies** — only Go standard library
- 📦 **Single binary** — compile and run anywhere
- 🔄 **Offline capable** — no internet required

## Installation

```bash
# Clone
git clone https://github.com/AetherCodeHQ/Secret-Vault-Manager.git
cd Secret-Vault-Manager

# Build
go build -o secret-vault-manager .

# Run
./secret-vault-manager <vault-file> <key-hex> [get|set|list] [name] [value]
```

### Or directly with `go run`:
```bash
go run main.go <vault-file> <key-hex> [get|set|list] [name] [value]
```

## Usage

```bash
# Basic usage
./secret-vault-manager <vault-file> <key-hex> [get|set|list] [name] [value]

# With flags
./secret-vault-manager <vault-file> <key-hex> [get|set|list] [name] [value] value <vault-file> <key-hex> [get|set|list] [name] [value]
```

### Example Output

```
$ ./secret-vault-manager <vault-file> <key-hex> [get|set|list] [name] [value]
<vault-file> <key-hex> [get|set|list] [name] [value]
%d secrets:\n
  %s = %s\n
```

## Project Structure

```
Secret-Vault-Manager/
  main.go          # Entry point (116 lines)
  go.mod            # Go module definition
  go.sum            # Dependency checksums
  README.md         # This file
  LICENSE           # MIT License
  CHANGELOG.md      # Version history
```

## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

MIT License - see [LICENSE](LICENSE) for details.

---

Built with ❤️ by [AetherCodeHQ](https://github.com/AetherCodeHQ)
