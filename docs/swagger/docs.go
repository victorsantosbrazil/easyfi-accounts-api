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
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/v1/institutions": {
            "get": {
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "institutions"
                ],
                "summary": "list institutions",
                "parameters": [
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": " ",
                        "name": "size",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/usecase.ListInstitutionsUseCaseResponse"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/errors.ApiError"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "errors.ApiError": {
            "type": "object",
            "properties": {
                "detail": {
                    "type": "string",
                    "example": "Could not process request due ..."
                },
                "title": {
                    "type": "string",
                    "example": "Example error"
                },
                "type": {
                    "type": "string",
                    "example": "api_error"
                }
            }
        },
        "pagination.Pagination": {
            "type": "object",
            "properties": {
                "page": {
                    "type": "integer",
                    "example": 1
                },
                "size": {
                    "type": "integer",
                    "example": 10
                },
                "totalElements": {
                    "type": "integer",
                    "example": 50
                },
                "totalPages": {
                    "type": "integer",
                    "example": 5
                }
            }
        },
        "usecase.ListInstitutionsUseCaseResponse": {
            "type": "object",
            "properties": {
                "items": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/usecase.ListInstitutionsUseCaseResponseItem"
                    }
                },
                "pagination": {
                    "$ref": "#/definitions/pagination.Pagination"
                }
            }
        },
        "usecase.ListInstitutionsUseCaseResponseItem": {
            "type": "object",
            "properties": {
                "id": {
                    "type": "integer",
                    "example": 1
                },
                "name": {
                    "type": "string",
                    "example": "Brazil Bank"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "0.1.0",
	Host:             "",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Accounts API",
	Description:      "API for managing record of bank accounts and credit cards",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
