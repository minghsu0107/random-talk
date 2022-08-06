// Package chat GENERATED BY SWAG; DO NOT EDIT
// This file was generated by swaggo/swag
package chat

import "github.com/swaggo/swag"

const docTemplatechat = `{
    "schemes": {{ marshal .Schemes }},
    "swagger": "2.0",
    "info": {
        "description": "{{escape .Description}}",
        "title": "{{.Title}}",
        "contact": {
            "name": "Ming Hsu",
            "email": "minghsu0107@gmail.com"
        },
        "version": "{{.Version}}"
    },
    "host": "{{.Host}}",
    "basePath": "{{.BasePath}}",
    "paths": {
        "/chat": {
            "get": {
                "description": "Websocket initialization endpoint for starting a chat",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Start a chat",
                "parameters": [
                    {
                        "type": "integer",
                        "description": "user id",
                        "name": "uid",
                        "in": "query",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "access token of the channel",
                        "name": "access_token",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    }
                }
            }
        },
        "/chat/channel": {
            "delete": {
                "description": "Delete a channel",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Delete channel",
                "parameters": [
                    {
                        "type": "string",
                        "description": "channel authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "id of the user that performs the deletion",
                        "name": "delby",
                        "in": "query",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "No Content",
                        "schema": {
                            "$ref": "#/definitions/common.SuccessMessage"
                        }
                    },
                    "400": {
                        "description": "Bad Request",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    }
                }
            }
        },
        "/chat/channel/messages": {
            "get": {
                "description": "List messages of a channel",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "List channel messages",
                "parameters": [
                    {
                        "type": "string",
                        "description": "channel authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    },
                    {
                        "type": "string",
                        "description": "page state",
                        "name": "ps",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/chat.MessagesPresenter"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    }
                }
            }
        },
        "/chat/chanusers": {
            "get": {
                "description": "Get all users of a channel",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Get channel users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "channel authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/chat.UserIDsPresenter"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    }
                }
            }
        },
        "/chat/chanusers/online": {
            "get": {
                "description": "Get all online users of a channel",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "chat"
                ],
                "summary": "Get online users",
                "parameters": [
                    {
                        "type": "string",
                        "description": "channel authorization",
                        "name": "Authorization",
                        "in": "header",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/chat.UserIDsPresenter"
                        }
                    },
                    "401": {
                        "description": "Unauthorized",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "404": {
                        "description": "Not Found",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "$ref": "#/definitions/common.ErrResponse"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "chat.MessagePresenter": {
            "type": "object",
            "properties": {
                "event": {
                    "type": "integer"
                },
                "message_id": {
                    "type": "string"
                },
                "payload": {
                    "type": "string"
                },
                "seen": {
                    "type": "boolean"
                },
                "time": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                }
            }
        },
        "chat.MessagesPresenter": {
            "type": "object",
            "properties": {
                "messages": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/chat.MessagePresenter"
                    }
                },
                "next_ps": {
                    "type": "string"
                }
            }
        },
        "chat.UserIDsPresenter": {
            "type": "object",
            "properties": {
                "user_ids": {
                    "type": "array",
                    "items": {
                        "type": "string"
                    }
                }
            }
        },
        "common.ErrResponse": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string"
                }
            }
        },
        "common.SuccessMessage": {
            "type": "object",
            "properties": {
                "msg": {
                    "type": "string",
                    "example": "ok"
                }
            }
        }
    }
}`

// SwaggerInfochat holds exported Swagger Info so clients can modify it
var SwaggerInfochat = &swag.Spec{
	Version:          "2.0",
	Host:             "",
	BasePath:         "/api",
	Schemes:          []string{},
	Title:            "Chat Service Swagger API",
	Description:      "Chat service API",
	InfoInstanceName: "chat",
	SwaggerTemplate:  docTemplatechat,
}

func init() {
	swag.Register(SwaggerInfochat.InstanceName(), SwaggerInfochat)
}