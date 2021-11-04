load(
    "//protobuf:deps.bzl",
    "protobuf",
)

load(
    "//:deps.bzl",
    "io_bazel_rules_go",
)

def sql_proto_compile(**kwargs):
    protobuf(**kwargs)
    io_bazel_rules_go(**kwargs)
