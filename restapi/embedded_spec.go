// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for managing users",
    "title": "User API",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/users": {
      "get": {
        "tags": [
          "Users"
        ],
        "summary": "Get all users",
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "Users"
        ],
        "summary": "Create a new user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "User object",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad request"
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "tags": [
          "Users"
        ],
        "summary": "Get a user by ID",
        "operationId": "getUserByID",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "Not found"
          }
        }
      },
      "put": {
        "tags": [
          "Users"
        ],
        "summary": "Update a user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "User object",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad request"
          },
          "404": {
            "description": "Not found"
          }
        }
      },
      "delete": {
        "tags": [
          "Users"
        ],
        "summary": "Delete a user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "No content"
          },
          "404": {
            "description": "Not found"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "required": [
        "name",
        "email"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http",
    "https"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for managing users",
    "title": "User API",
    "version": "1.0.0"
  },
  "basePath": "/v1",
  "paths": {
    "/users": {
      "get": {
        "tags": [
          "Users"
        ],
        "summary": "Get all users",
        "operationId": "getUsers",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "type": "array",
              "items": {
                "$ref": "#/definitions/User"
              }
            }
          }
        }
      },
      "post": {
        "tags": [
          "Users"
        ],
        "summary": "Create a new user",
        "operationId": "createUser",
        "parameters": [
          {
            "description": "User object",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "201": {
            "description": "Created"
          },
          "400": {
            "description": "Bad request"
          }
        }
      }
    },
    "/users/{id}": {
      "get": {
        "tags": [
          "Users"
        ],
        "summary": "Get a user by ID",
        "operationId": "getUserByID",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/User"
            }
          },
          "404": {
            "description": "Not found"
          }
        }
      },
      "put": {
        "tags": [
          "Users"
        ],
        "summary": "Update a user",
        "operationId": "updateUser",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          },
          {
            "description": "User object",
            "name": "user",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/User"
            }
          }
        ],
        "responses": {
          "200": {
            "description": "OK"
          },
          "400": {
            "description": "Bad request"
          },
          "404": {
            "description": "Not found"
          }
        }
      },
      "delete": {
        "tags": [
          "Users"
        ],
        "summary": "Delete a user",
        "operationId": "deleteUser",
        "parameters": [
          {
            "type": "string",
            "description": "User ID",
            "name": "id",
            "in": "path",
            "required": true
          }
        ],
        "responses": {
          "204": {
            "description": "No content"
          },
          "404": {
            "description": "Not found"
          }
        }
      }
    }
  },
  "definitions": {
    "User": {
      "type": "object",
      "required": [
        "name",
        "email"
      ],
      "properties": {
        "email": {
          "type": "string"
        },
        "id": {
          "type": "integer",
          "format": "int64"
        },
        "name": {
          "type": "string"
        }
      }
    }
  }
}`))
}
