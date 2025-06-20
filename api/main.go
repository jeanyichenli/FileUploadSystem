package api

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jeanyichenli/FileUploadSystem/api/routes"
)

func StartHttpServer(port string) {
	router := gin.Default()

	routes.SetRouter(router)

	addr := fmt.Sprintf(":%s", port)
	if err := router.Run(addr); err != nil {
		log.Fatalf("Failed to run http server, err: %s", err.Error())
	}

	fmt.Println("Start running http server..")
}
