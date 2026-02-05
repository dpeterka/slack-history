#!/bin/bash
# Test script to verify birthday message on January 8th

# Force today's date to January 8th for testing
# Note: This doesn't actually change the system date, but if you have a TEST_DATE
# environment variable in your code, you could use that instead

echo "Testing birthday message functionality..."
echo ""
echo "Note: The bot checks if today is January 8th"
echo "Today's date is: $(date '+%B %d, %Y')"
echo ""

# Build the bot
echo "Building bot..."
go build -o bin/history-slackbot cmd/bot/main.go

if [ $? -ne 0 ]; then
    echo "Build failed!"
    exit 1
fi

echo "Build successful!"
echo ""

# Run tests
echo "Running birthday module tests..."
go test -v ./internal/birthday/

echo ""
echo "To manually test the birthday posting:"
echo "  1. Temporarily modify birthday.IsBotBirthday() to return true"
echo "  2. Run: RUN_ONCE=true ./bin/history-slackbot"
echo "  3. Check your Slack channel for the birthday message"
echo ""
echo "Or wait until January 8th and run the bot normally!"
