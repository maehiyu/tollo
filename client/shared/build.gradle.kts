plugins {
  id("com.apollographql.apollo") version "4.3.3"
  kotlin("multiplatform") version "1.9.22"
}

kotlin {
  // A target is required by the Kotlin Multiplatform plugin
  jvm() 

  sourceSets {
    commonMain.dependencies {
      implementation("com.apollographql.apollo:apollo-runtime:4.3.3")
    }
  }
}

apollo {
  service("tollo") {
    packageName.set("io.github.maehiyu.tollo.client.shared")
    schemaFiles.from(file("src/commonMain/graphql/schema.graphqls"))
    srcDir("src/commonMain/graphql")
  }
}