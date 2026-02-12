# Quick Start Guide - Todo CLI

Get up and running with Todo CLI in 5 minutes!

## Installation

### Option 1: Install via Homebrew (macOS)

The easiest way to install on macOS:

```bash
# Add the tap
brew tap gourangadassamrat/tap

# Install the app
brew install todo-cli
```

That's it! You can now run the app using the `todo` command.

### Option 2: Manual Installation

#### Step 1: Download the Application

1. If you're on Mac, Linux, or Windows, you can download the latest build for your OS from the [Release page](https://github.com/gourangadassamrat/todo-cli-go/releases).

#### Step 2: Test It Works

```bash
./todo --help
```

You should see the help menu with all available commands.

#### Step 3 (Optional): Install Globally

```bash
# Linux/macOS
sudo cp todo /usr/local/bin/

# Windows
# Copy todo.exe to a directory in your PATH
```

## Your First Tasks

### 1. Add Your First Task

```bash
todo add -t "Learn Todo CLI" -p high
```

### 2. Add a Task with Details

```bash
todo add \
  -t "Complete project report" \
  -d "Include Q1 metrics and forecasts" \
  -p medium \
  --project "Work" \
  --due "2024-03-20"
```

### 3. View All Tasks

```bash
todo list
```

### 4. Try Interactive Mode

```bash
todo interactive
```

Use arrow keys to navigate, Enter to select!

## Essential Commands

| What You Want | Command                                |
| ------------- | -------------------------------------- |
| Add a task    | `todo add -t "Task name"`              |
| See all tasks | `todo list`                            |
| Search tasks  | `todo search "keyword"`                |
| Mark as done  | `todo complete -i <task-id>`           |
| Edit a task   | `todo edit -i <task-id> -t "New name"` |
| Delete a task | `todo delete -i <task-id>`             |

## Quick Tips

### 🎯 Use Priorities

```bash
todo add -t "Critical bug fix" -p high
todo add -t "Code review" -p medium
todo add -t "Update docs" -p low
```

### 📅 Set Due Dates

```bash
todo add -t "Submit report" --due "2024-03-15"
todo add -t "Team meeting" --due "2024-03-10 14:00"
```

### 🏷️ Organize with Tags & Projects

```bash
todo add -t "Fix login bug" --project "Website" --tags "bug,urgent"
todo add -t "Design mockup" --project "Mobile App" --tags "design"
```

### 🔍 Find Tasks Quickly

```bash
# By status
todo list --status pending
todo list --status completed
todo list --status overdue

# By priority
todo list --priority high

# By project
todo list --project "Website"

# Search
todo search "bug"
```

### 📊 Sort Your View

```bash
# By priority (high to low)
todo list --sort priority

# By due date (soonest first)
todo list --sort date

# Oldest first
todo list --sort created --asc
```

## Common Workflows

### Morning Routine

```bash
# What's on today's agenda?
todo list --status pending --sort date
```

### Adding Multiple Tasks

```bash
todo add -t "Task 1" -p high
todo add -t "Task 2" -p medium --project "Work"
todo add -t "Task 3" -p low --due "2024-03-25"
```

### Completing Tasks

```bash
# Mark as done
todo complete -i <task-id>

# Oops, need to reopen it
todo complete -i <task-id> --incomplete
```

### Weekly Review

```bash
# See what was completed
todo list --status completed

# Check overdue items
todo list --status overdue

# High priority pending
todo list --status pending --priority high
```

## Shortcuts & Aliases

Save typing with command aliases:

```bash
todo a -t "Quick add"        # instead of 'add'
todo l                        # instead of 'list'
todo ls --status pending      # another list alias
todo e -i 123 -t "Edit"      # instead of 'edit'
todo c -i 123                # instead of 'complete'
todo s "search term"         # instead of 'search'
```

## Backup Your Data

```bash
# Create a backup
todo backup

# See available backups
todo restore

# Restore if needed
todo restore tasks_backup_2024-03-10_14-30-00.json
```

## Next Steps

1. Read the full [README.md](../README.md) for detailed documentation
2. Check out [EXAMPLES.md](EXAMPLES.md) for more use cases
3. Try all the commands: `todo --help`
4. Set up tab completion (see README.md)

## Getting Help

```bash
# Overall help
todo --help

# Command-specific help
todo add --help
todo list --help
todo edit --help
```

## Tips for Success

✅ **Start Simple**: Just use `add` and `list` at first
✅ **Use Interactive Mode**: Great for beginners - `todo interactive`
✅ **Set Priorities**: Helps focus on what matters
✅ **Regular Reviews**: Check `todo list` daily
✅ **Backup Often**: Run `todo backup` regularly

## Troubleshooting

**Can't find the todo command?**

- If installed via Homebrew, try `brew doctor` to check for issues
- If manually installed, make sure you're in the right directory, or install it globally (see Option 2, Step 3 above)

**Tasks not saving?**

- Check that `~/.todo-cli/` directory exists
- Verify you have write permissions

**Colors not showing?**

- Ensure your terminal supports ANSI colors
- Try a different terminal if needed

---

**You're ready to go! Start managing your tasks like a pro! 🚀**

For more details, see:

- [README.md](../README.md) - Full documentation
- [EXAMPLES.md](./EXAMPLES.md) - Detailed examples
- [CONTRIBUTING.md](./CONTRIBUTING.md) - How to contribute
