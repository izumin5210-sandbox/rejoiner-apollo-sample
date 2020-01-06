package dev.izumin.sandbox.rejoiner.bff.client

import com.google.inject.AbstractModule
import dev.izumin5210.sandbox.github.UserServiceGrpc
import io.grpc.ManagedChannelBuilder

class GitHubUserClientModule : AbstractModule() {
    companion object {
        const val HOST = "localhost"
        const val PORT = 50052
    }

    override fun configure() {
        val ch = ManagedChannelBuilder.forAddress(HOST, PORT).usePlaintext().build()
        bind(UserServiceGrpc.UserServiceFutureStub::class.java)
                .toInstance(UserServiceGrpc.newFutureStub(ch))
    }
}
