# URLed

URLed is an open-source self-hosted lightweight URL shortener server written in Go

## What URLed is built on

- Go 1.22
    - Fiber (HTTP web framework)
    - Gorm (ORM)
- SQLite

## Why did I build this?

I built URLed because I was sick of so many other self-hosted URL shorteners being based on PHP and being hard to set
up. Even those that had Docker images available didn't seem to offer one compatible with Raspberry Pi 2 and Zero, which use ARMv6. I wanted a URL shortener 
I could run without much setup and self-host on super minimal hardware.

## Getting Started

### Minimum Requirements

URLed is available for these operating systems/architectures:

- Linux (386)
- Linux (amd64)
- Linux (arm64)
- Linux (armv7)
- Linux (armv6)
- Linux (armv5)
- FreeBSD (386)
- FreeBSD (amd64)
- FreeBSD (armhf)
- FreeBSD (arm64)
- MacOS (amd64)
- MacOS (arm64)
- Windows (386)
- Windows (amd64)
- Windows (arm64)

URLed uses ~<20MB of RAM when running (in my testing), and the binary is <20MB in size.

I want to make this project available to run on whatever hardware you have, so if you don't see your architecture
listed above, please open an issue and I will try to add support for it.

### Installation

Create a file called `.env` in the same directory as where your binary will go with the following contents:

```BASE_URL=[your base URL]```

Note: Make sure you include the protocol (http/https) in the base URL

#### Stable

Download the binary for your architecture from the [releases page](https://github.com/masoncfrancis/urled/releases)
and put it in the same directory as your `.env` file.

If you have an issue in UNIX-style OSes with an "access denied" error when you run it, run the following command to add execution permission
and then try again:

```bash
chmod +x ./[your URLed executable]
```

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

#### Adding/Removing Shortened URLs

- `-add [long url]` - Shorten a new URL
- `-remove [short url suffix]` - Remove a URL
- `-remove-long [long url]` - Remove all instances of a URL
- `-list` - List all URLs being shortened

Note: You can modify URLs being served by an already running server by opening a new terminal window and running
the binary separately using these flags

## Future Plans

Currently the project is in it's infancy, but I am planning on adding these features in the future:

- Add a web interface for managing URLs
- Docker image
- Add support for custom short URLs
- Add support for URL expiration
- Add support for URL analytics
- Install as a background service that runs on boot
- Automatic updates

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
