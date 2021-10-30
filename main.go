package main

import (
	"AltaStore/api"
	"AltaStore/api/middleware"

	// Controller
	adminController "AltaStore/api/v1/admin"
	adminAuthController "AltaStore/api/v1/adminauth"
	userController "AltaStore/api/v1/user"
	userAuthController "AltaStore/api/v1/userauth"

	// Service
	adminService "AltaStore/business/admin"
	adminAuthService "AltaStore/business/adminauth"
	loggerService "AltaStore/business/logger"
	userService "AltaStore/business/user"
	userAuthService "AltaStore/business/userauth"

	"AltaStore/config"

	// Repository
	adminRepository "AltaStore/modules/admin"
	loggerRepo "AltaStore/modules/logger"
	"AltaStore/modules/migration"
	userRepository "AltaStore/modules/user"

	"context"
	"time"

	"fmt"

	log "github.com/sirupsen/logrus"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"github.com/go-redis/redis/v7"
	"github.com/labstack/echo/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func newDatabaseConnection(cfg *config.ConfigApp) *gorm.DB {
	stringConnection := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s",
		cfg.DbHost, cfg.DbPort, cfg.DbUsername, cfg.DbPassword, cfg.DbName,
	)
	db, err := gorm.Open(postgres.Open(stringConnection), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	migration.TableMigration(db)

	return db
}

func newMongoDBConnection(cfg *config.ConfigApp) *mongo.Database {
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf("mongodb://%s:%s@%s:%d", cfg.MongoUsername, cfg.MongoPassword, cfg.MongoHost, cfg.MongoPort),
	)

	client, err := mongo.NewClient(clientOptions)
	if err != nil {
		panic(err)
	}

	err = client.Connect(context.Background())
	if err != nil {
		panic(err)
	}

	return client.Database(cfg.MongoDbName)
}

func newRedisConnection(cfg *config.ConfigApp) *redis.Client {
	// stringConnection := fmt.Sprintf(
	// 	"%s:%d",
	// 	cfg.RedisHost, cfg.RedisPort,
	// )
	// client := redis.NewClient(&redis.Options{
	// 	Addr:     stringConnection, // redis port
	// 	Password: "",               // no password set
	// 	DB:       0,                // use default DB
	// })
	// _, err := client.Ping().Result()
	// if err != nil {
	// 	panic(err)
	// }
	// return client
	return nil
}

func main() {
	// retrieves application configuration and returns common values when there is a problem
	config := config.GetConfig()

	// Open mongodb logger
	mongoConnection := newMongoDBConnection(config)

	// Register repository
	logrRepo := loggerRepo.NewRepository(mongoConnection)

	// Register service
	logeService := loggerService.NewService(logrRepo)

	// Register logs
	log.SetFormatter(&log.JSONFormatter{})
	log.SetOutput(logeService)

	// open database server base session
	dbConnection := newDatabaseConnection(config)

	// open redis connection
	//redisConnection := newRedisConnection(config)
	_ = newRedisConnection(config)

	//initiate user repository
	user := userRepository.NewDBRepository(dbConnection)

	//initiate user service
	userService := userService.NewService(user)

	//initiate user controller
	userController := userController.NewController(userService)

	// Initiate Respository Category
	//_ = authRepository.NewRepository(redisConnection)

	//initiate admin repository
	admin := adminRepository.NewDBRepository(dbConnection)

	//initiate admin service
	adminService := adminService.NewService(admin)

	//initiate admin controller
	adminController := adminController.NewController(adminService)

	//initiate auth service
	userAuthService := userAuthService.NewService(userService)

	//initiate auth controller
	userAuthController := userAuthController.NewController(userAuthService)

	//initiate auth service
	adminAuthService := adminAuthService.NewService(adminService)

	//initiate auth controller
	adminAuthController := adminAuthController.NewController(adminAuthService)

	// create echo http
	e := echo.New()

	// Register API Path and Controller
	api.RegisterPath(e,
		userController,
		adminController,
		userAuthController,
		adminAuthController,
	)

	lock := make(chan error)

	go func(lock chan error) {
		address := fmt.Sprintf(":%d", config.AppPort)
		lock <- e.Start(address)
	}(lock)

	time.Sleep(1 * time.Millisecond)
	middleware.MakeLogEntry(nil).Info(fmt.Sprintf("Application Start In Port => ::%d", config.AppPort))

	err := <-lock
	if err != nil {
		middleware.MakeLogEntry(nil).Panic("Shutdown Echo Service")
	}
}
