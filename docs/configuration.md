# Configuration Guide

Learn how to configure tmpo for your projects and workflow.

## Storage Location

All time tracking data and configuration is stored locally on your machine:

```text
~/.tmpo/
  ├── tmpo.db          # SQLite database with time entries
  ├── config.yaml      # Global configuration (optional)
  └── projects.yaml    # Global projects registry (optional)
```

Your data never leaves your machine. All files can be backed up, copied, or version controlled if desired.

> [!NOTE]
> **Contributors**, when developing tmpo with `TMPO_DEV=1` or `TMPO_DEV=true`, both files are stored in `~/.tmpo-dev/` instead to keep development work separate from your production data.

## Global Configuration

### The `tmpo config` Command

Use `tmpo config` to set user-wide preferences that apply across all projects:

```bash
tmpo config
```

This launches an interactive configuration wizard where you can set:

- **Currency** - Your preferred currency for displaying billing rates and earnings
- **Date Format** - Choose between MM/DD/YYYY, DD/MM/YYYY, or YYYY-MM-DD
- **Time Format** - Choose between 24-hour (15:30) or 12-hour (3:30 PM)
- **Timezone** - IANA timezone for your location (e.g., America/New_York, Europe/London)
- **Export Path** - Default directory for exported files (type "clear" to remove)

### Global Settings

Global preferences are stored in `~/.tmpo/config.yaml`:

```yaml
currency: USD
date_format: MM/DD/YYYY
time_format: 12-hour (AM/PM)
timezone: America/New_York
export_path: ~/Documents/timesheets
```

These settings affect how tmpo displays times and currencies throughout the application:

#### Currency

Your currency choice determines the symbol displayed for all billing information across all projects:

**Supported Currencies:**

tmpo supports 30+ currencies including:

- **Americas:** USD ($), CAD (CA$), BRL (R$), MXN (MX$)
- **Europe:** EUR (€), GBP (£), CHF (Fr), SEK (kr), NOK (kr)
- **Asia:** JPY (¥), CNY (¥), INR (₹), KRW (₩), SGD (S$)
- **Oceania:** AUD (A$), NZD (NZ$)

See the [full currency code list](https://en.wikipedia.org/wiki/ISO_4217#Active_codes).

#### Date & Time Formats

Choose how dates and times are displayed and entered throughout tmpo:

**Date Formats:**

- `MM/DD/YYYY` - US format (01/15/2024)
- `DD/MM/YYYY` - European format (15/01/2024)
- `YYYY-MM-DD` - ISO format (2024-01-15)

> [!NOTE]
> Your date format setting affects both display output (in logs, stats, etc.) and input prompts (when using `tmpo manual` or `tmpo edit`). The prompts will show and accept dates in your configured format.

**Time Formats:**

- `24-hour` - Military time (14:30, 23:45)
- `12-hour (AM/PM)` - Standard time (2:30 PM, 11:45 PM)

#### Timezone

Set your IANA timezone for accurate time tracking when working across time zones. Common examples:

- North America: `America/New_York`, `America/Chicago`, `America/Los_Angeles`
- Europe: `Europe/London`, `Europe/Paris`, `Europe/Berlin`
- Asia: `Asia/Tokyo`, `Asia/Singapore`, `Asia/Hong_Kong`
- Oceania: `Australia/Sydney`, `Pacific/Auckland`

Full list: [IANA Time Zone Database](https://en.wikipedia.org/wiki/List_of_tz_database_time_zones)

#### Export Path

Set a default directory where exported files (CSV, JSON) will be saved. This can be overridden per-project in `.tmporc` files.

**Setting the export path:**

```bash
tmpo config
# Export path (press Enter to keep current): ~/Documents/timesheets
```

**Clearing the export path:**

To remove the export path setting and revert to saving in the current directory:

```bash
tmpo config
# Export path (press Enter to keep current): clear
```

**How it works:**

- **Global setting** (`~/.tmpo/config.yaml`): Applies to all projects unless overridden
- **Project setting** (`.tmporc`): Overrides global setting for that project
- **If not set**: Files are exported to your current working directory
- **Supports `~`**: Use `~/Documents` instead of `/Users/yourname/Documents`

**Examples:**

```yaml
# Export to home directory
export_path: ~/exports

# Export to specific folder
export_path: /Users/dylan/Dropbox/timesheets

# No default (export to current directory)
export_path: ""
```

## Global Projects

### What Are Global Projects?

Global projects allow you to track time for any project from any directory without needing `.tmporc` files or Git repositories. They're perfect for:

- **Consulting work** - Track multiple clients without directory structure
- **Non-code projects** - Meetings, research, administrative tasks
- **Flexible workflows** - Switch projects without changing directories
- **Project portfolios** - Manage many small projects easily

### Creating Global Projects

Use `tmpo init --global` to create a global project:

```bash
tmpo init --global
# [tmpo] Initialize Global Project
# Project name: Client Consulting
# Hourly rate (press Enter to skip): 175
# Description (press Enter to skip): Hourly consulting for Acme Corp
# Export path (press Enter to skip): ~/Documents/acme-timesheets
```

Or use `--accept-defaults` for quick setup:

```bash
tmpo init --global --accept-defaults
# Uses current directory name as project name with default values
```

### Using Global Projects

Once created, track global projects from anywhere:

```bash
# Track from any directory
cd /tmp
tmpo start --project "Client Consulting" "Architecture review"

# Resume from anywhere
cd /
tmpo resume --project "Client Consulting"

# View logs from anywhere
tmpo log --project "Client Consulting"

# Export from anywhere
tmpo export --project "Client Consulting" --format csv
```

### The `projects.yaml` File

Global projects are stored in `~/.tmpo/projects.yaml`:

```yaml
projects:
  - name: "Client Consulting"
    hourly_rate: 175.0
    description: "Hourly consulting for Acme Corp"
    export_path: "~/Documents/acme-timesheets"
  - name: "Side Project"
    hourly_rate: 50.0
    description: "Personal side project"
  - name: "Research"
    description: "General research and learning"
```

### Configuration Fields

#### `name` (required)

The project name used when tracking time with the `--project` flag. Must be unique.

```yaml
name: "Client Work - Q1 2024"
```

#### `hourly_rate` (optional)

Your billing rate per hour for this project. The currency symbol is determined by your global currency setting (`tmpo config`).

```yaml
hourly_rate: 150.00
```

Omit or set to `0` to disable rate tracking:

```yaml
# No hourly rate
projects:
  - name: "Personal Project"
    description: "My side project"
```

#### `description` (optional)

Notes or details about the project for your reference.

```yaml
description: "Web development for Acme Corp. Contact: john@acme.com"
```

#### `export_path` (optional)

Default export directory for this project's data.

```yaml
export_path: "~/Documents/client-exports"
```

### Managing Global Projects

You can manually edit `~/.tmpo/projects.yaml` to:

- **Add projects** - Add a new entry to the `projects` list
- **Update projects** - Modify name, rate, description, or export path
- **Remove projects** - Delete an entry from the list

**Example manual edit:**

```yaml
projects:
  - name: "Old Project Name"
    hourly_rate: 100.0
  - name: "New Project"  # Added manually
    hourly_rate: 125.0
    description: "Newly added project"
```

> [!NOTE]
> After manually editing, validate the YAML syntax. Invalid YAML will cause errors when loading projects.

## Project Configuration

### The `.tmporc` File

Place a `.tmporc` file in your project root to customize tracking settings for that project. When you run tmpo commands from within the project directory (or any subdirectory), it will automatically use these settings.

### Creating a Configuration File

Use `tmpo init` to create a `.tmporc` file using an interactive form:

```bash
cd ~/projects/my-project
tmpo init
# You'll be prompted for:
# - Project name (defaults to auto-detected name)
# - Hourly rate (optional, press Enter to skip)
# - Description (optional, press Enter to skip)
# - Export path (optional, press Enter to skip)
```

For quick setup without prompts, use the `--accept-defaults` flag:

```bash
tmpo init --accept-defaults
# Creates .tmporc with auto-detected project name and default values
```

This creates a `.tmporc` file in the current directory.

### File Format

The `.tmporc` file uses YAML format:

```yaml
# tmpo project configuration
# This file configures time tracking settings for this project

# Project name (used to identify time entries)
project_name: My Awesome Project

# [OPTIONAL] Hourly rate for billing calculations (set to 0 to disable)
hourly_rate: 125.50

# [OPTIONAL] Description for this project
description: Client project for Acme Corp

# [OPTIONAL] Default export path for this project (overrides global export path)
export_path: ~/Documents/acme-timesheets
```

### Configuration Fields

#### `project_name` (required)

The name used to identify time entries for this project. This overrides automatic detection from git or directory names.

**Example:**

```yaml
project_name: Client Website Redesign
```

#### `hourly_rate` (optional)

Your billing rate per hour. When set, tmpo will calculate estimated earnings based on tracked time. The currency symbol displayed is determined by your global currency setting (see `tmpo config`).

**Example:**

```yaml
hourly_rate: 150.00
```

Set to `0` or omit to disable rate tracking:

```yaml
hourly_rate: 0
```

#### `description` (optional)

A longer description or notes about the project. This is for your reference and doesn't affect time tracking.

**Example:**

```yaml
description: Q1 2024 website redesign for Acme Corp. Main contact: john@acme.com
```

#### `export_path` (optional)

Default directory for exported files (CSV, JSON) for this project. This overrides the global export path setting from `tmpo config`.

**Example:**

```yaml
export_path: ~/Documents/client-timesheets
```

**How priority works:**

1. **Project `.tmporc` export path** - Highest priority (used if set)
2. **Global config export path** - Used if no project-specific path
3. **Current directory** - Default if neither is set

**Supports home directory expansion:**

```yaml
export_path: ~/Dropbox/timesheets     # Expands to /Users/yourname/Dropbox/timesheets
export_path: /absolute/path/exports   # Absolute paths work too
```

Set to empty string to export to current directory for this project:

```yaml
export_path: ""
```

## Project Detection Priority

When you run `tmpo start`, the project name is determined in this order:

1. **`--project` flag** - Explicitly specified global project (highest priority)
2. **`.tmporc` file** - If present in current directory or any parent directory
3. **Git repository name** - The name of the git repository root folder
4. **Current directory name** - The name of your current working directory (fallback)

This means you can:

- Use `--project` to explicitly track a global project from anywhere
- Override automatic detection by adding a `.tmporc` file
- Let tmpo auto-detect from Git or directory name

### Example Scenarios

#### **Scenario 1:** Explicit global project (highest priority)

```bash
# Directory: /tmp (any directory)
# Global project "Client Work" exists in projects.yaml
tmpo start --project "Client Work"
# → Tracks to global project "Client Work"
```

#### **Scenario 2:** With .tmporc file

```bash
# Directory: ~/code/website-2024/
# .tmporc contains: project_name: "Acme Website"
tmpo start
# → Tracks to project "Acme Website"
```

#### **Scenario 3:** Git repo name

```bash
# Directory: ~/code/website-2024/
# Git repo name: website-2024
# No .tmporc file, no --project flag
tmpo start
# → Tracks to project "website-2024"
```

#### **Scenario 4:** Subdirectory detection

```bash
# Directory: ~/code/my-project/src/components/
# .tmporc exists at: ~/code/my-project/.tmporc
tmpo start
# → Uses .tmporc from project root
```

#### **Scenario 5:** Override local with global

```bash
# Directory: ~/code/website-2024/
# .tmporc contains: project_name: "Website"
# But you want to track to a global project instead
tmpo start --project "Client Work"
# → Tracks to global project "Client Work" (--project overrides .tmporc)
```

## Multi-Project Setup

### Choosing Your Approach

You have three options for managing multiple projects:

1. **Global Projects** - Track projects from any directory (best for consulting, non-code work)
2. **Local .tmporc Files** - Directory-based tracking (best for code projects)
3. **Mix Both** - Use global for flexible work, local for specific codebases

### Option 1: Global Projects

Create global projects once, use them anywhere:

```bash
# Create global projects
tmpo init --global
# Project name: Client A Consulting
# Hourly rate: 150

tmpo init --global
# Project name: Client B Development
# Hourly rate: 175

tmpo init --global
# Project name: Internal Projects
# Hourly rate: 100

# Track from anywhere
cd /tmp
tmpo start --project "Client A Consulting" "Architecture review"
tmpo start --project "Client B Development" "Feature implementation"
```

**Best for:**

- Consulting and freelance work
- Multiple small projects
- Non-code tasks (meetings, research, admin)
- Working across many directories

### Option 2: Local .tmporc Files

Create a `.tmporc` in each project directory using `tmpo init`:

```bash
# Client A - $150/hour
cd ~/projects/client-a
tmpo init
# Project name: Client A - Web Development
# Hourly rate: 150
# Description: [press Enter to skip]

# Client B - different rate
cd ~/projects/client-b
tmpo init
# Project name: Client B - Game Development
# Hourly rate: 175
# Description: [press Enter to skip]

# Personal project - no billing
cd ~/projects/my-app
tmpo init --accept-defaults  # Quick setup with defaults
```

To change currency display (affects all projects):

```bash
tmpo config
# Select your preferred currency (USD, EUR, GBP, etc.)
```

Alternatively, you can manually create `.tmporc` files:

```bash
# Client configuration
cat > ~/projects/client-project/.tmporc << EOF
project_name: Client Project - Web Development
hourly_rate: 150.00
EOF
```

**Best for:**

- Code projects in specific directories
- Team projects with shared .tmporc files
- Projects with consistent directory structure

### Option 3: Mix Global and Local

Combine both approaches for maximum flexibility:

```bash
# Global projects for flexible work
tmpo init --global
# Project name: Consulting Calls
# Hourly rate: 200

tmpo init --global
# Project name: Research & Planning
# Hourly rate: 0

# Local .tmporc for main code projects
cd ~/projects/client-website
tmpo init
# Project name: Client Website
# Hourly rate: 150

cd ~/projects/internal-tool
tmpo init
# Project name: Internal Dashboard
# Hourly rate: 100
```

**Usage example:**

```bash
# Morning: consulting call (global project, from anywhere)
tmpo start --project "Consulting Calls" "Client strategy session"
tmpo stop

# Afternoon: code work (local .tmporc, auto-detected)
cd ~/projects/client-website
tmpo start "Implementing new feature"
tmpo stop

# Evening: research (global project, from anywhere)
cd ~/Downloads
tmpo start --project "Research & Planning" "Exploring new frameworks"
```

**Best for:**

- Mixed work types (code + consulting + meetings)
- Flexibility without losing structure
- Large project portfolios

### Monorepo with Sub-Projects

In a monorepo, you can track different sub-projects separately:

```bash
# Main repo tracks to "My Company Platform"
cd ~/monorepo
echo "project_name: My Company Platform" > .tmporc

# But frontend team tracks separately
cd ~/monorepo/frontend
echo "project_name: Platform - Frontend" > .tmporc

# And backend team tracks separately
cd ~/monorepo/backend
echo "project_name: Platform - Backend" > .tmporc
```

## Version Control

### Should I commit `.tmporc`?

**Yes, commit it** *if*:

- Your team wants shared project naming
- It's an open source project and contributors might want to track time
- The configuration has no sensitive information

**Don't commit it** *if*:

- The hourly rate is personal/confidential
- Each team member prefers their own project naming

### Using `.gitignore`

To keep `.tmporc` files local:

```bash
echo ".tmporc" >> .gitignore
```

Or create a global gitignore:

```bash
echo ".tmporc" >> ~/.gitignore_global
git config --global core.excludesfile ~/.gitignore_global
```

## Migrating Data

### Backing Up Your Data

```bash
# Create a backup of your time tracking database
cp ~/.tmpo/tmpo.db ~/backups/tmpo-backup-$(date +%Y%m%d).db

# Backup your global config
cp ~/.tmpo/config.yaml ~/backups/tmpo-config-backup-$(date +%Y%m%d).yaml

# Backup your global projects registry (if you have global projects)
cp ~/.tmpo/projects.yaml ~/backups/tmpo-projects-backup-$(date +%Y%m%d).yaml
```

### Moving to a New Machine

```bash
# On old machine - backup all files
cp ~/.tmpo/tmpo.db ~/tmpo-export.db
cp ~/.tmpo/config.yaml ~/tmpo-config.yaml
cp ~/.tmpo/projects.yaml ~/tmpo-projects.yaml  # If you have global projects

# Transfer files to new machine, then:
mkdir -p ~/.tmpo
cp ~/tmpo-export.db ~/.tmpo/tmpo.db
cp ~/tmpo-config.yaml ~/.tmpo/config.yaml
cp ~/tmpo-projects.yaml ~/.tmpo/projects.yaml  # If you have global projects
```

### Exporting for External Tools

Use `tmpo export` to get your data in portable formats:

```bash
# Export everything to CSV
tmpo export --output all-time-data.csv

# Export to JSON for programmatic access
tmpo export --format json --output all-time-data.json
```

See the [Usage Guide](usage.md#tmpo-export) for more export options.
