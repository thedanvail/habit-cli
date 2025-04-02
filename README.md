# Habits CLI

A simple command-line habit tracker that helps you build and maintain daily habits.

## Screenshots

### Main Habit Grid (Month View)

![Month Grid View](assets/images/month-grid.png)
*Monthly view of all habit completions*

### 30-Day Tracker View

![30-Day Tracker](assets/images/thirty-day-view.png)
*Last 30 days of habit tracking*

### Habit List

![Habit List](assets/images/habit-list.png)
*List of all tracked habits*

### Single Habit Detail

![Single Habit Detail](assets/images/single-habit.png)
*Detailed view of a single habit's tracking*

### Statistics View

![Statistics View](assets/images/stats-view.png)
*Completion statistics and streaks*

### Undone Habits View

![Undone Habits](assets/images/undone-view.png)
*List of habits not yet completed today*

### Command Help

![Command Help](assets/images/command-help.png)
*Available commands and usage*

## Features

- Track multiple habits from the command line
- View your habits in a visual grid by day, week, month, or custom ranges
- Get statistics on your streaks and completion rates
- Import and export your data
- Minimal and fast interface
- Cross-platform: works on Windows, macOS, and Linux

## Installation

### Prerequisites

- [Go](https://golang.org/doc/install) 1.20 or higher

### macOS and Linux Installation

```bash
# Clone the repository
git clone https://github.com/yourusername/habit-cli.git
cd habit-cli

# Run the installation script
chmod +x install.sh
./install.sh
```

The install script will:

1. Build the binary
2. Install it to `~/.local/bin/`
3. Add this directory to your PATH if using fish shell
4. Provide instructions for other shells

### Windows Installation

```powershell
# Clone the repository
git clone https://github.com/yourusername/habit-cli.git
cd habit-cli

# Run the PowerShell installation script (in an elevated PowerShell terminal)
.\install.ps1
```

The Windows installation script will:

1. Build the binary
2. Install it to `%USERPROFILE%\AppData\Local\Programs\habits-cli\`
3. Add this directory to your user PATH
4. You'll need to restart any open command prompts for the PATH change to take effect

### Manual Installation

If the installation scripts don't work for you, you can manually build and install:

```bash
# Clone the repository
git clone https://github.com/yourusername/habit-cli.git
cd habit-cli

# Build the binary
go build -o habits habits.go  # On Windows: go build -o habits.exe habits.go

# Move it to a directory in your PATH
# Linux/macOS:
mv habits /usr/local/bin/  # or another directory in your PATH

# Windows (PowerShell):
# Create destination folder
mkdir -Force -Path "$env:USERPROFILE\AppData\Local\Programs\habits-cli"
# Copy binary
Copy-Item -Path .\habits.exe -Destination "$env:USERPROFILE\AppData\Local\Programs\habits-cli"
# Add to PATH (optional)
$env:Path += ";$env:USERPROFILE\AppData\Local\Programs\habits-cli"
```

## Usage

```
habits [command] [args]
```

### Basic Commands

- `habits list` - List all your tracked habits
- `habits add "Habit Name"` - Add a new habit to track
- `habits done <habit>` - Mark a habit as completed for today
- `habits delete <habit>` - Delete a habit from tracking
- `habits tracker [habit]` - View habit tracker (for a specific habit or all habits)
- `habits stats [habit]` - Show statistics about your habits

### Advanced Commands

- `habits export --file <filename>` - Export your habits data to JSON
- `habits import --file <filename>` - Import habits data from JSON
- `habits edit <habit> --name "New Name"` - Edit a habit's name
- `habits undone` - List habits not completed today

Run `habits help` to see all available commands.

## Demo Data

To try the application with sample data, you can use the included seed file:

```bash
# Import the seed data
habits import --file seed_data.json
```

## Terminal Compatibility Notes

### Windows

- For the best experience on Windows, we recommend using Windows Terminal or PowerShell
- The default Windows Command Prompt (cmd.exe) has limited support for colors and formatting
- If you see garbled text or missing colors, try using Windows Terminal instead

### macOS and Linux

- The application should work in most terminal emulators on macOS and Linux
- For the best experience, use a terminal that supports 256 colors

## License

MIT
