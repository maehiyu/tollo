package io.github.maehiyu.tollo.client.shared.network

import com.apollographql.apollo.api.http.HttpRequest
import com.apollographql.apollo.api.http.HttpResponse
import com.apollographql.apollo.network.http.HttpInterceptor
import com.apollographql.apollo.network.http.HttpInterceptorChain
import io.github.maehiyu.tollo.client.shared.domain.auth.AuthContext

/**
 * Apollo Client のインターセプター
 * リクエストに X-User-ID ヘッダーを自動的に追加する
 */
class AuthInterceptor(private val authContext: AuthContext) : HttpInterceptor {
    override suspend fun intercept(
        request: HttpRequest,
        chain: HttpInterceptorChain
    ): HttpResponse {
        val userId = authContext.getUserId()

        val newRequest = if (userId != null) {
            request.newBuilder()
                .addHeader("X-User-ID", userId)
                .build()
        } else {
            request
        }

        return chain.proceed(newRequest)
    }
}