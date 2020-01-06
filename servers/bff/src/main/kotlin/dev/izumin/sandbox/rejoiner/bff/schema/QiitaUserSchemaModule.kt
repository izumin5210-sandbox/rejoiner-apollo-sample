package dev.izumin.sandbox.rejoiner.bff.schema

import com.google.api.graphql.rejoiner.Query
import com.google.api.graphql.rejoiner.SchemaModification
import com.google.api.graphql.rejoiner.SchemaModule
import com.google.common.base.Function
import com.google.common.util.concurrent.Futures
import com.google.common.util.concurrent.ListenableFuture
import com.google.common.util.concurrent.MoreExecutors
import dev.izumin5210.sandbox.qiita.GetUserRequest
import dev.izumin5210.sandbox.qiita.ListFolloweesRequest
import dev.izumin5210.sandbox.qiita.ListFolloweesResponse
import dev.izumin5210.sandbox.qiita.ListFollowersRequest
import dev.izumin5210.sandbox.qiita.ListFollowersResponse
import dev.izumin5210.sandbox.qiita.User
import dev.izumin5210.sandbox.qiita.UserServiceGrpc

class QiitaUserSchemaModule : SchemaModule() {
    @Query("getUser")
    fun getUser(req: GetUserRequest, client: UserServiceGrpc.UserServiceFutureStub): ListenableFuture<User> {
        return client.getUser(req)
    }

    @SchemaModification(addField = "followees", onType = User::class)
    fun userToFollowees(user: User, client: UserServiceGrpc.UserServiceFutureStub): ListenableFuture<List<User>> {
        val req = ListFolloweesRequest.newBuilder().setUserId(user.id).build()
        return Futures.transform(
                client.listFollowees(req),
                Function { resp: ListFolloweesResponse? -> resp!!.usersList },
                MoreExecutors.directExecutor()
        )
    }

    @SchemaModification(addField = "followers", onType = User::class)
    fun userToFollowers(user: User, client: UserServiceGrpc.UserServiceFutureStub): ListenableFuture<List<User>> {
        val req = ListFollowersRequest.newBuilder().setUserId(user.id).build()
        return Futures.transform(
                client.listFollowers(req),
                Function { resp: ListFollowersResponse? -> resp!!.usersList },
                MoreExecutors.directExecutor()
        )
    }
}
