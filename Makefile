.PHONY: proto
proto:
	docker build -t example-proto:latest .
	docker run --rm --name=example-proto -v `pwd`:/build -e module_name=github.com/skunkie/grpc-gateway-example -e out_path=./internal example-proto:latest
