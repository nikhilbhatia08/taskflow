PROTOC = protoc
PROTOC_GEN_GO = protoc-gen-go
PROTOC_GEN_GO_GRPC = protoc-gen-go-grpc

PROTO_SRC_DIR = pkg/grpc
PROTO_OUT_DIR = generatedproto
PROTO_OUT_SDK_DIR = taskflow-gosdk/generatedproto

PROTO_FILES := $(wildcard $(PROTO_SRC_DIR)/*.proto)

$(PROTO_OUT_DIR):
	mkdir -p $(PROTO_OUT_DIR)
	mkdir -p $(PROTO_OUT_SDK_DIR)

generate: $(PROTO_OUT_DIR)
	$(PROTOC) --proto_path=$(PROTO_SRC_DIR) --go_out=$(PROTO_OUT_DIR) --go-grpc_out=$(PROTO_OUT_DIR) $(PROTO_FILES)
	$(PROTOC) --proto_path=$(PROTO_SRC_DIR) --go_out=$(PROTO_OUT_SDK_DIR) --go-grpc_out=$(PROTO_OUT_SDK_DIR) $(PROTO_FILES)

clean:
	rm -rf $(PROTO_OUT_DIR)/*
	rm -rf $(PROTO_OUT_SDK_DIR)/*

all: clean generate

.PHONY: generate clean all
