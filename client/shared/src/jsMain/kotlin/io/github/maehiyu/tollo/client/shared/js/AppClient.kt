@file:OptIn(ExperimentalJsExport::class)

package io.github.maehiyu.tollo.client.shared.js

import com.apollographql.apollo.api.Optional
import io.github.maehiyu.tollo.client.shared.createApolloClient
import io.github.maehiyu.tollo.client.shared.data.repository.UserRepositoryImpl
import io.github.maehiyu.tollo.client.shared.data.repository.ChatRepositoryImpl
import io.github.maehiyu.tollo.client.shared.domain.repository.UserRepository
import io.github.maehiyu.tollo.client.shared.domain.repository.ChatRepository
import io.github.maehiyu.tollo.client.shared.domain.service.UserService
import io.github.maehiyu.tollo.client.shared.domain.service.ChatService
import io.github.maehiyu.tollo.client.shared.type.GeneralProfileInput
import io.github.maehiyu.tollo.client.shared.type.ProfessionalProfileInput
import kotlinx.coroutines.promise
import kotlinx.coroutines.Dispatchers
import kotlin.coroutines.EmptyCoroutineContext
import kotlinx.coroutines.CoroutineScope
import kotlin.js.Promise
import kotlin.js.JsExport
import kotlin.js.JsName

import io.github.maehiyu.tollo.client.shared.data.auth.DevAuthContext
import io.github.maehiyu.tollo.client.shared.domain.service.AuthService

// DI Container: 依存を解決
private val apolloClient = createApolloClient(DevAuthContext)
private val userRepository: UserRepository = UserRepositoryImpl(apolloClient)
private val chatRepository: ChatRepository = ChatRepositoryImpl(apolloClient)
private val authServiceInstance: AuthService = AuthService(DevAuthContext)
private val userServiceInstance: UserService = UserService(userRepository)
private val chatServiceInstance: ChatService = ChatService(chatRepository)

// AuthService wrapper functions
@JsExport
fun signUp(email: String, password: String): String? {
    return authServiceInstance.signUp(email, password)
}

@JsExport
fun login(email: String, password: String): String? {
    return authServiceInstance.login(email, password)
}

@JsExport
fun logout() {
    authServiceInstance.logout()
}

@JsExport
fun isLoggedIn(): Boolean {
    return authServiceInstance.isLoggedIn()
}

@JsExport
fun getCurrentUserId(): String? {
    return authServiceInstance.getCurrentUserId()
}

@JsExport
data class JsGeneralProfile(
  val points: Int,
  val introduction: String
)

@JsExport
data class JsProfessionalProfile(
  val proBadgeUrl: String,
  val biography: String
)

@JsExport
data class JsUser(
    @JsName("id") val id: String,
    @JsName("name") val name: String,
    @JsName("email") val email: String,
    @JsName("createdAt") val createdAt: String // GraphQL の Time スカラーは String に変換されると仮定
)

@JsExport
fun getUser(id: String): Promise<JsUser?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        userServiceInstance.getUser(id = id, email = null)?.let { user ->
            JsUser(user.id, user.name, user.email, user.createdAt.toString())
        }
    }

@JsExport
fun getUserByEmail(email: String): Promise<JsUser?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        userServiceInstance.getUser(id = null, email = email)?.let { user ->
            JsUser(user.id, user.name, user.email, user.createdAt.toString())
        }
    }

@JsExport
fun getCurrentUser(): Promise<JsUser?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        userServiceInstance.getCurrentUser()?.let { user ->
            JsUser(user.id, user.name, user.email, user.createdAt.toString())
        }
    }

@JsExport
fun createUser(
  name: String,
  description: String?,
  generalProfile: JsGeneralProfile?,
  professionalProfile: JsProfessionalProfile?
): Promise<JsUser?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
      val generalInput = generalProfile?.let {
        GeneralProfileInput(it.points, it.introduction)
      }
      val professionalInput = professionalProfile?.let {
        ProfessionalProfileInput(it.proBadgeUrl, it.biography)
      }
      userServiceInstance.createUser(name, description, generalInput, professionalInput)?.let { user ->
          JsUser(user.id, user.name, user.email, user.createdAt.toString())
      }
    }

// ChatService wrapper types and functions
@JsExport
data class JsChat(
    @JsName("id") val id: String,
    @JsName("generalUserId") val generalUserId: String,
    @JsName("professionalUserId") val professionalUserId: String,
    @JsName("createdAt") val createdAt: String
)

@JsExport
data class JsMessage(
    @JsName("id") val id: String,
    @JsName("chatId") val chatId: String,
    @JsName("senderId") val senderId: String,
    @JsName("sentAt") val sentAt: String,
    @JsName("type") val type: String, // "__typename" from payload
    @JsName("content") val content: String?,
    @JsName("tags") val tags: Array<String>?,
    @JsName("questionId") val questionId: String?,
    @JsName("title") val title: String?,
    @JsName("body") val body: String?,
    @JsName("actionUrl") val actionUrl: String?,
    @JsName("imageUrl") val imageUrl: String?
)

@JsExport
fun getUserChats(userId: String): Promise<Array<JsChat>?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        chatServiceInstance.getUserChats(userId)?.map { chat ->
            JsChat(
                id = chat.id,
                generalUserId = chat.generalUserID,
                professionalUserId = chat.professionalUserID,
                createdAt = chat.createdAt.toString()
            )
        }?.toTypedArray()
    }

@JsExport
fun getChatMessages(chatId: String): Promise<Array<JsMessage>?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        chatServiceInstance.getChatMessages(chatId)?.map { msg ->
            val payload = msg.payload
            JsMessage(
                id = msg.id,
                chatId = msg.chatId,
                senderId = msg.senderId,
                sentAt = msg.sentAt.toString(),
                type = payload?.__typename ?: "Unknown",
                content = payload?.onStandardMessage?.content
                    ?: payload?.onQuestionMessage?.content
                    ?: payload?.onAnswerMessage?.content
                    ?: null,
                tags = payload?.onQuestionMessage?.tags?.toTypedArray(),
                questionId = payload?.onAnswerMessage?.questionId,
                title = payload?.onPromotionalMessage?.title,
                body = payload?.onPromotionalMessage?.body,
                actionUrl = payload?.onPromotionalMessage?.actionUrl,
                imageUrl = payload?.onPromotionalMessage?.imageUrl
            )
        }?.toTypedArray()
    }

@JsExport
fun sendStandardMessage(chatId: String, content: String): Promise<JsMessage?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        chatServiceInstance.sendStandardMessage(chatId, content)?.let { msg ->
            val payload = msg.payload
            JsMessage(
                id = msg.id,
                chatId = msg.chatId,
                senderId = msg.senderId,
                sentAt = msg.sentAt.toString(),
                type = payload?.__typename ?: "StandardMessage",
                content = payload?.onStandardMessage?.content,
                tags = null,
                questionId = null,
                title = null,
                body = null,
                actionUrl = null,
                imageUrl = null
            )
        }
    }

@JsExport
fun sendQuestionMessage(chatId: String, content: String, tags: Array<String>): Promise<JsMessage?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        chatServiceInstance.sendQuestionMessage(chatId, content, tags.toList())?.let { msg ->
            val payload = msg.payload
            JsMessage(
                id = msg.id,
                chatId = msg.chatId,
                senderId = msg.senderId,
                sentAt = msg.sentAt.toString(),
                type = payload?.__typename ?: "QuestionMessage",
                content = payload?.onQuestionMessage?.content,
                tags = payload?.onQuestionMessage?.tags?.toTypedArray(),
                questionId = null,
                title = null,
                body = null,
                actionUrl = null,
                imageUrl = null
            )
        }
    }

@JsExport
fun sendAnswerMessage(chatId: String, content: String, questionId: String): Promise<JsMessage?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        chatServiceInstance.sendAnswerMessage(chatId, content, questionId)?.let { msg ->
            val payload = msg.payload
            JsMessage(
                id = msg.id,
                chatId = msg.chatId,
                senderId = msg.senderId,
                sentAt = msg.sentAt.toString(),
                type = payload?.__typename ?: "AnswerMessage",
                content = payload?.onAnswerMessage?.content,
                tags = null,
                questionId = payload?.onAnswerMessage?.questionId,
                title = null,
                body = null,
                actionUrl = null,
                imageUrl = null
            )
        }
    }

@JsExport
fun sendPromotionalMessage(
    chatId: String,
    title: String,
    body: String,
    actionUrl: String,
    imageUrl: String?
): Promise<JsMessage?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        chatServiceInstance.sendPromotionalMessage(chatId, title, body, actionUrl, imageUrl)?.let { msg ->
            val payload = msg.payload
            JsMessage(
                id = msg.id,
                chatId = msg.chatId,
                senderId = msg.senderId,
                sentAt = msg.sentAt.toString(),
                type = payload?.__typename ?: "PromotionalMessage",
                content = null,
                tags = null,
                questionId = null,
                title = payload?.onPromotionalMessage?.title,
                body = payload?.onPromotionalMessage?.body,
                actionUrl = payload?.onPromotionalMessage?.actionUrl,
                imageUrl = payload?.onPromotionalMessage?.imageUrl
            )
        }
    }

@JsExport
fun createChat(generalUserId: String, professionalUserId: String): Promise<JsChat?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        chatServiceInstance.createChat(generalUserId, professionalUserId)?.let { chat ->
            JsChat(
                id = chat.id,
                generalUserId = chat.generalUserID,
                professionalUserId = chat.professionalUserID,
                createdAt = chat.createdAt.toString()
            )
        }
    }

