// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"Swagger-Sqlite-DB/models"
	"crypto/tls"
	"log"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"

	"Swagger-Sqlite-DB/restapi/operations"
	"Swagger-Sqlite-DB/restapi/operations/users"

	_ "github.com/mattn/go-sqlite3"
)

//go:generate swagger generate server --target ../../Swagger-Sqlite-DB --name UserAPI --spec ../swagger.yml --principal interface{}

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

	InitDB()
	api.UsersCreateUserHandler = users.CreateUserHandlerFunc(createUser)
	api.UsersUpdateUserHandler = users.UpdateUserHandlerFunc(updateUser)
	api.UsersDeleteUserHandler = users.DeleteUserHandlerFunc(deleteUser)
	api.UsersGetUsersHandler = users.GetUsersHandlerFunc(getUsers)
	api.UsersGetUserByIDHandler = users.GetUserByIDHandlerFunc(getUsersByID)

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

func getUsersByID(params users.GetUserByIDParams) middleware.Responder {
	userID := params.ID

	row := DB.QueryRow("SELECT id, username, email FROM users where id = ?", userID)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Email)
	if err != nil {
		return users.NewGetUserByIDNotFound()
	}
	userList := []*models.User{&user}
	return users.NewGetUsersOK().WithPayload(userList)
}
func getUsers(params users.GetUsersParams) middleware.Responder {

	rows, err := DB.Query("SELECT id, username, email FROM users")
	if err != nil {
		// Handle the error and return an appropriate response
		log.Printf("Error creating table: %v", err)
		return users.NewGetUsersInternalServerError()

	}
	defer rows.Close()

	var userList []*models.User

	for rows.Next() {
		var user models.User // Use the correct package here
		err := rows.Scan(&user.ID, &user.Name, &user.Email)
		if err != nil {
			// Handle the error and return an appropriate response
			return users.NewGetUsersBadRequest()
		}
		userList = append(userList, &user)
	}

	// Return the list of users
	return users.NewGetUsersOK().WithPayload(userList)
}
func createUser(params users.CreateUserParams) middleware.Responder {
	// Extract the user data from params
	user := params.User

	_, errInsert := DB.Exec("INSERT INTO users (username,email) VALUES (?, ?)", user.Name, user.Email)
	if errInsert != nil {
		log.Printf("Error inserting user into the database: %v", errInsert) // Handle the error and return an appropriate response
		return users.NewCreateUserInternalServerError()
	}
	// Return a success response
	return users.NewCreateUserCreated()
}
func updateUser(params users.UpdateUserParams) middleware.Responder {
	// Extract the user data from params
	userID := params.ID
	user := params.User

	_, err := DB.Exec("UPDATE users SET username = ?, email = ? WHERE id = ?", user.Name, user.Email, userID)
	if err != nil {
		log.Printf("Error when updating user :%v", err)
		return users.NewUpdateUserBadRequest()
	} else {
		log.Printf("Updated Correctly")
	}
	return users.NewUpdateUserOK()
}
func deleteUser(params users.DeleteUserParams) middleware.Responder {
	// Extract the user ID from params
	userID := params.ID

	_, err := DB.Exec("DELETE FROM users WHERE id = ?", userID)
	if err != nil {
		log.Printf("Error while deleting user from the database: %v", err)
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
