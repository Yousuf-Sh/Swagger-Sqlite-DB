// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"crypto/tls"
	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"log"
	"net/http"
	"os"

	"Swagger-Sqlite-DB/restapi/operations"
	"Swagger-Sqlite-DB/restapi/operations/users"
	_ "github.com/mattn/go-sqlite3"
)

//go:generate swagger generate server --target ../../Swagger-Sqlite-DB --name UserAPI --spec ../swagger.yml --principal interface{}
var logger *log.Logger

func init() {
	// Open or create the log file for writing
	logFile, err := os.OpenFile("Delete.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}

	// Create a custom logger that writes to the log file
	logger = log.New(logFile, "", log.Ldate|log.Ltime|log.Lshortfile)
}
func configureFlags(api *operations.UserAPIAPI) {
	// api.CommandLineOptionsGroups = []swag.CommandLineOptionsGroup{ ... }
}

func configureAPI(api *operations.UserAPIAPI) http.Handler {
	// configure the api here
	api.ServeError = errors.ServeError

	// Set your custom logger if needed. Default one is log.Printf
	// Expected interface func(string, ...interface{})
	//
	// Example:
	// api.Logger = log.Printf

	api.UseSwaggerUI()
	// To continue using redoc as your UI, uncomment the following line
	// api.UseRedoc()

	api.JSONConsumer = runtime.JSONConsumer()

	api.JSONProducer = runtime.JSONProducer()

	InitDB("../userDB.db")
	api.UsersDeleteUserHandler = users.DeleteUserHandlerFunc(deleteUser)

	if api.UsersCreateUserHandler == nil {
		api.UsersCreateUserHandler = users.CreateUserHandlerFunc(func(params users.CreateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation users.CreateUser has not yet been implemented")
		})
	}
	if api.UsersDeleteUserHandler == nil {
		api.UsersDeleteUserHandler = users.DeleteUserHandlerFunc(func(params users.DeleteUserParams) middleware.Responder {
			return middleware.NotImplemented("operation users.DeleteUser has not yet been implemented")
		})
	}
	if api.UsersGetUserByIDHandler == nil {
		api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(func(params users.GetUserByIDParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUserByID has not yet been implemented")
		})
	}
	if api.UsersGetUsersHandler == nil {
		api.UsersGetUsersHandler = users.GetUsersHandlerFunc(func(params users.GetUsersParams) middleware.Responder {
			return middleware.NotImplemented("operation users.GetUsers has not yet been implemented")
		})
	}
	if api.UsersUpdateUserHandler == nil {
		api.UsersUpdateUserHandler = users.UpdateUserHandlerFunc(func(params users.UpdateUserParams) middleware.Responder {
			return middleware.NotImplemented("operation users.UpdateUser has not yet been implemented")
		})
	}

	api.PreServerShutdown = func() {}

	api.ServerShutdown = func() {}

	return setupGlobalMiddleware(api.Serve(setupMiddlewares))
}

func deleteUser(params users.DeleteUserParams) middleware.Responder {
	// Extract the user ID from params
	userID := params.ID

	// Implement logic to delete the user from the database by ID
	// You can use the global DB variable to interact with the database

	// Example: Delete the user from the "users" table by ID
	_, err := DB.Exec("DELETE FROM usersF WHERE id = ?", userID)
	if err != nil {
		logger.Printf("Error while deleting user from the database: %v", err) // Handle the error and return an appropriate response
		return users.NewDeleteUserNotFound()
	}

	// Return a success response
	return users.NewDeleteUserOK()
}

// The TLS configuration before HTTPS server starts.
func configureTLS(tlsConfig *tls.Config) {
	// Make all necessary changes to the TLS configuration here.
}

// As soon as server is initialized but not run yet, this function will be called.
// If you need to modify a config, store server instance to stop it individually later, this is the place.
// This function can be called multiple times, depending on the number of serving schemes.
// scheme value will be set accordingly: "http", "https" or "unix".
func configureServer(s *http.Server, scheme, addr string) {
}

// The middleware configuration is for the handler executors. These do not apply to the swagger.json document.
// The middleware executes after routing but before authentication, binding and validation.
func setupMiddlewares(handler http.Handler) http.Handler {
	return handler
}

// The middleware configuration happens before anything, this middleware also applies to serving the swagger.json document.
// So this is a good place to plug in a panic handling middleware, logging and metrics.
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	return handler
}
