package main

//go:generate swagger generate spec -o ../../api/swagger.yaml --scan-models
import (
	_ "github.com/dyjwl/gin-web-plugin-demo/swagger/docs"
)

func main() {
}
