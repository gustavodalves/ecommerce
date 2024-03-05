package main

import (
	"database/sql"
	"net/http"
	"os"

	"github.com/go-chi/chi"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gustavodalves/ecommerce/internal/app/service"
	repository_mysql "github.com/gustavodalves/ecommerce/internal/infra/repository/mysql"
	web "github.com/gustavodalves/ecommerce/internal/infra/web/handlers"
)

func main() {
	db, err := sql.Open("mysql", "root:admin@tcp(localhost:3308)/ecommerce")

	if err != nil {
		panic(err)
	}

	defer db.Close()

	repository := repository_mysql.NewOrderRepository(db)
	orderService := service.OrderService{Repository: repository}

	orderHandler := web.NewHandler(orderService)

	r := chi.NewRouter()

	r.Post("/order", orderHandler.CreateOrder)

	stopServer := make(chan os.Signal, 1)
	go func() {
		if err := http.ListenAndServe(":3000", r); err != nil {
			panic(err)
		}
	}()
	<-stopServer
}
