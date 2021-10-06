package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"log"
	"strconv"
	"strings"
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

/*
Пускай будет формат
<div id='html_data'>
	<span id='id'><span>
	<span id='timestamp'><span>
	<span id='hash'><span>
</div>
*/
func HTML_encode(r_id uint, r_time string, r_hash string) string {
	var html string = "<div id='html_data'>"
	html += "<span id='id'>" + fmt.Sprint(r_id) + "</span>"
	html += "<span id='timestamp'>" + fmt.Sprint(r_time) + "</span>"
	html += "<span id='hash'>" + fmt.Sprint(r_hash) + "</span></div>"
	return html
}

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

func HTML_decode(html string) (int, string, string) {
	var id int
	var time, hash string
	html = strings.ReplaceAll(html, "<div id='html_data'><span id='id'>", "")

	id, _ = strconv.Atoi(strings.TrimRight(html, "</span>"))
	log.Print(html)
	html = strings.TrimRight(html, "</span>")
	log.Print(html)
	html = strings.ReplaceAll(html, "<span id='timestamp'>", "")
	log.Print(html)

	time = strings.TrimRight(html, "</span>")

	html = strings.TrimLeft(html, "</span>")
	html = strings.ReplaceAll(html, "<span id='hash'>", "")
	log.Print(html)

	hash = strings.TrimRight(html, "</span>")

	return id, time, hash
}
