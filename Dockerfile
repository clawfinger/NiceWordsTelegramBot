FROM golang:1.22-alpine AS build
WORKDIR /build
COPY . .
RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build

FROM alpine
ARG CHANNEL_ID
ARG BOT_TOKEN

ENV BOT_TOKEN ${BOT_TOKEN}
ENV CHANNEL_ID ${CHANNEL_ID}
ENV MSG_TIMEOUT_MINS 1

COPY --from=build /build/nicebot /nicebot
COPY --from=build /build/nicewords.json /nicewords.json

ENTRYPOINT ["/nicebot"]