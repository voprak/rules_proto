load("@io_bazel_rules_go//go:def.bzl", "go_binary", "go_library")

go_library(
    name = "protoc-gen-sql_lib",
    srcs = ["main.go"],
    importpath = "github.com/pgga-es/lumada/data-access/generator/cmd/protoc-gen-sql",
    visibility = ["//visibility:private"],
    deps = ["@org_golang_google_protobuf//compiler/protogen"],
)

go_binary(
    name = "protoc-gen-sql",
    embed = [":protoc-gen-sql_lib"],
    visibility = ["//visibility:public"],
)
