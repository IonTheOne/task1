// Package swagger Code generated by swaggo/swag. DO NOT EDIT
package swagger

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
        "/news": {
            "get": {
                "description": "Fetch all news items from the database",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "summary": "Get all news",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/httpentity.NewsItem"
                            }
                        }
                    }
                }
            }
        },
        "/news/{newsItemId}": {
            "get": {
                "description": "Fetch a single news item from the database by ID",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "news"
                ],
                "summary": "Get news by ID",
                "parameters": [
                    {
                        "type": "string",
                        "description": "News Item ID",
                        "name": "newsItemId",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/httpentity.NewsItem"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "httpentity.CustomTime": {
            "type": "object",
            "properties": {
                "time.Time": {
                    "type": "string"
                }
            }
        },
        "httpentity.NewsItem": {
            "type": "object",
            "properties": {
                "articleURL": {
                    "type": "string"
                },
                "isPublished": {
                    "type": "boolean"
                },
                "lastUpdateDate": {
                    "$ref": "#/definitions/httpentity.CustomTime"
                },
                "newsArticleID": {
                    "type": "string"
                },
                "optaMatchId": {
                    "type": "string"
                },
                "publishDate": {
                    "$ref": "#/definitions/httpentity.CustomTime"
                },
                "taxonomies": {
                    "type": "string"
                },
                "teaserText": {
                    "type": "string"
                },
                "thumbnailImageURL": {
                    "type": "string"
                },
                "title": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:8100",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "Task1 API",
	Description:      "API for task1 service",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}