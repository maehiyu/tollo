package io.github.maehiyu.tollo.client.shared.domain.auth

interface AuthContext {
    fun setUser(userId: String, email: String)
    fun getUserId(): String?
    fun getEmail(): String?
    fun clearUser()
}