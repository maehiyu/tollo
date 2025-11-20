package io.github.maehiyu.tollo.client.shared.data.repository

import com.apollographql.apollo.ApolloClient
import com.apollographql.apollo.api.Optional
import io.github.maehiyu.tollo.client.shared.ChatMessagesQuery
import io.github.maehiyu.tollo.client.shared.CreateChatMutation
import io.github.maehiyu.tollo.client.shared.SendMessageMutation
import io.github.maehiyu.tollo.client.shared.UserChatsQuery
import io.github.maehiyu.tollo.client.shared.domain.repository.ChatRepository
import io.github.maehiyu.tollo.client.shared.type.AnswerMessageInput
import io.github.maehiyu.tollo.client.shared.type.CreateChatInput
import io.github.maehiyu.tollo.client.shared.type.PromotionalMessageInput
import io.github.maehiyu.tollo.client.shared.type.QuestionMessageInput
import io.github.maehiyu.tollo.client.shared.type.SendMessageInput
import io.github.maehiyu.tollo.client.shared.type.StandardMessageInput

class ChatRepositoryImpl(private val apolloClient: ApolloClient) : ChatRepository {
    override suspend fun getUserChats(userId: String): List<UserChatsQuery.UserChat>? {
        return try {
            val response = apolloClient.query(UserChatsQuery(userId = userId)).execute()
            if (response.hasErrors()) {
                println("GraphQL Errors: ${response.errors}")
                return null
            }
            println("UserChats Response Data: ${response.data}")
            response.data?.userChats
        } catch (e: Exception) {
            println("Network or other exception: ${e.message}")
            return null
        }
    }
    override suspend fun getChatMessages(chatId: String): List<ChatMessagesQuery.ChatMessage>? {
        return try {
            val response = apolloClient.query(ChatMessagesQuery(chatId = chatId)).execute()
            if (response.hasErrors()) {
                println("GraphQL Errors: ${response.errors}")
                return null
            }
            println("ChatMessages Response Data: ${response.data}")
            response.data?.chatMessages
        } catch (e: Exception) {
            println("Network or other exception: ${e.message}")
            return null
        }
    }
    override suspend fun sendMessage(
        chatId: String,
        standard: StandardMessageInput?,
        question: QuestionMessageInput?,
        answer: AnswerMessageInput?,
        promotional: PromotionalMessageInput?
    ): SendMessageMutation.SendMessage? {
        return try {
            val input = SendMessageInput(
                chatId = chatId,
                standard = Optional.presentIfNotNull(standard),
                question = Optional.presentIfNotNull(question),
                answer = Optional.presentIfNotNull(answer),
                promotional = Optional.presentIfNotNull(promotional)
            )
            val response = apolloClient.mutation(SendMessageMutation(input = input)).execute()
            if (response.hasErrors()) {
                println("GraphQL Errors: ${response.errors}")
                return null
            }
            println("SendMessage Response Data: ${response.data}")
            response.data?.sendMessage
        } catch (e: Exception) {
            println("Network or other exception: ${e.message}")
            return null
        }
    }
    override suspend fun createChat(
        generalUserId: String,
        professionalUserId: String
    ): CreateChatMutation.CreateChat? {
        return try {
            val input = CreateChatInput(
                generalUserID = generalUserId,
                professionalUserID = professionalUserId
            )
            val response = apolloClient.mutation(CreateChatMutation(input = input)).execute()
            if (response.hasErrors()) {
                println("GraphQL Errors: ${response.errors}")
                return null
            }
            println("CreateChat Response Data: ${response.data}")
            response.data?.createChat
        } catch (e: Exception) {
            println("Network or other exception: ${e.message}")
            return null
        }
    }
}