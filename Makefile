
# Copyright 2016 Michal Witkowski. All Rights Reserved.
# See LICENSE for licensing terms.

install_deps:
	go install -v github.com/gogo/protobuf/protoc-gen-gogo
	go install -v github.com/golang/protobuf/protoc-gen-go

install:
	go install -v ./...

regenerate_test_gogo:
	@echo "Regenerating test .proto files with gogo imports"
	(cd test && \
	protoc \
		--proto_path=. \
		--proto_path=.. \
		--proto_path=../vendor/ \
		--gogo_out=gogo \
		--govalidators_out=gogoimport=true:gogo *.proto)

regenerate_test_golang:
	@echo "--- Regenerating test .proto files with golang imports"
	(cd test && \
	 protoc  \
		--proto_path=. \
		--proto_path=.. \
		--proto_path=../vendor/ \
		--go_out=golang \
		--govalidators_out=golang *.proto)

regenerate_example: install
	@echo "--- Regenerating example directory"
	(cd examples && \
	 protoc  \
		--proto_path=. \
		--proto_path=.. \
		--go_out=. \
		--govalidators_out=. *.proto)

test: regenerate install regenerate_test_gogo regenerate_test_golang
	@echo "Running tests"
	go test -v ./...

regenerate: install_deps
	@echo "--- Regenerating validator.proto"
	(protoc \
		--proto_path=. \
		--proto_path=vendor/github.com/gogo/protobuf/ \
		--gogo_out=Mgoogle/protobuf/descriptor.proto=github.com/gogo/protobuf/protoc-gen-gogo/descriptor:. \
		validator.proto )

clean:
	(find examples test -name "*.pb.go" -delete)
	rm -f validator.pb.go
