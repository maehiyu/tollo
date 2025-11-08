pluginManagement {
    repositories {
        google()
        mavenCentral()
        gradlePluginPortal()
    }
}

// 依存ライブラリを探す場所を定義します
dependencyResolutionManagement {
    repositoriesMode.set(RepositoriesMode.FAIL_ON_PROJECT_REPOS)
    repositories {
        google()
        mavenCentral() // <- これでkotlin-scripting-compilerが見つかるようになります
    }
}

rootProject.name = "tollo"
include(":client:shared")
