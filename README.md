# Go Web Project Demo

This is a web application built using the Go programming language.

## Table of Contents

- [Installation](#installation)
- [Tips](#tips)
- [Usage](#usage)
- [Features](#features)
- [License](#license)

## Installation

To install the project, clone the repository and navigate to the project directory:

```sh
git clone git@github.com:sai-shiyuu/go-twirp-demo.git
cd go-twirp-demo
```

First, download and install Golang
https://go.dev/dl/

extract
```sh
sudo tar -C /usr/local -xzf go1.23.5.linux-amd64.tar.gz
```
Add to system env
```sh
sudo nvim /etc/profile
```
add the following script
```
export GOBIN=/home/yourname/go
export PATH=$PATH:/usr/local/go/bin:$GOBIN
```
reboot system

Then install the protobuf and twirp
```sh
sudo pacman -Syu protobuf
go install github.com/twitchtv/twirp/protoc-gen-twirp@latest
go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
```
Then, build the project:

```sh
go build
```
## Tips

To generate twirp framework code

```sh
protoc --go_out=. --twirp_out=. rpc/user/service.proto
```

Cron Common Examples:

- */1 * * * *: Every minute
- 0 0 * * *: Every day at midnight
- 0 10 * * MON-FRI: At 10:00 AM on weekdays
- 0 0 1 * *: On the 1st of every month at midnight
- 0 0 * * 0: Every Sunday at midnight

## Usage

To run the application, execute the following command:

```sh
go run .
```

Open your web browser and navigate to `http://localhost:8080` to see the application in action.

API test file
```
project/common/request.http
```
## Features

- Protocol: twirp rpc
- Database connection: MariaDB
- Cache: Redis
- Logging: logrus
- Password encrypt: sha256
- Job: cron

## License

This project is just a demo.