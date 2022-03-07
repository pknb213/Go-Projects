package io.youngjo.service

import io.youngjo.common.SafeRequest
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import reactor.util.Loggers

object TestService {
    private val log = Loggers.getLogger(this.javaClass)

    fun get(req: SafeRequest): Mono<Map<String, Any>> {

        return Flux.just("test").collectList().map { mapOf(it.toString() to 1)}
    }
}