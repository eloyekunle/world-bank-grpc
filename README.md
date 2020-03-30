### World Bank gRPC

[![Build Status](https://travis-ci.com/eloyekunle/world-bank-grpc.svg?token=gE7ya6SNZs6o39nbGJAd&branch=master)](https://travis-ci.com/eloyekunle/world-bank-grpc)
[![GoDoc](https://godoc.org/github.com/eloyekunle/world-bank-grpc?status.svg)](https://godoc.org/github.com/eloyekunle/world-bank-grpc)
[![](https://images.microbadger.com/badges/image/playmice/world-bank-grpc.svg)](https://microbadger.com/images/playmice/world-bank-grpc "Get your own image badge on microbadger.com")
[![Go Report Card](https://goreportcard.com/badge/github.com/eloyekunle/world-bank-grpc)](https://goreportcard.com/report/github.com/eloyekunle/world-bank-grpc)

This is a simple application to get country data from the [World Bank API](https://datahelpdesk.worldbank.org/knowledgebase/articles/889386-developer-information-overview).
Countries can be filtered based on their:

1. Income Levels
2. Regions
3. Lending Types

#### Building

To build `world-bank-grpc`, run the following `make` command:

```shell script
make build
```

The binary will be compiled and placed in `bin/world-bank-grpc`.

You could also choose to install it for the user. Run the following `make` command:

```shell script
make install
```

This places the binary in `$GOPATH/bin/world-bank-grpc`. If `$GOPATH/bin` is in your `PATH` already, you can simply run
`world-bank-grpc` to use the application.

#### Starting the Server

#### Running the Client
