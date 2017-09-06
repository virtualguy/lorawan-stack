# Copyright © 2017 The Things Network Foundation, distributed under the MIT license (see LICENSE file)

# Go
GO_PROTO_TYPES = any duration empty struct timestamp
GO_PROTO_TYPE_CONVERSIONS = $(subst $(SPACE),$(COMMA),$(foreach type,$(GO_PROTO_TYPES),Mgoogle/protobuf/$(type).proto=$(GOGO_REPO)/types))
GO_PROTOC_FLAGS ?= \
	--gogottn_out=plugins=grpc,$(GO_PROTO_TYPE_CONVERSIONS):$(GO_PATH)/src \
	--grpc-gateway_out=:$(GO_PATH)/src

go.protos: $(wildcard $(PROTO_DIR)/*.proto)
	$(PROTOC) $(GO_PROTOC_FLAGS) $(PROTO_DIR)/*.proto
	$(MAKE_DIR)/protos/fix-grpc-gateway-names.sh $(PROTO_DIR)

# vim: ft=make
