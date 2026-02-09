<div align="center">

# Couik

**The modern, fast, and beautiful TUI typing speed test.**

[![Go Report Card](https://goreportcard.com/badge/github.com/fadilix/couik)](https://goreportcard.com/report/github.com/fadilix/couik)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![AUR version](https://img.shields.io/aur/version/couik-bin)](https://aur.archlinux.org/packages/couik-bin)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/fadilix/couik)](https://github.com/fadilix/couik/releases)


<img width="1920" height="1200" alt="image" src="https://github.com/user-attachments/assets/683b8898-9d75-48ee-888f-e808a3c738ef" />


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

### Windows (Not stable yet. You can use wsl)
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


### Controls

- **Start Typing**: The test begins immediately.
- `ESC` / `Ctrl+C`: Quit the application.
- `Tab`: Restart the test.
- `Ctrl+L`: Restart the test with the same text.
- `Shift+Tab`: Change game mode.

## Configuration

You can customize `couik` using the `config` subcommand.

```bash
couik config set [key] [value]
```

### Available Settings

| Key | Valid Values | Description | Example |
| :--- | :--- | :--- | :--- |
| `mode` | `quote`, `time`, `words` | Sets the default game mode. | `couik config set mode quote` |
| `time` | `15s`, `30s`, `60s`, `120s`| Sets the test duration or word count limit. | `couik config set time 30s` |
| `quote_type` | `small`, `mid`, `thicc` | Adjusts the length/volume of the quotes. | `couik config set quote_type mid` |
| `dashboard_ascii` | path to `.txt` file | Sets a custom ASCII art for the dashboard. | `couik config set dashboard_ascii ~/art.txt` |


### Custom Dashboard Art

You can personalize the dashboard by adding your own ASCII art. Simply create a text file with your art and point `couik` to it.

<img width="1920" height="1200" alt="image" src="https://github.com/user-attachments/assets/d0b9898c-08a9-4f4c-a6a6-06f25d1c1764" />

To set a custom ASCII logo, use the following command:

```bash
couik config set dashboard_ascii /path/to/your/ascii.txt
```

Alternatively, you can manually edit the configuration file (typically located at `/home/username/.config/couik/config.yaml` on Linux):

```yaml
mode: ""
dashboard_ascii: /home/username/.config/couik/logo.txt
quote_type: ""
time: ""
```

### Subcommands

- `config`: Configure your preferences.
- `completion`: Generate autocompletion scripts.


## Contributing

Contributions are welcome! Feel free to open issues or submit pull requests.

1. Fork the repository.
2. Create your feature branch (`git checkout -b feature/amazing-feature`).
3. Commit your changes (`git commit -m 'Add some amazing feature'`).
4. Push to the branch (`git push origin feature/amazing-feature`).
5. Open a Pull Request.

