// Package docs provides auto-generated swagger documentation
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "swagger": "2.0",
    "info": {
        "title": "Todo API",
        "description": "RESTful API for Todo management using Gin and GORM",
        "contact": {},
        "version": "1.0.0"
    },
    "host": "localhost:8080",
    "basePath": "/api/v1",
    "paths": {
        "/todos": {
            "get": {
                "description": "Get all todos from database",
                "produces": ["application/json"],
                "tags": ["Todos"],
                "summary": "Get all todos",
                "responses": {
                    "200": {
                        "description": "List of todos",
                        "schema": {
                            "$ref": "#/definitions/TodoListResponse"
                        }
                    }
                }
            },
            "post": {
                "description": "Create a new todo item",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["Todos"],
                "summary": "Create a new todo",
                "parameters": [
                    {
                        "description": "Todo to create",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/CreateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Todo created successfully",
                        "schema": {
                            "$ref": "#/definitions/SuccessResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid request body",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        },
        "/todos/{id}": {
            "get": {
                "description": "Get a todo by its ID",
                "produces": ["application/json"],
                "tags": ["Todos"],
                "summary": "Get todo by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Todo details",
                        "schema": {
                            "$ref": "#/definitions/SuccessResponse"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "put": {
                "description": "Update an existing todo",
                "consumes": ["application/json"],
                "produces": ["application/json"],
                "tags": ["Todos"],
                "summary": "Update todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Todo update data",
                        "name": "todo",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/UpdateTodoRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Todo updated successfully",
                        "schema": {
                            "$ref": "#/definitions/SuccessResponse"
                        }
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            },
            "delete": {
                "description": "Delete a todo by its ID",
                "produces": ["application/json"],
                "tags": ["Todos"],
                "summary": "Delete todo",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Todo ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Todo deleted successfully"
                    },
                    "404": {
                        "description": "Todo not found",
                        "schema": {
                            "$ref": "#/definitions/ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "CreateTodoRequest": {
            "type": "object",
            "required": ["title"],
            "properties": {
                "title": {
                    "type": "string",
                    "example": "Belajar Golang"
                },
                "description": {
                    "type": "string",
                    "example": "Belajar Gin dan GORM"
                }
            }
        },
        "UpdateTodoRequest": {
            "type": "object",
            "properties": {
                "title": {
                    "type": "string"
                },
                "description": {
                    "type": "string"
                },
                "completed": {
                    "type": "boolean"
                }
            }
        },
        "TodoResponse": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "title": {
                    "type": "string",
                    "example": "Belajar Golang"
                },
                "description": {
                    "type": "string",
                    "example": "Belajar Gin dan GORM"
                },
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "created_at": {
                    "type": "string"
                },
                "updated_at": {
                    "type": "string"
                }
            }
        },
        "TodoListResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "message": {
                    "type": "string"
                },
                "data": {
                    "type": "object",
                    "properties": {
                        "todos": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/TodoResponse"
                            }
                        },
                        "total": {
                            "type": "integer"
                        }
                    }
                }
            }
        },
        "SuccessResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean",
                    "example": true
                },
                "message": {
                    "type": "string"
                },
                "data": {
                    "$ref": "#/definitions/TodoResponse"
                }
            }
        },
        "ErrorResponse": {
            "type": "object",
            "properties": {
                "success": {
                    "type": "boolean",
                    "example": false
                },
                "message": {
                    "type": "string"
                },
                "error": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info
var SwaggerInfo = &swag.Spec{
	Version:          "1.0.0",
	Host:             "localhost:8080",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Todo API",
	Description:      "RESTful API for Todo management using Gin and GORM",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
