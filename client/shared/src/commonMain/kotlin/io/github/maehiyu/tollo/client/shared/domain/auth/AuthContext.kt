package io.github.maehiyu.tollo.client.shared.domain.auth

object AuthContext {
    private var userId: String? = null
    fun setUserId(userId: String) {
        this.userId = userId
    }
    fun getUserId(): String? {
        return userId
    }

    fun clearUserId() {
        userId = null
    }
}