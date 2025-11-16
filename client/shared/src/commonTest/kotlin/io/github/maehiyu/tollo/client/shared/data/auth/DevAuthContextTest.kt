package io.github.maehiyu.tollo.client.shared.data.auth

import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.BeforeTest
import kotlin.test.assertNull

class DevAuthContextTest {
    @BeforeTest
    fun setUp() {
        DevAuthContext.clearUserId()
    }

    @Test
    fun testSetAndGetUserId() {
        // Given
        val userId = "user-123"

        // When
        DevAuthContext.setUserId(userId)

        // Then
        assertEquals(userId, DevAuthContext.getUserId())
    }

    @Test
    fun testClearUserId() {
        val userId = "user-123"

        DevAuthContext.setUserId(userId)
        DevAuthContext.clearUserId()
        assertEquals(null, DevAuthContext.getUserId())
    }

    @Test
    fun testUserIdIsNull() {
        val userId = DevAuthContext.getUserId()
        assertNull(userId)
    }
}
