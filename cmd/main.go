package main

import "RPN/cmd/services"

func main() {
	token,err := services.GenerateJWT("erfan")
	print(token,err)
}