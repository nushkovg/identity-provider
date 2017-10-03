//go:generate goagen bootstrap -d github.com/JormungandrK/identity-provider/design

package main

import (
	"net/http"
	"os"

	"github.com/JormungandrK/identity-provider/app"
	"github.com/JormungandrK/identity-provider/db"
	"github.com/JormungandrK/microservice-tools/gateway"
	"github.com/goadesign/goa"
	"github.com/goadesign/goa/middleware"
)

func main() {
	// Gateway self-registration
	unregisterService := registerMicroservice()
	defer unregisterService() // defer the unregister for after main exits

	// Create service
	service := goa.New("identity provider")

	// Mount middleware
	service.Use(middleware.RequestID())
	service.Use(middleware.LogRequest(true))
	service.Use(middleware.ErrorHandler(service, true))
	service.Use(middleware.Recover())

	// Load MongoDB ENV variables
	host, username, password, database := loadMongnoSettings()
	// Create new session to MongoDB
	session := db.NewSession(host, username, password, database)

	// At the end close session
	defer session.Close()

	// Create metadata collection and indexes
	indexesMetadata := []string{"name"}
	metadataCollection := db.PrepareDB(session, database, "services", indexesMetadata)

	// Create type collection and indexes
	indexesType := []string{"id"}
	typeCollection := db.PrepareDB(session, database, "sessions", indexesType)

	// Create type collection and indexes
	indexesDataSource := []string{"name"}
	dataSourceCollection := db.PrepareDB(session, database, "users", indexesDataSource)

	// Mount "idp" controller
	c := NewIdpController(service, &db.MongoCollections{
		Services: metadataCollection,
		Sessions: typeCollection,
		Users:    dataSourceCollection,
	})

	app.MountIdpController(service, c)
	// Mount "swagger" controller
	c2 := NewSwaggerController(service)
	app.MountSwaggerController(service, c2)

	// Start service
	if err := service.ListenAndServe(":8080"); err != nil {
		service.LogError("startup", "err", err)
	}

}

func loadMongnoSettings() (string, string, string, string) {
	host := os.Getenv("MONGO_URL")
	username := os.Getenv("MS_USERNAME")
	password := os.Getenv("MS_PASSWORD")
	database := os.Getenv("MS_DBNAME")

	if host == "" {
		host = "127.0.0.1:27017"
	}
	if username == "" {
		username = "restapi"
	}
	if password == "" {
		password = "restapi"
	}
	if database == "" {
		database = "identity-provider"
	}

	return host, username, password, database
}

func loadGatewaySettings() (string, string) {
	gatewayURL := os.Getenv("API_GATEWAY_URL")
	serviceConfigFile := os.Getenv("SERVICE_CONFIG_FILE")

	if gatewayURL == "" {
		gatewayURL = "http://localhost:8001"
	}
	if serviceConfigFile == "" {
		serviceConfigFile = "config.json"
	}

	return gatewayURL, serviceConfigFile
}

func registerMicroservice() func() {
	gatewayURL, configFile := loadGatewaySettings()
	registration, err := gateway.NewKongGatewayFromConfigFile(gatewayURL, &http.Client{}, configFile)
	if err != nil {
		panic(err)
	}
	err = registration.SelfRegister()
	if err != nil {
		panic(err)
	}

	return func() {
		registration.Unregister()
	}
}