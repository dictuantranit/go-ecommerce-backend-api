package initialize

import "github.com/gin-gonic/gin"

func Run() *gin.Engine {
	LoadConfig()
	InitLogger()
	InitMysql()
	InitMySqlC()
	InitRedis()
	InitKafka()

	r := InitRouter()
	r.Run(":8002")
	return r
}
