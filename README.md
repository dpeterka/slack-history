# History Slackbot

A Go-based Slack bot that posts interesting "Today in History" events and fun holidays daily. The bot fetches historical events from RSS feeds, uses Claude AI to select the most interesting ones, and posts them to Slack with beautiful formatting.

## Features

### Core Features
- Fetches historical events from multiple RSS feeds
- Uses Anthropic's Claude AI with **enhanced diversity scoring** to select the most interesting events
- **AI-powered diversity requirements**: Ensures variety across time periods, categories, geographic regions, and cultural perspectives
- Posts beautifully formatted messages to Slack with rich formatting
- **Variable content by day of week** - Posts different numbers of events based on the day (Monday: 3, Tuesday: 2, Wednesday: 1, Thursday: 2, Friday: 3)
- **Weekday-only posting** - Automatically skips weekends (Saturday/Sunday)
- Configurable scheduling (default: daily at 9 AM)
- Support for multiple RSS feed sources
- Run-once mode for testing
- Containerized with Docker

### Variety Features (New!)
- **Quote of the Day** - Inspirational quotes from historical figures (via ZenQuotes API with fallback quotes)
- **Emo Comments** - Daily philosophical/emo observations about work, life, and relationships (44 unique comments)
- **Mr Blobby Facts** - Daily fascinating facts about the legendary Mr Blobby (47 unique facts covering Origins, Music, Theme Parks, TV, Legacy, and Controversies)
- **WikiHow Articles** - Random WikiHow article suggestions
- **Hot Tub Care Tips** - Practical hot tub maintenance advice with personality
- **Gardening Tips** - Vegetable and hydroponic gardening wisdom
- **3D Printing Tips** - Practical 3D printing advice covering bed adhesion, filament, troubleshooting, and more (50 unique tips)
- **Notable Births & Deaths** - Highlights significant people born or died on this day (via Wikipedia API)
- **Historical Events** - Interesting historical events from this day in history (AI-curated for uniqueness)
- **Content Rotation System** - Prevents repeating events for configurable weeks (default: 6 weeks)
- **Event History Cache** - Tracks posted events to ensure fresh content
- **Geographic Diversity** - AI prioritizes events from different countries and regions (not just US-centric)
- **Time Period Diversity** - Events span ancient, medieval, modern, and contemporary eras
- **Category Diversity** - Mix of Science, Arts, Politics, Sports, Culture, Technology, and Social movements

### Holiday Features
- Displays major holidays for India, Canada, and the US (e.g., "Today is Christmas Day")
- Fetches fun/unusual holidays (filtered to exclude serious observances)

## Architecture

The project follows Go best practices with a clean architecture:

- `cmd/bot/` - Main application entry point
- `internal/config/` - Configuration management
- `internal/rss/` - RSS feed parsing (with Brotli compression support)
- `internal/llm/` - LLM integration for event selection with diversity scoring
- `internal/slack/` - Slack webhook integration with rich formatting
- `internal/scheduler/` - Job scheduling
- `internal/holidays/` - Major holiday detection for India, Canada, and the US
- `internal/quotes/` - Quote of the day fetcher (ZenQuotes API)
- `internal/emo/` - Emo/philosophical comment generator (44 unique comments)
- `internal/blobby/` - Mr Blobby fact generator (47 unique facts)
- `internal/wikihow/` - WikiHow article selector
- `internal/hottub/` - Hot tub care tip generator
- `internal/gardening/` - Gardening tip generator
- `internal/printing3d/` - 3D printing tip generator (50 unique tips)
- `internal/people/` - Notable births/deaths extraction and filtering (Wikipedia API)
- `internal/wikipedia/` - Wikipedia API integration for people and event links
- `internal/funfacts/` - Daily content rotation coordinator
- `internal/cache/` - Event history tracking for content rotation

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
RSS_FEED_URLS=https://www.onthisday.com/rss/today-in-history.xml,https://unbelievablefactsblog.com/rss,http://feeds.feedburner.com/FutilityCloset,https://www.kickassfacts.com/feed/,https://www.mentalfloss.com/feed
HOLIDAY_FEED_URL=https://api.checkiday.com/rss?tz=America/New_York
SCHEDULE_CRON=0 9 * * *  # 9 AM daily
MAX_EVENTS=1
MAX_HOLIDAYS=2
RUN_ONCE=false
SKIP_INITIAL_RUN=false  # Set to true to skip running on container startup
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
| **Required** | | |
| `SLACK_WEBHOOK_URL` | Slack incoming webhook URL | Required |
| `CLAUDE_API_KEY` | Anthropic Claude API key | Required |
| **Core Settings** | | |
| `CLAUDE_MODEL` | Claude model to use | `claude-sonnet-4-5` |
| `RSS_FEED_URLS` | Comma-separated list of historical events/facts RSS feed URLs | `https://www.onthisday.com/rss/today-in-history.xml,...` |
| `HOLIDAY_FEED_URL` | Fun holidays RSS feed URL | `https://api.checkiday.com/rss?tz=America/New_York` |
| `SCHEDULE_CRON` | Cron expression for scheduling | `0 9 * * *` (9 AM daily) |
| `MAX_EVENTS` | Number of historical events to select | `1` |
| `MAX_HOLIDAYS` | Number of fun holidays to display | `2` |
| `RUN_ONCE` | Run once and exit | `false` |
| **Variety Features** | | |
| `INCLUDE_QUOTE` | Include quote of the day | `true` |
| `INCLUDE_EMO_COMMENT` | Include emo/philosophical comment | `true` |
| `INCLUDE_BLOBBY_FACT` | Include Mr Blobby fact | `true` |
| `INCLUDE_WIKIHOW` | Include WikiHow articles | `true` |
| `INCLUDE_HOTTUB` | Include hot tub care tips | `true` |
| `INCLUDE_GARDENING` | Include gardening tips | `true` |
| `INCLUDE_PRINTING3D` | Include 3D printing tips | `true` |
| `INCLUDE_PEOPLE` | Include notable births/deaths | `true` |
| `INCLUDE_EVENTS` | Include historical events | `true` |
| `MAX_PEOPLE` | Max number of notable people to display | `2` |
| `CACHE_DIR` | Directory for event history cache | `.cache` |
| `CONTENT_ROTATION_WEEKS` | Weeks before repeating an event | `6` |
| `SKIP_INITIAL_RUN` | Skip running on startup, only use schedule | `false` |
| **Advanced** | | |
| `EVENT_SELECTION_PROMPT` | Custom LLM prompt (overrides diversity scoring) | Default prompt |

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
│   ├── scheduler/
│   │   └── scheduler.go      # Job scheduling
│   └── holidays/
│       ├── holidays.go       # Major holiday detection
│       └── holidays_test.go  # Holiday tests
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
2. **Weekend Check** - Skips posting if today is Saturday or Sunday (weekday-only for work Slack)
3. **Day-of-Week Logic** - Determines how many events to post:
   - Monday: 3 events (start the week strong!)
   - Tuesday: 2 events (keep momentum)
   - Wednesday: 1 event (hump day - keep it light)
   - Thursday: 2 events (building to the weekend)
   - Friday: 3 events (end the week with a bang!)
4. **Event History Cache** - Initializes cache and cleans up old entries
5. **RSS Parser** - Fetches historical events and fun holidays from configured RSS feeds (with Brotli compression support)
6. **Content Filtering** - Removes recently posted events (within `CONTENT_ROTATION_WEEKS`)
7. **LLM Selector** - Sends events to Claude AI with enhanced diversity requirements:
   - **Time Period Diversity**: Ancient, medieval, modern, contemporary eras
   - **Category Diversity**: Science, Arts, Politics, Sports, Culture, Technology, Social movements
   - **Geographic Diversity**: Events from Asia, Europe, Africa, Latin America, Middle East (not just US)
   - **Cultural Diversity**: Different cultures and perspectives
   - **Impact Diversity**: Mix of major turning points and fascinating lesser-known events
8. **Cache Update** - Stores selected events in cache to prevent future repeats
9. **Major Holiday Detection** - Checks if today is a major holiday in India, Canada, or the US
10. **Holiday Filter** - Filters out serious/political holidays, keeping only fun ones
11. **Quote Fetcher** - Fetches inspirational quote of the day (if enabled)
12. **People Extractor** - Identifies and filters notable births/deaths from events (if enabled)
13. **Emo Comment Selector** - Selects daily emo/philosophical observation (if enabled)
14. **Mr Blobby Fact Selector** - Selects daily fascinating Mr Blobby fact (if enabled)
15. **Slack Poster** - Formats and posts complete message with:
    - Major holiday (if applicable)
    - Emo comment (if enabled)
    - Mr Blobby fact (if enabled)
    - Quote of the day (if enabled)
    - Fun holidays
    - Notable births/deaths (if enabled)
    - Selected historical events
    - All with rich Slack formatting

## Example Output

The bot posts messages to Slack with this enhanced format:

```
From the desk of the Grant - Wednesday, December 25

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

💭 Another day, another existential crisis at the office.
At least the coffee is consistent.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎀 Mr Blobby Fact of the Day
In 1993, Mr Blobby released a single called 'Mr Blobby' which reached #1 on
the UK Singles Chart, beating Take That's 'Babe' for the Christmas #1 spot.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Today is Christmas Day

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

💭 Quote of the Day
"The only way to do great work is to love what you do."
— Steve Jobs

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎉 Today's Fun Holidays
• National Pumpkin Pie Day
• A'Phabet Day

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎂 Born on This Day
Isaac Newton (1643)
English physicist and mathematician who formulated the laws of motion and universal gravitation

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

1776 • Military
Washington Crosses the Delaware
General George Washington and his troops crossed the Delaware River on
Christmas night, launching a surprise attack on Hessian forces at Trenton,
New Jersey. This bold move became a turning point in the Revolutionary War.

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

```

**Notes:**
- Header changed to "From the desk of the Grant"
- Emo comment appears when `INCLUDE_EMO_COMMENT=true` (44 unique comments rotating daily)
- Mr Blobby fact appears when `INCLUDE_BLOBBY_FACT=true` (47 unique facts covering 13 categories)
- The "Today is [Major Holiday]" section only appears on major holidays for India, Canada, and the US
- Quote of the Day appears when `INCLUDE_QUOTE=true`
- Notable births/deaths appear when `INCLUDE_PEOPLE=true`
- Content sections are automatically included/excluded based on configuration
- Events won't repeat within the configured `CONTENT_ROTATION_WEEKS` period
- Emo comments cover three categories: Work, Life, and Relationships
- Mr Blobby facts cover: Origins, Music, Theme Park, TV, Legacy, Controversies, Commercial, Modern, Design, Academic, Almost Happened, Influence, and Creator

## License

MIT License

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.
