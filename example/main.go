package main

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
)

func main() {
	//mySigningKey := []byte("AllYourBase")
	//
	//type CustomClaims struct {
	//	//Name string `json:"name"`
	//	//Role string `json:"role"`
	//	jwt.StandardClaims
	//}
	//
	//customClaims := &CustomClaims{
	//	//Name:           "denis",
	//	//Role:           "admin",
	//	StandardClaims: jwt.StandardClaims{
	//		ExpiresAt: -1,
	//		Issuer:    "test",
	//	},
	//}
	//
	//// Create the Claims
	////claims := &jwt.StandardClaims{
	////	ExpiresAt: 15000,
	////	Issuer:    "test",
	////}
	//
	//token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	//ss, err := token.SignedString(mySigningKey)
	//fmt.Printf("%v %v", ss, err)
	//Output:
	//eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE1MDAwLCJpc3MiOiJ0ZXN0In0.QsODzZu3lUZMVdhbO76u3Jv02iYCvEHcYVUI1kOWEU0 <nil>

	mySigningKey := []byte("AllYourBase")

	// Create the Claims
	claims := &jwt.StandardClaims{
		ExpiresAt: 15000,
		Issuer:    "test",
	}

	token := jwt.NewWithClaims(jwt.GetSigningMethod("HS256"), claims)
	ss, err := token.SignedString(mySigningKey)
	fmt.Printf("%v %v", ss, err)

	tok, err := jwt.ParseWithClaims(ss, &jwt.StandardClaims{}, func(tok *jwt.Token) (interface{}, error) {
		return []byte("AllYourBase"), nil })
	if claims, ok := tok.Claims.(*jwt.StandardClaims); ok && tok.Valid {
		fmt.Printf("инфо - %v", claims.ExpiresAt)
	} else {
		fmt.Println(err)
	}
}
