# Discord History Bot

A simple Discord bot written in Go that allows users to export chat history to text files using slash commands.

## Features

- Export recent chat messages to text files
- Configurable number of messages (1-100)
- Filters out bot messages and empty messages
- Optional local file saving
- Timestamps included with each message
- Uses Discord slash commands

## Prerequisites

- Go 1.21 or higher
- A Discord Bot Token
- A Discord server with proper permissions

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
TEST_GUILD_ID=your_test_server_id_here  # Optional, for development
```

2. Set up your Discord bot:
   - Go to [Discord Developer Portal](https://discord.com/developers/applications)
   - Create a new application
   - Go to the "Bot" section and create a bot
   - Under "Privileged Gateway Intents", enable:
     - Message Content Intent
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
     - Use Slash Commands
   - The total permissions integer should be: `117760`

3. Development Setup (Optional):
   - Enable Developer Mode in Discord (User Settings -> App Settings -> Advanced -> Developer Mode)
   - Right-click your test server and "Copy ID"
   - Add this ID to your `.env` file as `TEST_GUILD_ID`
   - This enables instant command updates during development

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

## Troubleshooting

1. "Missing Access" error:
   - Ensure the bot has all required permissions listed above
   - Check if the bot can see the channel you're using
   - Verify the Message Content Intent is enabled

2. Commands not appearing:
   - If using global commands (no TEST_GUILD_ID), wait up to an hour
   - For test servers, commands should appear instantly
   - Make sure the bot has the "Use Slash Commands" permission

3. Can't save files:
   - Check if the bot has "Attach Files" permission
   - Ensure the bot has write permissions in its directory

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

## Contributing

1. Fork the repository
2. Create your feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add some amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## License

This project is licensed under the MIT License - see the LICENSE file for details.
