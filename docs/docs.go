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
        "/persons": {
            "get": {
                "description": "Get a list of all persons",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Retrieve all persons",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Person"
                            }
                        }
                    }
                }
            }
        },
        "/persons/search": {
            "get": {
                "description": "Search for persons using a text query",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "persons"
                ],
                "summary": "Search for persons",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Text to search for",
                        "name": "text",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/domain.Person"
                            }
                        }
                    },
                    "400": {
                        "description": "Bad Request",
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
        "domain.ContactInfo": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "phone": {
                    "type": "string"
                }
            }
        },
        "domain.Person": {
            "type": "object",
            "properties": {
                "about": {
                    "type": "string"
                },
                "contacts": {
                    "$ref": "#/definitions/domain.ContactInfo"
                },
                "department": {
                    "type": "string"
                },
                "division": {
                    "type": "string"
                },
                "head": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "jobtitle": {
                    "type": "string"
                },
                "middle_name_rus": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "status": {
                    "type": "string"
                },
                "surname": {
                    "type": "string"
                },
                "team": {
                    "type": "string"
                },
                "working_hour": {
                    "type": "string"
                },
                "workplace": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "",
	Description:      "",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
