FROM gcr.io/distroless/static

ADD bin/world-bank-grpc /world-bank-grpc

ENTRYPOINT ["/world-bank-grpc", "server"]
