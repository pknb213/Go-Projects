package io.youngjo.common

import reactor.util.Loggers
import java.io.StringReader
import java.nio.file.Files
import java.nio.file.Paths
import java.util.*
import kotlin.io.path.absolutePathString

object Config {
    private val prop = Properties()
    private val env = System.getProperty("env") ?: "dev"
    private val log = Loggers.getLogger(this.javaClass)
    val isDevMode = "dev" == env

    val a = let {
        Util.getResource("config.common.properties") {
            prop.load(StringReader(Files.readString(it)))
        }
        val fileName ="config.$env.properties"
        Util.getResource(fileName){
            prop.load(StringReader(Files.readString(it)))

            var p = it.absolutePathString()
            val p2 = Paths.get(prop.getProperty("config.path")) // 지정된 경로에 파일이 있을 경우
            if (p2.toFile().isFile) {
                p = p2.toUri().path
                Files.newBufferedReader(p2).use { prop.load(it) }
            } else {
                log.warn("[$p2] File not found in config.properties file")
            }
            log.info("config.properties path: $p")
            log.debug(prop.toString())
        }
    }

    @JvmStatic
    fun get(key: String): String {
        return prop.getProperty(key).trim();
    }

    @JvmStatic
    fun getInt(key: String): Int {
        return prop.getProperty(key).trim().toInt()
    }
}