package storage

import (
	"database/sql"
	"errors"
	"flag"
	"fmt"

	_ "github.com/lib/pq"
)

// Для расширения
// 1) Изменить DataStorage
// 2) Написать соответствующий код по интерфесам
// 3)Изменить NewDataStorage

// Структура, в которой методом встраивания можно заменить реализацию
type DataStorage struct {
	//Хеш-таблица для хранения в оперативной памяти.
	UrlMap map[string]string
	//Проверка на наличия ключа в хранилище
	ChekShort func(string, *DataStorage) error
	//Установка новой пары ключ-значение
	SetShort func(string, string, *DataStorage)
	//Взятие значения по ключу
	GetLong func(string, *DataStorage) (string, error)
	PostgresStore
}

// Создание экземпляра хранилища с определением типа хранения в зависимости от флага
func NewDataStorage() (*DataStorage, error) {
	//Парсинг флагов
	useDatabase := flag.Bool("d", false, "Use database storage")
	flag.Parse()
	//////////////////////////////////////////////////////////
	///                     POSTGRES                       ///
	//////////////////////////////////////////////////////////
	if *useDatabase {
		ds, err := setStoragePostgres()
		return ds, err
	}
	//////////////////////////////////////////////////////////
	///                     ПАМЯТЬ                       ///
	//////////////////////////////////////////////////////////
	ds, err := setStorageMap()
	return ds, err
}
func (ds *DataStorage) CreateShortUrl(longUrl string) (string, error) {
	shortUrl := ds.GenerateShortURL()
	ds.SetShort(shortUrl, longUrl, ds)
	return "http://localhost:8080/" + shortUrl, nil
}
func (ds *DataStorage) GetLongUrl(shortUrl string) (string, error) {
	longUrl, err := ds.GetLong(shortUrl, ds)
	if err != nil {
		return "", errors.New("the URL does not exist")
	}
	return longUrl, nil
}

func (ds *DataStorage) GenerateShortURL() string {
	const alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	length := 2 // Начинаем с генерации URL длиной 2 символа

	for {
		// Перебор всех возможных комбинаций
		for i := 0; i < len(alphabet); i++ {
			for j := 0; j < len(alphabet); j++ {
				shortURL := string(alphabet[i]) + string(alphabet[j])
				err := ds.ChekShort(shortURL, ds)
				if err != nil {
					return shortURL
				}
			}
		}

		// Увеличиваем длину URL для следующей попытки
		length++
	}
}

// ///////////////////////////////////////////////////////////////////////////////////////////////
// /                                          MAP                                              ///
// ///////////////////////////////////////////////////////////////////////////////////////////////
func setStorageMap() (*DataStorage, error) {
	ds := DataStorage{}
	ds.UrlMap = make(map[string]string)
	ds.ChekShort = chekShorlUrldMap
	ds.SetShort = setShorlUrldMap
	ds.GetLong = getLongUrlMap
	return &ds, nil
}
func chekShorlUrldMap(shortUrl string, ds *DataStorage) error {
	//Если url есть в БД, то возвращаем nil
	_, ok := ds.UrlMap[shortUrl]
	if ok {
		return nil
	} else {
		return errors.New("the URL does not exist")
	}
}
func setShorlUrldMap(shortUrl string, longUrl string, ds *DataStorage) {
	chekLongUrlMap(longUrl, ds)
	ds.UrlMap[shortUrl] = longUrl
}
func getLongUrlMap(shortUrl string, ds *DataStorage) (string, error) {
	longUrl, ok := ds.UrlMap[shortUrl]
	if ok {
		return longUrl, nil
	} else {
		return longUrl, errors.New("the URL does not exist")
	}
}
func chekLongUrlMap(LongUrl string, ds *DataStorage) bool {
	for k, v := range ds.UrlMap {
		if v == LongUrl {
			delete(ds.UrlMap, k)
			return true
		}
	}
	return false
}

// ///////////////////////////////////////////////////////////////////////////////////////////////
// /                                       POSTGRESS                                           ///
// ///////////////////////////////////////////////////////////////////////////////////////////////
// Структура для работы с postgres Sql
type PostgresStore struct {
	//Адресс
	Host string
	Port int
	//Имя пользователя/роли
	User     string
	Password string
	//postgres
	Dbname string
	//Ссылка на БД. Удобный интерфейс для работы с ним
	Db *sql.DB
}

const (
	Host     = "localhost"
	Port     = 5432
	User     = "microservice"
	Password = "123456"
	Dbname   = "postgres"
)

func setStoragePostgres() (*DataStorage, error) {
	fmt.Println("Программа запущена с флагом -d")
	var err error
	ds := DataStorage{}
	ds.Db, err = sql.Open("postgres", fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", Host, Port, User, Password, Dbname))
	if err != nil {
		return &ds, err
	}
	ds.ChekShort = checkShortURLPostgres
	ds.SetShort = setShorlUrldPostgres
	ds.GetLong = getLongUrlPostGres
	return &ds, nil
}
func checkShortURLPostgres(shortURL string, ds *DataStorage) error {
	var count int
	row := ds.Db.QueryRow("SELECT COUNT(*) FROM url_store WHERE key = $1", shortURL)
	err := row.Scan(&count)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	//Если url есть в БД, то возвращаем nil
	if count > 0 {
		fmt.Println("Ссылка уже в БД под ключом ", shortURL)
		return nil
	}
	return errors.New("")
}
func setShorlUrldPostgres(shortUrl string, longUrl string, ds *DataStorage) {
	chekLongUrlPostgres(longUrl, ds)
	insertStatement := `
	INSERT INTO url_store (key, value)
	VALUES ($1, $2)
`
	_, err := ds.Db.Exec(insertStatement, shortUrl, longUrl)
	if err != nil {
		fmt.Println(err.Error())
	}
}
func getLongUrlPostGres(shortUrl string, ds *DataStorage) (string, error) {
	var value string
	row := ds.Db.QueryRow("SELECT value FROM url_store WHERE key = $1", shortUrl)
	err := row.Scan(&value)
	if err != nil {
		///////////////
		fmt.Println(err.Error())
		return "", err
	}
	return value, nil
}
func chekLongUrlPostgres(LongUrl string, ds *DataStorage) bool {
	_, err := ds.Db.Exec("DELETE FROM url_store WHERE value = $1", LongUrl)
	return nil != err
}
