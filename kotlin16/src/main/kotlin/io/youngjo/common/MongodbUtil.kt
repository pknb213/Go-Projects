package io.youngjo.common

import com.mongodb.reactivestreams.client.ClientSession
import com.mongodb.reactivestreams.client.MongoClients
import com.mongodb.reactivestreams.client.MongoCollection
import org.bson.Document
import reactor.core.publisher.Mono
import reactor.util.Loggers

object MongodbUtil {
    private val log = Loggers.getLogger(this.javaClass)
    val SAAS_DB_NAME = "youngjo_db"

    private val client by lazy {

//		ClusterSettings
//			.builder()
//			.hosts(
//				Config.get("mongodb.addresses").split(",").map { ServerAddress(it, Config.getInt("mongodb.port")) }
//			)
//			.build()
//			.let { cs ->
//				MongoClientSettings
//					.builder()
//					.applyToClusterSettings {
//						it.applySettings(cs)
//					}.build()
//			}
//			.let { mcs ->
//					MongoClients.create(mcs)
//			}

        val addresses = Config.get("mongodb.addresses")
        val replicaSet = Config.get("mongodb.replicaSet")
        val readPreference = Config.get("mongodb.readPreference")
        if (addresses.contains(",")) MongoClients.create("mongodb://${addresses}/?replicaSet=${replicaSet}&readPreference=${readPreference}")
        else MongoClients.create("mongodb://${addresses}")
    }


    private val db by lazy {
        client.getDatabase(SAAS_DB_NAME)
    }

    fun getCollection(collectionName: String): MongoCollection<Document> {
        return db.getCollection(collectionName)
    }

    fun startSession(): Mono<ClientSession> {
        return Mono.from(client.startSession())
    }

    // when the server starts, it reads the json files and upserts the collection data into mongodb
    init {
//        val initColl = System.getProperty("initcoll")
//        runMongoServer()
//        if (initColl == "true") {
//            initCollection()
//        }
//        BatchUtil.start() // 배치 초기화
    } // end of init()
}