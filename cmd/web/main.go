package main

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Ashis-007/hms/internal/model"
	loggerPkg "github.com/Ashis-007/hms/pkg/logger"
	"github.com/Ashis-007/hms/pkg/response"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var logger = loggerPkg.GetLogger()

func initConfig() (*Config, error) {
	envType := os.Getenv("GO_ENV")

	configFileName := "config"
	if envType == "development" {
		configFileName = "config.local"
	}

	// Initialize Viper
	viper.SetConfigName(configFileName) // The name of your configuration file (without extension)
	viper.AddConfigPath(".")            // Search in the current directory
	viper.SetConfigType("yaml")         // Set the configuration file type

	// Read the configuration file
	if err := viper.ReadInConfig(); err != nil {
		logger.Error("error while reading config file", zap.Error(err))
		return nil, err
	}

	config := &Config{}

	// Unmarshal the configuration into a struct
	if err := viper.Unmarshal(config); err != nil {
		logger.Error("error while unmarshalling config file", zap.Error(err))
		return nil, err
	}

	return config, nil
}

func connectDB(config *Config) (*sqlx.DB, error) {
	dsn := config.DBDSN
	if dsn == "" {
		dsn = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s", config.DBHost, config.DBPort, config.DBUser, config.DBPassword, config.DBName)
	}
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		logger.Error("error while connecting to db", zap.Error(err))
		return nil, err
	}

	return db, nil
}

func logRoutesMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now() // Start timer
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Fill the params
		param := gin.LogFormatterParams{}

		param.TimeStamp = time.Now() // Stop timer
		param.Latency = param.TimeStamp.Sub(start)
		if param.Latency > time.Minute {
			param.Latency = param.Latency.Truncate(time.Second)
		}

		param.ClientIP = c.ClientIP()
		param.Method = c.Request.Method
		param.StatusCode = c.Writer.Status()
		param.ErrorMessage = c.Errors.ByType(gin.ErrorTypePrivate).String()
		param.BodySize = c.Writer.Size()
		if raw != "" {
			path = path + "?" + raw
		}
		param.Path = path

		// contextUser, _ := pkg.GetContextUser(c)
		// if contextUser.UserID.Valid {
		// 	log.Info("", zap.Any("method", param.Method), zap.Any("path", param.Path), zap.Any("status", param.StatusCode), zap.Any("latency", param.Latency.Round(time.Millisecond).String()), zap.Any("uid", contextUser.UserID))
		// } else {
		// 	log.Info("", zap.Any("method", param.Method), zap.Any("path", param.Path), zap.Any("status", param.StatusCode), zap.Any("latency", param.Latency.Round(time.Millisecond).String()))
		// }
	}
}

func main() {
	startTime := time.Now()

	fmt.Println(logoAscii)

	// load config
	config, err := initConfig()
	if err != nil {
		panic(err)
	}

	db, err := connectDB(config)
	if err != nil {
		panic(err)
	}
	logger.Info("[DATABASE] successfully connected to database")

	model.SetPostgresDB(db)

	fmt.Println(db)

	// initialise gin app
	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.RemoveExtraSlash = true
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8080", "*"}, // TODO: restrict origins
		AllowMethods:     []string{http.MethodGet, http.MethodPatch, http.MethodPost, http.MethodPut, http.MethodDelete, http.MethodHead, http.MethodOptions},
		AllowHeaders:     []string{"Content-Type", "X-XSRF-TOKEN", "Accept", "Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Use(logRoutesMiddleware())
	r.Use(gin.Recovery())

	r.GET("/healthcheck", func(c *gin.Context) {
		response.OkResponse(c, &response.Response{
			Msg: "Backend server up and running",
		})
	})

	r.GET("/uptime", func(c *gin.Context) {
		response.OkResponse(c, &response.Response{
			Msg: "Backend server up and running",
			Data: map[string]string{
				"uptime": time.Duration(time.Since(startTime)).Round(time.Second).String(),
			},
		})
	})

	serverUrl := fmt.Sprintf(":%s", config.Port)

	logger.Info(fmt.Sprintf("🚀 [SERVER INIT] successfully started server at port %s", config.Port))

	err = r.Run(serverUrl)
	if err != nil {
		logger.Error("error while starting server", zap.Error(err))
		panic(err)
	}
}
