<div align="center">

# Couik

**The modern, fast, and beautiful TUI typing speed test.**

[![Go Report Card](https://goreportcard.com/badge/github.com/fadilix/couik)](https://goreportcard.com/report/github.com/fadilix/couik)
[![License: MIT](https://img.shields.io/badge/License-MIT-yellow.svg)](https://opensource.org/licenses/MIT)
[![AUR version](https://img.shields.io/aur/version/couik-bin)](https://aur.archlinux.org/packages/couik-bin)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/fadilix/couik)](https://github.com/fadilix/couik/releases)

<img width="1920" height="1200" alt="image" src="https://github.com/user-attachments/assets/d8bf2f78-28b2-4988-8806-8f353a0cb6a0" />

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

#### Quick Install (recommended)

```bash
curl -fsSL https://raw.githubusercontent.com/Fadilix/couik/main/scripts/install.sh | bash
```

> Automatically detects your CPU architecture (`x86_64` / `arm64`) and installs the right binary to `/usr/local/bin`.  
> On **Arch Linux** it will use your AUR helper (`yay`, `paru`, …) if one is found.

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

| Flag        | Short | Description                            | Example                             |
| ----------- | ----- | -------------------------------------- | ----------------------------------- |
| `--time`    | `-t`  | Run a timed test (seconds)             | `couik -t 60`                       |
| `--words`   | `-w`  | Run a word count test                  | `couik -w 50`                       |
| `--file`    | `-f`  | Use a custom text file                 | `couik -f ./code.txt`               |
| `--custom`  | `-c`  | Use a custom string                    | `couik -c "Hello World"`            |
| `--lang`    | `-l`  | Set the language (`english`, `french`) | `couik -l french`                   |
| `--host`    | `-p`  | Host a multiplayer game on a port      | `couik -p 4242 -n Alice`            |
| `--join`    | `-j`  | Join a multiplayer game (`ip:port`)    | `couik -j 192.168.1.10:4242 -n Bob` |
| `--name`    | `-n`  | Your display name in multiplayer       | `couik -n Alice`                    |
| `--history` | `-i`  | View your past results                 | `couik -i`                          |
| `--help`    | `-h`  | Show help message                      | `couik -h`                          |

### Controls

#### General

| Key              | Action                                             |
| ---------------- | -------------------------------------------------- |
| Start typing     | The test begins immediately                        |
| `ESC` / `Ctrl+C` | Quit the application                               |
| `Ctrl+R`         | Restart the test with a new text                   |
| `Tab`            | Restart the test (on results screen)               |
| `Ctrl+L`         | Restart with the **same** text (on results screen) |
| `Shift+Tab`      | Open / close the mode selector                     |
| `←` / `h`        | Move selector left                                 |
| `→` / `l`        | Move selector right                                |
| `Enter`          | Confirm selection                                  |
| `Ctrl+N`         | Toggle language (English ↔ French)                 |
| `Ctrl+E`         | Open / close the quote type selector               |
| `Ctrl+P`         | Open command palette                               |
| `Ctrl+G`         | Open configuration view                            |

#### Multiplayer (lobby)

| Key              | Action                                         |
| ---------------- | ---------------------------------------------- |
| `Ctrl+J`         | **(Host only)** Start the game for all players |
| `ESC` / `Ctrl+C` | Leave the lobby and quit                       |

## Configuration

You can customize `couik` using the `config` subcommand.

```bash
couik config set [key] [value]
```

### Available Settings

| Key               | Valid Values                | Description                                 | Example                                      |
| :---------------- | :-------------------------- | :------------------------------------------ | :------------------------------------------- |
| `mode`            | `quote`, `time`, `words`    | Sets the default game mode.                 | `couik config set mode quote`                |
| `time`            | `15s`, `30s`, `60s`, `120s` | Sets the test duration or word count limit. | `couik config set time 30s`                  |
| `quote_type`      | `small`, `mid`, `thicc`     | Adjusts the length/volume of the quotes.    | `couik config set quote_type mid`            |
| `dashboard_ascii` | path to `.txt` file         | Sets a custom ASCII art for the dashboard.  | `couik config set dashboard_ascii ~/art.txt` |

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

## Multiplayer

Couik supports real-time multiplayer typing races over TCP. One player hosts the session and others join using the host's local IP address.

### Hosting a game

```bash
couik --host <port> --name <YourName>
# short form:
couik -p 4242 -n Alice
```

This starts a server on the given port, puts you in the lobby, and waits for players to connect. Once everyone is ready, press **`Ctrl+J`** to send the text to all players and start the countdown.

### Joining a game

```bash
couik --join <host-ip>:<port> --name <YourName>
# short form:
couik -j 192.168.1.42:4242 -n Bob
```

Connect to a running host using their local IP address and the port they chose. You will be placed in the lobby automatically. The game starts when the host presses `Ctrl+J`.

### Multiplayer flow

1. **Host** runs `couik -p <port> -n <name>`.
2. **Guests** run `couik -j <host-ip>:<port> -n <name>`.
3. All players appear in the lobby. The host sees `ctrl+j  start game`.
4. Host presses `Ctrl+J` — a 5-second countdown begins for everyone.
5. All players type the same text simultaneously; progress bars update in real time.
6. After the round, the host can press `Ctrl+J` again to start a new round.

> **Note:** Multiplayer uses a direct TCP connection. Make sure the port is open on the host's firewall and that all players are on the same network (or the host has port-forwarded for internet play).

## Contributing

First off, thanks for taking the time to contribute! We welcome all contributions, big or small.

### Development Setup

1. **Fork** the repository on GitHub.
2. **Clone** your fork locally:

   ```bash
   git clone https://github.com/YOUR_USERNAME/couik.git
   cd couik
   ```

3. **Run** the project locally to ensure everything works:

   ```bash
   go run cmd/couik/main.go
   ```

### Pull Requests

When you are ready to make a change, follow these steps:

1. **Open an Issue**: Every new feature or bug fix should be linked to an issue. Please open an issue first to discuss your proposed changes before writing code.
2. **Create a branch**: Create a new branch for your feature or bugfix.

   ```bash
   git checkout -b feat/my-cool-feature
   ```

3. **Make your changes**: Write your code and make sure to use clear, descriptive commit messages (e.g., `feat: add awesome feature (#20)`).
4. **Test your code**: Ensure the code builds and runs properly without breaking existing functionality.
5. **Submit a PR**: Push to your fork and submit a Pull Request against the main repository.

### Issues

If you find a bug or have a feature request, [open an issue](https://github.com/fadilix/couik/issues) on GitHub. Include as much detail as possible to help us understand the problem or request.
