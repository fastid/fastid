# syntax=docker/dockerfile:1
FROM --platform=$BUILDPLATFORM golang:alpine as builder

ARG TARGETPLATFORM
ARG BUILDPLATFORM
ARG TARGETOS
ARG TARGETARCH

RUN echo "TARGETPLATFORM $TARGETPLATFORM" > /log
RUN echo "BUILDPLATFORM $BUILDPLATFORM" > /log
RUN echo "TARGETOS $TARGETOS" > /log
RUN echo "TARGETARCH $TARGETARCH" > /log

RUN apk update

WORKDIR $GOPATH/fastid/
COPY . .
COPY configs/fastid.yml /fastid.yml

RUN go mod download

RUN CGO_ENABLED=0 GOOS=$TARGETOS GOARCH=$TARGETARCH go build -ldflags="-s -w" -o /fastid cmd/fastid.go

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder fastid /fastid/fastid
COPY --from=builder fastid.yml /fastid/fastid.yml
ENTRYPOINT ["/fastid/fastid", "-config", "/fastid/fastid.yml", "-run"]
