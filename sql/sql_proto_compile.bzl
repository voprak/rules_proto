load("//:compile.bzl", "proto_compile")

def sql_proto_compile(**kwargs):
    proto_compile(
        plugins = [
            str(Label("//sql:sql")),
        ],
        **kwargs
    )
