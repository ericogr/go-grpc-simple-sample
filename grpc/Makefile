PROTO_DIR=proto
OUT_DIR=chatpb
PROTO_FILE=$(PROTO_DIR)/chat.proto

PROTOC=protoc
PROTOC_GEN_GO=$(shell which protoc-gen-go || echo "$(shell go env GOPATH)/bin/protoc-gen-go")
PROTOC_GEN_GO_GRPC=$(shell which protoc-gen-go-grpc || echo "$(shell go env GOPATH)/bin/protoc-gen-go-grpc")

check-tools:
	@command -v $(PROTOC) >/dev/null 2>&1 || { echo "Erro: protoc não encontrado. Instale o protocolo buffer compiler."; exit 1; }
	@command -v $(PROTOC_GEN_GO) >/dev/null 2>&1 || { echo "Erro: protoc-gen-go não encontrado. Execute 'make install'."; exit 1; }
	@command -v $(PROTOC_GEN_GO_GRPC) >/dev/null 2>&1 || { echo "Erro: protoc-gen-go-grpc não encontrado. Execute 'make install'."; exit 1; }

generate: check-tools
	@$(PROTOC) \
		-I=$(PROTO_DIR) \
		--go_out=$(OUT_DIR) --go_opt=paths=source_relative \
		--go-grpc_out=$(OUT_DIR) --go-grpc_opt=paths=source_relative \
		$(PROTO_FILE)

install:
	@sudo apt install -y protobuf-compiler
	@go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	@go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

clean:
	@rm -rf $(OUT_DIR)/*.go

.PHONY: all generate install clean check-tools