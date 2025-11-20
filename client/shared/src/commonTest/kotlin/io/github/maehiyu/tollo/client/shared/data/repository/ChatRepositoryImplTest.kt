package io.github.maehiyu.tollo.client.shared.data.repository

import io.github.maehiyu.tollo.client.shared.type.StandardMessageInput
import kotlin.test.BeforeTest
import kotlin.test.Test
import kotlin.test.assertEquals
import kotlin.test.assertNotNull

class ChatRepositoryImplTest {
    private lateinit var chatRepository: ChatRepositoryImpl

    @BeforeTest
    fun setUp() {
        // TODO: Mock ApolloClient
    }

    @Test
    fun testGetUserChats(){
        // TODO: Test getUserChats
    }

    @Test
    fun testGetChatMessages() {
        // TODO: Test getChatMessages
    }

    @Test
    fun testSendStandardMessage() {
        // TODO: Test sendMessage with StandardMessageInput
    }

    @Test
    fun testCreateChat() {
        // TODO: Test createChat
    }
}