module main

go 1.19

require (
	github.com/golang/glog v1.1.1
	github.com/google/go-attestation v0.4.4-0.20220404204839-8820d49b18d9
	github.com/google/go-tpm v0.3.3
	github.com/google/go-tpm-tools v0.3.12
	github.com/gorilla/mux v1.7.3
	github.com/salrashid123/tls_ak/verifier v0.0.0
	golang.org/x/net v0.10.0
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c
	google.golang.org/grpc v1.55.0
)

require (
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/certificate-transparency-go v1.1.2 // indirect
	github.com/google/go-sev-guest v0.6.1 // indirect
	github.com/google/go-tspi v0.2.1-0.20190423175329-115dea689aad // indirect
	github.com/google/logger v1.1.1 // indirect
	github.com/google/uuid v1.3.0 // indirect
	github.com/pborman/uuid v1.2.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	golang.org/x/crypto v0.0.0-20220525230936-793ad666bf5e // indirect
	golang.org/x/sys v0.8.0 // indirect
	golang.org/x/text v0.9.0 // indirect
	google.golang.org/genproto v0.0.0-20230306155012-7f2fa6fef1f4 // indirect
	google.golang.org/protobuf v1.30.0 // indirect
)

replace github.com/salrashid123/tls_ak/verifier => ../verifier
