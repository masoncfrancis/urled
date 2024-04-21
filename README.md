# URLed

URLed is a lightweight URL shortener server written in Go

## Tech Stack

- Go
- SQLite
- Gin (HTTP web framework)
- Gorm (ORM)

## Getting Started

### Prerequisites

None. Just download the binary and run it.

### Installation

Download the binary for your architecture from the [releases page](https://github.com/masoncfrancis/urled/releases)
and put it wherever you want.

### Usage

To start the server, run the following command, replacing `urled` with whatever you named the binary:

```bash
$ ./urled -server
```

To edit URls being shortened, see the flags below.

#### Editing Shortened URLs

- `-add [long url]` - Shorten a new URL
- `-remove [short url suffix]` - Remove a URL
- `-remove-long [long url]` - Remove all instances of a URL
- `-list` - List all URLs being shortened

Note: You can modify URLs being served by an already running server by opening a separate terminal window and running
the binary using these flags
 
