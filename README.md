# Secret Vault Manager

![CI](https://github.com/Qyroxen/Secret-Vault-Manager/actions/workflows/ci.yml/badge.svg)
![CodeQL](https://github.com/Qyroxen/Secret-Vault-Manager/actions/workflows/codeql.yml/badge.svg)
![Go](https://img.shields.io/badge/Go-1.23+-00ADD8?style=flat&logo=go)
![License](https://img.shields.io/badge/License-MIT-yellow.svg)
![Stars](https://img.shields.io/github/stars/Qyroxen/Secret-Vault-Manager?style=social)
![Issues](https://img.shields.io/github/issues/Qyroxen/Secret-Vault-Manager)
![PRs](https://img.shields.io/github/issues-pr/Qyroxen/Secret-Vault-Manager)

> A production-ready CLI tool built with Go

[![Star Badge](https://img.shields.io/github/stars/Qyroxen/Secret-Vault-Manager?style=social)](https://github.com/Qyroxen/Secret-Vault-Manager/stargazers)

## What is it?

Secret Vault Manager is a production-ready CLI tool built with Go. It provides powerful functionality with a beautiful terminal interface.

## Features

- Fast and efficient (written in Go)
- Beautiful CLI with colored output
- Comprehensive documentation
- GitHub Actions CI/CD
- CodeQL security analysis
- Dependabot for dependency updates
- MIT Licensed
- Fully offline - zero cloud dependency

## Quick Start

```bash
# Install
git clone https://github.com/Qyroxen/Secret-Vault-Manager.git
cd Secret-Vault-Manager
go build -o secretvaultmanager .

# Run
./secretvaultmanager --help
```

## CLI Usage

```bash
# Basic usage
./secretvaultmanager

# With flags
./secretvaultmanager --verbose --output json

# Get help
./secretvaultmanager --help
```

## Examples

```bash
# Example 1
./secretvaultmanager example1

# Example 2
./secretvaultmanager example2 --flag value
```

## Development

```bash
# Run tests
go test ./...

# Build
go build -o secretvaultmanager .

# Lint
golangci-lint run

# Security scan
codeql analyze
```

## Contributing

Contributions are welcome! Please see [CONTRIBUTING.md](CONTRIBUTING.md) for details.

## Security

For security vulnerabilities, please see [SECURITY.md](SECURITY.md).

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

---

<p align="center">
  <a href="https://github.com/Qyroxen/Secret-Vault-Manager/stargazers">
    <img src="https://img.shields.io/github/stars/Qyroxen/Secret-Vault-Manager?style=social" alt="Star this repo">
  </a>
  <a href="https://github.com/Qyroxen/Secret-Vault-Manager/forks">
    <img src="https://img.shields.io/github/forks/Qyroxen/Secret-Vault-Manager?style=social" alt="Fork this repo">
  </a>
  <a href="https://github.com/Qyroxen/Secret-Vault-Manager/issues">
    <img src="https://img.shields.io/github/issues/Qyroxen/Secret-Vault-Manager" alt="Issues">
  </a>
  <a href="https://github.com/Qyroxen/Secret-Vault-Manager/pulls">
    <img src="https://img.shields.io/github/issues-pr/Qyroxen/Secret-Vault-Manager" alt="Pull Requests">
  </a>
</p>
