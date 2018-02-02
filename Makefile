grpc-generate:
	protoc -I./meetup\proto \
		meetup.proto \
		--go_out=plugins=grpc:.