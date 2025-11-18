package io.github.maehiyu.tollo.client.shared.domain.service

import io.github.maehiyu.tollo.client.shared.data.auth.DevAuthContext
import kotlin.test.BeforeTest
import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertNotNull
import kotlin.test.assertNull

class AuthServiceTest {
    private lateinit var authService: AuthService
    @BeforeTest
    fun setUp() {
        DevAuthContext.clearUser()
        authService = AuthService(DevAuthContext)
    }

    @Test
    fun testSignUpAndLogin() {
        val generatedUserId = authService.signUp("test@sample.com", "password")
        val userId = authService.login("test@sample.com", "password")
        assertEquals(generatedUserId, userId)
    }
    @Test
    fun testLogout() {
        val generatedUserId = authService.signUp("test@sample.com" ,"password")
        assertNotNull(generatedUserId,DevAuthContext.getUserId())
        authService.logout()
        assertNull(DevAuthContext.getUserId())
    }
    @Test
    fun testisLoggedIn() {
        authService.signUp("test@sample.com", "password")
        authService.login("test@sample.com", "password")
        val isLoggedIn = authService.isLoggedIn()
        assertEquals(true, isLoggedIn)
    }
    @Test
    fun testGetCurrentUserId() {
        authService.signUp("test@sample.com", "password")
        val userId = authService.login("test@sample.com", "password")
        assertEquals(userId,authService.getCurrentUserId())
    }
    @Test
    fun testGetEmail() {
        authService.signUp("test@sample.com", "password")
        assertEquals("test@sample.com", authService.getCurrentUserEmail())
    }

}