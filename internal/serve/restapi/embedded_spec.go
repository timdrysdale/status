// Code generated by go-swagger; DO NOT EDIT.

package restapi

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
)

var (
	// SwaggerJSON embedded version of the swagger document used at generation time
	SwaggerJSON json.RawMessage
	// FlatSwaggerJSON embedded flattened version of the swagger document used at generation time
	FlatSwaggerJSON json.RawMessage
)

func init() {
	SwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for accessing status of experiments",
    "title": "Status",
    "contact": {
      "name": "Timothy Drysdale",
      "url": "https://github.com/timdrysdale",
      "email": "timothy.d.drysdale@gmail.com"
    },
    "version": "0.0"
  },
  "host": "localhost",
  "basePath": "/",
  "paths": {
    "/experiments": {
      "get": {
        "description": "Get the status of all experiments",
        "produces": [
          "application/json"
        ],
        "summary": "Get the status of all experiments",
        "operationId": "statusExperiments",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ExperimentStatuses"
            }
          },
          "401": {
            "$ref": "#/responses/Unauthorized"
          },
          "404": {
            "$ref": "#/responses/NotFound"
          },
          "500": {
            "$ref": "#/responses/InternalError"
          }
        }
      }
    },
    "/jumps": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get the status of all jump connections",
        "produces": [
          "application/json"
        ],
        "summary": "Get the status of all jump connections",
        "operationId": "statusJumps",
        "responses": {
          "200": {
            "description": "Status",
            "schema": {
              "$ref": "#/definitions/JumpStatuses"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ExperimentStatus": {
      "type": "object",
      "title": "Status of an experiment",
      "required": [
        "data",
        "name",
        "test",
        "video"
      ],
      "properties": {
        "data": {
          "$ref": "#/definitions/StreamStatus"
        },
        "name": {
          "description": "name of the experiment",
          "type": "string",
          "example": "spin30"
        },
        "test": {
          "$ref": "#/definitions/TestStatus"
        },
        "video": {
          "$ref": "#/definitions/StreamStatus"
        }
      }
    },
    "ExperimentStatuses": {
      "type": "array",
      "title": "List of experiment statuses",
      "items": {
        "$ref": "#/definitions/ExperimentStatus"
      }
    },
    "JumpStatus": {
      "type": "object",
      "title": "Status of the jump connection for an experiment",
      "required": [
        "name",
        "active",
        "clients",
        "connected",
        "last"
      ],
      "properties": {
        "active": {
          "description": "is the experiment currently actively sending?",
          "type": "boolean"
        },
        "clients": {
          "description": "number of clients connected (0 if just the experiment)",
          "type": "number",
          "format": "integer"
        },
        "connected": {
          "description": "is the experiment currently connected to jump?",
          "type": "boolean"
        },
        "last": {
          "description": "duration since last send by experiment",
          "type": "string"
        },
        "name": {
          "description": "name of the experiment",
          "type": "string",
          "example": "spin30"
        }
      }
    },
    "JumpStatuses": {
      "type": "array",
      "title": "List of jump statuses",
      "items": {
        "$ref": "#/definitions/JumpStatus"
      }
    },
    "StreamStatus": {
      "type": "object",
      "title": "Status of a stream",
      "required": [
        "active",
        "clients",
        "connected",
        "last",
        "required"
      ],
      "properties": {
        "active": {
          "description": "is the experiment currently actively sending?",
          "type": "boolean"
        },
        "clients": {
          "description": "number of clients connected (0 if just the experiment)",
          "type": "number",
          "format": "integer"
        },
        "connected": {
          "description": "is the experiment currently connected to relay?",
          "type": "boolean"
        },
        "last": {
          "description": "duration since last send by experiment",
          "type": "string"
        },
        "required": {
          "description": "does the experiment require this stream?",
          "type": "boolean"
        }
      }
    },
    "TestStatus": {
      "type": "object",
      "title": "Status of a test",
      "required": [
        "passed",
        "last"
      ],
      "properties": {
        "attempts": {
          "description": "number of tests attempted to date",
          "type": "number",
          "format": "integer"
        },
        "last": {
          "description": "RFC3339 datetime of last test",
          "type": "string"
        },
        "passed": {
          "description": "did the experiment pass the test last time?",
          "type": "boolean"
        },
        "score": {
          "description": "number of tests passed to date",
          "type": "number",
          "format": "integer"
        }
      }
    }
  },
  "responses": {
    "InternalError": {
      "description": "Internal Error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
	FlatSwaggerJSON = json.RawMessage([]byte(`{
  "consumes": [
    "application/json"
  ],
  "produces": [
    "application/json"
  ],
  "schemes": [
    "http"
  ],
  "swagger": "2.0",
  "info": {
    "description": "API for accessing status of experiments",
    "title": "Status",
    "contact": {
      "name": "Timothy Drysdale",
      "url": "https://github.com/timdrysdale",
      "email": "timothy.d.drysdale@gmail.com"
    },
    "version": "0.0"
  },
  "host": "localhost",
  "basePath": "/",
  "paths": {
    "/experiments": {
      "get": {
        "description": "Get the status of all experiments",
        "produces": [
          "application/json"
        ],
        "summary": "Get the status of all experiments",
        "operationId": "statusExperiments",
        "responses": {
          "200": {
            "description": "OK",
            "schema": {
              "$ref": "#/definitions/ExperimentStatuses"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "404": {
            "description": "The specified resource was not found",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          },
          "500": {
            "description": "Internal Error",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    },
    "/jumps": {
      "get": {
        "security": [
          {
            "Bearer": []
          }
        ],
        "description": "Get the status of all jump connections",
        "produces": [
          "application/json"
        ],
        "summary": "Get the status of all jump connections",
        "operationId": "statusJumps",
        "responses": {
          "200": {
            "description": "Status",
            "schema": {
              "$ref": "#/definitions/JumpStatuses"
            }
          },
          "401": {
            "description": "Unauthorized",
            "schema": {
              "$ref": "#/definitions/Error"
            }
          }
        }
      }
    }
  },
  "definitions": {
    "Error": {
      "type": "object",
      "required": [
        "code",
        "message"
      ],
      "properties": {
        "code": {
          "type": "string"
        },
        "message": {
          "type": "string"
        }
      }
    },
    "ExperimentStatus": {
      "type": "object",
      "title": "Status of an experiment",
      "required": [
        "data",
        "name",
        "test",
        "video"
      ],
      "properties": {
        "data": {
          "$ref": "#/definitions/StreamStatus"
        },
        "name": {
          "description": "name of the experiment",
          "type": "string",
          "example": "spin30"
        },
        "test": {
          "$ref": "#/definitions/TestStatus"
        },
        "video": {
          "$ref": "#/definitions/StreamStatus"
        }
      }
    },
    "ExperimentStatuses": {
      "type": "array",
      "title": "List of experiment statuses",
      "items": {
        "$ref": "#/definitions/ExperimentStatus"
      }
    },
    "JumpStatus": {
      "type": "object",
      "title": "Status of the jump connection for an experiment",
      "required": [
        "name",
        "active",
        "clients",
        "connected",
        "last"
      ],
      "properties": {
        "active": {
          "description": "is the experiment currently actively sending?",
          "type": "boolean"
        },
        "clients": {
          "description": "number of clients connected (0 if just the experiment)",
          "type": "number",
          "format": "integer"
        },
        "connected": {
          "description": "is the experiment currently connected to jump?",
          "type": "boolean"
        },
        "last": {
          "description": "duration since last send by experiment",
          "type": "string"
        },
        "name": {
          "description": "name of the experiment",
          "type": "string",
          "example": "spin30"
        }
      }
    },
    "JumpStatuses": {
      "type": "array",
      "title": "List of jump statuses",
      "items": {
        "$ref": "#/definitions/JumpStatus"
      }
    },
    "StreamStatus": {
      "type": "object",
      "title": "Status of a stream",
      "required": [
        "active",
        "clients",
        "connected",
        "last",
        "required"
      ],
      "properties": {
        "active": {
          "description": "is the experiment currently actively sending?",
          "type": "boolean"
        },
        "clients": {
          "description": "number of clients connected (0 if just the experiment)",
          "type": "number",
          "format": "integer"
        },
        "connected": {
          "description": "is the experiment currently connected to relay?",
          "type": "boolean"
        },
        "last": {
          "description": "duration since last send by experiment",
          "type": "string"
        },
        "required": {
          "description": "does the experiment require this stream?",
          "type": "boolean"
        }
      }
    },
    "TestStatus": {
      "type": "object",
      "title": "Status of a test",
      "required": [
        "passed",
        "last"
      ],
      "properties": {
        "attempts": {
          "description": "number of tests attempted to date",
          "type": "number",
          "format": "integer"
        },
        "last": {
          "description": "RFC3339 datetime of last test",
          "type": "string"
        },
        "passed": {
          "description": "did the experiment pass the test last time?",
          "type": "boolean"
        },
        "score": {
          "description": "number of tests passed to date",
          "type": "number",
          "format": "integer"
        }
      }
    }
  },
  "responses": {
    "InternalError": {
      "description": "Internal Error",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "NotFound": {
      "description": "The specified resource was not found",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    },
    "Unauthorized": {
      "description": "Unauthorized",
      "schema": {
        "$ref": "#/definitions/Error"
      }
    }
  },
  "securityDefinitions": {
    "Bearer": {
      "type": "apiKey",
      "name": "Authorization",
      "in": "header"
    }
  }
}`))
}
