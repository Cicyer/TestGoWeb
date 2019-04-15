package main

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/kataras/iris"
	"github.com/iris-contrib/middleware/cors"
	"fmt"
	jwtmiddleware "github.com/iris-contrib/middleware/jwt"
	"time"
)

type Token struct {
	Token string `json:"token"`
}

func myHandler(ctx iris.Context) {
	user := ctx.Values().Get("jwt").(*jwt.Token)
	ctx.JSON(user)
}

func main(){
	fmt.Println("server intialing...")
	app:=iris.New()
	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // allows everything, use that to change the hosts.
		AllowCredentials: true,
		AllowedMethods:   []string{"PUT", "PATCH", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Origin", "Authorization"},
		ExposedHeaders:   []string{"Accept", "Content-Type", "Content-Length", "Accept-Encoding", "X-CSRF-Token", "Authorization"},
	})
	v1 := app.Party("/api", crs).AllowMethods(iris.MethodOptions)
	{
		v1.Post("/login", func(ctx iris.Context) {
			token := jwt.New(jwt.SigningMethodHS256)
			claims := make(jwt.MapClaims)
			claims["exp"] = time.Now().Add(time.Hour * time.Duration(1)).Unix()
			claims["iat"] = time.Now().Unix()
			token.Claims = claims
			tokenString, err := token.SignedString([]byte("My Secret"))
			if err != nil {
			}
			response := Token{tokenString}

			ctx.JSON(response)
		})
		v1.Get("/about", func(ctx iris.Context) {
			ctx.JSON(result)
		})
		v1.Post("/send", func(ctx iris.Context) {
			ctx.WriteString("sent")
		})
		v1.Put("/updated", func(ctx iris.Context) {
			ctx.WriteString("updated")
		})
		v1.Delete("/deleted", func(ctx iris.Context) {
			ctx.WriteString("deleted")
		})
		jwtHandler := jwtmiddleware.New(jwtmiddleware.Config{
			ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
				return []byte("My Secret"), nil
			},
			SigningMethod: jwt.SigningMethodHS256,
		})

		v1.Use(jwtHandler.Serve)
		v1.Get("/ping", myHandler)
	}
	app.Run(iris.Addr(":8080"))
	fmt.Println("server started successfully")
}
