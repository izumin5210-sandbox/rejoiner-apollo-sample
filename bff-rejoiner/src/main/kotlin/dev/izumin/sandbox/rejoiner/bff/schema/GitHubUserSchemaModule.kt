package dev.izumin.sandbox.rejoiner.bff.schema

import com.google.api.graphql.rejoiner.Query
import com.google.api.graphql.rejoiner.SchemaModification
import com.google.api.graphql.rejoiner.SchemaModule
import com.google.common.base.Function
import com.google.common.util.concurrent.Futures
import com.google.common.util.concurrent.ListenableFuture
import com.google.common.util.concurrent.MoreExecutors
import dev.izumin5210.sandbox.github.GetUserRequest
import dev.izumin5210.sandbox.github.ListUsersRequest
import dev.izumin5210.sandbox.github.User
import dev.izumin5210.sandbox.github.UserServiceGrpc
import graphql.schema.DataFetchingEnvironment
import net.javacrumbs.futureconverter.java8guava.FutureConverter
import org.dataloader.DataLoaderRegistry


class GitHubUserSchemaModule : SchemaModule() {
    @Query("getGithubUser")
    fun getUser(req: GetUserRequest, client: UserServiceGrpc.UserServiceFutureStub): ListenableFuture<User> {
        return client.getUser(req)
    }

    @Query("listGithubUsers")
    fun listUsers(req: ListUsersRequest, client: UserServiceGrpc.UserServiceFutureStub): ListenableFuture<List<User>> {
        return Futures.transform(
                client.listUsers(req),
                Function { it!!.usersList },
                MoreExecutors.directExecutor()
        )
    }

    @SchemaModification(addField = "githubAccount", onType = dev.izumin5210.sandbox.qiita.User::class)
    fun qiitaUserToGitHubUser(user: dev.izumin5210.sandbox.qiita.User, environment: DataFetchingEnvironment): ListenableFuture<User?> {
        return FutureConverter.toListenableFuture(
                environment
                        .getContext<DataLoaderRegistry>()
                        .getDataLoader<String, User?>("github/users")
                        .load(user.githubLoginName)
        )
    }
}
