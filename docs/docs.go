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
        "/staff": {
            "post": {
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "staff"
                ],
                "summary": "createStaff",
                "parameters": [
                    {
                        "description": "body",
                        "name": "RequestBody",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/domain.StaffRequest"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/domain.StaffResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "domain.StaffRequest": {
            "type": "object",
            "properties": {
                "cityName": {
                    "type": "string",
                    "example": "Almaty"
                },
                "firstname": {
                    "type": "string",
                    "example": "Azamat"
                },
                "lastname": {
                    "type": "string",
                    "example": "Azamatov"
                },
                "middlename": {
                    "type": "string",
                    "example": "Azamatuly"
                },
                "phoneNumber": {
                    "type": "string",
                    "example": "8(777)777-77-77"
                }
            }
        },
        "domain.StaffResponse": {
            "type": "object",
            "properties": {
                "cityName": {
                    "type": "string",
                    "example": "Almaty"
                },
                "firstname": {
                    "type": "string",
                    "example": "Azamat"
                },
                "id": {
                    "type": "string",
                    "example": "bbb7b76f-7bb5-11ee-9ef7-9828a62c67e0"
                },
                "lastname": {
                    "type": "string",
                    "example": "Azamatov"
                },
                "middlename": {
                    "type": "string",
                    "example": "Azamatuly"
                },
                "phoneNumber": {
                    "type": "string",
                    "example": "8(777)777-77-77"
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
