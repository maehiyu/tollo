plugins {
  id("com.apollographql.apollo") version "4.3.3"
  kotlin("multiplatform") version "2.2.21"
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
    commonMain.dependencies {
      implementation("com.apollographql.apollo:apollo-runtime:4.3.3")
    }
  }
}

apollo {
  service("api") {
    packageName.set("io.github.maehiyu.tollo.client.shared")
    schemaFiles.from(file("src/commonMain/graphql/schema.graphqls"))
    srcDir("src/commonMain/graphql")
  }
}
