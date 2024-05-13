module main

go 1.21

require (
	github.com/golang/glog v1.2.1
	github.com/google/go-attestation v0.5.1
	github.com/google/go-tpm-tools v0.4.4
	github.com/google/uuid v1.6.0
	github.com/gorilla/mux v1.8.1
	github.com/salrashid123/gcp-tpm/parser v0.0.0-20220930151352-4675346f7ef3
	github.com/salrashid123/tls_ak/verifier v0.0.0
	golang.org/x/net v0.25.0
	golang.org/x/sync v0.7.0
	google.golang.org/grpc v1.63.2
)

require (
	github.com/golang/protobuf v1.5.4 // indirect
	github.com/google/certificate-transparency-go v1.1.8 // indirect
	github.com/google/go-configfs-tsm v0.2.2 // indirect
	github.com/google/go-sev-guest v0.11.1 // indirect
	github.com/google/go-tdx-guest v0.3.1 // indirect
	github.com/google/go-tpm v0.9.0 // indirect
	github.com/google/go-tspi v0.3.0 // indirect
	github.com/google/logger v1.1.1 // indirect
	github.com/pborman/uuid v1.2.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	go.uber.org/multierr v1.11.0 // indirect
	golang.org/x/crypto v0.23.0 // indirect
	golang.org/x/sys v0.20.0 // indirect
	golang.org/x/text v0.15.0 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20240509183442-62759503f434 // indirect
	google.golang.org/protobuf v1.34.1 // indirect
)

replace github.com/salrashid123/tls_ak/verifier => ./verifier