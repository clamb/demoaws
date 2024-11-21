package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"

	"github.com/gofiber/fiber/v3"
)

func main() {
	var dbserver string
	var port string

	flag.StringVar(&dbserver, "dbserver", "", "the hostname of the server (env DB_SERVER)")
	flag.StringVar(&port, "port", "", "the port to listen to (SRV_PORT)")

	flag.Parse()

	if dbserver == "" {
		dbserver = os.Getenv("DB_SERVER")
		if dbserver == "" {
			flag.Usage()
			os.Exit(0)
		}
	}

	if port == "" {
		port = os.Getenv("SRV_PORT")
		if port == "" {
			flag.Usage()
			os.Exit(0)
		}
	}

	app := fiber.New()

	app.Get("/", func(c fiber.Ctx) error {
		return c.SendString("Hello SOI")
	})

	app.Get("/db", func(c fiber.Ctx) error {
		db, err := sql.Open("mysql", "root:password1234@tcp("+dbserver+")/hello")
		if err != nil {
			fmt.Printf("error: %v", err)
			return c.SendString(fmt.Sprintf("%v", err))
		}

		err = db.Ping()
		if err != nil {
			return c.SendString(fmt.Sprintf("cannot connect to db: %v", err))
		}
		return c.SendString("ok")
	})

	log.Printf("listening on port %s, using db `hello` on host %s", port, dbserver)
	log.Fatal(app.Listen(fmt.Sprintf(":%s", port)))
}
