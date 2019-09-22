module github.com/bitum-project/bitumwallet

require (
	github.com/boltdb/bolt v1.3.1
	github.com/decred/slog v1.0.0
	github.com/golang/protobuf v1.3.1
	github.com/gorilla/websocket v1.4.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/jrick/bitset v1.0.0
	github.com/jrick/logrotate v1.0.0
	golang.org/x/crypto v0.0.0-20190605123033-f99c8df09eb5
	golang.org/x/net v0.0.0-20190607181551-461777fb6f67
	golang.org/x/sync v0.0.0-20190423024810-112230192c58
	google.golang.org/genproto v0.0.0-20190219182410-082222b4a5c5 // indirect
	google.golang.org/grpc v1.18.0
)

replace github.com/bitum-project/bitumd => ../bitumd
