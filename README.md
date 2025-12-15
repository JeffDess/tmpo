# tmpo CLI

> Set the `tmpo` — A minimal CLI time tracker for developers.

![screenshot of tmpo start and tmpo stats](https://github.com/user-attachments/assets/6480b22a-e148-4142-9cb8-0b4ef1007430)

`tmpo` allows you to track time effortlessly with automatic project detection and simple commands that live in your terminal.

## Table of Contents

1. [Installation](#installation)
2. [Usage](#usage)
3. [Features](#features)
4. [Configuration](#configuration)
5. [License](#license)

---

## About

**tmpo** is a lightweight, developer-friendly time tracking tool designed to integrate seamlessly with your terminal workflow. It automatically detects your project context from Git repositories or configuration files, making time tracking as simple as `tmpo start` and `tmpo stop`.

### Why tmpo?

- **🚀 Fast & Lightweight** - Built in Go, tmpo starts instantly and uses minimal resources
- **🎯 Automatic Project Detection** - Detects project names from Git repos or `.tmporc` files
- **💾 Local Storage** - All data stored locally in SQLite - your time tracking stays private
- **📊 Rich Reporting** - View stats, export to CSV/JSON, and track hourly rates
- **⚡ Zero Configuration** - Works out of the box, configure only when you need to

## Installation

### Homebrew (macOS/Linux)

```bash
# Coming soon
brew install tmpo
```

### Download Pre-built Binaries

Download the latest release for your platform from the [releases page](https://github.com/DylanDevelops/tmpo/releases).

### Build from Source

```bash
git clone https://github.com/DylanDevelops/tmpo.git
cd tmpo
go build -o tmpo .
```

## Quick Start

```bash
# Initialize an optional configuration file
tmpo init

# Start tracking time (auto-detects project from git/directory)
tmpo start

# Start with a description
tmpo start "Implementing user authentication"

# Check current status
tmpo status

# Stop tracking
tmpo stop

# View your time log
tmpo log

# See statistics
tmpo stats
```

## Usage

### Basic Commands

#### `tmpo start [description]`

Start tracking time for the current project. Automatically detects the project name from:

1. `.tmporc` configuration file (if present)
2. Git repository name
3. Current directory name

```bash
tmpo start                              # Start tracking
tmpo start "Fix authentication bug"    # Start with description
```

#### `tmpo stop`

Stop the currently running time entry.

```bash
tmpo stop
```

#### `tmpo status`

View the current tracking session with elapsed time.

```bash
tmpo status
# Output:
# [tmpo] Currently tracking: my-project
#     Started: 2:30 PM
#     Duration: 1h 23m
#     Description: Implementing feature
```

#### `tmpo log`

View your time tracking history.

```bash
tmpo log           # Show recent entries
tmpo log --limit 50  # Show more entries
```

#### `tmpo stats`

Display statistics about your tracked time.

```bash
tmpo stats         # All-time stats
tmpo stats --today  # Today's stats
tmpo stats --week   # This week's stats
```

### Project Configuration

#### `tmpo init`

Create a `.tmporc` configuration file for the current project.

```bash
tmpo init                                    # Auto-detect project name
tmpo init --name "My Project"               # Specify name
tmpo init --name "Client Work" --rate 150   # Set hourly rate
```

This creates a `.tmporc` file:

```yaml
# tmpo project configuration
# This file configures time tracking settings for this project

# Project name (used to identify time entries)
project_name: My Project

# [OPTIONAL] Hourly rate for billing calculations (set to 0 to disable)
hourly_rate: 150.00

# [OPTIONAL] Description for this project
description: ""
```

### Advanced Features

#### `tmpo manual`

Create manual time entries for past work using an interactive prompt.

```bash
tmpo manual
# Prompts for:
# - Project name
# - Start date and time
# - End date and time
# - Description
```

#### `tmpo export`

Export your time tracking data to CSV or JSON.

```bash
tmpo export                              # Export all as CSV
tmpo export --format json                # Export as JSON
tmpo export --project "My Project"       # Filter by project
tmpo export --today                      # Export today's entries
tmpo export --week                       # Export this week
tmpo export --output timesheet.csv       # Specify output file
```

## Features

### Automatic Project Detection

tmpo intelligently detects your project context:

- **`.tmporc` files**: Place a config file in your project root for explicit naming
- **Git repositories**: Automatically uses the repository name
- **Directory fallback**: Uses the current directory name

### Hourly Rate Tracking

Track billable hours with automatic earnings calculations:

```bash
tmpo init --rate 150
tmpo start "Client consultation"
# When you stop, tmpo shows estimated earnings
```

### Flexible Reporting

- **Daily/Weekly stats**: `tmpo stats --today` or `tmpo stats --week`
- **Project summaries**: See time breakdown by project
- **Export options**: CSV and JSON formats for integration with other tools

### Local & Private

All data is stored locally in `~/.tmpo/tmpo.db` using SQLite. Your time tracking data never leaves your machine.

## Configuration

### Global Storage Location

```
~/.tmpo/
  └── tmpo.db          # SQLite database with all time entries
```

### Project Configuration (`.tmporc`)

Place a `.tmporc` file in your project root for custom settings:

```yaml
project_name: My Awesome Project
hourly_rate: 125.50
description: Client project for Acme Corp
```

The `.tmporc` file is automatically detected when you run `tmpo start` from within the project directory or any subdirectory.

## Development

### Building

```bash
# Build for local development
go build -o tmpo .

# Run tests
go test -v ./...

# Build with goreleaser (for releases)
goreleaser build --snapshot --clean
```

### Project Structure

```
tmpo/
├── cmd/                 # CLI commands (Cobra)
│   ├── start.go
│   ├── stop.go
│   ├── status.go
│   └── ...
├── internal/
│   ├── config/         # Configuration management
│   ├── storage/        # SQLite database layer
│   ├── project/        # Project detection logic
│   └── export/         # Export functionality
└── main.go
```

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.

## Acknowledgments

Built with:

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [modernc.org/sqlite](https://gitlab.com/cznic/sqlite) - Pure Go SQLite driver
- [promptui](https://github.com/manifoldco/promptui) - Interactive prompts

---

<p align="center">Made with ❤️ by <a href="https://github.com/DylanDevelops">Dylan Ravel</a> and you!</p>
