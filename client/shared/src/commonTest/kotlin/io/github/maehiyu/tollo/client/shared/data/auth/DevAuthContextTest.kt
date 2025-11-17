package io.github.maehiyu.tollo.client.shared.data.auth

import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.BeforeTest
import kotlin.test.assertNull

class DevAuthContextTest {
    @BeforeTest
    fun setUp() {
        DevAuthContext.clearUser()
    }

    @Test
    fun testSetAndGetUserIdAndEmail() {
        // Given
        val userId = "user-123"
        val email = "test@sample.com"

        // When
        DevAuthContext.setUser(userId, email)

        // Then
        assertEquals(userId, DevAuthContext.getUserId())
        assertEquals(email, DevAuthContext.getEmail())
    }

    @Test
    fun testClearUser() {
        val userId = "user-123"
        val email = "test@sample.com"
        DevAuthContext.setUser(userId, email)
        DevAuthContext.clearUser()
        assertEquals(null, DevAuthContext.getUserId())
        assertEquals(null, DevAuthContext.getEmail())
    }

    @Test
    fun testUserIdAndEmailAreNull() {
        val userId = DevAuthContext.getUserId()
        val email = DevAuthContext.getEmail()
        assertNull(userId)
        assertNull(email)
    }
}
