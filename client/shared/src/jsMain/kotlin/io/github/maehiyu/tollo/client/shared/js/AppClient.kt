@file:OptIn(ExperimentalJsExport::class)

package io.github.maehiyu.tollo.client.shared.js

import com.apollographql.apollo.api.Optional
import io.github.maehiyu.tollo.client.shared.createApolloClient
import io.github.maehiyu.tollo.client.shared.data.repository.UserRepositoryImpl
import io.github.maehiyu.tollo.client.shared.domain.repository.UserRepository
import io.github.maehiyu.tollo.client.shared.domain.service.UserService
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

private val apolloClient = createApolloClient(DevAuthContext)
private val userRepository: UserRepository = UserRepositoryImpl(apolloClient)

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
val userService: UserService = UserService(userRepository)

@JsExport
fun getUser(id: String): Promise<JsUser?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        userService.getUser(id = id, email = null)?.let { user ->
            JsUser(user.id, user.name, user.email, user.createdAt.toString())
        }
    }

@JsExport
fun getUserByEmail(email: String): Promise<JsUser?> =
    CoroutineScope(Dispatchers.Unconfined).promise {
        userService.getUser(id = null, email = email)?.let { user ->
            JsUser(user.id, user.name, user.email, user.createdAt.toString())
        }
    }

@JsExport
fun createUser(
  name: String,
  email: String,
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
      userService.createUser(name, email, description, generalInput, professionalInput)?.let { user ->
          JsUser(user.id, user.name, user.email, user.createdAt.toString())
      }
    }

