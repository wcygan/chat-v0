# Quickstart

Generate proto files

```
buf generate proto
```

Run server

```
cd server
go run cmd/main.go
```

Run client

```
cd client
go run cmd/main.go
```

# Chat Application Examples

- [chat-v0](https://github.com/wcygan/chat-v0) - server that iterates over internal state to distribute messages to clients
- [chat-v1](https://github.com/wcygan/chat-v1) - server that uses nats (pub/sub) to distribute messages to clients