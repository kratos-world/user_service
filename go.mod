module github.com/kratos-world/user_service

go 1.16

require (
	github.com/go-kratos/kratos/contrib/registry/etcd/v2 v2.0.0-20220310144244-ac99a5c877c4
	github.com/go-kratos/kratos/v2 v2.2.0
	github.com/google/wire v0.5.0
	go.etcd.io/etcd/client/v3 v3.5.2
	go.opentelemetry.io/otel v1.5.0
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.5.0
	go.opentelemetry.io/otel/sdk v1.5.0
	golang.org/x/net v0.0.0-20220127200216-cd36cc0744dd // indirect
	golang.org/x/sys v0.0.0-20220114195835-da31bd327af9 // indirect
	google.golang.org/genproto v0.0.0-20220126215142-9970aeb2e350
	google.golang.org/grpc v1.44.0
	google.golang.org/protobuf v1.27.1
)
