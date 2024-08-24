package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

const (
	// DebugMode indicates service mode is debug.
	DebugMode = "debug"
	// TestMode indicates service mode is test.
	TestMode = "test"
	// ReleaseMode indicates service mode is release.
	ReleaseMode = "release"
)

type Config struct {
	ServiceName string
	Environment string // debug, test, release
	Version     string


	HTTPPort   string
	HTTPScheme string
	
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDatabase string

	SaleServiceHost string
	SaleGRPCPort    string

	StocktentServiceHost string
	StockGRPCPort    string

	ContentServiceHost string
	ContentGRPCPort    string

	DefaultOffset string
	DefaultLimit  string


	PostgresMaxConnections int32
}

// Load ...
func Load() Config {
	if err := godotenv.Load("/article_article_service.env"); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}


	config.HTTPPort = cast.ToString(getOrReturnDefaultValue("HTTP_PORT", ":8081"))
	config.HTTPScheme = cast.ToString(getOrReturnDefaultValue("HTTP_SCHEME", "http"))

	config.ServiceName = cast.ToString(getOrReturnDefaultValue("SERVICE_NAME", "market-api-gateway"))
	config.Environment = cast.ToString(getOrReturnDefaultValue("ENVIRONMENT", DebugMode))
	config.Version = cast.ToString(getOrReturnDefaultValue("VERSION", "1.0"))

	config.PostgresHost = cast.ToString(getOrReturnDefaultValue("POSTGRES_HOST", "localhost"))
	config.PostgresPort = cast.ToInt(getOrReturnDefaultValue("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(getOrReturnDefaultValue("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(getOrReturnDefaultValue("POSTGRES_PASSWORD", "postgres"))
	config.PostgresDatabase = cast.ToString(getOrReturnDefaultValue("POSTGRES_DATABASE", "saleio"))

	config.SaleServiceHost = cast.ToString(getOrReturnDefaultValue("SALE_SERVICE_HOST", "localhost"))
	config.SaleGRPCPort = cast.ToString(getOrReturnDefaultValue("SALE_GRPC_PORT", ":9104"))


	config.StocktentServiceHost = cast.ToString(getOrReturnDefaultValue("CONTENT_SERVICE_HOST", "localhost"))
	config.StockGRPCPort = cast.ToString(getOrReturnDefaultValue("CONTENT_GRPC_PORT", ":9103"))
	

	config.ContentServiceHost = cast.ToString(getOrReturnDefaultValue("CONTENT_SERVICE_HOST", "localhost"))
	config.ContentGRPCPort = cast.ToString(getOrReturnDefaultValue("CONTENT_GRPC_PORT", ":9102"))

	config.DefaultOffset = cast.ToString(getOrReturnDefaultValue("DEFAULT_OFFSET", "0"))
	config.DefaultLimit = cast.ToString(getOrReturnDefaultValue("DEFAULT_LIMIT", "10"))

	config.PostgresMaxConnections = cast.ToInt32(getOrReturnDefaultValue("POSTGRES_MAX_CONNECTIONS", 30))

	return config
}

func getOrReturnDefaultValue(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
