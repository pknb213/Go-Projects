plugins {
    kotlin("jvm") version "1.6.0"
    application
}

group = "org.example"
version = "1.0-SNAPSHOT"

repositories {
    mavenCentral()
}

dependencies {
    implementation(kotlin("stdlib"))
//    implementation("org.jetbrains.kotlinx.spark", "kotlin-spark-api-common", "1.0.1")
    implementation("org.jetbrains.kotlinx.spark", "kotlin-spark-api-3.0", "1.0.1")
    compileOnly("org.apache.spark", "spark-sql_2.12", "3.+")
    implementation("io.projectreactor.netty" , "reactor-netty" , "1.+")
    implementation("ch.qos.logback" , "logback-classic" , "1.+")
//    implementation("org.slf4j" , "slf4j-log4j12" , "1.7.30")
    implementation("org.mongodb" , "mongodb-driver-reactivestreams" , "4.+")
    implementation("com.fasterxml.jackson.core" , "jackson-databind" , "2.10.4")
    implementation(kotlin("reflect"))
    implementation("com.auth0" , "java-jwt" , "3.10.3")
    implementation("com.google.code.gson" , "gson" , "2.+")
}

tasks {
    compileKotlin {
        kotlinOptions.jvmTarget = "16"
    }
    compileTestKotlin {
        kotlinOptions.jvmTarget = "16"
    }
}

application {
    mainClassName = "Server"

    val taskNames = gradle.startParameter.taskNames

    if( taskNames.contains("dev") ){
        applicationDefaultJvmArgs = project.ext.properties.map { "-D${it.key}=${it.value}" } + "-Denv=dev"
        println(">> D E V")
    }
}

task("dev"){
    dependsOn("run")
}