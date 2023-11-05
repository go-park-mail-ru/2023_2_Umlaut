// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {},
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/api/v1/dialogs": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "dialog"
                ],
                "summary": "get user dialogs",
                "operationId": "dialog",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-array_model_Dialog"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        },
        "/api/v1/feed": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feed"
                ],
                "summary": "get user for feed",
                "operationId": "feed",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-model_User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        },
        "/api/v1/feed/users": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "feed"
                ],
                "summary": "get users for feed",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-array_model_User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        },
        "/api/v1/like": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "like"
                ],
                "summary": "create user like",
                "operationId": "like",
                "parameters": [
                    {
                        "description": "Like data to update",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.Like"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        },
        "/api/v1/tags": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "tag"
                ],
                "summary": "get all tags",
                "operationId": "tag",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-array_string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        },
        "/api/v1/user": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "get user information",
                "operationId": "user",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-model_User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            },
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "update user",
                "operationId": "user",
                "parameters": [
                    {
                        "description": "User data to update",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/model.User"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-model_User"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/photo": {
            "post": {
                "consumes": [
                    "multipart/form-data"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "update user photo",
                "parameters": [
                    {
                        "type": "file",
                        "description": "file",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            },
            "delete": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "delete user photo",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        },
        "/api/v1/user/{id}/photo": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "tags": [
                    "user"
                ],
                "summary": "get user photo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "User ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        },
        "/auth/login": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "log in to account",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "Sign-in input parameters",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.signInInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "get": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "log out of account",
                "operationId": "logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        },
        "/auth/sign-up": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "auth"
                ],
                "summary": "sign up account",
                "operationId": "create-account",
                "parameters": [
                    {
                        "description": "Sign-up input user",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/handler.signUpInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-handler_idResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handler.ClientResponseDto-string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "handler.ClientResponseDto-array_model_Dialog": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "payload": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.Dialog"
                    }
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "handler.ClientResponseDto-array_model_User": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "payload": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/model.User"
                    }
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "handler.ClientResponseDto-array_string": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "payload": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "handler.ClientResponseDto-handler_idResponse": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "payload": {
                    "$ref": "#/definitions/handler.idResponse"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "handler.ClientResponseDto-model_User": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "payload": {
                    "$ref": "#/definitions/model.User"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "handler.ClientResponseDto-string": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
                },
                "payload": {
                    "type": "string"
                },
                "status": {
                    "type": "integer"
                }
            }
        },
        "handler.idResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                }
            }
        },
        "handler.signInInput": {
            "type": "object",
            "required": [
                "mail",
                "password"
            ],
            "properties": {
                "mail": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "handler.signUpInput": {
            "type": "object",
            "required": [
                "mail",
                "name",
                "password"
            ],
            "properties": {
                "mail": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "model.Dialog": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer"
                },
                "user1_id": {
                    "type": "integer"
                },
                "user2_id": {
                    "type": "integer"
                }
            }
        },
        "model.Like": {
            "type": "object",
            "properties": {
                "liked_to_user_id": {
                    "type": "integer"
                }
            }
        },
        "model.User": {
            "type": "object",
            "required": [
                "mail",
                "name"
            ],
            "properties": {
                "age": {
                    "type": "integer"
                },
                "birthday": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "education": {
                    "type": "string"
                },
                "hobbies": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "image_path": {
                    "type": "string"
                },
                "looking": {
                    "type": "string"
                },
                "mail": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "online": {
                    "type": "boolean"
                },
                "prefer_gender": {
                    "type": "integer"
                },
                "tags": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                },
                "user_gender": {
                    "type": "integer"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "umlaut-bmstu.me:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Umlaut API",
	Description:      "API Server for Umlaut Application",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
