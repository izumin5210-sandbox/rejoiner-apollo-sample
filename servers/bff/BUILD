load("@io_grpc_grpc_java//:java_grpc_library.bzl", "java_grpc_library")
load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_jvm_binary")

load("@rules_proto//proto:defs.bzl", "proto_library")

java_proto_library(
    name = "api_qiita_user_proto",
    deps = ["//api/qiita:user_proto"],
)

java_grpc_library(
    name = "api_qiita_user_grpc",
    srcs = ["//api/qiita:user_proto"],
    deps = [":api_qiita_user_proto"],
)

java_proto_library(
    name = "api_qiita_item_proto",
    deps = ["//api/qiita:item_proto"],
)

java_grpc_library(
    name = "api_qiita_item_grpc",
    srcs = ["//api/qiita:item_proto"],
    deps = [":api_qiita_item_proto"],
)

java_library(
    name = "api",
    deps = [
        ":api_qiita_user_grpc",
        ":api_qiita_user_proto",
        ":api_qiita_item_grpc",
        ":api_qiita_item_proto",
    ]
)

java_library(
    name = "deps",
    exports = [
        "@maven//:io_grpc_grpc_netty_shaded",
        "@maven//:io_grpc_grpc_protobuf",
        "@maven//:io_grpc_grpc_stub",
    ]
)

kt_jvm_binary(
    name = "BFF",
    main_class = "dev.izumin.sandbox.bff.MainKt",
    srcs = glob(["src/main/kotlin/dev/izumin/sandbox/bff/*.kt"]),
    runtime_deps = [
        ":api",
        ":deps",
    ]
)
