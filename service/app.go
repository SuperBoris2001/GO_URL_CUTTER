package service

import (
	"URL_CUTTER/storage"
	"flag"
	"log"
	"net/http"
	"os"
)

func CreateWebApp(ip, port string, storage *storage.DataStorage) {
	infoLog := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	errorLog := log.New(os.Stderr, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	addr := flag.String("addr", (ip + ":" + port), "Сетевой адрес веб-сервера")
	flag.Parse()
	app := &Application{
		ErrorLog: errorLog,
		InfoLog:  infoLog,
	}
	app.DataStorage = storage
	app.Routes()
	srv := &http.Server{
		Addr:     *addr,
		ErrorLog: errorLog,
		Handler:  app.Routes(),
	}

	infoLog.Printf("Запуск сервера на %s", *addr)
	err := srv.ListenAndServe()
	errorLog.Fatal(err)
}
