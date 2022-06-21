MAKE_PATH = .

PROTOC_GRPC = protoc -I . \
			--go_out . \
    		--go-grpc_out . \
			--go-grpc_opt=paths=source_relative,\
			--go-grpc_opt=require_unimplemented_servers=false,\
			./proto/*.proto

PROTOC_GATEWAY = protoc -I . \
				--grpc-gateway_out ./${MAKE_PATH} \
				--grpc-gateway_opt logtostderr=true \
				--grpc-gateway_opt paths=source_relative \
				--grpc-gateway_opt generate_unbound_methods=true \
				./${MAKE_PATH}/proto/*.proto

# TODO: run all services rule
all:
	make db
	make user
	make auth
	make food
	make recipe
	make gateway
	docker compose up

gen-protoc:
# cd ${MAKE_PATH} && $(PROTOC_GRPC)
	$(PROTOC_GRPC)

gen-gateway:
	$(PROTOC_GATEWAY)


.PHONY: gen-gateway gen-protoc
