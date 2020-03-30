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

Depending on how you installed the application, you can run either `bin/world-bank-grpc` (if built) or `world-bank-grpc` (if installed).

To see available commands, run this command:

```shell script
bin/world-bank-grpc --help
```

You can also view help for a sub-command by adding `help` after it. E.g.:

```shell script
bin/world-bank-grpc client --help
```

#### Starting the Server

To start the gRPC server, run this command:

```shell script
bin/world-bank-grpc server
```

The server listens on port `50001` by default but this can be overridden by the `PORT` environment variable.

#### Running the Client

To view available client commands, run this:

```shell script
bin/world-bank-grpc client --help
```

Example Usages:

1. List all regions

```shell script
bin/world-bank-grpc client regions
```

2. List countries in Sub-Saharan Africa (ID: SSF)

```shell script
bin/world-bank-grpc client countries --region SSF
```

3. List High Income (ID: HIC) countries with IBRD lending type (ID: IBD)

```shell script
bin/world-bank-grpc client countries --lending-type IBD --income-level HIC
```

4. Get details about Nigeria

```shell script
bin/world-bank-grpc client country --id NGA
```

#### Environment Variables

|     | Variable | Default         | Components |
| --- | -------- | --------------- | ---------- |
| 1   | PORT     | 50001           | server     |
| 1   | HOST     | localhost:50001 | client     |
