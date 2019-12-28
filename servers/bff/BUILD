# load("@rules_java//java:defs.bzl", "java_binary")
load("@io_bazel_rules_kotlin//kotlin:kotlin.bzl", "kt_jvm_binary")

kt_jvm_binary(
    name = "BFF",
    main_class = "dev.izumin.sandbox.bff.MainKt",
    srcs = glob(["src/main/kotlin/dev/izumin/sandbox/bff/*.kt"]),
)
