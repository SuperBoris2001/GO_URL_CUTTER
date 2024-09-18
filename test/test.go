package test

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"time"
)

func sendPOST() {
	url := "http://localhost:8080/"

	// Тело запроса
	body := []byte("http://cjdr17afeihmk.biz/123/kdni9/z9d112423421")

	// Отправка POST запроса
	resp, err := http.Post(url, "application/octet-stream", bytes.NewBuffer(body))
	if err != nil {
		//fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Получение тела ответа
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("Ошибка при чтении тела ответа:", err)
		return
	}

	// Вывод ответа
	fmt.Println("Ответ сервера:", string(respBody))
}
func sendGet() {
	url := "http://localhost:8080/aa"

	// Отправка GET запроса
	resp, err := http.Get(url)
	if err != nil {
		//fmt.Println("Ошибка при отправке запроса:", err)
		return
	}
	defer resp.Body.Close()

	// Получение тела ответа
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		//fmt.Println("Ошибка при чтении тела ответа:", err)
		return
	}

	// Вывод ответа
	fmt.Println("Ответ сервера:", string(respBody))
}
func StartTesting() {
	time.Sleep(5 * time.Second)
	sendPOST()
	time.Sleep(5 * time.Second)
	sendGet()
	time.Sleep(5 * time.Second)
	sendGet()
	time.Sleep(5 * time.Second)
	sendPOST()
}
