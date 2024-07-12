run:
		go run cmd/main.go
proto:
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/admin/admin.proto
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/doctor/doctor.proto
		protoc --go_out=. --go-grpc_out=. ./pkg/pb/patient/patient.proto