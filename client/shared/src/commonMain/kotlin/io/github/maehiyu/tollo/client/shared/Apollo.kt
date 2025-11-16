package io.github.maehiyu.tollo.client.shared

import com.apollographql.apollo.ApolloClient
import io.github.maehiyu.tollo.client.shared.domain.auth.AuthContext
import io.github.maehiyu.tollo.client.shared.network.AuthInterceptor

fun createApolloClient(authContext: AuthContext): ApolloClient {
    return ApolloClient.Builder()
        .serverUrl("http://localhost:8080/query")
        .addHttpInterceptor(AuthInterceptor(authContext))
        .build()
}

