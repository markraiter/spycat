// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Mark Raiter",
            "email": "raitermark@proton.me"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/auth/login": {
            "post": {
                "description": "Logs user in",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Login",
                "operationId": "login",
                "parameters": [
                    {
                        "description": "credentials",
                        "name": "input",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.LoginRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Token",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        },
        "/auth/logout": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Logs user out",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Logout",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        },
        "/auth/register": {
            "post": {
                "description": "Register user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Auth"
                ],
                "summary": "Register user",
                "parameters": [
                    {
                        "description": "User data",
                        "name": "Register_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.UserRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "User ID",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        },
        "/cats": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get all cats",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cat"
                ],
                "summary": "Get all cats",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Cat"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create cat",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cat"
                ],
                "summary": "Create cat",
                "parameters": [
                    {
                        "description": "Cat data",
                        "name": "Create_cat_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CatRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Cat ID",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        },
        "/cats/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get cat by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cat"
                ],
                "summary": "Get cat by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cat ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Cat"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Update cat by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cat"
                ],
                "summary": "Update cat by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cat ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Cat data",
                        "name": "Update_cat_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.CatRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Delete cat by ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Cat"
                ],
                "summary": "Delete cat by ID",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Cat ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        },
        "/missions": {
            "post": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Create mission",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Mission"
                ],
                "summary": "Create mission",
                "parameters": [
                    {
                        "description": "Mission data",
                        "name": "Create_mission_request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.MissionRequest"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Mission ID",
                        "schema": {
                            "type": "integer"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "403": {
                        "description": "Forbidden",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "406": {
                        "description": "Not Acceptable",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/domain.Response"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.Cat": {
            "type": "object",
            "required": [
                "breed",
                "name"
            ],
            "properties": {
                "breed": {
                    "type": "string",
                    "example": "Siamese"
                },
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string",
                    "example": "Tom"
                },
                "salary": {
                    "type": "integer",
                    "example": 1000
                },
                "years_of_experience": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "domain.CatRequest": {
            "type": "object",
            "required": [
                "breed",
                "name"
            ],
            "properties": {
                "breed": {
                    "type": "string",
                    "example": "Siamese"
                },
                "name": {
                    "type": "string",
                    "example": "Tom"
                },
                "salary": {
                    "type": "integer",
                    "example": 1000
                },
                "years_of_experience": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "domain.LoginRequest": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@example.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8,
                    "example": "Password12345!"
                }
            }
        },
        "domain.MissionRequest": {
            "type": "object",
            "required": [
                "cat_id",
                "completed",
                "targets"
            ],
            "properties": {
                "cat_id": {
                    "type": "integer",
                    "example": 1
                },
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "notes": {
                    "type": "string",
                    "example": "Lorem ipsum"
                },
                "targets": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/domain.Target"
                    }
                }
            }
        },
        "domain.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "response message"
                }
            }
        },
        "domain.Target": {
            "type": "object",
            "required": [
                "completed",
                "country",
                "mission_id",
                "name"
            ],
            "properties": {
                "completed": {
                    "type": "boolean",
                    "example": false
                },
                "country": {
                    "type": "string",
                    "example": "USA"
                },
                "id": {
                    "type": "integer"
                },
                "mission_id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "John Doe"
                },
                "notes": {
                    "type": "string",
                    "example": "Lorem ipsum"
                }
            }
        },
        "domain.UserRequest": {
            "type": "object",
            "required": [
                "email",
                "password",
                "username"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "email@example.com"
                },
                "password": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 8,
                    "example": "Password12345!"
                },
                "username": {
                    "type": "string",
                    "maxLength": 50,
                    "minLength": 3,
                    "example": "username"
                }
            }
        }
    },
    "securityDefinitions": {
        "ApiKeyAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8888",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "SpyCat API",
	Description:      "Docs for SpyCat API",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
