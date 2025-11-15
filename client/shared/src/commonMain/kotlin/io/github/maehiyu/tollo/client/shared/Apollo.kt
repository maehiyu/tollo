package io.github.maehiyu.tollo.client.shared

import com.apollographql.apollo.ApolloClient

fun createApolloClient(): ApolloClient {
    return ApolloClient.Builder()
        .serverUrl("http://localhost:8080/query")
        .build()
}

