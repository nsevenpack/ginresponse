# ginresponse


[![Go Reference](https://pkg.go.dev/badge/github.com/nsevenpack/ginresponse.svg)](https://pkg.go.dev/github.com/nsevenpack/ginresponse)
[![Version](https://img.shields.io/github/v/tag/nsevenpack/ginresponse?label=version&sort=semver)](https://github.com/nsevenpack/ginresponse/releases)
[![CI](https://github.com/nsevenpack/ginresponse/actions/workflows/release.yml/badge.svg)](https://github.com/nsevenpack/ginresponse/actions/workflows/release.yml)
[![License](https://img.shields.io/github/license/nsevenpack/ginresponse)](https://github.com/nsevenpack/ginresponse/blob/main/LICENSE)


Petit package de response api, par default en JSON
utilisant le framework [gin-gonic/gin]

## Installation

```bash
# installe la derniere version 0.x.x
go get github.com/nsevenpack/ginresponse

# liste les versions disponibles pour 0.x.x
go list -m -versions github.com/nsevenpack/ginresponse

# installe une version précise
go get github.com/nsevenpack/ginresponse@v0.1.0
```

## Utilisation 

- pour pouvoir utiliser les helpers du package, il faut initialiser ginresponse
pour cela il faut appeller la fonction `ginresponse.SetFormatter(ginresponse.JsonFormatter{})`  
et lui donner en parametre un `ginresponse.Formatter` ginresponse possede un formatter json par defaut.

```go
func init() {
    ginresponse.SetFormatter(ginresponse.JsonFormatter{})
}
```

- ensuite dans vos handlers vous pouvez utiliser les helpers de ginresponse
donner en parametre le context de gin et les parametres de la reponse
    - data : des données souvent de type struct ou map
    - message : un message de reponse

```go
Success(c *gin.Context, message string, data interface{})
Created(c *gin.Context, message string, data interface{})
BadRequest(c *gin.Context, message string, err interface{})
Unauthorized(c *gin.Context, message string, err interface{})
Forbidden(c *gin.Context, message string, err interface{})
NotFound(c *gin.Context, message string, err interface{})
InternalServerError(c *gin.Context, message string, err interface{})
NoContent(c *gin.Context, message string)

// exemple d'utilisatoin
func GetUser(c *gin.Context) {
    // vous pouvez utiliser des structs, maps ou des slices
	user := map[string]string{"name": "Alice"}
	ginresponse.Success(c, "User found", user)
}
```

- voici le retour de la réponse (map)
```go
// la map utilisé en interne
resp := map[string]interface{}{
	"message":   message,
	"data":     nil,
	"error":    nil,
	"meta": map[string]interface{}{
		"status":    status,
		"path":      c.Request.URL.Path,
		"method":    c.Request.Method,
		"timestamp": time.Now().Format(time.RFC3339), 
	},
}
```

- voici le retour de la réponse (json)
```json
// success
{
  "message": "Utilisateur récupéré avec succès",
  "data": {
    "id": 42,
    "name": "Alice Dupont",
    "email": "alice@example.com",
    "role": "admin"
  },
  "error": null,
  "meta": {
    "status": 200,
    "path": "/api/users/42",
    "method": "GET",
    "timestamp": "2025-04-17T20:10:00Z"
  }
}

// erreur simple
{
  "message": "Aucun utilisateur trouvé avec cet ID",
  "data": null,
  "error": "User not found",
  "meta": {
    "status": 404,
    "path": "/api/users/999",
    "method": "GET",
    "timestamp": "2025-04-17T20:10:00Z"
  }
}

// erreur utilisant une slice du struct d'erreur
{
  "message": "Erreur de validation",
  "data": null,
  "error": [
	{
		"message": "Le champ 'email' est requis",
		"type": "validation",
		"field": "email",
		"detail": "L'email doit être au format valide"
	},
	{
		"message": "Le champ 'password' est requis",
		"type": "validation",
		"field": "password",
		"detail": ""
	}
 ],
  "meta": {
	"status": 400,
	"path": "/api/users/42",
	"method": "POST",
	"timestamp": "2025-04-17T20:10:00Z"
  }
}
```

- vous pouvez creer vos propre formatter en implementant l'interface `ginresponse.Formatter`
```go
// creer votre propre formatter
// comme ci dessous

type MyFormatter struct{}

func (f MyFormatter) Format(c *gin.Context, status int, message string, data interface{}, err interface{}) map[string]interface{} {
	return map[string]interface{}{
		"code":    status,
		"msg":     message,
		"result":  data,
		"error":   err,
		"request": c.Request.Method + " " + c.Request.URL.Path,
	}
}

func (f MyFormatter) Render(c *gin.Context, status int, payload map[string]interface{}) {
	c.JSON(status, payload)
}

// et ensuite donner le à SetFormatter
func init() {
    ginresponse.SetFormatter(MyFormatter{})
}

// maintenant ginresponse utilise votre propre format de données de réponse (dans notre exemple en JSON)
```

- ginresponse vous donne un struct d'erreur `ginresponse.ErrorModel` qui contient les champs suivants
```go
type ErrorModel struct {
	Message string `json:"message"`
	Type   string `json:"type"`
	Field  string `json:"field"`
	Detail string `json:"detail"`
}
```
vous pouvez l'utiliser pour renvoyer des erreurs plus détaillées