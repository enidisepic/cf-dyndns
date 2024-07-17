FROM golang:1.22.5-alpine AS build
WORKDIR /build

ARG MAKE_VERSION=4.4.1-r2
RUN apk add --no-cache make=${MAKE_VERSION}

COPY . .
RUN OUTPUT_DIRECTORY=. make build

FROM alpine:3
WORKDIR /app

COPY --from=build /build/cf-dyndns cf-dyndns

HEALTHCHECK NONE
CMD ["./cf-dyndns"]
