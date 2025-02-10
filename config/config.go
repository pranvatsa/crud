package config

import (
	"os"
)

const (
	DefaultStorageMode = "json"
	DefaultMongoURI    = "mongodb://localhost:27017"
	DefaultMongoDBName = "mydatabase"
)

var (
	StorageMode string
	MongoURI    string
	MongoDBName string
)

func LoadConfig() {
	StorageMode = GetEnv("STORAGE_MODE", DefaultStorageMode)
	MongoURI = GetEnv("MONGO_URI", DefaultMongoURI)
	MongoDBName = GetEnv("MONGO_DB_NAME", DefaultMongoDBName)
}

func GetEnv(key, defaultValue string) string {
	value, exists := os.LookupEnv(key)
	if !exists {
		return defaultValue
	}
	return value
}
