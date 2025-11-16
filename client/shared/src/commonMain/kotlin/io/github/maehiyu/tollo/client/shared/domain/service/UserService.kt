package io.github.maehiyu.tollo.client.shared.domain.service

import io.github.maehiyu.tollo.client.shared.CreateUserMutation
import io.github.maehiyu.tollo.client.shared.GetUserQuery
import io.github.maehiyu.tollo.client.shared.domain.repository.UserRepository
import io.github.maehiyu.tollo.client.shared.type.GeneralProfileInput
import io.github.maehiyu.tollo.client.shared.type.ProfessionalProfileInput

class UserService(private val userRepository: UserRepository) {
    suspend fun getUser(id: String?, email: String?): GetUserQuery.User? {
        return userRepository.getUser(id, email)
    }

    suspend fun createUser(
      name: String, 
      email: String, 
      description: String?,
      general: GeneralProfileInput?,
      professional: ProfessionalProfileInput?
    ): CreateUserMutation.CreateUser? {
        return userRepository.createUser(name, email, description, general, professional)
    }
}
