package dev.izumin.sandbox.rejoiner.bff.client

import com.google.inject.AbstractModule
import dev.izumin5210.sandbox.qiita.ItemServiceGrpc
import io.grpc.ManagedChannelBuilder

class QiitaItemClientModule : AbstractModule() {
    companion object {
        const val HOST = "localhost"
        const val PORT = 50051
    }

    override fun configure() {
        val ch = ManagedChannelBuilder.forAddress(HOST, PORT).usePlaintext().build()
        bind(ItemServiceGrpc.ItemServiceFutureStub::class.java)
                .toInstance(ItemServiceGrpc.newFutureStub(ch))
    }
}
