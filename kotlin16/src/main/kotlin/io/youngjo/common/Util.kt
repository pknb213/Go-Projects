package io.youngjo.common

import com.google.gson.*
import org.bson.Document
import org.bson.json.JsonMode
import org.bson.json.JsonWriterSettings
import java.net.InetAddress
import java.nio.file.FileSystems
import java.nio.file.Files
import java.nio.file.Path
import java.util.*
import kotlin.io.path.isDirectory
import kotlin.io.path.toPath

object Util {
    // TODO Gson()을 사용하면 json number 값을 파싱할 때 지수로 표현한다. (ex 1640050762904 >> 1.640050762904E12)
    private val gson = GsonBuilder()
        .setObjectToNumberStrategy(ToNumberPolicy.LONG_OR_DOUBLE)
        .create()
    private val gsonPretty = GsonBuilder().setPrettyPrinting().create()
    private val classMap = mutableMapOf<String, Any>().javaClass
    private val classList = mutableListOf<Map<String, Any>>().javaClass
    val hostName = InetAddress.getLocalHost().getHostName()

    /**
     *
     */
    fun toJsonString(list: Any): String {
        return gson.toJson(list)
    }

    /**
     *
     */
    fun toMapJsonString(listMap: List<MutableMap<String, Any>>?): String {
        return gson.toJson(listMap)
    }

    /**
     *
     */
    fun jsonToMap(json: String): Map<String, Any> {
        return gson.fromJson(json, classMap)
    }

    /**
     *
     */
    fun jsonToMutableMap(json: String): MutableMap<String, Any> {
        return gson.fromJson(json, classMap)
    }

    /**
     *
     */
    fun jsonToList(json: String): List<Map<String, Any>> {
        return gson.fromJson(json, classList)
    }

    /**
     *
     */
    fun jsonToMap(json: ByteArray): Map<String, Any> {
        return gson.fromJson(String(json), classMap)
    }

    /**
     *
     */
    fun jsonToList(json: ByteArray): List<Map<String, Any>> {
        return gson.fromJson(String(json), classList)
    }

    /**
     *
     */
    fun toPretty(json: Any): String {
        return gsonPretty.toJson(json)
    }

    /**
     *
     */
    fun <R> getResource(path: String, func: (Path) -> R): R {
        val uri = ClassLoader.getSystemResource(path).toURI()
        return	if(uri.scheme == "jar") FileSystems.newFileSystem(uri, Collections.emptyMap<String, Object>()).use{
            func(it.getPath("/${path}"))
        }
        else func(uri.toPath()) as R
    }

    /**
     *
     */
    fun <R> getResourceList(path: String, func: (Path) -> R): List<R> {
        val uri = ClassLoader.getSystemResource(path).toURI()
        return if(uri.scheme == "jar") FileSystems.newFileSystem(uri, Collections.emptyMap<String, Object>()).use {
            val p = it.getPath("/${path}")
            Files.walk(p, 1)
                .filter { !it.isDirectory() }
                .map {
                    func(it)
                }
                .toList()
        }
        else uri.toPath().toFile().listFiles().filter { !it.isDirectory }.map { func(it.toPath()) }
    }

    private val algorithm = "HmacSHA256"
    private val base64UrlEncoder = Base64.getUrlEncoder().withoutPadding()
    val base64UrlDecoder = Base64.getUrlDecoder()!!
}