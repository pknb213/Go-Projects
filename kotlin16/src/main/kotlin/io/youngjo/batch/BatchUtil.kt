package io.youngjo.batch

import reactor.util.Loggers

object BatchUtil {
    private val log = Loggers.getLogger(BatchUtil.javaClass)
    @Suppress("UNCHECKED_CAST")
    @Throws(Exception::class)
    @JvmStatic
    fun main(args: Array<String>) {
        Server.logInit()

        
    }
}