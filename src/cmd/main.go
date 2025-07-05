// @title           Product API
// @version         1.0
// @description     Online Shop REST API
// @termsOfService  http://swagger.io/terms/
// @contact.name    Support
// @contact.email   example@example.com
// @license.name    MIT
// @host            202.133.88.175:8090
// @BasePath        /

package main

import (
	httpserver "RestGoTest/src"
	"RestGoTest/src/config"
	"RestGoTest/src/database"
	"RestGoTest/src/repository"
	"fmt"
	"time"
)

func main() {

	cfg := config.GetConfig()

	fmt.Println("═══════════════════════════════════════════════")
	fmt.Println("🧩      API Server      ")
	fmt.Println("🚀   OnlineShop REST API   ")
	fmt.Println("═══════════════════════════════════════════════")

	InternalPort := fmt.Sprintf(":%s", cfg.Server.InternalPort)

	fmt.Printf("✅ Server successfully started on port %s\n", InternalPort)
	fmt.Println("🟢 Running... Press Ctrl+C to stop")
	fmt.Printf("📅 Startup time: %s\n\n", time.Now().Format("2006-01-02 15:04:05"))

	/*Start Point of Service*/
	/*═══════════════════════════════════════════════*/
	database.InitDatabase()
	defer repository.DB.Close()
	a := &httpserver.App{Port: InternalPort}
	a.Init()
	a.Run()
	/*═══════════════════════════════════════════════*/

}
