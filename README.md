# Discord History Bot

A simple Discord bot written in Go that allows users to export chat history to text files using slash commands.

## Features

- Export recent chat messages to text files
- Configurable number of messages (1-100)
- Filters out bot messages and empty messages
- Timestamps included with each message
- Uses Discord slash commands

## Prerequisites

- Go 1.21 or higher
- A Discord Bot Token
- Proper bot permissions set up

## Installation

```
git clone https://github.com/yourusername/discord-history-bot
cd discord-history-bot
go mod tidy
```

## Configuration

1. Create a `.env` file in the root directory:

```
DISCORD_TOKEN=your_discord_bot_token_here
```

2. Set up your Discord bot:
   - Go to [Discord Developer Portal](https://discord.com/developers/applications)
   - Create a new application
   - Go to the "Bot" section and create a bot
   - Copy the token and paste it in your `.env` file
   - Go to OAuth2 -> URL Generator
   - Select the following scopes:
     - `bot`
     - `applications.commands`
   - Select the following bot permissions:
     - Read Messages/View Channels
     - Send Messages
     - Read Message History
     - Attach Files
   - Use the generated URL to invite the bot to your server

## Running the Bot

```
go run main.go
```

## Usage

The bot provides the following slash command:

- `/history [number] (save)`: Retrieves the specified number of messages (1-100) and sends them as a text file
  - Example: `/history 50` will retrieve the last 50 messages and send the file to chat
  - Example: `/history 50 save` will also save the file locally
  - The save parameter is optional and defaults to false

The output file will be formatted as:
```
[19 Dec 24 21:48 +0000] Username: Message content
[19 Dec 24 21:47 +0000] AnotherUser: Another message
```

Note: Bot messages and empty messages are automatically filtered out.

## File Structure

```
discord-history-bot/
├── .env                 # Environment variables
├── .gitignore          # Git ignore file
├── go.mod              # Go module file
├── main.go             # Main application file
├── commands/
│   └── history.go      # History command handler
└── utils/
    └── file.go         # File utility functions
```

## Error Handling

The bot includes error handling for common issues:
- Invalid message count requests
- Missing permissions
- File operation errors

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
