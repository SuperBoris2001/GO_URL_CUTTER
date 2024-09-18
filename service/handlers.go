package service

import (
	"fmt"
	"io"
	"net/http"
)

// Обработка главной страницы (корня)
func (app *Application) home(w http.ResponseWriter, r *http.Request) {
	fmt.Println("")
	app.InfoLog.Output(2, "Попытка запроса с методом: "+r.Method)
	//Отказ обработки, при неверном запросе.
	if (r.Method != "POST") && (r.Method != "GET") {
		app.InfoLog.Output(2, "Запрос с неверным методом.")
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}
	//Тело ответа
	responseBodyStr := ""
	var err error
	//////////////////////////////////////////////////////////
	///                        POST                        ///
	//////////////////////////////////////////////////////////
	if r.Method == "POST" {
		var body []byte
		body, err = io.ReadAll(r.Body)
		fmt.Println("Тело запроса: " + string(body))
		if err != nil {
			app.InfoLog.Output(2, "Ошибка чтения тела запроса")
			http.Error(w, "Ошибка чтения тела запроса", http.StatusInternalServerError)
			return
		}
		responseBodyStr, err = app.CreateShortUrl(string(body))
	} else {
		//////////////////////////////////////////////////////////
		///                        GET                        ///
		//////////////////////////////////////////////////////////
		fmt.Println("URL с которым произошел запрос GET: ", r.URL.String()[1:])
		responseBodyStr, err = app.GetLongUrl(r.URL.String()[1:])
	}
	//////////////////////////////////////////////////////////
	///                        BOTH                        ///
	//////////////////////////////////////////////////////////
	if err != nil {
		if err.Error() == "the URL does not exist" {
			app.InfoLog.Output(2, "Запрос на несуществующий короткий url")
			http.Error(w, "Request for a non-existent short urlr", http.StatusInternalServerError)
		} else {
			app.InfoLog.Output(2, "Произошла ошибка на сервере")
			http.Error(w, "An error occurred on the server", http.StatusInternalServerError)
		}
		return
	}
	_, err = w.Write([]byte(responseBodyStr))
	if err != nil {
		app.InfoLog.Output(2, "Ошибка записи тела ответа")
		http.Error(w, "Ошибка записи тела ответа", http.StatusInternalServerError)
		return
	}
	app.InfoLog.Output(2, "Совершен успешный запрос: "+r.Method)
}
