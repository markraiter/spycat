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
        }
    },
    "definitions": {
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
        "domain.Response": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string",
                    "example": "response message"
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
