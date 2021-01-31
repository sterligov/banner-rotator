//go:generate protoc -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.0.1/third_party/googleapis -I ./../../../ --go_out ./pb --go-grpc_out ./pb ./../../../api/banner_service.proto
//go:generate protoc -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.0.1/third_party/googleapis -I ./../../../ --grpc-gateway_out ./pb --grpc-gateway_opt logtostderr=true --grpc-gateway_opt generate_unbound_methods=true ./../../../api/banner_service.proto

//go:generate protoc -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.0.1/third_party/googleapis -I ./../../../ --go_out ./pb --go-grpc_out ./pb ./../../../api/group_service.proto
//go:generate protoc -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.0.1/third_party/googleapis -I ./../../../ --grpc-gateway_out ./pb --grpc-gateway_opt logtostderr=true --grpc-gateway_opt generate_unbound_methods=true ./../../../api/group_service.proto

//go:generate protoc -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.0.1/third_party/googleapis -I ./../../../ --go_out ./pb --go-grpc_out ./pb ./../../../api/slot_service.proto
//go:generate protoc -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.0.1/third_party/googleapis -I ./../../../ --grpc-gateway_out ./pb --grpc-gateway_opt logtostderr=true --grpc-gateway_opt generate_unbound_methods=true ./../../../api/slot_service.proto

//go:generate protoc -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.0.1/third_party/googleapis -I ./../../../ --go_out ./pb --go-grpc_out ./pb ./../../../api/health_service.proto
//go:generate protoc -I${GOPATH}/pkg/mod/github.com/grpc-ecosystem/grpc-gateway/v2@v2.0.1/third_party/googleapis -I ./../../../ --grpc-gateway_out ./pb --grpc-gateway_opt logtostderr=true --grpc-gateway_opt generate_unbound_methods=true ./../../../api/health_service.proto
package grpc
