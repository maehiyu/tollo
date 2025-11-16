package io.github.maehiyu.tollo.client.shared.domain.service

import com.benasher44.uuid.uuid4
import io.github.maehiyu.tollo.client.shared.domain.auth.AuthContext

class AuthService(val authContext: AuthContext) {
    fun login(email: String, password: String): String? {
        authContext.setUserId(email)
        return email
    }
    fun logout() {
        authContext.clearUserId()
    }
    fun isLoggedIn(): Boolean {
        return authContext.getUserId() != null
    }
    fun getCurrentUserId(): String? {
        return authContext.getUserId()
    }
}