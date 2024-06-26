// Package docs Code generated by swaggo/swag. DO NOT EDIT
package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "termsOfService": "http://swagger.io/terms/",
        "contact": {
            "name": "Pedro ramos",
            "email": "pedrohscramos@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/suppliers": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "get all suppliers",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "suppliers"
                ],
                "summary": "List suppliers",
                "parameters": [
                    {
                        "type": "string",
                        "description": "page number",
                        "name": "page",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/entity.Supplier"
                            }
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/suppliers/{id}": {
            "get": {
                "security": [
                    {
                        "ApiKeyAuth": []
                    }
                ],
                "description": "Get a supplier",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "suppliers"
                ],
                "summary": "Get a supplier",
                "parameters": [
                    {
                        "type": "string",
                        "format": "int",
                        "description": "supplier ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/entity.Supplier"
                        }
                    },
                    "404": {
                        "description": "Not Found"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/users": {
            "post": {
                "description": "Create user",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Create user",
                "parameters": [
                    {
                        "description": "user request",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.CreateUserInput"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created"
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        },
        "/users/generate_token": {
            "post": {
                "description": "Get a user JWT",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "users"
                ],
                "summary": "Get a user JWT",
                "parameters": [
                    {
                        "description": "user credentials",
                        "name": "request",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/dto.GetJWTInput"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/dto.GetJWTOutput"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/handlers.Error"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "dto.CreateUserInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.GetJWTInput": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                }
            }
        },
        "dto.GetJWTOutput": {
            "type": "object",
            "properties": {
                "access_token": {
                    "type": "string"
                }
            }
        },
        "entity.Supplier": {
            "type": "object",
            "properties": {
                "additional_info": {
                    "type": "string"
                },
                "availability_id": {
                    "type": "string"
                },
                "benefit_id": {
                    "type": "string"
                },
                "benefit_ordert_id": {
                    "type": "string"
                },
                "benefit_type": {
                    "type": "string"
                },
                "billing_currency": {
                    "type": "string"
                },
                "billing_pre_tax_total": {
                    "type": "number"
                },
                "charge_end_date": {
                    "type": "string"
                },
                "charge_start_date": {
                    "type": "string"
                },
                "charge_type": {
                    "type": "string"
                },
                "consumed_service": {
                    "type": "string"
                },
                "credit_percentage": {
                    "type": "integer"
                },
                "credit_type": {
                    "type": "string"
                },
                "customer_country": {
                    "type": "string"
                },
                "customer_domain_name": {
                    "type": "string"
                },
                "customer_id": {
                    "type": "string"
                },
                "customer_name": {
                    "type": "string"
                },
                "effective_unit_price": {
                    "type": "number"
                },
                "entilement_description": {
                    "type": "string"
                },
                "entilement_id": {
                    "type": "string"
                },
                "id": {
                    "type": "integer"
                },
                "invoice_number": {
                    "type": "string"
                },
                "meter_category": {
                    "type": "string"
                },
                "meter_id": {
                    "type": "string"
                },
                "meter_name": {
                    "type": "string"
                },
                "meter_region": {
                    "type": "string"
                },
                "meter_sub_category": {
                    "type": "string"
                },
                "meter_type": {
                    "type": "string"
                },
                "mpn_id": {
                    "type": "integer"
                },
                "partner_earned_credit_percentage": {
                    "type": "integer"
                },
                "partner_id": {
                    "type": "string"
                },
                "partner_name": {
                    "type": "string"
                },
                "pc_to_bc_exchange_rate": {
                    "type": "integer"
                },
                "pc_to_bc_exchange_rate_date": {
                    "type": "string"
                },
                "pricing_currency": {
                    "type": "string"
                },
                "pricing_pre_tax_total": {
                    "type": "number"
                },
                "product_id": {
                    "type": "string"
                },
                "product_name": {
                    "type": "string"
                },
                "publisher_id": {
                    "type": "integer"
                },
                "publisher_name": {
                    "type": "string"
                },
                "quantity": {
                    "type": "number"
                },
                "resource_group": {
                    "type": "string"
                },
                "resource_location": {
                    "type": "string"
                },
                "resource_uri": {
                    "type": "string"
                },
                "service_info1": {
                    "type": "string"
                },
                "service_info2": {
                    "type": "string"
                },
                "sku_id": {
                    "type": "string"
                },
                "sku_name": {
                    "type": "string"
                },
                "subscription_description": {
                    "type": "string"
                },
                "subscription_id": {
                    "type": "string"
                },
                "tags": {
                    "type": "string"
                },
                "tier2_mpn_id": {
                    "type": "integer"
                },
                "unit": {
                    "type": "string"
                },
                "unit_price": {
                    "type": "number"
                },
                "unit_type": {
                    "type": "string"
                },
                "usage_date": {
                    "type": "string"
                }
            }
        },
        "handlers.Error": {
            "type": "object",
            "properties": {
                "message": {
                    "type": "string"
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
	Host:             "localhost:8000",
	BasePath:         "/",
	Schemes:          []string{},
	Title:            "API Enube",
	Description:      "Api para gerenciamento de fornecedores",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
