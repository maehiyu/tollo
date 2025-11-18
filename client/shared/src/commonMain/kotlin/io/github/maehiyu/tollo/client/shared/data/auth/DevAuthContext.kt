package io.github.maehiyu.tollo.client.shared.data.auth

import io.github.maehiyu.tollo.client.shared.domain.auth.AuthContext

/**
 * 開発用の認証コンテキスト実装
 * ユーザーID をメモリ内に保持する
 */
object DevAuthContext : AuthContext {
    private var userId: String? = null
    private var email: String? = null

    override fun setUser(userId: String, email: String) {
        this.userId = userId
        this.email = email
    }

    override fun getUserId(): String? {
        return userId
    }

    override fun getEmail(): String? {
        return email
    }

    override fun clearUser() {
        userId = null
        email = null
    }
}