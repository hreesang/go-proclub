# Go ProClub

## Description

Go ProClub is a Discord bot that integrates with the EAFC API to retrieve Pro Club information.

### Features:

- **Event Management**: Discord event handlers are wrapped with `sync.WaitGroup`, ensuring the app waits for all events to complete before shutting down.
- **Simplified Slash Commands**: Easily add and manage slash commands.
- **Streamlined EAFC API Requests**: Construct EAFC API requests effortlessly using helper functions.

## Prerequisites

- Go 1.16 or higher
- Git

## Setup

1. **Clone the repository**:

```sh
git clone https://github.com/yourusername/go-proclub.git
cd go-proclub
```

2. Install dependencies:

```sh
go mod tidy
```

3. Set up environment variables:

```sh
export BOT_TOKEN=your_bot_token
export MONGODB_URI=your_mongodb_uri
export MONGODB_DATABASE=your_mongodb_database
```

For Windows, use:

```sh
set BOT_TOKEN=your_bot_token
set MONGODB_URI=your_mongodb_uri
set MONGODB_DATABASE=your_mongodb_database
```

## Build

To build the project, run:

### On Linux:

```sh
go build -o bin/go-proclub
```

### On Windows:

```sh
go build -o bin\go-proclub.exe
```

## Run

To run the project, execute:

### On Linux:

```sh
./bin/go-proclub
```

### On Windows:

```sh
bin\go-proclub.exe
```

## Contributing

1. Fork the repository
2. Create a new branch (`git checkout -b feature-branch`)
3. Commit your changes (`git commit -m 'Add some feature'`)
4. Push to the branch (`git push origin feature-branch`)
5. Open a pull request

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
