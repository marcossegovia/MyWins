[![Build Status](https://travis-ci.org/MarcosSegovia/MyWins.svg?branch=master)](https://travis-ci.org/MarcosSegovia/MyWins)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarcosSegovia/MyWins)](https://goreportcard.com/report/github.com/MarcosSegovia/MyWins)

<p align="center">
	<img alt="MyWins" src="logo.png?raw=true">
</p>
---

MyWins is a service to track your daily routines whatever it is, just to cheer you up in the pursuit of your dreams.

##How it works

1. You should download the app form `url` and create an account.
2. After you've already log in, you'll be asked to submit your success or your fail of the day.
3. Keep your wins green!

##Installation

You just `go get github.com/MarcosSegovia/MyWins` so you'll get the repository inside your go workspace, be sure to get [glide](https://github.com/Masterminds/glide) to be able to get dependencies. If so, then run `glide install`

##How to build and run it locally

1. `go build -o bin/mywins src/*.go`
2. `./bin/mywins`
3. MyWins will be running on 0.0.0.0:8080

##How to build and run it with Docker locally

1. First uncomment the `Local Development` Dockerfile code and comment the `Production/Staging` one.
2. Build it by the provided Dockerfile `docker build -t marcossegovia/mywins .`
3. Run it by executing `docker run -d -P marcossegovia/mywins`
4. You will see how Docker will get a free port to expose the server: so run a `docker ps` and you'll see the app running on the 0.0.0.0:{port}

