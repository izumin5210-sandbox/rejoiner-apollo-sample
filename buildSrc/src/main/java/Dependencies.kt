object Dependencies {
    object Versions {
        const val KOTLIN = "1.3.60"
        const val PROTOBUF = "3.11.1"
        const val GRPC_JAVA = "1.26.0"
        const val GOOGLE_APIS_COMMON_PROTOS = "0.0.3"
        const val JAVAX_ANNOTATION = "1.3.2"

        const val GUAVA = "28.2-jre"
        const val PROTO_GOOGLE_COMMON_PROTOS = "1.16.0"

        const val REJOINER = "0.3.0"
        const val JETTY = "9.3.8.v20160314"
        const val SLF4J = "1.6.2"
        const val GUICE = "4.2.0"
        const val FUTURE_CONVERTER_JAVA8_GUAVA = "1.1.0"

        const val KTOR = "1.2.6"
    }

    const val KOTLIN_STD_LIB = "org.jetbrains.kotlin:kotlin-stdlib:${Versions.KOTLIN}"
    const val PROTOC = "com.google.protobuf:protoc:${Versions.PROTOBUF}"
    const val GRPC_JAVA = "io.grpc:protoc-gen-grpc-java:${Versions.GRPC_JAVA}"
    const val GRPC_NETTY = "io.grpc:grpc-netty:${Versions.GRPC_JAVA}"
    const val GRPC_PROTOBUF = "io.grpc:grpc-protobuf:${Versions.GRPC_JAVA}"
    const val GRPC_STUB = "io.grpc:grpc-stub:${Versions.GRPC_JAVA}"
    const val GOOGLE_APIS_COMMON_PROTOS =
        "com.google.api.grpc:googleapis-common-protos:${Versions.GOOGLE_APIS_COMMON_PROTOS}"
    const val JAVAX_ANNOTATION = "javax.annotation:javax.annotation-api:${Versions.JAVAX_ANNOTATION}"

    const val PROTOBUF_JAVA_UTIL = "com.google.protobuf:protobuf-java-util:${Versions.PROTOBUF}"
    const val GUAVA = "com.google.guava:guava:${Versions.GUAVA}"
    const val PROTO_GOOGLE_COMMON_PROTOS = "com.google.api.grpc:proto-google-common-protos:${Versions.PROTO_GOOGLE_COMMON_PROTOS}"

    const val REJOINER = "com.google.api.graphql:rejoiner:${Versions.REJOINER}"
    const val JETTY = "org.eclipse.jetty:jetty-server:${Versions.JETTY}"
    const val JETTY_SERVLET = "org.eclipse.jetty:jetty-servlet:${Versions.JETTY}"
    const val SLF4J = "org.slf4j:slf4j-simple:${Versions.SLF4J}"
    const val GUICE_SERVLET = "com.google.inject.extensions:guice-servlet:${Versions.GUICE}"
    const val FUTURE_CONVERTER_JAVA8_GUAVA = "net.javacrumbs.future-converter:future-converter-java8-guava:${Versions.FUTURE_CONVERTER_JAVA8_GUAVA}"

    const val KTOR_NETTY = "io.ktor:ktor-server-netty:${Versions.KTOR}"
    const val KTOR_GSON = "io.ktor:ktor-gson:${Versions.KTOR}"
}
