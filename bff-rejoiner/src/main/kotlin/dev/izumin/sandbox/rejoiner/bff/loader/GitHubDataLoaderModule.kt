package dev.izumin.sandbox.rejoiner.bff.loader

import com.google.inject.AbstractModule
import com.google.inject.Provides
import dev.izumin5210.sandbox.github.ListUsersRequest
import dev.izumin5210.sandbox.github.User
import dev.izumin5210.sandbox.github.UserServiceGrpc
import net.javacrumbs.futureconverter.java8guava.FutureConverter
import org.dataloader.DataLoader
import org.dataloader.DataLoaderOptions
import org.dataloader.DataLoaderRegistry
import org.dataloader.MappedBatchLoader

class GitHubDataLoaderModule : AbstractModule() {
    @Provides
    fun dataLoaderRegistry(userClient: UserServiceGrpc.UserServiceFutureStub): DataLoaderRegistry {
        val userBatchLoader = MappedBatchLoader<String, User> { keys ->
            val req = ListUsersRequest.newBuilder()
                    .addAllLogins(keys.filterNot { it.isEmpty() })
                    .build()
            FutureConverter.toCompletableFuture(userClient.listUsers(req))
                    .thenApply { it!!.usersList.associateBy { it.login } }
        }

        val registry = DataLoaderRegistry()
        registry.register("github/users", DataLoader.newMappedDataLoader(
                userBatchLoader, DataLoaderOptions.newOptions().setMaxBatchSize(5)))
        return registry
    }
}
