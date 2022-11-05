package main

import (
	"BE_WEB_BEM_Proker/infrastructure/database_connection"
	"BE_WEB_BEM_Proker/infrastructure/database_driver"
	"BE_WEB_BEM_Proker/route"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	e := gin.Default()
	env, err := database_driver.ReadEnvSupabase()
	if err != nil {
		panic(err)
	}
	db, err := database_connection.MakeConnection(env)
	if err != nil {
		panic(err)
	}
	e.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	e.Use(cors.Default())
	route.InitRouteAll(e, db)
	if err := e.Run(":8070"); err != nil {
		panic(err)
	}
	// client := storage_go.NewClient("https://jgjyjvyldoamqndazixl.supabase.co/storage/v1", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJzdXBhYmFzZSIsInJlZiI6ImpnanlqdnlsZG9hbXFuZGF6aXhsIiwicm9sZSI6ImFub24iLCJpYXQiOjE2NDc4MzQ0MDQsImV4cCI6MTk2MzQxMDQwNH0.WVMjJIRoK_cnyfRXdYvTokNWBCCqLWfbeu7xXeZrs6I", nil)
	// // Get buckets
	// fmt.Println(client.ListBuckets())
	// file, err := os.OpenFile("./1.jpg", os.O_RDWR, 0666)
	// if err != nil {
	// 	panic(err)
	// }
	// resp := client.UploadFile("foto-proker", "data4", file)
	// //imageData, imageType, err := image.Decode(file)
	// fmt.Println(resp)
	// fmt.Println(client.CreateSignedUrl("foto-proker", "data3", int(20*time.Hour.Hours())))
	// data := fmt.Sprintf("https://jgjyjvyldoamqndazixl.supabase.co/storage/v1/object/public/foto-proker/%s", "data4")
	// fmt.Println(data)
}
