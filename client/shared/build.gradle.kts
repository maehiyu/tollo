plugins {
  id("com.apollographql.apollo") version "4.3.3"
  kotlin("multiplatform") version "2.2.21"
  kotlin("plugin.serialization") version "2.2.21"
}

kotlin {
  // A target is required by the Kotlin Multiplatform plugin
  jvm()
  js(IR) {
    useEsModules() // Add this line
    browser {
      binaries.library()
    }

  }

  sourceSets {
    val commonMain by getting{
      dependencies {
        implementation("com.apollographql.apollo:apollo-runtime:4.3.3")
        implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core:1.9.0")
      }
    }
    val commonTest by getting {
      dependencies {
        implementation(kotlin("test"))
      }
    }
    val jvmMain by getting
    val jsMain by getting {
      dependencies {
        implementation("org.jetbrains.kotlinx:kotlinx-coroutines-core-js:1.9.0")
      }
    }
  }
}

apollo {
    service("api") {
        packageName.set("io.github.maehiyu.tollo.client.shared")
        schemaFiles.from(file("src/commonMain/graphql/schema.graphqls"))
    }
}