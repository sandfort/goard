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

## The `db` package
The database package requires a running SQL database server. I have only tested
with MySQL 8.0, but others may work. The package also includes migrations
intended for use with Flyway.

The database can be used by adding the `-db` flag when running the app.
Otherwise it will default to in-memory storage. Using this option requires the
following environment variables to be set:
- `DB_USER`
- `DB_PASSWORD`
- `DB_NAME`

### Migrations
Running the migrations requires the Flyway CLI. I have tested with Flyway
Community Edition 5.2.0. To run the migrations, you'll need to give Flyway the
credentials for your database. This can be done by command-line arguments,
config files, or by environment variables. I have provided an example Flyway
configuration file that you can fill in with your own credentials. When you
have all that sorted, you can run `flyway migrate` and you should be set.

### Testing
Similarly, credentials will have to be provided to the `db` package to run
the tests. I recommend having a separate test database (and yes, this means
you would have to manage those credentials for migrations as well). The tests
rely on their own set of credentials which must be provided via the following
environment variables:
- `TEST_DB_USER`
- `TEST_DB_PASSWORD`
- `TEST_DB_NAME`