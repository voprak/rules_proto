load(
    "@build_stack_rules_proto//rules:nodejs_grpc_compile.bzl",
    "nodejs_grpc_compile",
)
load(
    "@build_stack_rules_proto//rules:proto_compile_js_library.bzl",
    "proto_compile_js_library",
)

GRPC_DEPS = [
    "@google_protobuf_node_modules//google-protobuf",
    "@grpc_js_node_modules//@grpc/grpc-js",
]

def nodejs_grpc_library(**kwargs):
    name_pb = kwargs.get("name") + "_pb"
    js_deps = kwargs.pop("js_deps", GRPC_DEPS)

    nodejs_grpc_compile(
        name = name_pb,
        **{k: v for (k, v) in kwargs.items() if k in ("deps", "verbose")} # Forward args
    )

    proto_compile_js_library(
        name = kwargs.get("name"),
        deps = [name_pb],
        js_deps = js_deps,
        visibility = kwargs.get("visibility", []),
        tags = kwargs.get("tags", []),
    )