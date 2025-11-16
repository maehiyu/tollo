package io.github.maehiyu.tollo.client.shared.domain.service

import io.github.maehiyu.tollo.client.shared.data.auth.DevAuthContext
import kotlin.test.BeforeTest
import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertNotNull
import kotlin.test.assertNull

class AuthServiceTest {
    val authService = AuthService(DevAuthContext)
    @BeforeTest
    fun setUp() {
        DevAuthContext.clearUserId()
    }

    @Test
    fun testLogin() {
        val userId = authService.login("test@sample.com", "password")
        assertNotNull(userId)
        assertEquals("test@sample.com",DevAuthContext.getUserId())
    }

    @Test
    fun testLogout() {
        authService.login("test@sample.com", "password")
        authService.logout()
        assertNull(DevAuthContext.getUserId())
    }
    @Test
    fun testisLoggedIn() {
        authService.login("test@sample.com", "password")
        val isLoggedIn = authService.isLoggedIn()
        assertEquals(true, isLoggedIn)
    }
    @Test
    fun testGetCurrentUserId() {
        val userId = authService.login("test@sample.com", "password")
        assertEquals(userId,authService.getCurrentUserId())
    }
}