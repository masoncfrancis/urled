# URLed

URLed is an open-source lightweight URL shortener server written in Go

## Tech Stack

- Go
- - Gin (HTTP web framework)
- Gorm (ORM)
- SQLite

## Getting Started

### Prerequisites

Set up an env file in the same directory as your binary with the following contents:

```BASE_URL=[your base URL]```

Note: Make sure you include the protocol (http/https) in the base URL

### Installation

Download the binary for your architecture from the [releases page](https://github.com/masoncfrancis/urled/releases)
and put it wherever you want.

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
