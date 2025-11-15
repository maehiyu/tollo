package io.github.maehiyu.tollo.client.shared.domain.repository

import io.github.maehiyu.tollo.client.shared.CreateUserMutation
import io.github.maehiyu.tollo.client.shared.GetUserQuery

import io.github.maehiyu.tollo.client.shared.type.GeneralProfileInput
import io.github.maehiyu.tollo.client.shared.type.ProfessionalProfileInput

interface UserRepository {
    suspend fun getUser(id: String?, email: String?): GetUserQuery.User?
    suspend fun createUser(
      name: String,
      email: String, 
      description: String?,
      general: GeneralProfileInput?,
      professional: ProfessionalProfileInput?
    ): CreateUserMutation.CreateUser?
}
