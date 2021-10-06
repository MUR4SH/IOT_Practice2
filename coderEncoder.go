package main

import (
	"encoding/json"
	"encoding/xml"
	"log"
)

// Структура для кодирования/декодирования данных
type Format struct {
	Id        uint
	Timestamp string
	Hash      string
}

// Кодирует трое данных в json (возвращает массив byte)
func JSON_encode(r_id uint, r_time string, r_hash string) []byte {
	g_json, err := json.Marshal(Format{r_id, r_time, r_hash}) // Создаем объект и делаем из него json
	if err != nil {
		log.Fatal(err)
	}
	return g_json
}

func XML_encode(r_id uint, r_time string, r_hash string) []byte {
	g_xml, err := xml.Marshal(Format{r_id, r_time, r_hash})
	if err != nil {
		log.Fatal(err)
	}
	return g_xml
}

/*func HTML_encode() []byte {
	return
}*/

// Декодирует json в трое данных - id, timestamp, hash (их и возвращает - int, string, string)
func JSON_decode(r_json []byte) (uint, string, string) {
	var m Format //Создаем переменную, в которую будет записан расшифрованный json
	err := json.Unmarshal(r_json, &m)
	if err != nil {
		log.Fatal(err)
	}
	return m.Id, m.Timestamp, m.Hash
}

func XML_decode(r_xml []byte) (uint, string, string) {
	var m Format
	err := xml.Unmarshal(r_xml, &m)
	if err != nil {
		log.Fatal(err)
	}
	return m.Id, m.Timestamp, m.Hash
}

/*func HTML_decode() (uint, string, string) {

}*/
