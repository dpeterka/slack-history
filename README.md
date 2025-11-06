# History Slackbot

A Go-based Slack bot that posts interesting "Today in History" events and fun holidays daily. The bot fetches historical events from RSS feeds, uses Claude AI to select the most interesting ones, and posts them to Slack with beautiful formatting.

## Features

- Fetches historical events from RSS feeds
- Fetches fun/unusual holidays (filtered to exclude serious observances)
- Uses Anthropic's Claude AI to intelligently select interesting, rare, or significant events
- Posts beautifully formatted messages to Slack
- Configurable scheduling (default: daily at 9 AM)
- Support for multiple RSS feed sources
- Run-once mode for testing
- Containerized with Docker

## Architecture

The project follows Go best practices with a clean architecture:

- `cmd/bot/` - Main application entry point
- `internal/config/` - Configuration management
- `internal/rss/` - RSS feed parsing
- `internal/llm/` - LLM integration for event selection
- `internal/slack/` - Slack webhook integration
- `internal/scheduler/` - Job scheduling

## Prerequisites

- Go 1.21 or later
- Anthropic Claude API key
- Slack incoming webhook URL

## Installation

### Clone the repository

```bash
git clone https://github.com/dpeterka/history-slackbot.git
cd history-slackbot
```

### Install dependencies

```bash
go mod download
```

### Configure environment variables

Copy the example environment file and configure it:

```bash
cp .env.example .env
```

Edit `.env` with your configuration:

```bash
# Required
SLACK_WEBHOOK_URL=https://hooks.slack.com/services/YOUR/WEBHOOK/URL
CLAUDE_API_KEY=sk-ant-api03-xxx

# Optional (defaults shown)
CLAUDE_MODEL=claude-sonnet-4-5
RSS_FEED_URL=https://www.onthisday.com/rss/today-in-history.xml
HOLIDAY_FEED_URL=https://api.checkiday.com/rss?tz=America/New_York
SCHEDULE_CRON=0 9 * * *  # 9 AM daily
MAX_EVENTS=1
MAX_HOLIDAYS=2
RUN_ONCE=false
```

## Setup Guide

### 1. Create a Slack Incoming Webhook

1. Go to [Slack API](https://api.slack.com/apps)
2. Click "Create New App" → "From scratch"
3. Name your app (e.g., "History Bot") and select your workspace
4. In the app settings, go to "Incoming Webhooks"
5. Activate incoming webhooks
6. Click "Add New Webhook to Workspace"
7. Select the channel where you want posts to appear
8. Copy the webhook URL and add it to your `.env` file

### 2. Get an Anthropic Claude API Key

1. Go to [Anthropic Console](https://console.anthropic.com/)
2. Sign up or log in
3. Navigate to API Keys
4. Create a new API key
5. Copy the key and add it to your `.env` file

## Usage

### Build the application

```bash
go build -o bin/history-slackbot cmd/bot/main.go
```

### Run locally

```bash
# Load environment variables and run
source .env
./bin/history-slackbot
```

Or use the Makefile:

```bash
make run
```

### Run once (for testing)

To test without scheduling:

```bash
RUN_ONCE=true go run cmd/bot/main.go
```

Or:

```bash
make test-run
```

### Run with Docker

Build the Docker image:

```bash
docker build -t history-slackbot .
```

Run the container:

```bash
docker run --env-file .env history-slackbot
```

Or use docker-compose:

```bash
docker-compose up -d
```

## Configuration Options

| Variable | Description | Default |
|----------|-------------|---------|
| `SLACK_WEBHOOK_URL` | Slack incoming webhook URL | Required |
| `CLAUDE_API_KEY` | Anthropic Claude API key | Required |
| `CLAUDE_MODEL` | Claude model to use | `claude-sonnet-4-5` |
| `RSS_FEED_URL` | Historical events RSS feed URL | `https://www.onthisday.com/rss/today-in-history.xml` |
| `HOLIDAY_FEED_URL` | Fun holidays RSS feed URL | `https://api.checkiday.com/rss?tz=America/New_York` |
| `SCHEDULE_CRON` | Cron expression for scheduling | `0 9 * * *` (9 AM daily) |
| `MAX_EVENTS` | Number of historical events to select | `1` |
| `MAX_HOLIDAYS` | Number of fun holidays to display | `2` |
| `RUN_ONCE` | Run once and exit | `false` |
| `EVENT_SELECTION_PROMPT` | Custom LLM prompt | Default prompt |

### Cron Schedule Format

The `SCHEDULE_CRON` variable uses a simplified cron format: `minute hour * * *`

Examples:
- `0 9 * * *` - 9:00 AM daily (default)
- `30 8 * * *` - 8:30 AM daily
- `0 12 * * *` - 12:00 PM (noon) daily

## Development

### Project structure

```
history-slackbot/
├── cmd/
│   └── bot/
│       └── main.go           # Application entry point
├── internal/
│   ├── config/
│   │   └── config.go         # Configuration management
│   ├── rss/
│   │   └── parser.go         # RSS feed parsing
│   ├── llm/
│   │   └── selector.go       # LLM event selection
│   ├── slack/
│   │   └── poster.go         # Slack posting
│   └── scheduler/
│       └── scheduler.go      # Job scheduling
├── .env.example              # Example environment variables
├── .gitignore
├── Dockerfile
├── Makefile
├── README.md
└── go.mod
```

### Run tests

```bash
go test ./...
```

Or use the Makefile:

```bash
make test
```

### Format code

```bash
gofmt -w .
```

Or use the Makefile:

```bash
make fmt
```

### Lint code

```bash
golint ./...
```

## How It Works

1. **Scheduler** - Runs the job at the configured time (or immediately if `RUN_ONCE=true`)
2. **RSS Parser** - Fetches historical events and fun holidays from configured RSS feeds
3. **Holiday Filter** - Filters out serious/political holidays, keeping only fun ones
4. **LLM Selector** - Sends events to Claude AI to select the most interesting ones based on:
   - Historical significance
   - Rarity or uniqueness
   - General audience interest
   - Variety across time periods and categories
5. **Slack Poster** - Formats and posts the holidays and selected events to Slack with rich formatting

## Example Output

The bot posts messages to Slack with this format:

```
📅 On This Day in History - Wednesday, November 6

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎉 Today's Fun Holidays
• National Nachos Day
• National Saxophone Day

━━━━━━━━━━━━━━━━━━━━━━━━━━━

1860 • Politics
Abraham Lincoln Elected President
Abraham Lincoln was elected as the 16th President of the United States,
becoming the first Republican to win the presidency. His election triggered
the secession of Southern states and ultimately led to the Civil War.

━━━━━━━━━━━━━━━━━━━━━━━━━━━

Curated by AI from today's historical events
```

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
