package config

import (
	"os"
)

var MongoURI = os.Getenv("MONGO_URI")
