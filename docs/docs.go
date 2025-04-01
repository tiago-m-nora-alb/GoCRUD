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
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/image": {
            "post": {
                "description": "Upload an image to MinIO bucket \"imagens\"",
                "consumes": [
                    "multipart/form-data"
                ],
                "tags": [
                    "image"
                ],
                "summary": "Upload an image to MinIO",
                "parameters": [
                    {
                        "type": "file",
                        "description": "File to upload",
                        "name": "file",
                        "in": "formData",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Success message with file name",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/image/{fileName}": {
            "get": {
                "description": "Get an image from MinIO bucket \"imagens\"",
                "tags": [
                    "image"
                ],
                "summary": "Get an image from MinIO",
                "parameters": [
                    {
                        "type": "string",
                        "description": "File name to retrieve",
                        "name": "fileName",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Image file",
                        "schema": {
                            "type": "file"
                        }
                    },
                    "400": {
                        "description": "Error message",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "description": "Find all products",
                "tags": [
                    "Products"
                ],
                "summary": "Find all products",
                "responses": {}
            },
            "post": {
                "description": "Create a new product with name, description, price, and stock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Create a new product",
                "parameters": [
                    {
                        "description": "Product data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ProductRequest"
                        }
                    }
                ],
                "responses": {}
            }
        },
        "/product/{id}": {
            "get": {
                "description": "Find a product by ID",
                "tags": [
                    "Products"
                ],
                "summary": "Find a product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            },
            "put": {
                "description": "Update a product by ID with name, description, price, and stock",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Products"
                ],
                "summary": "Update a product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    },
                    {
                        "description": "Product data",
                        "name": "product",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/schemas.ProductRequest"
                        }
                    }
                ],
                "responses": {}
            },
            "delete": {
                "description": "Delete a product by ID",
                "tags": [
                    "Products"
                ],
                "summary": "Delete a product by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {}
            }
        },
        "/status/health": {
            "get": {
                "description": "Responds with a hello message from the Product Controller",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Status"
                ],
                "summary": "Greets from Product Controller",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "object",
                            "additionalProperties": {
                                "type": "string"
                            }
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "schemas.ProductRequest": {
            "type": "object",
            "required": [
                "description",
                "name",
                "price",
                "stock"
            ],
            "properties": {
                "description": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "number"
                },
                "stock": {
                    "type": "integer"
                }
            }
        }
    },
    "securityDefinitions": {
        "bearerToken": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8000",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Swagger GO API",
	Description:      "This is a example GO API.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
