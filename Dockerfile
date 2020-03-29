FROM gcr.io/distroless/static

ARG BIN

ADD bin/$BIN /$BIN

ENTRYPOINT ["/{$BIN}"]
