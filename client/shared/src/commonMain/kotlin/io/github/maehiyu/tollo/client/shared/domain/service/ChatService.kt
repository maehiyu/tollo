package io.github.maehiyu.tollo.client.shared.domain.service

import com.apollographql.apollo.api.Optional
import io.github.maehiyu.tollo.client.shared.ChatMessagesQuery
import io.github.maehiyu.tollo.client.shared.CreateChatMutation
import io.github.maehiyu.tollo.client.shared.SendMessageMutation
import io.github.maehiyu.tollo.client.shared.UserChatsQuery
import io.github.maehiyu.tollo.client.shared.domain.repository.ChatRepository
import io.github.maehiyu.tollo.client.shared.type.AnswerMessageInput
import io.github.maehiyu.tollo.client.shared.type.PromotionalMessageInput
import io.github.maehiyu.tollo.client.shared.type.QuestionMessageInput
import io.github.maehiyu.tollo.client.shared.type.StandardMessageInput

class ChatService(
    private val chatRepository: ChatRepository
) {
    suspend fun getUserChats(userId: String): List<UserChatsQuery.UserChat>? {
        return chatRepository.getUserChats(userId)
    }

    suspend fun getChatMessages(chatId: String): List<ChatMessagesQuery.ChatMessage>? {
        return chatRepository.getChatMessages(chatId)
    }

    suspend fun sendStandardMessage(
        chatId: String,
        content: String
    ): SendMessageMutation.SendMessage? {
        return chatRepository.sendMessage(
            chatId = chatId,
            standard = StandardMessageInput(content),
            question = null,
            answer = null,
            promotional = null
        )
    }

    suspend fun sendQuestionMessage(
        chatId: String,
        content: String,
        tags: List<String>
    ): SendMessageMutation.SendMessage? {
        return chatRepository.sendMessage(
            chatId = chatId,
            standard = null,
            question = QuestionMessageInput(content, Optional.presentIfNotNull(tags)),
            answer = null,
            promotional = null
        )
    }

    suspend fun sendAnswerMessage(
        chatId: String,
        content: String,
        questionId: String
    ): SendMessageMutation.SendMessage? {
        return chatRepository.sendMessage(
            chatId = chatId,
            standard = null,
            question = null,
            answer = AnswerMessageInput(content, questionId),
            promotional = null
        )
    }

    suspend fun sendPromotionalMessage(
        chatId: String,
        title: String,
        body: String,
        actionUrl: String,
        imageUrl: String? = null
    ): SendMessageMutation.SendMessage? {
        return chatRepository.sendMessage(
            chatId = chatId,
            standard = null,
            question = null,
            answer = null,
            promotional = PromotionalMessageInput(title, body, actionUrl, Optional.presentIfNotNull(imageUrl))
        )
    }

    suspend fun createChat(
        generalUserId: String,
        professionalUserId: String
    ): CreateChatMutation.CreateChat? {
        return chatRepository.createChat(generalUserId, professionalUserId)
    }
}