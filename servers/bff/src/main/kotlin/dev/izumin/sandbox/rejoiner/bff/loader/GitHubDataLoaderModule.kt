package dev.izumin.sandbox.rejoiner.bff.loader

import com.google.common.base.Function
import com.google.common.util.concurrent.Futures
import com.google.common.util.concurrent.ListenableFuture
import com.google.common.util.concurrent.MoreExecutors
import com.google.inject.AbstractModule
import com.google.inject.Provides
import dev.izumin5210.sandbox.github.ListUsersRequest
import dev.izumin5210.sandbox.github.User
import dev.izumin5210.sandbox.github.UserServiceGrpc
import net.javacrumbs.futureconverter.java8guava.FutureConverter
import org.dataloader.BatchLoader
import org.dataloader.DataLoader
import org.dataloader.DataLoaderOptions
import org.dataloader.DataLoaderRegistry

class GitHubDataLoaderModule : AbstractModule() {
    @Provides
    fun dataLoaderRegistry(userClient: UserServiceGrpc.UserServiceFutureStub): DataLoaderRegistry {
        val userBatchLoader = BatchLoader<String, User> { keys ->
            val req = ListUsersRequest.newBuilder()
                    .addAllLogins(keys)
                    .build()
            val lf: ListenableFuture<List<User?>> =
                    Futures.transform(
                            userClient.listUsers(req),
                            Function { resp ->
                                val m = resp!!.usersList.map { it.login to it }.toMap()
                                keys.map { m[it] }
                            },
                            MoreExecutors.directExecutor()
                    );
            FutureConverter.toCompletableFuture(lf)
        }

        val registry = DataLoaderRegistry()
        registry.register("github/users", DataLoader(userBatchLoader, DataLoaderOptions.newOptions().setMaxBatchSize(5)))
        return registry
    }
}
