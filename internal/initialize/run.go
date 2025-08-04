package initialize

func Run() {
	// load configuration
	LoadConfig()
	InitLogger()
	InitMysql()
	InitRedis()

	r := InitRouter()

	r.Run(":8002")
}
