package io.github.maehiyu.tollo.client.shared.domain.service

import com.benasher44.uuid.uuid4
import io.github.maehiyu.tollo.client.shared.domain.auth.AuthContext

class AuthService(val authContext: AuthContext) {
    private val users = mutableMapOf<String, UserCredentials>()

    data class UserCredentials(
        val userId: String,
        val email: String,
        val password: String,
    )

    fun signUp(email: String, password: String): String? {
        if (users.containsKey(email)) {
            return null  // 既に存在
        }
        val userId = uuid4().toString()
        users[email] = UserCredentials(userId, email, password)
        authContext.setUser(userId, email)
        return userId
    }
    fun login(email: String, password: String): String? {
        val user = users[email] ?: return null
        if (user.password != password) {
            return null  // パスワード不一致
        }
        authContext.setUser(user.userId, user.email)
        return user.userId
    }
    fun logout() {
        authContext.clearUser()
    }
    fun isLoggedIn(): Boolean {
        return authContext.getUserId() != null
    }
    fun getCurrentUserId(): String? {
        return authContext.getUserId()
    }
    fun getCurrentUserEmail(): String? {
        return authContext.getEmail()
    }
}