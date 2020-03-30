### World Bank gRPC

[![Build Status](https://travis-ci.com/eloyekunle/world-bank-grpc.svg?token=gE7ya6SNZs6o39nbGJAd&branch=master)](https://travis-ci.com/eloyekunle/world-bank-grpc)
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
| 2   | HOST     | localhost:50001 | client     |

#### Tests

We use [mocking](./pkg/mock_worldbank) to test client-side logic without the overhead of connecting to a real server.
Mocking enables users to write light-weight unit tests to check functionalities on client-side without invoking RPC calls to a server.

When tests have passed, the application is built and automatically deployed to [Docker Hub](https://hub.docker.com/repository/docker/playmice/world-bank-grpc).

#### Kubernetes

Kubernetes manifests have been provided in [kubernetes](./kubernetes).

To make the service available to an external client
from the cluster, the user can setup an Ingress to the provided service with the `NGINX` provider.

Optionally, the user can also provision TLS encryption. In this case, the client will have to be updated to trust the
certificate authority (CA).

#### Event Store

The benefit of having an event store will be to ensure that every change to the state of the application is captured as an event,
and that the events are stored in the same sequence they were applied for the same lifetime as the application itself.

This application right now has immutable state, because the countries, regions, income levels and lending types do not
change over the course of the application's lifetime.

However, if the application were to have mutable state, e.g. ability to edit countries etc, then the first step in integrating
an event store is to define possible events. For example: DELETE_COUNTRY, ADD_INCOME_LEVEL, EDIT_REGION etc. Next step is to define
the events object themselves, as well as a processor for each event.

#### Cloud-Native Understanding

Some key attributes of cloud-native applications can be found on [TheNewStack.io](https://thenewstack.io/10-key-attributes-of-cloud-native-applications/).
They include services managed through agile DevOps processes, centered around APIs for interaction, deployment on cloud infrastructure, etc.

This application successfully implements cloud native understanding.

#### 12factor app best practices

[12-factor app best practices](https://12factor.net/) include - storing config in the environment, one codebase tracked
in version control, separate build and run stages, services exposed via port binding, logs as event streams, etc.

This application successfully implements 12 factor app best practices.
