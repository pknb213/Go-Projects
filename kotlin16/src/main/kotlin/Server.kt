import com.mongodb.reactivestreams.client.MongoClients
import reactor.core.publisher.Flux
import reactor.core.publisher.Mono
import reactor.core.scheduler.Schedulers
import reactor.netty.http.server.HttpServer
import MongoUtil
import ch.qos.logback.classic.Level
import ch.qos.logback.classic.encoder.PatternLayoutEncoder
import ch.qos.logback.core.ConsoleAppender
import com.mongodb.client.model.UpdateOneModel
import com.mongodb.client.model.UpdateOptions
import io.youngjo.common.MongodbUtil
import io.youngjo.common.Util
import org.bson.Document
import org.bson.types.ObjectId
import reactor.util.Loggers.getLogger
import reactor.util.function.Tuples
import java.lang.Math.random
import java.time.Duration
import java.util.*
import kotlin.concurrent.timer
import org.jetbrains.kotlinx.spark.api.*
import org.jetbrains.kotlinx.spark.api.SparkLogLevel.*
import org.slf4j.Logger
import org.slf4j.LoggerFactory
import reactor.netty.http.server.HttpServerRequest
import reactor.util.Loggers


object Server {
    private val log = Loggers.getLogger(this.javaClass)
    init {
        logInit()
        val TEST_DB_NAME = "userhabit"
        val TEST_COLLECTION_NAME = "app"
    }
    /** Test */
    fun get(): String {
        val client = MongoClients.create("mongodb://localhost:27017")
        val db = client.getDatabase("testdb")
        val coll_ls = db.listCollections()
        val coll = db.getCollection("test")
        val doc = coll.find()
        val ls = client.listDatabases()
        val mono = Flux
            .from(doc)
            .map {
                println(">> $it")
                it
            }
            .subscribe()
        return mono.toString()
    }

    @JvmStatic
    fun main(args: Array<String>) {
        log.debug(System.getProperties().map { it.toString().plus("\n") }.sorted().toString())
        log.info("hostname : ${Util.hostName}")

        val logger = LoggerFactory.getLogger(Logger.ROOT_LOGGER_NAME) as ch.qos.logback.classic.Logger
        logger.level = if(System.getProperty("log.level") == "debug") Level.DEBUG else Level.INFO

        println("<< Welcome >>")
//        val log = Loggers.getLogger(this.javaClass)
        val server = HttpServer.create()
            .route{
                it
                    .get("/ping") { req, res -> res.sendString(Mono.just("Pong\n")) }
                    .get("/mongo") { req, res -> res.sendString(Mono.just(get())) }
                    .get("/test") { req, res -> res.sendString(Mono.just("1")) }
<<<<<<< rebert
=======
//                    .post("/flume") { req, res ->
//                        res.sendString(req.receive().aggregate().map {
//                            println(it.toString(Charset.defaultCharset()))
////                            Mono.from(MongodbUtil.getCollection("test").bulkWrite(
////                                listOf(
////                                    UpdateOneModel<Document>(
////                                        Document("_id", ObjectId() ),
////                                        Document("\$set", Document("data", it)),
////                                        UpdateOptions().upsert(true)
////                                    )
////                                )
////                            )).subscribe()
//                            it.release()
//                            it.toString()
//                        })
//                    }
                    .get("/flume2", ServiceHelper.http(TestService::get))
                    .post("/flume", ServiceHelper.http(TestService::post))
>>>>>>> update
            }
            .port(8000)
            .compress(true)
            .bindNow()

        server.onDispose().block()
    }

    fun logInit(){
        // TODO delete (for log4j)
//		Configurator.setRootLevel(if(Util.isDevMode ) Level.DEBUG else Level.INFO)
//		Loggers.resetLoggerFactory()

        val logger = LoggerFactory.getLogger(Logger.ROOT_LOGGER_NAME) as ch.qos.logback.classic.Logger
        logger.level = Level.INFO
//		logger.isAdditive = true
        (logger.getAppender("console") as ConsoleAppender).let { app ->
            app.encoder = PatternLayoutEncoder().apply{
//				pattern = "%d{yyyy-MM-dd HH:mm:ss} [%thread] %highlight(%-5level) %logger{36} - %msg%n"
                pattern = "%d{HH:mm:ss} [%thread] %highlight(%-5level) %logger{36} - %msg%n"
                context = app.context
                start()
            }
        }
    }
}