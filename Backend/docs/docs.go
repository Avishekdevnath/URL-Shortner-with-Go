// Code generated by swaggo/swag. DO NOT EDIT.
// S:\SDE\Hard Core\Learn\Golang\Projects\URL-Shortner-with-Go\Backend\docs\docs.go
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
        "/api/shorten": {
            "post": {
                "description": "Takes a long URL and returns a shortened version",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "Shorten a URL",
                "parameters": [
                    {
                        "description": "URL to shorten",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/models.ShortURLRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Shortened URL",
                        "schema": {
                            "$ref": "#/definitions/models.ShortURLResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid input",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/urls": {
            "get": {
                "description": "Returns a list of all shortened URLs with their original URLs",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "List all shortened URLs",
                "responses": {
                    "200": {
                        "description": "List of shortened URLs",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/models.URLListResponse"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal server error",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        },
        "/{code}": {
            "get": {
                "description": "Redirects to the original URL based on the short code",
                "produces": [
                    "text/plain"
                ],
                "tags": [
                    "URL"
                ],
                "summary": "Redirect to original URL",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Short URL code",
                        "name": "code",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "301": {
                        "description": "Redirects to the original URL"
                    },
                    "404": {
                        "description": "URL not found",
                        "schema": {
                            "$ref": "#/definitions/models.ErrorResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "models.ErrorResponse": {
            "type": "object",
            "properties": {
                "error": {
                    "type": "string"
                }
            }
        },
        "models.ShortURLRequest": {
            "type": "object",
            "required": [
                "url"
            ],
            "properties": {
                "url": {
                    "type": "string"
                }
            }
        },
        "models.ShortURLResponse": {
            "type": "object",
            "properties": {
                "original_url": {
                    "description": "Added OriginalURL field",
                    "type": "string"
                },
                "short_url": {
                    "type": "string"
                }
            }
        },
        "models.URLListResponse": {
            "type": "object",
            "properties": {
                "original_url": {
                    "type": "string"
                },
                "short_code": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8080",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "URL Shortener API",
	Description:      "A simple URL shortening service built with Go and Gin.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
