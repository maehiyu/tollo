package io.github.maehiyu.tollo.client.shared.data.repository

import com.apollographql.apollo.ApolloClient
import com.apollographql.apollo.api.Optional
import io.github.maehiyu.tollo.client.shared.CreateUserMutation
import io.github.maehiyu.tollo.client.shared.GetUserQuery
import io.github.maehiyu.tollo.client.shared.MeQuery
import io.github.maehiyu.tollo.client.shared.domain.repository.UserRepository
import io.github.maehiyu.tollo.client.shared.type.GeneralProfileInput
import io.github.maehiyu.tollo.client.shared.type.ProfessionalProfileInput

class UserRepositoryImpl(private val apolloClient: ApolloClient) : UserRepository {

    override suspend fun getUser(id: String?, email: String?): GetUserQuery.User? {
        return try {
            val response = apolloClient.query(
                GetUserQuery(
                    id = Optional.presentIfNotNull(id),
                    email = Optional.presentIfNotNull(email)
                )
            ).execute()
            if (response.hasErrors()) {
                println("GraphQL Errors: ${response.errors}")
                return null
            }
            println("GetUser Response Data: ${response.data}")
            println("Searched User: ${response.data?.user}")
            response.data?.user
        } catch (e: Exception) {
            println("Network or other exception: ${e.message}")
            return null
        }
    }

    override suspend fun getCurrentUser(): MeQuery.Me? {
        return try {
            val response = apolloClient.query(MeQuery()).execute()
            if (response.hasErrors()) {
                println("GraphQL Errors: ${response.errors}")
                return null
            }
            println("Me Response Data: ${response.data}")
            println("Current User: ${response.data?.me}")
            response.data?.me
        } catch (e: Exception) {
            println("Network or other exception: ${e.message}")
            return null
        }
    }

    override suspend fun createUser(
      name: String,
      description: String?,
      general: GeneralProfileInput?,
      professional: ProfessionalProfileInput?
    ): CreateUserMutation.CreateUser? {
        return try {
            val response = apolloClient.mutation(
                CreateUserMutation(
                    name = name,
                    description = Optional.presentIfNotNull(description),
                    general = Optional.presentIfNotNull(general),
                    professional = Optional.presentIfNotNull(professional)
                )
            ).execute()
            if (response.hasErrors()) {
                println("GraphQL Errors: ${response.errors}")
                return null
            }
            println("CreateUser Response Data: ${response.data}")
            println("Created User: ${response.data?.createUser}")
            response.data?.createUser
        } catch (e: Exception) {
            println("Network or other exception: ${e.message}")
            return null
        }
    }
}
