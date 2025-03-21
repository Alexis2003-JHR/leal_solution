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
        "/branches": {
            "get": {
                "description": "Recuperar las sucursales asociadas a una empresa específica.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Branches"
                ],
                "summary": "Obtener sucursales de comercio",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID tributario de la empresa",
                        "name": "tax_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "Endpoint to create a new branch for a business. Requires branch details in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Branches"
                ],
                "summary": "Crear nueva sucursal",
                "parameters": [
                    {
                        "description": "Branch data",
                        "name": "branch",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateBranch"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/business": {
            "post": {
                "description": "Endpoint to register a new business. Requires business details in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Businesses"
                ],
                "summary": "Crear nuevo comercio",
                "parameters": [
                    {
                        "description": "Business data",
                        "name": "business",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateBusiness"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/campaigns": {
            "get": {
                "description": "Recuperar las campañas asociadas a una empresa específica.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Campaigns"
                ],
                "summary": "Obtener campañas de comercio",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "ID tributario de la empresa",
                        "name": "tax_id",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            },
            "post": {
                "description": "Endpoint to create a new marketing campaign for a business. Requires campaign details in the request body.",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Campaigns"
                ],
                "summary": "Crear nueva campaña",
                "parameters": [
                    {
                        "description": "Campaign data",
                        "name": "campaign",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateCampaign"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/redemptions/points": {
            "post": {
                "description": "Permitir a un usuario redimir puntos por un premio específico.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rewards"
                ],
                "summary": "Redimir puntos",
                "parameters": [
                    {
                        "description": "Datos para redimir puntos",
                        "name": "redeem",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.RedeemPoints"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/rewards": {
            "post": {
                "description": "Guardar los datos de un nuevo premio en la base de datos.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Rewards"
                ],
                "summary": "Crear nuevo premio",
                "parameters": [
                    {
                        "description": "Datos del premio",
                        "name": "reward",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateReward"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/transactions": {
            "post": {
                "description": "Procesar y encolar una nueva transacción.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Transactions"
                ],
                "summary": "Procesar transacción",
                "parameters": [
                    {
                        "description": "Datos de la transacción",
                        "name": "transaction",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.ProcessTransaction"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Guardar los datos de un nuevo usuario en la base de datos.",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Users"
                ],
                "summary": "Crear nuevo usuario",
                "parameters": [
                    {
                        "description": "Datos del usuario",
                        "name": "user",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CreateUser"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK"
                    }
                }
            }
        }
    },
    "definitions": {
        "request.ConversionFactor": {
            "type": "object",
            "properties": {
                "cashback_por_unidad": {
                    "type": "number"
                },
                "puntos_por_unidad": {
                    "type": "number"
                },
                "valor_minimo": {
                    "type": "number"
                }
            }
        },
        "request.CreateBranch": {
            "type": "object",
            "properties": {
                "nit_empresa": {
                    "type": "integer"
                },
                "nombre_sucursal": {
                    "type": "string"
                },
                "valor_conversion": {
                    "$ref": "#/definitions/request.ConversionFactor"
                }
            }
        },
        "request.CreateBusiness": {
            "type": "object",
            "properties": {
                "correo": {
                    "type": "string"
                },
                "nit": {
                    "type": "integer"
                },
                "razon_social": {
                    "type": "string"
                },
                "telefono": {
                    "type": "integer"
                },
                "valor_conversion": {
                    "$ref": "#/definitions/request.ConversionFactor"
                }
            }
        },
        "request.CreateCampaign": {
            "type": "object",
            "properties": {
                "compra_minima": {
                    "type": "number"
                },
                "fecha_fin": {
                    "type": "string"
                },
                "fecha_inicio": {
                    "type": "string"
                },
                "id_sucursal": {
                    "type": "integer"
                },
                "multiplicador_cashback": {
                    "type": "number"
                },
                "multiplicador_puntos": {
                    "type": "number"
                },
                "nit_empresa": {
                    "type": "integer"
                }
            }
        },
        "request.CreateReward": {
            "type": "object",
            "properties": {
                "descripcion": {
                    "type": "string"
                },
                "nit_empresa": {
                    "type": "integer"
                },
                "nombre": {
                    "type": "string"
                },
                "puntos_requeridos": {
                    "type": "integer"
                }
            }
        },
        "request.CreateUser": {
            "type": "object",
            "properties": {
                "correo": {
                    "type": "string"
                },
                "nombre": {
                    "type": "string"
                },
                "numero_documento": {
                    "type": "integer"
                }
            }
        },
        "request.ProcessTransaction": {
            "type": "object",
            "properties": {
                "id_sucursal": {
                    "type": "integer"
                },
                "usuario": {
                    "$ref": "#/definitions/request.CreateUser"
                },
                "valor": {
                    "type": "number"
                }
            }
        },
        "request.RedeemPoints": {
            "type": "object",
            "properties": {
                "id_premio": {
                    "type": "integer"
                },
                "nit_empresa": {
                    "type": "integer"
                },
                "usuario": {
                    "$ref": "#/definitions/request.CreateUser"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:6060",
	BasePath:         "/api/v1",
	Schemes:          []string{},
	Title:            "Tag Service API",
	Description:      "A Tag service API in Go using Gin Framework",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
