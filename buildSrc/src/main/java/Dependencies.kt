object Dependencies {
    object Versions {
        const val KOTLIN = "1.3.60"
        const val PROTOC = "3.9.1"
        const val GRPC_JAVA = "1.24.0"
        const val GOOGLE_APIS_COMMON_PROTOS = "0.0.3"
        const val JAVAX_ANNOTATION = "1.3.2"
    }

    const val KOTLIN_STD_LIB = "org.jetbrains.kotlin:kotlin-stdlib:${Versions.KOTLIN}"
    const val PROTOC = "com.google.protobuf:protoc:${Versions.PROTOC}"
    const val GRPC_JAVA = "io.grpc:protoc-gen-grpc-java:${Versions.GRPC_JAVA}"
    const val GRPC_NETTY = "io.grpc:grpc-netty:${Versions.GRPC_JAVA}"
    const val GRPC_PROTOBUF = "io.grpc:grpc-protobuf:${Versions.GRPC_JAVA}"
    const val GRPC_STUB = "io.grpc:grpc-stub:${Versions.GRPC_JAVA}"
    const val GOOGLE_APIS_COMMON_PROTOS =
        "com.google.api.grpc:googleapis-common-protos:${Versions.GOOGLE_APIS_COMMON_PROTOS}"
    const val JAVAX_ANNOTATION = "javax.annotation:javax.annotation-api:${Versions.JAVAX_ANNOTATION}"
}
