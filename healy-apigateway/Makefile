run:
		swag init -g cmd/main.go -o ./cmd/docs
		go run cmd/main.go
proto:
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/patient/patient.proto
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/doctor/doctor.proto
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/admin/admin.proto
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/chat/chat.proto
swag:
		swag init -g cmd/main.go -o ./cmd/docs