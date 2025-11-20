package io.github.maehiyu.tollo.client.shared.domain.repository

import io.github.maehiyu.tollo.client.shared.ChatMessagesQuery
import io.github.maehiyu.tollo.client.shared.CreateChatMutation
import io.github.maehiyu.tollo.client.shared.SendMessageMutation
import io.github.maehiyu.tollo.client.shared.UserChatsQuery
import io.github.maehiyu.tollo.client.shared.type.StandardMessageInput
import io.github.maehiyu.tollo.client.shared.type.QuestionMessageInput
import io.github.maehiyu.tollo.client.shared.type.AnswerMessageInput
import io.github.maehiyu.tollo.client.shared.type.PromotionalMessageInput

interface ChatRepository {
    suspend fun getUserChats(userId: String): List<UserChatsQuery.UserChat>?
    suspend fun getChatMessages(chatId: String): List<ChatMessagesQuery.ChatMessage>?
    suspend fun sendMessage(
        chatId: String,
        standard: StandardMessageInput? = null,
        question: QuestionMessageInput? = null,
        answer: AnswerMessageInput? = null,
        promotional: PromotionalMessageInput? = null
    ): SendMessageMutation.SendMessage?
    suspend fun createChat(
        generalUserId: String,
        professionalUserId: String
    ): CreateChatMutation.CreateChat?
}