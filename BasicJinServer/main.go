package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/assert/v2"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
	"time"
)

type nameForm struct {
	Firstname string
}

type myInfo struct {
	phone string
	age   int
	city  string
}

// 개인정보 샘플 데이터
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// 샘플 변수 설정
		c.Set("example", "12345")

		// Request 이전

		c.Next()

		// Request 이후
		latency := time.Since(t)
		log.Print(latency)

		// 송신할 상태 코드에 접근
		status := c.Writer.Status()
		log.Println(status)
	}
}

// Booking은 유효성 검사 후 바인딩 된 데이터를 갖습니다.
type Booking struct {
	CheckIn  time.Time `form:"check_in" binding:"required,bookabledate" time_format:"2006-01-02"`
	CheckOut time.Time `form:"check_out" binding:"required,gtfield=CheckIn,bookabledate" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
	r.GET("/myHandler", func(c *gin.Context) {
		objA := nameForm{}
		objB := myInfo{}
		// 아래 c.ShouldBind는 c.Request.Body를 소모하며, 재 사용이 불가능하다.
		if errA := c.ShouldBind(&objA); errA == nil {
			c.String(http.StatusOK, `the body should be nameForm`)
			// 아래 c.Reqeust.Body가 EOF이므로 에러가 발생.
		} else if errB := c.ShouldBind(&objB); errB == nil {
			c.String(http.StatusOK, `the body should be myInfo`)
		} else {
		}
	})
	r.GET("/myHandlerFix", func(c *gin.Context) {
		objA := nameForm{}
		objB := myInfo{}
		if errA := c.ShouldBindBodyWith(&objA, binding.JSON); errA == nil {
			c.String(http.StatusOK, `the body should be nameForm JSOn`)
		} else if errB := c.ShouldBindBodyWith(&objB, binding.JSON); errB == nil {
			c.String(http.StatusOK, `the body should be myInfo JSON`)
		} else if errC := c.ShouldBindBodyWith(&objB, binding.JSON); errC == nil {
			c.String(http.StatusOK, `the body should be myInfo XML`)
		} else {
		}
	})
	r.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("paeg", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
	})
	r.GET("/user/:id", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "Hello %s", id)
	})
	r.GET("/user/:id/*action", func(c *gin.Context) {
		id := c.Param("id")
		action := c.Param("action")
		message := id + " is " + action
		c.String(http.StatusOK, "Hello %s", message)
	})
	// gin.BasicAuth() 미들웨어를 사용하여 그룹화 하기
	// gin.Accounts는 map[string]string 타입입니다.
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))
	// 엔드포인트 /admin/secrets
	// "localhost:8080/admin/secrets"로 접근합니다.
	authorized.GET("/secrets", func(c *gin.Context) {
		// BasicAuth 미들웨어를 통해 유저를 설정합니다.
		user := c.MustGet(gin.AuthUserKey).(string)
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// v1 그룹
	v1 := r.Group("/v1")
	{
		v1.POST("/login", func(context *gin.Context) {
			// loginEndpoint
		})
		v1.POST("/submit", func(context *gin.Context) {
			// submitEndpoint
		})
		v1.POST("/read", func(context *gin.Context) {
			// readEndpoint
		})
	}

	// v2 그룹
	v2 := r.Group("/v2")
	{
		v2.POST("/login", func(context *gin.Context) {
			// loginEndpoint
		})
		v2.POST("/submit", func(context *gin.Context) {
			// submitEndpoint
		})
		v2.POST("/read", func(context *gin.Context) {
			// readEndpoint
		})
	}
	// Todo: Goroutine example !!
	r.GET("/long_async", func(c *gin.Context) {
		// Go루틴 내부에서 사용하기 위한 복사본을 작성합니다.
		cCp := c.Copy()
		go func() {
			// time.Sleep()를 사용하여 장시간(5초) 작업을 시뮬레이션 합니다.
			time.Sleep(5 * time.Second)

			// 중요! 복사된 context인 "cCp"를 사용하세요.
			log.Println("Done! in path " + cCp.Request.URL.Path)
		}()
	})

	r.GET("/long_sync", func(c *gin.Context) {
		// time.Sleep()를 사용하여 장시간(5초) 작업을 시뮬레이션 합니다.
		time.Sleep(5 * time.Second)

		// Go루틴을 사용하지 않는다면, context를 복사할 필요가 없습니다.
		log.Println("Done! in path " + c.Request.URL.Path)
	})
	r.Use(Logger())
	r.GET("/test", func(c *gin.Context) {
		example := c.MustGet("example").(string)
		// 출력내용: "12345"
		log.Println(example)
	})
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("bookabledate", func(fl validator.FieldLevel) bool {
			return true
		}) // Todo: 왜 안되는지 모르겟네: bookableDate
	}
	r.GET("/bookable", getBookable)
	// 내가 만든 쿠키~
	r.GET("/cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("gin_cookie")
		if err != nil {
			cookie = "NotSet"
			c.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s \n", cookie)
	})
	// 멀티파트 폼에 대한 최저 메모리 설정 (기본값 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 단일 파일
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		// 특정 경로(dst)에 파일을 업로드 합니다.
		c.SaveUploadedFile(file, "./")
		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})
	// 멀티파트 폼에 대한 최저 메모리 설정 (기본값 32 MiB)
	r.MaxMultipartMemory = 8 << 20 // 8 MiB
	r.POST("/upload", func(c *gin.Context) {
		// 멀티파트 폼
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]

		for _, file := range files {
			log.Println(file.Filename)

			// 특정 경로(dst)에 파일을 업로드 합니다.
			c.SaveUploadedFile(file, "./")
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})
	r.Run()
	r2 := setupRouter()
	r2.Run("8081")
}
func getBookable(c *gin.Context) {
	var b Booking
	if err := c.ShouldBindWith(&b, binding.Query); err == nil {
		c.JSON(http.StatusOK, gin.H{"message": "Booking dates are valid!"})
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
}

func bookableDate(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if date, ok := field.Interface().(time.Time); ok {
		today := time.Now()
		if today.Year() > date.Year() || today.YearDay() > date.YearDay() {
			return false
		}
	}
	return true
}
func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
