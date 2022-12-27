.PHONY: build
build:
	bazel build //cmd/istream

.PHONY: test
test:
	bazel test //...

.PHONY: protos
protos:
	bazel run //build/stack/inputstream/v1beta1:v1beta1_go_compiled_sources.update
	mv build/stack/inputstream/v1beta1/build/stack/inputstream/v1beta1/*.pb.go build/stack/inputstream/v1beta1
	rm -rf build/stack/inputstream/v1beta1/build

	bazel run //build/stack/auth/v1beta1:v1beta1_go_compiled_sources.update
	mv build/stack/auth/v1beta1/build/stack/auth/v1beta1/*.pb.go build/stack/auth/v1beta1
	rm -rf build/stack/auth/v1beta1/build

.PHONY: mocks
mocks:
	mockery --output mocks --dir=build/stack/auth/v1beta1 --name=AuthServiceClient 
	mockery --output mocks --dir=build/stack/inputstream/v1beta1 --name=UsersClient 
	mockery --output mocks --dir=build/stack/inputstream/v1beta1 --name=InputsClient 

.PHONY: tidy
tidy:
	bazel run @go_sdk//:bin/go -- mod tidy
	bazel run //:update_go_repositories
	bazel run //:gazelle

