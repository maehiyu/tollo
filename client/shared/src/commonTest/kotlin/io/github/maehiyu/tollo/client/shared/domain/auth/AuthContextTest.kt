package io.github.maehiyu.tollo.client.shared.domain.auth

import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.BeforeTest
import kotlin.test.assertNull

class AuthContextTest {
    @BeforeTest
    fun setUp() {
        AuthContext.clearUserId()
    }
    @Test
    fun testSetAndGetUserId() {
        // Given
        val userId = "user-123"

        // When
        AuthContext.setUserId(userId)

        // Then
        assertEquals(userId, AuthContext.getUserId())
    }

    @Test
    fun testClearUserId() {
        val userId = "user-123"

        AuthContext.setUserId(userId)
        AuthContext.clearUserId()
        assertEquals(null, AuthContext.getUserId())
    }

    @Test
    fun testUserIdIsNull() {
        val userId = AuthContext.getUserId()
        assertNull(userId)
    }
}