package io.github.maehiyu.tollo.client.shared.domain.auth

interface AuthContext {
    fun setUserId(userId: String)
    fun getUserId(): String?
    fun clearUserId()
}