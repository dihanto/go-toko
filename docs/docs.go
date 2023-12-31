// Code generated by swaggo/swag. DO NOT EDIT.

package docs

import "github.com/swaggo/swag"

const docTemplate = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Dihanto",
            "email": "dihanto2306@gmail.com"
        },
        "license": {
            "name": "Apache 2.0",
            "url": "http://www.apache.org/licenses/LICENSE-2.0.html"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/customer": {
            "put": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Update customer",
                "tags": [
                    "Customer"
                ],
                "summary": "Update customer",
                "parameters": [
                    {
                        "description": "Update Customer",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CustomerUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Delete customer",
                "tags": [
                    "Customer"
                ],
                "summary": "Delete customer",
                "parameters": [
                    {
                        "description": "Delete customer",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CustomerDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/customer/login": {
            "post": {
                "description": "Login customer",
                "tags": [
                    "Customer"
                ],
                "summary": "Login customer",
                "parameters": [
                    {
                        "description": "Login Customer",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CustomerLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/customer/register": {
            "post": {
                "description": "Register Customer",
                "tags": [
                    "Customer"
                ],
                "summary": "Register customer",
                "parameters": [
                    {
                        "description": "Register Customer",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.CustomerRegister"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/order": {
            "post": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Add order",
                "tags": [
                    "Order"
                ],
                "summary": "Add order",
                "parameters": [
                    {
                        "description": "Request Body",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AddOrder"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/order/{id}": {
            "post": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Find order",
                "tags": [
                    "Order"
                ],
                "summary": "Find order",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Order ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/product": {
            "get": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Get product",
                "tags": [
                    "Product"
                ],
                "summary": "Get product",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Add product",
                "tags": [
                    "Product"
                ],
                "summary": "Add product",
                "parameters": [
                    {
                        "description": "Add Product",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AddProduct"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/product/": {
            "get": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Search product",
                "tags": [
                    "Product"
                ],
                "summary": "Search product",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Search",
                        "name": "search",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Offset",
                        "name": "offset",
                        "in": "path",
                        "required": true
                    },
                    {
                        "type": "integer",
                        "description": "Limit",
                        "name": "limit",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/product/{id}": {
            "get": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Find product",
                "tags": [
                    "Product"
                ],
                "summary": "Find product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Update product",
                "tags": [
                    "Product"
                ],
                "summary": "Update product",
                "parameters": [
                    {
                        "description": "Update Product",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateProduct"
                        }
                    },
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Delete product",
                "tags": [
                    "Product"
                ],
                "summary": "Delete product",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "Product ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/product/{id}/wishlist": {
            "post": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Add product to wishlist",
                "tags": [
                    "Product"
                ],
                "summary": "Add product to wishlist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Delete product from wishlist",
                "tags": [
                    "Product"
                ],
                "summary": "Delete product from wishlist",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Product Id",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/seller": {
            "put": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Update seller",
                "tags": [
                    "Seller"
                ],
                "summary": "Update seller",
                "parameters": [
                    {
                        "description": "Update Seller",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SellerUpdate"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Delete seller",
                "tags": [
                    "Seller"
                ],
                "summary": "Delete seller",
                "parameters": [
                    {
                        "description": "Delete Seller",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SellerDelete"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/seller/login": {
            "post": {
                "description": "Register seller",
                "tags": [
                    "Seller"
                ],
                "summary": "Register seller",
                "parameters": [
                    {
                        "description": "Login Seller",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SellerLogin"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/seller/register": {
            "post": {
                "description": "Register Seller",
                "tags": [
                    "Seller"
                ],
                "summary": "Register seller",
                "parameters": [
                    {
                        "description": "Register Seller",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.SellerRegister"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        },
        "/wallet": {
            "get": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Get wallet",
                "tags": [
                    "Wallet"
                ],
                "summary": "Get wallet",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            },
            "put": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Update wallet",
                "tags": [
                    "Wallet"
                ],
                "summary": "Update wallet",
                "parameters": [
                    {
                        "description": "Update Wallet",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.UpdateWallet"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            },
            "post": {
                "security": [
                    {
                        "JWTAuth": []
                    }
                ],
                "description": "Add wallet",
                "tags": [
                    "Wallet"
                ],
                "summary": "Add wallet",
                "parameters": [
                    {
                        "description": "Add Wallet",
                        "name": "request_body",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/request.AddWallet"
                        }
                    }
                ],
                "responses": {
                    "201": {
                        "description": "Created",
                        "schema": {
                            "$ref": "#/definitions/response.WebResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "request.AddOrder": {
            "type": "object",
            "required": [
                "id_customer",
                "id_product",
                "quantity"
            ],
            "properties": {
                "id_customer": {
                    "type": "string"
                },
                "id_product": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "request.AddProduct": {
            "type": "object",
            "required": [
                "id_seller",
                "name",
                "price",
                "quantity"
            ],
            "properties": {
                "id_seller": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "request.AddWallet": {
            "type": "object",
            "required": [
                "balance",
                "id_customer"
            ],
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "id_customer": {
                    "type": "string"
                }
            }
        },
        "request.CustomerDelete": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "request.CustomerLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "request.CustomerRegister": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string",
                    "example": "dihanto@gmail.com"
                },
                "name": {
                    "type": "string",
                    "minLength": 5,
                    "example": "dihanto"
                },
                "password": {
                    "type": "string",
                    "minLength": 8,
                    "example": "awesdkoire"
                }
            }
        },
        "request.CustomerUpdate": {
            "type": "object",
            "required": [
                "email",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "request.SellerDelete": {
            "type": "object",
            "required": [
                "email"
            ],
            "properties": {
                "email": {
                    "type": "string"
                }
            }
        },
        "request.SellerLogin": {
            "type": "object",
            "required": [
                "email",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "request.SellerRegister": {
            "type": "object",
            "required": [
                "email",
                "name",
                "password"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "minLength": 5
                },
                "password": {
                    "type": "string",
                    "minLength": 8
                }
            }
        },
        "request.SellerUpdate": {
            "type": "object",
            "required": [
                "email",
                "name"
            ],
            "properties": {
                "email": {
                    "type": "string"
                },
                "name": {
                    "type": "string",
                    "minLength": 5
                }
            }
        },
        "request.UpdateProduct": {
            "type": "object",
            "required": [
                "id",
                "name",
                "price",
                "quantity"
            ],
            "properties": {
                "id": {
                    "type": "integer"
                },
                "name": {
                    "type": "string"
                },
                "price": {
                    "type": "integer"
                },
                "quantity": {
                    "type": "integer"
                }
            }
        },
        "request.UpdateWallet": {
            "type": "object",
            "required": [
                "balance",
                "id_customer"
            ],
            "properties": {
                "balance": {
                    "type": "integer"
                },
                "id_customer": {
                    "type": "string"
                }
            }
        },
        "response.WebResponse": {
            "type": "object",
            "properties": {
                "code": {
                    "type": "integer"
                },
                "data": {},
                "message": {
                    "type": "string"
                }
            }
        }
    }
}`

// SwaggerInfo holds exported Swagger Info so clients can modify it
var SwaggerInfo = &swag.Spec{
	Version:          "1.0",
	Host:             "localhost:2000",
	BasePath:         "",
	Schemes:          []string{},
	Title:            "Go-Toko API",
	Description:      "This is a simple API for marketplace.",
	InfoInstanceName: "swagger",
	SwaggerTemplate:  docTemplate,
	LeftDelim:        "{{",
	RightDelim:       "}}",
}

func init() {
	swag.Register(SwaggerInfo.InstanceName(), SwaggerInfo)
}
