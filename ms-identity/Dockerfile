FROM golang:1.18-alpine as build_stage

WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 go build -v -o build

FROM alpine:3.14

COPY --from=build_stage /app/build /app/build

CMD /app/build
