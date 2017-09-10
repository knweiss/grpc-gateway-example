package echopb 

const (
swagger = `{
  "swagger": "2.0",
  "info": {
    "title": "service.proto",
    "version": "version not set"
  },
  "schemes": [
    "http",
    "https"
  ],
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "paths": {
    "/v1/echo": {
      "post": {
        "summary": "Echo returns the received message back to the caller.",
        "operationId": "Echo",
        "responses": {
          "200": {
            "description": "",
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        },
        "parameters": [
          {
            "name": "body",
            "in": "body",
            "required": true,
            "schema": {
              "$ref": "#/definitions/echopbEchoMessage"
            }
          }
        ],
        "tags": [
          "EchoService"
        ]
      }
    }
  },
  "definitions": {
    "echopbEchoMessage": {
      "type": "object",
      "properties": {
        "value": {
          "type": "string"
        }
      },
      "description": "EchoMessage is the message text of an Echo() request and response."
    }
  }
}
`
)
