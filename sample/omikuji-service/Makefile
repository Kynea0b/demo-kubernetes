.PHONY: proto
proto:
	protoc --proto_path=../ \
		--go_out=./pkg/grpc --go_opt=paths=source_relative \
		--go-grpc_out=./pkg/grpc --go-grpc_opt=paths=source_relative \
		../proto/*.proto

.PHONY: draw
draw:
	grpcurl --plaintext localhost:8080 omikuji.OmikujiService.Draw

IMAGE_NAME="omikuji-service:1.0.0"

.PHONY: build
build:
	@if docker image inspect $(IMAGE_NAME) > /dev/null 2>&1; then \
		echo "Image $(IMAGE_NAME) exists. Deleting...";\
		docker image rm $(IMAGE_NAME);\
	fi
	docker build -t $(IMAGE_NAME) .

dockerbuild:
	docker build . -t ${IMAGE_NAME}
