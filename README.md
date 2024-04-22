# URLed

URLed is an open-source lightweight URL shortener server written in Go

## Tech Stack

- Go
- Gin (HTTP web framework)
- Gorm (ORM)
- SQLite

## Getting Started

### Minimum Requirements

URLed is available for these operating systems/architectures:

- Linux (amd64)
- Linux (386)
- Linux (arm64)
- Linux (armv7)
- Linux (armv6)
- Linux (armv5)
- FreeBSD (amd64)
- FreeBSD (386)
- FreeBSD (armhf)
- MacOS (x86_64)
- MacOS (arm64)
- Windows (arm64)

URLed uses <10MB of RAM when running (in my testing), and the binary is <20MB in size.

I want to make this project available to run on whatever architecture you have, so if you don't see your architecture
listed above, please open an issue and I will try to add support for it.

### Installation

Set up an env file in the same directory as your binary with the following contents:

```BASE_URL=[your base URL]```

Note: Make sure you include the protocol (http/https) in the base URL

#### Stable

Download the binary for your architecture from the [releases page](https://github.com/masoncfrancis/urled/releases)
and put it in the same directory as your `.env` file.

#### Unstable

If you want to run the latest version of URLed, you can clone the repository and build the binary yourself.

```bash
go build ./cmd/urled
```

### Usage

To start the server, run the following command:

```bash
$ ./[your URLed executable] -server
```

The server will run on port `4567`. Your shortened URLs will only be available at your base URL if 
you configure traffic at that URL to be forwarded to the server at port `4567`.

To edit URls being shortened, see the flags below.

#### Editing Shortened URLs

- `-add [long url]` - Shorten a new URL
- `-remove [short url suffix]` - Remove a URL
- `-remove-long [long url]` - Remove all instances of a URL
- `-list` - List all URLs being shortened

Note: You can modify URLs being served by an already running server by opening a separate terminal window and running
the binary using these flags

## Future Plans

Currently the project is in it's infancy, but I am planning on adding these features in the future:

- Add a web interface for managing URLs
- Runnable in a Docker container
- Add support for custom short URLs
- Add support for URL expiration
- Add support for URL analytics
- Install as a background service that runs on boot

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
