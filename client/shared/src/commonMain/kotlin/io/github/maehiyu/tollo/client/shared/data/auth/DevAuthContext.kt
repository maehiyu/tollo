package io.github.maehiyu.tollo.client.shared.data.auth

import io.github.maehiyu.tollo.client.shared.domain.auth.AuthContext

/**
 * 開発用の認証コンテキスト実装
 * ユーザーID をメモリ内に保持する
 */
object DevAuthContext : AuthContext {
    private var userId: String? = null

    override fun setUserId(userId: String) {
        this.userId = userId
    }

    override fun getUserId(): String? {
        return userId
    }

    override fun clearUserId() {
        userId = null
    }
}