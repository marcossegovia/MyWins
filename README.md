[![Build Status](https://travis-ci.org/MarcosSegovia/MyWins.svg?branch=master)](https://travis-ci.org/MarcosSegovia/MyWins)
[![Go Report Card](https://goreportcard.com/badge/github.com/MarcosSegovia/MyWins)](https://goreportcard.com/report/github.com/MarcosSegovia/MyWins)
[![GitHub release](https://img.shields.io/badge/release-v0.1-blue.svg)](https://github.com/MarcosSegovia/MyWins/releases/tag/v0.1)

<p align="center">
	<img alt="MyWins" src="logo.png?raw=true">
</p>
---

MyWins is a service to track your daily routines whatever it is, just to cheer you up in the pursuit of your dreams.

##How it works

1. You should download the app form `url` and create an account. (To be done)
2. After you've already log in, you'll be asked to submit your success or your fail of the day.
3. Keep your wins green!

##Installation

1. You just `go get github.com/MarcosSegovia/MyWins` so you'll get the repository inside your go workspace.
2. Be sure to get [glide](https://github.com/Masterminds/glide) to be able to get dependencies. 
3. If so, then run `glide install`


##How to build and run it locally

1. `go build -o bin/mywins src/*.go`
2. `./bin/mywins`
3. MyWins will be running on 0.0.0.0:8080

> To be able to check the app functionality you would have to provide a mongodb installation and set the environment variables inside `mongo.go` with your own host and port from your mongodb client.

##How to build and run it with Docker locally

1. First uncomment the `Local Development` Dockerfile code and comment the `Production/Staging` one.
2. Build it by the provided Dockerfile `docker build -t marcossegovia/mywins .`
3. Run the docker-compose file like the following `docker-compose up -d mywins`
4. You will see how Docker will get a free port to expose the server: so run a `docker ps` and you'll see the app running on the 0.0.0.0:{port}

##Roadmap

When reaching stability to release the v1.0, MyWins is going to push forward and provide friendly frontend clients to consume from different devices.

iOS app is going to take place to be able to directly get track of your current MyWins streaks !
