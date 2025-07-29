package main

import (
	controller_user "SlotGameServer/pkgs/controller/user"
	i18n_repo "SlotGameServer/pkgs/dao/i18n/repo"
	repo_pgsql "SlotGameServer/pkgs/dao/postgresql/repo"
	repo_redis "SlotGameServer/pkgs/dao/redis/repo"
	service_user "SlotGameServer/pkgs/service/user"
	"SlotGameServer/utils"
	utils_middleware "SlotGameServer/utils/middleware"
	"context"
	"flag"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

var (
	verFlag    = flag.Bool("version", false, "Print version information")
	configPath = flag.String("config", "./server.yaml", "The path to the config file. For more information, see the config file in this repository.")
	version    string
	buildTime  string
	branch     string
	commitId   string
)

func main() {
	flag.Parse()

	if *verFlag {
		logrus.Infof("Version: %s\nBuildTime: %s\nBranch: %s\nCommitId: %s\n", version, buildTime, branch, commitId)
		os.Exit(0)
	}

	if *configPath == "" {
		logrus.Fatal("--config must be supplied")
	}

	// 初始化配置
	utils.SetupConfig(*configPath)
	// 日志配置
	utils.SetupLogging()
	// 初始化邮箱
	utils.SetupEmail()
	// 初始化数据库
	utils.SetupPostgreSQL()
	// 初始化Redis
	utils.SetupRedis()
	// 初始化Geo
	utils.SetupGeo()

	// 链路追踪
	closer, err := utils.SetupTracing()
	if err != nil {
		logrus.Fatal("failed to start opentracing")
	}
	defer closer.Close()

	r, cleanup := setupRouter()
	defer cleanup()
	// 创建HTTP服务器
	srv := &http.Server{
		Addr:    utils.Conf.Server.Domain,
		Handler: r,
	}

	go func() {
		logrus.Infof("Server is running at %s", utils.Conf.Server.Domain)
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	<-quit

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("Server forced to shutdown: ", err)
	}

	logrus.Info("Server exiting")
}

func setupRouter() (*gin.Engine, func()) {
	gin.SetMode(gin.ReleaseMode)
	r := gin.New()

	r.Use(gin.Recovery())

	// 添加CORS中间件
	// r.Use(cors.Default())
	// 自定义CORS配置
	corsConfig := cors.Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type", "Authorization"},
		AllowCredentials: false,
		AllowAllOrigins:  true,
		ExposeHeaders:    []string{"Content-Length", "Content-Type", "Date", "Uber-Trace-Id"},
		MaxAge:           12 * time.Hour,
	}
	r.Use(cors.New(corsConfig))

	r.Use(utils_middleware.AgentMiddleware())

	if utils.Conf.Tracing.Enabled {
		r.Use(utils_middleware.TracerMiddleware())
	}

	tmpUserService := &service_user.UserService{
		VerifyCodeRedisRepo: repo_redis.NewVerifyCodeRedisRepo(utils.RedisClient),
		UserDBRepo:          repo_pgsql.NewUserDBRepo(utils.PostgresDB[utils.PostgresClub]),
		UserRedisRepo:       repo_redis.NewUserRedisRepo(utils.RedisClient),
		EmailI18NRepo:       i18n_repo.NewEmailI18NRepo(utils.I18nTranslator),
	}
	tmpUserService.CheckInitialization()

	// API v1 路由组
	v1 := r.Group("/v1")
	{
		//用户
		userController := controller_user.NewUserController(tmpUserService)
		userNoAuth := v1.Group("/user")
		{
			userNoAuth.POST("/signup", userController.SignUp)
			userNoAuth.POST("/signin", userController.SignIn)
			userNoAuth.POST("/verify/code", userController.VerifyCode)
		}
		userAuth := v1.Group("/user", utils_middleware.ForceAuthMiddleware())
		{
			// userAuth.DELETE("/delete", userController.DeleteUser)
			userAuth.GET("/profile", userController.GetProfile)
			// userAuth.PUT("/username", userController.SetUsername)
			// userAuth.PUT("/avatar", userController.SetAvatar)
			// userAuth.PUT("/gender", userController.SetGender)
			// userAuth.PUT("/country", userController.SetCountry)
			// userAuth.PUT("/language", userController.SetLanguage)
			// userAuth.PUT("/birthday", userController.SetBirthday)
		}
	}

	// 定义清理函数
	cleanup := func() {
		if err := utils.ClosePostgreSQL(); err != nil {
			logrus.Error("Failed to close db: ", err)
		}
		if err := utils.CloseRedis(); err != nil {
			logrus.Error("Failed to close redis: ", err)
		}
		if err := utils.CloseGeo(); err != nil {
			logrus.Error("Failed to close geo: ", err)
		}
		logrus.Info("Cleaning up resources")
	}

	return r, cleanup
}
