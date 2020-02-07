FROM golang:alpine as builder

COPY [ ".", "/go/src/github.com/deletescape/noise" ]

WORKDIR /go/src/github.com/deletescape/noise

RUN CGO_ENABLED=0 GOOS=linux go build -a -ldflags '-extldflags "-static"' ./cmd/blast

#----------#

FROM scratch

COPY --from=builder [ "/etc/ssl/certs/ca-certificates.crt", "/etc/ssl/certs/ca-certificates.crt" ]
COPY --from=builder [ "/go/src/github.com/deletescape/noise/blast", "/blast" ]

CMD [ "/blast" ]
