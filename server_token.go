// MIT License
//
// Copyright (c) 2018 BeeOnTheGo

package main

import (
	"fmt"
	"os"
	"crypto/rand"
	"encoding/base64"
	"github.com/dgrijalva/jwt-go"
)
// the name or random string to include as audience in standard jwt claims
var GiteaRootServerTokenAudience = getRootServerTokenAudience()

var GiteaRootServerTokenSigningByte = getServerTokenSigningByte()

type TokenClaims struct {
	Owner []string `json:"owner,omitempty"`
	Repo []string `json:"repo,omitempty"`
	Branch []string `json:"branch,omitempty"` //wildcard match
	Route []string `json:"route,omitempty"`  //wildcard match
	Method []string `json:"method,omitempty"`
	jwt.StandardClaims
	UID int64 `json:"uid,omitempty"` //user id
}

const nameLength = 12
const keySize = 64

func genRandBytes () ([]byte, error) {
	arr := new([nameLength]byte)
	clone := arr[:]
	_, err := rand.Read(clone)
	return clone, err
}

func genRandKey () ([]byte, error) {
	arr := new([keySize]byte)
	clone := arr[:]
	_, err := rand.Read(clone)
	return clone, err
}

func genRandString() string {
	strBytes, err := genRandBytes()
	if err != nil {
		panic ("Failed to generate random bytes. Please try again.")
	}
	return base64.RawURLEncoding.EncodeToString(strBytes)
}

func getRootServerTokenAudience() string {
	aud := os.Getenv("GITEA_ROOT_SERVER_TOKEN_AUDIENCE")
	if aud == "" {
		aud = genRandString()
	}
	return aud
}

func getServerTokenSigningByte() []byte {
	keyByte, err := base64.RawURLEncoding.DecodeString(os.Getenv("GITEA_ROOT_SERVER_TOKEN_SIGNING_SECRET"))
	if err != nil || len(keyByte) < 32 {
		newByte, err := genRandKey()
		if err != nil {
			panic ("Failed to generate random bytes for signing token. Please try again.")
		}
		return newByte
	}
	return keyByte
}

func SignServerToken(claims *TokenClaims) string {
	
	if claims.StandardClaims.Audience == "" {
		claims.StandardClaims.Audience = GiteaRootServerTokenAudience
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(GiteaRootServerTokenSigningByte)
	if err != nil {
		panic ("Failed to sign token. Please try again.")
	}
	return ss
}

func main(){
	claims := &TokenClaims{
		Owner: []string{},
		Repo: []string{},
		Branch: []string{},
		Route: []string{},
		Method: []string{},
		StandardClaims: jwt.StandardClaims{
			Issuer: "gitea",
			Audience: GiteaRootServerTokenAudience,
		},
		UID: 0,
	}
	key := base64.RawURLEncoding.EncodeToString(GiteaRootServerTokenSigningByte)
	fmt.Println("env GITEA_ROOT_SERVER_TOKEN_SIGNING_SECRET: ")
	fmt.Println(key)
	fmt.Println("env GITEA_ROOT_SERVER_TOKEN_AUDIENCE")
	fmt.Println(GiteaRootServerTokenAudience)
	fmt.Println("Server Token: ")
	fmt.Println(SignServerToken(claims))
}
