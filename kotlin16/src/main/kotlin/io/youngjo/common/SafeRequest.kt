package io.youngjo.common

import io.netty.buffer.ByteBuf
import io.netty.handler.codec.http.*
import io.netty.handler.codec.http.multipart.Attribute
import io.netty.handler.codec.http.multipart.FileUpload
import io.netty.handler.codec.http.multipart.HttpPostMultipartRequestDecoder
import io.netty.handler.codec.http.multipart.InterfaceHttpData
import reactor.netty.http.server.HttpServerRequest
import reactor.netty.http.websocket.WebsocketInbound
import reactor.util.Loggers
import java.net.URLDecoder
import java.nio.charset.Charset

class SafeRequest {
    private val log = Loggers.getLogger(this.javaClass)
    private lateinit var req: HttpServerRequest
    private lateinit var wsi: WebsocketInbound
    private lateinit var paramNew: Map<String, String>
    private lateinit var queryNew: Map<String, List<String>>
    private lateinit var bodyString: String
    private lateinit var file: Map<String, Map<String, Any>>
    private lateinit var bodyParam: Map<String, String>

    // TODO wsi
    private val param: Map<String, String> by lazy(LazyThreadSafetyMode.NONE) {
        if(::paramNew.isInitialized) paramNew else req.params() ?: mapOf()
    }
    // TODO wsi
    private val query: Map<String, List<String>> by lazy(LazyThreadSafetyMode.NONE) {
        if(::queryNew.isInitialized) queryNew else QueryStringDecoder(req.uri()).parameters()
    }
    private val headers: HttpHeaders by lazy(LazyThreadSafetyMode.NONE) {
        if(req != null) req.requestHeaders() else wsi.headers()
    }

    constructor(req: HttpServerRequest, byteBuf: ByteBuf) {
        this.req = req
//		log.info("1 ${byteBuf.refCnt()}")
        val contentType = headers.get(HttpHeaderNames.CONTENT_TYPE)

        if(contentType.isNullOrEmpty()){
            bodyString = byteBuf.toString(Charset.defaultCharset())
            bodyParam = mapOf()
            file = mapOf()
            byteBuf.release()
        }else if(contentType.startsWith(HttpHeaderValues.APPLICATION_JSON)){
            bodyString = byteBuf.toString(Charset.defaultCharset())
            bodyParam = if(bodyString.trim().first() == '{') Util.jsonToMap(bodyString) as Map<String, String> else mapOf()
            file = mapOf()
            byteBuf.release()
        }else if(contentType.startsWith(HttpHeaderValues.APPLICATION_X_WWW_FORM_URLENCODED)){
            bodyString = byteBuf.toString(Charset.defaultCharset())
            bodyParam = URLDecoder.decode(bodyString, Charset.defaultCharset())
                .split("&").map { it.split("=") }.associateBy({it[0]},{it[1]})
            file = mapOf()
            byteBuf.release()
        }else if(contentType.startsWith(HttpHeaderValues.MULTIPART_FORM_DATA)){
            val decoder = HttpPostMultipartRequestDecoder(
                DefaultFullHttpRequest(req.version(), req.method(), req.uri(), byteBuf, headers, headers)
            )

            val result = decoder.bodyHttpDatas
                .partition { it.httpDataType == InterfaceHttpData.HttpDataType.FileUpload }

            file = result.first.associateBy({it.name}, {
                it as FileUpload
                mapOf(
                    "body" to it.content().nioBuffer(),
                    "type" to it.contentType,
                    "name" to it.filename)
            })

            bodyParam = result.second.associateBy({it.name}, {
                (it as Attribute).content().toString(Charset.defaultCharset())
            })
            bodyString = ""

            byteBuf.release()
            decoder.destroy()
        }else{
            bodyString = byteBuf.toString(Charset.defaultCharset())
            bodyParam = mapOf()
            file = mapOf()
            byteBuf.release()
        }
//		log.info("2 ${byteBuf.refCnt()}")
    }
}