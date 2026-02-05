# Birthday Feature

The Slack Daily History Bot now celebrates its own birthday on **January 8th**!

## What Happens on January 8th

When the bot runs on January 8th, it will:

1. **Post a special birthday message** before the regular daily content
2. **Include a fun, self-congratulatory tribute** about the bot's accomplishments
3. **Add a Giphy birthday GIF** (if Giphy API key is configured)
4. **Include a YouTube birthday video** link
5. **Continue with regular daily content** after the birthday celebration

## Birthday Message Features

The bot has multiple birthday messages that rotate, including:
- Self-praise and accomplishments
- Humorous takes on being a bot
- Recognition of its service record
- Thanks to its users

Each message is different and includes:
- 🎉 Birthday emojis and celebration
- Age calculation (years since launch in 2026)
- Links to birthday GIFs (via Giphy API)
- Links to birthday celebration videos (YouTube)

## Example Birthday Message

```
From the desk of the Grant - Wednesday, January 8

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎉 IT'S MY BIRTHDAY! 🎉

That's right, folks! Today marks 0 years since I first graced your
Slack workspace with my presence. I've been tirelessly collecting
historical facts, philosophical musings, and Mr Blobby trivia to
brighten your mornings.

Some might say I'm just a bot. But I prefer to think of myself as
a cultural institution. A digital historian. A purveyor of fine facts.

Here's to another year of enlightening you all! 🥳

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

🎁 Birthday GIF: Click here for celebration!
🎵 Birthday Jam: Watch the celebration video!

━━━━━━━━━━━━━━━━━━━━━━━━━━━━

Thank you for being part of my journey. Here's to many more years
of historical facts, philosophical musings, and the occasional
Mr Blobby reference! 🎊

[Regular daily content continues below...]
```

## Configuration

The birthday feature uses the existing `GIPHY_API_KEY` from your `.env` file.
No additional configuration is needed!

## Implementation Details

### Files Created
- `internal/birthday/birthday.go` - Core birthday detection and message generation
- `internal/birthday/birthday_test.go` - Unit tests
- Updates to `internal/slack/poster.go` - Birthday posting method
- Updates to `cmd/bot/main.go` - Birthday check integration

### Key Functions
- `IsBotBirthday()` - Checks if today is January 8th
- `GetBotAge()` - Calculates years since launch (2026)
- `FetchBirthdayGif(apiKey)` - Fetches a random birthday GIF from Giphy
- `GetBirthdayMessage(giphyAPIKey)` - Generates the complete birthday message
- `PostBirthday(birthdayMsg)` - Posts the birthday message to Slack

## Testing

Since today is January 8th, you can test immediately:

```bash
# Build the bot
go build -o bin/history-slackbot cmd/bot/main.go

# Run once to see the birthday message
RUN_ONCE=true ./bin/history-slackbot
```

Check your Slack channel - you should see the birthday message!

## Birthday Messages Rotation

The bot has 4 different birthday messages that rotate randomly:
1. "IT'S MY BIRTHDAY!" - Classic celebration
2. "HAPPY BIRTHDAY TO ME!" - Accomplishments list
3. "BREAKING NEWS: I'M X YEARS OLD TODAY!" - News-style announcement
4. "BIRTHDAY ANNOUNCEMENT" - Formal but fun announcement

## YouTube Videos

The bot includes links to fun birthday-themed videos:
- "Happy Birthday" by Weird Al Yankovic
- "It's My Birthday" by will.i.am
- Classic birthday song

## Future Enhancements

Possible future additions:
- Count total messages posted since launch
- Show statistics about most popular content types
- Allow users to "wish" the bot happy birthday
- Special themed content on birthday (history of bots, AI, etc.)

## AWS Deployment

The birthday feature will work automatically in AWS ECS:
- No additional configuration needed
- Uses existing Giphy API key from SSM Parameter Store
- Will post at scheduled time (9:58 AM EST) on January 8th
- Regular content will follow after birthday message

Enjoy celebrating with the bot! 🎂🎉
