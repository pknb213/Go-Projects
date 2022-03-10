package io.youngjo.service

import com.mongodb.client.model.UpdateOneModel
import com.mongodb.client.model.UpdateOptions
import io.youngjo.common.MongodbUtil
import io.youngjo.common.SafeRequest
import org.bson.Document
import org.bson.types.ObjectId
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import reactor.util.Loggers

object TestService {
    private val log = Loggers.getLogger(this.javaClass)

    fun get(req: SafeRequest): Mono<Map<String, Any>> {
        println(banana(apple(30)))
        return Flux.just("test").collectList().map { mapOf(it.toString() to 1)}
    }

    fun post(req: SafeRequest): Mono<Map<String, Any>> {
        /**
        listing_id
        id
        date
        reviewer_id
        reviewer_name
        comments
         */
        val body = req.getBody().trim()
        val model = listOf(UpdateOneModel<Document>(
            Document("_id", ObjectId() ),
            Document("\$set", mapOf("data" to body)),
            UpdateOptions().upsert(true)
        ))
        return Mono.from(MongodbUtil.getCollection("test").bulkWrite(
            model
        )).map { mapOf("test" to body) }
    }

    interface flute {
        fun buy() = println("Please. I want Buy")
        fun sell() = println("Okay. I Sell")
        val count: Int
    }

    class rice: flute {
        override var count: Int = 10
        //            get() = TODO("Not yet implemented")
        fun hello(price: Int): Int{
            return count + price
        }
    }

    fun apple(fee: Int): flute {
        println(">> $fee")
        val r = rice()
        println(r.hello(fee))
        r.count = r.hello(fee)
        return r
    }

    fun banana(item: flute): Int {
        item.buy()
        item.sell()
        return item.count
    }
}