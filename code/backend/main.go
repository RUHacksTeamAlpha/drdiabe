package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	firebase "firebase.google.com/go"
	jwtMiddleware "github.com/auth0/go-jwt-middleware"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"google.golang.org/api/option"
)

var (
	app            *firebase.App
	corsMiddleware = cors.New(cors.Config{
		// 	// AllowOrigins:     []string{"https://drdiabe.tech", "http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Authorization", "Origin", "Content-Length", "Content-Type"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           12 * time.Hour,
	})
)

var (
	sqlConnString string
)

func main() {
	_ = godotenv.Load("secrets/.env")
	sqlConnString = os.Getenv("DB_CONN")

	fmt.Println("Starting Server")
	r := gin.Default()
	r.Use(corsMiddleware)

	opt := option.WithCredentialsFile("secrets/firebase-key.json") //TO DO UPDATE KEY
	app, err := firebase.NewApp(context.Background(), nil, opt)
	if err != nil {
		log.Fatalf("error initializing app: %v", err)
	}

	authMiddleware := func() gin.HandlerFunc {
		return func(c *gin.Context) {
			ctx := context.Background()
			idToken, _ := jwtMiddleware.FromAuthHeader(c.Request)

			// fmt.Println(c.Request, idToken)

			client, err := app.Auth(ctx)
			if err != nil {
				log.Printf("error getting Auth client: %v\n", err)
				c.JSON(401, err)
				return
			}

			token, err := client.VerifyIDToken(ctx, idToken)
			if err != nil {
				log.Printf("error verifying ID token: %v\n", err)
				c.JSON(401, err)
				return
			}

			// log.Printf("Verified ID token: %v\n", token)
			c.Set("token", token)
			c.Set("uid", token.UID)
			// log.Printf("UID %v", token.UID)
		}
	}

	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	r.GET("/auth", authMiddleware(), func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": c.GetString("token"),
		})
	})

	r.GET("/profile", authMiddleware(), getProfile)
	r.PUT("/profile", authMiddleware(), updateProfile)
	r.GET("/data", authMiddleware(), getData)
	r.POST("/data", authMiddleware(), postData)

	// r.POST("/endpoint", authMiddleware(), endpointDandler)

	r.Run(":8081") // listen and serve on 0.0.0.0:8081 (for windows "localhost:8081")
}
