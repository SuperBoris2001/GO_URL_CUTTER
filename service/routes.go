package service

import (
	"URL_CUTTER/storage"
	"log"
	"net/http"
)

// Структура приложения
type Application struct {
	//Логер для ошибок
	ErrorLog *log.Logger
	//Логер для информационных сообщений
	InfoLog *log.Logger
	//Структура для работы с хранилищем
	*storage.DataStorage
}

// Cоздает и возвращает маршрутизатор (ServeMux), который используется для определения обработчиков запросов для различных путей
func (app *Application) Routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.home)
	return mux
}
