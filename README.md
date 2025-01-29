# Go Web Project Demo

This is a web application built using the Go programming language.

## Table of Contents

- [Installation](#installation)
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

## Usage

To generate twirp framework code

```sh
protoc --go_out=. --twirp_out=. rpc/user/service.proto
```

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
- Job: cron(waitting)

## License

This project is just a demo.