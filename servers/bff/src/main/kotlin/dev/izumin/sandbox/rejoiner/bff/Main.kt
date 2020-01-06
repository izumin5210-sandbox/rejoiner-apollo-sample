package dev.izumin.sandbox.rejoiner.bff

import com.google.api.graphql.execution.GuavaListenableFutureSupport
import com.google.api.graphql.rejoiner.Schema
import com.google.api.graphql.rejoiner.SchemaProviderModule
import com.google.inject.AbstractModule
import com.google.inject.Guice
import com.google.inject.Inject
import dev.izumin.sandbox.rejoiner.bff.client.GitHubUserClientModule
import dev.izumin.sandbox.rejoiner.bff.client.QiitaItemClientModule
import dev.izumin.sandbox.rejoiner.bff.client.QiitaUserClientModule
import dev.izumin.sandbox.rejoiner.bff.schema.GitHubUserSchemaModule
import dev.izumin.sandbox.rejoiner.bff.schema.QiitaItemSchemaModule
import dev.izumin.sandbox.rejoiner.bff.schema.QiitaUserSchemaModule
import graphql.ExecutionInput
import graphql.GraphQL
import graphql.execution.instrumentation.ChainedInstrumentation
import graphql.execution.instrumentation.tracing.TracingInstrumentation
import graphql.schema.GraphQLSchema
import io.ktor.application.Application
import io.ktor.application.call
import io.ktor.application.install
import io.ktor.features.ContentNegotiation
import io.ktor.gson.gson
import io.ktor.http.HttpStatusCode
import io.ktor.http.content.resource
import io.ktor.http.content.static
import io.ktor.request.receive
import io.ktor.response.respond
import io.ktor.response.respondText
import io.ktor.routing.get
import io.ktor.routing.post
import io.ktor.routing.routing
import io.ktor.server.engine.embeddedServer
import io.ktor.server.netty.Netty
import java.util.logging.Logger

fun main(args: Array<String>) {
    embeddedServer(Netty, 8080, watchPaths = listOf<String>(), module = Application::module).start(true)
}

fun Application.module() {
    val logger = Logger.getLogger("bff")
    val port = 8080

    install(ContentNegotiation) {
        gson {}
    }

    Guice.createInjector(
            object : AbstractModule() {
                override fun configure() {
                    bind(Routes::class.java).asEagerSingleton()
                    bind(Application::class.java).toInstance(this@module)
                }
            },
            SchemaProviderModule(),
            object : AbstractModule() {
                override fun configure() {
                    install(QiitaUserSchemaModule())
                    install(QiitaItemSchemaModule())
                    install(GitHubUserSchemaModule())
                }
            },
            object : AbstractModule() {
                override fun configure() {
                    install(QiitaUserClientModule())
                    install(QiitaItemClientModule())
                    install(GitHubUserClientModule())
                }
            }
    )
}

data class GraphqlRequest(
        val query: String?,
        val operationName: String?,
        val variables: Map<String, Object>
) {
    fun newExecutionInput() =
            ExecutionInput.newExecutionInput()
                    .query(query)
                    .operationName(operationName)
                    .variables(variables ?: mapOf())
}

class Routes @Inject constructor(
        application: Application,
        @Schema schema: GraphQLSchema
) {
    init {
        application.routing {
            get("/echo") {
                call.respondText(call.request.queryParameters["message"]!!)
            }
            post("/graphql") {
                val instrumentation = ChainedInstrumentation(
                        arrayListOf(
                                GuavaListenableFutureSupport.listenableFutureInstrumentation(),
                                TracingInstrumentation()
                        )
                )
                val graphql = GraphQL.newGraphQL(schema).instrumentation(instrumentation).build()

                val req = call.receive(GraphqlRequest::class)
                val execInput = req.newExecutionInput()
                        .build()
                val result = graphql.execute(execInput)

                call.respond(HttpStatusCode.OK, result.toSpecification())
            }
            static {
                resource("/", "index.html")
            }
        }
    }
}
