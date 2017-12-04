FROM golang:alpine AS build
WORKDIR /app
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -v

FROM scratch
COPY --from=build /app/app /
ENTRYPOINT ["/app"]
