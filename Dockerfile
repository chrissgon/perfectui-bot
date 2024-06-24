FROM golang:1.22.4-alpine as build

WORKDIR /bot

COPY . .

RUN GOOS=linux go build .

FROM scratch

WORKDIR /bot

COPY --from=build bot/perfectui-bot .

CMD [ "./perfectui-bot" ]