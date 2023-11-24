# capturetheflag

A simple Golang app that exposes APIs for challengers to find info in them.

## Process

- User can register a username
- User can check if their username is saved
- User can get the secret attached to their username
- User can request the challenge
    - User can ask for a hint
- User can submit the result to the challenge

## Available Routes

```md
PATH                METHOD      BODY
/ping			    GET
/signup     	    POST        {"User" : "xxxx"}
/check      	    POST        {"User" : "xxxx"}
/secret     	    POST        {"User" : "xxxx"}
/getLevel			POST		AuthRequest
/getUserPoints		POST		AuthRequest
/getChallenge		POST		AuthRequest
/getHint			POST 		AuthRequest
/submitChallenge	POST		FullRequest
```

## Request Body

```go
Request struct {
	User string `json:"User"`
}
AuthRequest struct {
	User   string `json:"User"`
	Secret string `json:"Secret"`
}

FullRequest struct {
	User    string `json:"User"`
	Secret  string `json:"Secret"`
	Content struct {
		Level     uint `json:"Level"`
		Challenge struct {
			Username string `json:"Username"`
			Secret   string `json:"Secret"`
			Points   uint   `json:"Points"`
		} `json:"Challenge"`
		Protocol  string `json:"Protocol"`
		SecretKey string `json:"SecretKey"`
	} `json:"Content"`
}

```go