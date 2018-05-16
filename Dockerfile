# Build Stage
FROM golang:alpine AS build-env
ADD . /src
RUN apk update && apk add --no-cache git ca-certificates apache2-utils
RUN cd /src && go get -d -v && go build -o crt_sheep

# Final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /src/crt_sheep /app/
COPY --from=build-env /etc/ssl/certs /etc/ssl/certs
ENTRYPOINT ["./crt_sheep"]
