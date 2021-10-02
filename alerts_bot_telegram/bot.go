// программа подключается к servicenow, забирает инциденты и шлет их в телеграм бота

package main

import (
	"encoding/json"
	"fmt"
	"gopkg.in/telegram-bot-api.v4"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)
const (
	pass = "E:\\golang\\work\\src\\alerts_bot_telegram\\pass"
	incidentsfile = "E:\\golang\\work\\src\\alerts_bot_telegram\\incidents.txt"
	configJSON = "E:\\golang\\work\\src\\alerts_bot_telegram\\config.json"
)

type Config struct {
	TelegramBotToken string
}
var parsed map[string]interface{}
var oldIncList []string

func main() {
	file, _ := os.Open(configJSON)
	// fmt.Println(*file)
	decoder := json.NewDecoder(file)
	// fmt.Println(decoder)
	configuration := Config{}
	err := decoder.Decode(&configuration)
	if err != nil {
		log.Panic(err)
	}
	// fmt.Println(configuration.TelegramBotToken)

	bot, err := tgbotapi.NewBotAPI(configuration.TelegramBotToken)
	fmt.Println(bot)
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	// запуск таймера 
	ticker := time.NewTicker(time.Second * 300)
    for t := range ticker.C {
			fmt.Println("Tick at", t)
			NewIncList := getNewIncsList()
			fmt.Println(NewIncList)
			if NewIncList != nil {
			stringIncidents := strings.Join(NewIncList, ",")
			msg := tgbotapi.NewMessage(200128561, stringIncidents)
			bot.Send(msg)
			}
		}
}

func getNewIncsList() []string {
	oldIncList = nil
	fmt.Println(oldIncList)
	file, _ := ioutil.ReadFile(pass)
	pass := string(file)
	url := "https://login:" + pass + "@xxx.xxx.com/incident_list.do"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch:  %v\n", err)
	}
	data, err := ioutil.ReadAll(resp.Body)
	resp.Body.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "fetch:  чтение %s:  %v\n", url, err)
		os.Exit(1)
	}

	// перекодируем из json 
	err = json.Unmarshal(data, &parsed) 
	if err != nil {
		fmt.Println("error:", err)
	}
	jsonList := parsed["records"].([]interface{}) //убираем строку records
	incList := make([]string, 0)
	for i:=0; i < len(jsonList); i++{
		inc := jsonList[i].(map[string]interface{})
		incNumber := inc["number"].(string)
		incList = append(incList, incNumber)
	}
	ioutil.WriteFile(incidentsfile, data, 0644)

	return incList
}
