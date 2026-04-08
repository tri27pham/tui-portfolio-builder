FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /portfolio .

FROM alpine:3.20
RUN apk add --no-cache openssh-keygen
COPY --from=build /portfolio /portfolio
RUN mkdir -p /app/.ssh && ssh-keygen -t ed25519 -f /app/.ssh/id_ed25519 -N ""
WORKDIR /app
EXPOSE 2222
ENV SSH_MODE=1
CMD ["/portfolio"]
