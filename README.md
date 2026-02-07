<div align="center">

# Couik

**The modern, fast, and beautiful TUI typing speed test.**

[![Go Report Card](https://goreportcard.com/badge/github.com/fadilix/couik)](https://goreportcard.com/report/github.com/fadilix/couik)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![AUR version](https://img.shields.io/aur/version/couik-bin)](https://aur.archlinux.org/packages/couik-bin)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/fadilix/couik)](https://github.com/fadilix/couik/releases)

![Demo](https://vhs.charm.sh/vhs-2jdL3lZ9yX4vGk5n1vWqQ.gif)

</div>

---

## Overview

**Couik** is a terminal-based typing test designed for speed, aesthetics, and simplicity. Built with [Bubble Tea](https://github.com/charmbracelet/bubbletea), it provides a smooth and responsive experience right in your CLI.

Whether you want to warm up before coding, challenge your friends, or just track your WPM progress, Couik has you covered.

## Features

- **Timed Mode**: Test yourself against the clock (`15s`, `30s`, `60s`, `120s`).
- **Word Mode**: Practice with a set number of words (`10`, `25`, `50`, `100`).
- **Custom Text**: Load your own text files or paste custom strings.
- **History Tracking**: Keep track of your progress over time with built-in stats.
- **Beautiful UI**: Modern, clean interface with real-time feedback.
- **Cross-Platform**: Works on Linux, macOS, and Windows.

## Installation

### Linux

#### Arch Linux (AUR)
```bash
yay -S couik-bin
```

#### Debian / Ubuntu
Download the `.deb` file from [Releases](https://github.com/fadilix/couik/releases) and run:
```bash
sudo apt install ./couik_*_linux_amd64.deb
```

#### Fedora / RHEL
Download the `.rpm` file from [Releases](https://github.com/fadilix/couik/releases) and run:
```bash
sudo dnf install ./couik_*_linux_amd64.rpm
```

#### Binary (Generic)
```bash
tar -xzf couik_*_linux_amd64.tar.gz
sudo mv couik /usr/local/bin/
```

### Windows
1. Download the `.zip` archive from [Releases](https://github.com/fadilix/couik/releases).
2. Extract and run `couik.exe`.
3. (Optional) Add to your PATH for global access.

### macOS
Download the binary from [Releases](https://github.com/fadilix/couik/releases) or use Go:
```bash
go install github.com/fadilix/couik/cmd/couik@latest
```

### Build from Source
Requires [Go 1.21+](https://go.dev/dl/).
```bash
go install github.com/fadilix/couik/cmd/couik@latest
```

## Usage

Simply run `couik` to start a default session.

```bash
couik [flags]
```

### Command Flags

| Flag | Short | Description | Example |
|------|-------|-------------|---------|
| `--time` | `-t` | Run a timed test (seconds) | `couik -t 60` |
| `--words` | `-w` | Run a word count test | `couik -w 50` |
| `--file` | `-f` | Use a custom text file | `couik -f ./code.txt` |
| `--custom`| `-c` | Use a custom string | `couik -c "Hello World"` |
| `--history`| `-i`| View your past results | `couik -i` |
| `--help` | `-h` | Show help message | `couik -h` |

## Controls

- **Start Typing**: The test begins immediately.
- `ESC` / `Ctrl+C`: Quit the application.
- `Tab`: Restart the test (if supported by current mode).


## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/amazing-feature`).
3. Commit your changes (`git commit -m 'Add some amazing feature'`).
4. Push to the branch (`git push origin feature/amazing-feature`).
5. Open a Pull Request.

