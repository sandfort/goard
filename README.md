# goard
[![Build Status](https://travis-ci.org/sandfort/goard.svg?branch=master)](https://travis-ci.org/sandfort/goard)

## Three goals of this project

### 1. To improve my understanding of SOLID
I have worked on web apps that adhere to SOLID principles in familiar languages,
but these are usually the result of picking a design (e.g. hexagonal) that
happens to be based on one or more SOLID principles. I'd like to set pre-made
designs aside and work purely from SOLID principles. I'm using an unfamiliar
language to force myself to think about and understand how to implement them.

### 2. To make a product I would use
I've always loved forums and message boards. They were the first way I
interacted with strangers on the internet and eventually made friends. I think
this has more to do with community than software, but I have opinions about
what I do and don't want from a message board server, so I think I will have a
clear product direction.

### 3. To experiment with modularity
A lot of apps claim to be modular, plug-and-play, etc. My goal with this app is
to be able to swap out its web layer for an API, and its in-memory storage
module for a database-backed storage module.

## How to use
These instructions assume your `pwd` is `$GOPATH/src/github.com/sandfort/goard`.
Make sure the environment variable `PORT` is set to a number. I use 8080.
To run tests, run `go test ./...`
To build, run `go build`
To start the app, run `./goard` (after building)
To view the app, open a browser to `localhost:8080/threads`
