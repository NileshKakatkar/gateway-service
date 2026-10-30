package gateway

import "os"

type Config struct{ Port, UserURL, ProductURL, OrderURL, RedisAddr string }

func LoadConfig() Config {
	get := func(k, d string) string {
		if v := os.Getenv(k); v != "" {
			return v
		}
		return d
	}
	return Config{Port: get("GATEWAY_PORT", "8080"), UserURL: get("USER_SERVICE_URL", "http://localhost:8081"), ProductURL: get("PRODUCT_SERVICE_URL", "http://localhost:8082"), OrderURL: get("ORDER_SERVICE_URL", "http://localhost:8083"), RedisAddr: get("REDIS_ADDR", "localhost:6379")}
}
