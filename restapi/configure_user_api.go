// This file is safe to edit. Once it exists it will not be overwritten

package restapi

import (
	"Swagger-Sqlite-DB/models"
	"crypto/tls"
	"gorm.io/gorm"
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

type User struct {
	gorm.Model
	Username string `gorm:"column:name"`
	Email    string `gorm:"column:email"`
}

func (User) TableName() string {
	return "db_users"
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

	_, err := InitDB()
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}
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

	var user User
	result := DB.First(&user, userID)
	if result.Error != nil {
		log.Printf("Error retrieving user from the database: %v", result.Error)
		return users.NewGetUserByIDNotFound()
	}

	// Map GORM model fields to models.User fields
	userModel := &models.User{
		ID:    int64(user.ID),
		Name:  &user.Username, // Use &user.Username to get a pointer to the string
		Email: &user.Email,
	}

	// Return a single-element slice of the user
	userList := []*models.User{userModel}
	return users.NewGetUsersOK().WithPayload(userList)
}

func getUsers(params users.GetUsersParams) middleware.Responder {
	var userList []*models.User

	// Retrieve all users from the database using GORM
	var user []User
	result := DB.Find(&user)
	if result.Error != nil {
		log.Printf("Error retrieving users from the database: %v", result.Error)
		return users.NewGetUsersInternalServerError()
	}

	// Map GORM model fields to models.User fields for each user
	for _, user := range user {
		userModel := &models.User{
			ID:    int64(user.ID), // Convert uint to int64
			Name:  &user.Username, // Use &user.Username to get a pointer to the string
			Email: &user.Email,
		}
		userList = append(userList, userModel)
	}

	// Return the list of users
	return users.NewGetUsersOK().WithPayload(userList)
}
func createUser(params users.CreateUserParams) middleware.Responder {
	// Extract the user data from params
	user := params.User
	newUser := dbUser{
		Name:  user.Name,
		Email: user.Email,
	}

	result := DB.Create(&newUser)
	if result.Error != nil {
		log.Printf("Error inserting user into the database: %v", result.Error)
		return users.NewCreateUserInternalServerError()
	}

	// Return a success response
	return users.NewCreateUserCreated()
}

func updateUser(params users.UpdateUserParams) middleware.Responder {
	// Extract the user data from params
	userID := params.ID
	user := params.User

	// Find the user with the given ID
	var existingUser User
	result := DB.First(&existingUser, userID)
	if result.Error != nil {
		log.Printf("Error when finding user: %v", result.Error)
		return users.NewUpdateUserBadRequest()
	}

	// Update the user fields
	if user.Name != nil {
		existingUser.Username = *user.Name
	}
	if user.Email != nil {
		existingUser.Email = *user.Email
	}

	// Save the updated user to the database
	result = DB.Save(&existingUser)
	if result.Error != nil {
		log.Printf("Error when updating user: %v", result.Error)
		return users.NewUpdateUserBadRequest()
	}

	log.Printf("Updated user with ID %d", userID)
	return users.NewUpdateUserOK()
}

func deleteUser(params users.DeleteUserParams) middleware.Responder {
	// Extract the user ID from params
	userID := params.ID

	// Find the user with the given ID
	var user User
	result := DB.First(&user, userID)
	if result.Error != nil {
		log.Printf("Error when finding user: %v", result.Error)
		return users.NewDeleteUserNotFound()
	}

	// Delete the user from the database
	result = DB.Delete(&user)
	if result.Error != nil {
		log.Printf("Error while deleting user from the database: %v", result.Error)
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
