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
	golang.org/x/net v0.0.0-20220225172249-27dd8689420f // indirect
	golang.org/x/sys v0.0.0-20220317061510-51cd9980dadf // indirect
	google.golang.org/genproto v0.0.0-20220314164441-57ef72a4c106
	google.golang.org/grpc v1.45.0
	google.golang.org/protobuf v1.27.1
)
