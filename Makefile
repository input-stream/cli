.PHONY: build
build:
	bazel build //cmd/istream

.PHONY: test
test:
	bazel test //...

.PHONY: protos
protos:
	bazel run //stream/input/v1beta1:v1beta1_go_compiled_sources.update
	mv stream/input/v1beta1/stream/input/v1beta1/*.pb.go stream/input/v1beta1
	rm -rf stream/input/v1beta1/stream

.PHONY: tidy
tidy:
	bazel run @go_sdk//:bin/go -- mod tidy
	bazel run //:update_go_repositories
	bazel run //:gazelle

