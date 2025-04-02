# Habits CLI

A simple command-line habit tracker that helps you build and maintain daily habits.

## Screenshots

### Habit Tracker View

![Habit Tracker View](assets/images/tracker-view.png)

### Habit List

![Habit List](assets/images/habit-list.png)

### Monthly Statistics

![Monthly Statistics](assets/images/monthly-view.png)

## Features

- Track multiple habits from the command line
- View your habits in a visual grid by day, week, month, or custom ranges
- Get statistics on your streaks and completion rates
- Import and export your data
- Minimal and fast interface

## Installation

### Using the install script

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

### Manual installation

```bash
# Clone the repository
git clone https://github.com/yourusername/habit-cli.git
cd habit-cli

# Build the binary
go build -o habits habits.go

# Move it to a directory in your PATH
mv habits /usr/local/bin/  # or another directory in your PATH
```

## Usage

```
habits [command] [args]
```

### Basic Commands

- `habits list` - List all your tracked habits
- `habits add "Habit Name"` - Add a new habit to track
- `habits done <habit>` - Mark a habit as completed for today
- `habits remove <habit>` - Delete a habit from tracking
- `habits view [habit]` - View a habit grid (or all habits)
- `habits stats [habit]` - Show statistics about your habits

### Advanced Commands

- `habits export <filename>` - Export your habits data to JSON
- `habits import <filename>` - Import habits data from JSON
- `habits edit <habit>` - Edit a habit's name
- `habits undone <habit>` - Mark a habit as not done for today

Run `habits help` to see all available commands.

## License

MIT
