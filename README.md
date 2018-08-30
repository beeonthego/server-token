# server key and token generator
A command line tool to generate JWT access tokens and random signing keys.

# Purpose

This tool generates access tokens on the command line to test server API integration with [Gitea](https://gitea.io), a lightweight self-hosted git service. Quickly edit and run. Generate tokens for different scenarios to test server API response.

The token is signed with HS256 algorithm. The token can be used to check access on API routes, or for deployment to integrated servers.

It can also generate a pseudo random key and output the base64 url encoded key to the console.

This tool uses the following environment variables if set. 
```
GITEA_ROOT_SERVER_TOKEN_AUDIENCE 
```
the name to include in standard claims of JWT token.


```
GITEA_ROOT_SERVER_TOKEN_SIGNING_SECRET 
```
The secret key to sign the access tokens. The signing key should be base64 url encoded, and must be at least 32 bytes when decoded. If the key can not be decoded, or the key is less than 32 bytes, this tool will generate a random key and use the random key instead.


# Why creating this tool? Why not using Gitea GUI to create an access token?

Server token is a special access token. It has super user privilege, uses a different key to sign, and may not be associated with a real user on Gitea instance. 


# Installation

This tool is written in Go, and uses the [jwt-go](https://github.com/dgrijalva/jwt-go) library. Your computer needs Go installed and properly configured. 

To install the source of this tool, simply copy and paste the code to your favorite editor, or run this on the command line
```
go get github.com/beeonthego/server-token
```
or 
```
git clone https://github.com/beeonthego/server-token.git
```

To install jwt-go library, run this on the command line
```
go get github.com/dgrijalva/jwt-go
```

Then navigate to the folder with the code of this tool, edit and run 

```
go run server_token.go
```

# Caveats

This tool is meant to run on the command line to generate base64 URL encoded key and token for testing or deployment. It will simply throw out errors and exit if for some reason it can not generate random bytes. It does not parse or verify tokens.

The pseudo random bytes are generated with crypto/rand library included in Golang package. 

# License

MIT

