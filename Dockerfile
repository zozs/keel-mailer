# build
FROM golang:1.17 AS build

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY *.go .
RUN go build -o /keel-mailer

# deploy
FROM gcr.io/distroless/base-debian10
COPY --from=build /keel-mailer /
USER nonroot:nonroot
CMD ["/keel-mailer"]
