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
        "/exercise/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get an exercise by its ID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exercise"
                ],
                "summary": "Get exercise",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exercise ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.ExerciseGResUReq"
                        }
                    },
                    "400": {
                        "description": "Invalid exercise ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/exercises": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all exercises",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Exercise"
                ],
                "summary": "Get all exercises",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lesson ID",
                        "name": "lesson_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Exercise type",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.ExerciseGARes"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/leaderboard": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get the leaderboard sorted by XP",
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "leaderboard"
                ],
                "summary": "Get Leaderboard",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "type": "array",
                            "items": {
                                "$ref": "#/definitions/genprotos.LeadboardRes"
                            }
                        }
                    },
                    "500": {
                        "description": "Internal Server Error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/lesson/{id}": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get a lesson by it's UUID",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lesson"
                ],
                "summary": "Get lesson",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Leesson ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.LessonCReqGRes"
                        }
                    },
                    "400": {
                        "description": "Invalid lesson ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/lessons": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Get all lessons",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Lesson"
                ],
                "summary": "Get all lessons",
                "parameters": [
                    {
                        "type": "string",
                        "description": "name",
                        "name": "name",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "lang_1",
                        "name": "lang_1",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "lang_2",
                        "name": "lang_2",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "level",
                        "name": "level",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "order",
                        "name": "order",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.LessonGARes"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/mydata": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Retrieves user data",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "UserData"
                ],
                "summary": "Get User Data",
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/handlers.MyDataResponse"
                        }
                    },
                    "400": {
                        "description": "Invalid user ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/start-test": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Start a test by checking the answers of the provided quiz requests against the exercises of the specified lesson",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "quiz"
                ],
                "summary": "Start a test",
                "parameters": [
                    {
                        "description": "Test Check Request",
                        "name": "YourAnswers",
                        "in": "body",
                        "required": true,
                        "schema": {
                            "$ref": "#/definitions/genprotos.TestCheckReqForSwagger"
                        }
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.TestResultRes"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/vocabularies": {
            "get": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Gets vocabularies",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vocabulary"
                ],
                "summary": "Get Vocabularies",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Lesson ID",
                        "name": "lesson_id",
                        "in": "query"
                    },
                    {
                        "type": "string",
                        "description": "Exercise type",
                        "name": "type",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "limit",
                        "name": "limit",
                        "in": "query"
                    },
                    {
                        "type": "integer",
                        "description": "offset",
                        "name": "offset",
                        "in": "query"
                    }
                ],
                "responses": {
                    "200": {
                        "description": "OK",
                        "schema": {
                            "$ref": "#/definitions/genprotos.VocabulariesGARes"
                        }
                    },
                    "400": {
                        "description": "Invalid parameters",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        },
        "/vocabulary/{id}": {
            "post": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Adds an exercise to the vocabulary",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vocabulary"
                ],
                "summary": "Add to Vocabulary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exercise ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "200": {
                        "description": "Exercise added to vocabulary",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid exercise ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            },
            "delete": {
                "security": [
                    {
                        "BearerAuth": []
                    }
                ],
                "description": "Deletes an exercise from the vocabulary",
                "consumes": [
                    "application/json"
                ],
                "produces": [
                    "application/json"
                ],
                "tags": [
                    "Vocabulary"
                ],
                "summary": "Delete from Vocabulary",
                "parameters": [
                    {
                        "type": "string",
                        "description": "Exercise ID",
                        "name": "id",
                        "in": "path",
                        "required": true
                    }
                ],
                "responses": {
                    "204": {
                        "description": "Exercise deleted from vocabulary",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "400": {
                        "description": "Invalid exercise ID",
                        "schema": {
                            "type": "string"
                        }
                    },
                    "500": {
                        "description": "Server error",
                        "schema": {
                            "type": "string"
                        }
                    }
                }
            }
        }
    },
    "definitions": {
        "genprotos.CheckReq": {
            "type": "object",
            "properties": {
                "correct_answer": {
                    "type": "string"
                },
                "exercise_id": {
                    "type": "string"
                }
            }
        },
        "genprotos.ExerciseGARes": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "exercises": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.ExerciseGResUReq"
                    }
                }
            }
        },
        "genprotos.ExerciseGResUReq": {
            "type": "object",
            "properties": {
                "correct_answer": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lesson_id": {
                    "type": "string"
                },
                "options": {
                    "type": "string"
                },
                "question": {
                    "type": "string"
                },
                "type": {
                    "type": "string"
                }
            }
        },
        "genprotos.LeadboardRes": {
            "type": "object",
            "properties": {
                "users": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.LeadboardUserRes"
                    }
                }
            }
        },
        "genprotos.LeadboardUserRes": {
            "type": "object",
            "properties": {
                "daily_streak": {
                    "type": "integer"
                },
                "level": {
                    "type": "string"
                },
                "native_lang": {
                    "type": "string"
                },
                "played_games_count": {
                    "type": "integer"
                },
                "user_id": {
                    "type": "string"
                },
                "winning_percentage": {
                    "type": "number"
                },
                "xp": {
                    "type": "integer"
                }
            }
        },
        "genprotos.LessonCReqGRes": {
            "type": "object",
            "properties": {
                "content": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "lang_1": {
                    "type": "string"
                },
                "lang_2": {
                    "type": "string"
                },
                "level": {
                    "type": "string"
                },
                "name": {
                    "type": "string"
                },
                "order": {
                    "type": "integer"
                },
                "title": {
                    "type": "string"
                }
            }
        },
        "genprotos.LessonGARes": {
            "type": "object",
            "properties": {
                "count": {
                    "type": "integer"
                },
                "lessons": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.LessonCReqGRes"
                    }
                }
            }
        },
        "genprotos.TestCheckReqForSwagger": {
            "type": "object",
            "properties": {
                "lesson_id": {
                    "type": "string"
                },
                "requests": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.CheckReq"
                    }
                }
            }
        },
        "genprotos.TestResultRes": {
            "type": "object",
            "properties": {
                "correct_answers_count": {
                    "type": "integer"
                },
                "feedback": {
                    "type": "string"
                },
                "tests_count": {
                    "type": "integer"
                },
                "xp_given": {
                    "type": "integer"
                }
            }
        },
        "genprotos.UserDataGRes": {
            "type": "object",
            "properties": {
                "daily_streak": {
                    "type": "integer"
                },
                "level": {
                    "type": "string"
                },
                "native_lang": {
                    "type": "string"
                },
                "played_games_count": {
                    "type": "integer"
                },
                "winning_percentage": {
                    "type": "number"
                },
                "xp": {
                    "type": "integer"
                }
            }
        },
        "genprotos.VocabulariesGARes": {
            "type": "object",
            "properties": {
                "vocabularies": {
                    "type": "array",
                    "items": {
                        "$ref": "#/definitions/genprotos.ExerciseGResUReq"
                    }
                }
            }
        },
        "handlers.MyDataResponse": {
            "type": "object",
            "properties": {
                "user": {
                    "$ref": "#/definitions/models.User"
                },
                "user_data": {
                    "$ref": "#/definitions/genprotos.UserDataGRes"
                }
            }
        },
        "models.User": {
            "type": "object",
            "properties": {
                "email": {
                    "type": "string"
                },
                "id": {
                    "type": "string"
                },
                "password": {
                    "type": "string"
                },
                "username": {
                    "type": "string"
                }
            }
        }
    },
    "securityDefinitions": {
        "BearerAuth": {
            "type": "apiKey",
            "name": "Authorization",
            "in": "header"
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
