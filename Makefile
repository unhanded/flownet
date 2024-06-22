
proto:
	rm -rf ./pkg/fnetpb
	buf generate ./proto/fnet_addr.proto
	buf generate ./proto/fnet_msg.proto

test:
	go test ./internal/fnet -test.v
