package dev.izumin.sandbox.rejoiner.bff.schema

import com.google.api.graphql.rejoiner.GrpcSchemaModule
import com.google.api.graphql.rejoiner.Query
import com.google.api.graphql.rejoiner.SchemaModification
import com.google.common.base.Function
import com.google.common.util.concurrent.Futures
import com.google.common.util.concurrent.ListenableFuture
import com.google.common.util.concurrent.MoreExecutors
import dev.izumin5210.sandbox.qiita.Item
import dev.izumin5210.sandbox.qiita.ItemServiceGrpc
import dev.izumin5210.sandbox.qiita.ListItemsRequest
import dev.izumin5210.sandbox.qiita.ListItemsResponse
import dev.izumin5210.sandbox.qiita.ListStockersRequest
import dev.izumin5210.sandbox.qiita.ListStockersResponse
import dev.izumin5210.sandbox.qiita.User
import dev.izumin5210.sandbox.qiita.UserServiceGrpc

class ItemSchemaModule : GrpcSchemaModule() {
    @Query("listItems")
    fun listItems(req: ListItemsRequest, client: ItemServiceGrpc.ItemServiceFutureStub): ListenableFuture<List<Item>> {
        return Futures.transform(
                client.listItems(req),
                Function { resp: ListItemsResponse? -> resp!!.itemsList },
                MoreExecutors.directExecutor()
        )
    }

    @SchemaModification(addField = "items", onType = User::class)
    private fun userToItems(user: User, client: ItemServiceGrpc.ItemServiceFutureStub): ListenableFuture<List<Item>> {
        val req = ListItemsRequest.newBuilder().setUserId(user.id).build()
        return listItems(req, client)
    }

    @SchemaModification(addField = "followers", onType = Item::class)
    private fun userToStockers(item: Item, client: UserServiceGrpc.UserServiceFutureStub): ListenableFuture<List<User>> {
        val req = ListStockersRequest.newBuilder().setItemId(item.id).build()
        return Futures.transform(
                client.listStockers(req),
                Function { resp: ListStockersResponse? -> resp!!.usersList },
                MoreExecutors.directExecutor()
        )
    }
}
