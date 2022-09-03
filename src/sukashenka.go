package x90

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"
	"sync"
	"time"

	"github.com/bwmarrin/discordgo"
)

type Counter struct {
	Avg float64 `json:"avg"`
	Max float64 `json:"max"`
	Min float64 `json:"min"`
}

var lookup = map[string]string{
	"kaunas":      "https://remap.jrc.ec.europa.eu/api/maps/polygon?id=1234_591&zoom=8&_=",
	"vilnius":     "https://remap.jrc.ec.europa.eu/api/maps/polygon?id=1242_590&zoom=8&_=",
	"dieveniskes": "https://remap.jrc.ec.europa.eu/api/maps/polygon?id=1244_585&zoom=8&_=",
	"bulvydziai":  "https://remap.jrc.ec.europa.eu/api/maps/polygon?id=1245_590&zoom=8&_=",
	"osmany (BY)": "https://remap.jrc.ec.europa.eu/api/maps/polygon?id=1246_587&zoom=8&_=",
}

func checkCity(city string) Counter {
	res, err := http.Get(city)

	if err != nil {
		panic(err.Error())
	}

	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err.Error())
	}

	var rez Counter
	err = json.Unmarshal(body, &rez)
	if err != nil {
		panic(err)
	}

	return rez
}

func displayRez(rez Counter, city string, session *discordgo.Session, message *discordgo.MessageCreate) {
	var avg string = strconv.FormatFloat(rez.Avg, 'f', 2, 64)
	//avg, _ := strconv.ParseFloat(rez.Avg, 64)
	msg := city + ": " + avg + "nSv/h"
	sukashenka(session, message, msg)
}

func getSukaData(session *discordgo.Session, message *discordgo.MessageCreate) {
	var wg sync.WaitGroup

	wg.Add(len(lookup))

	for city, url := range lookup {
		go func(city string, url string) {
			defer wg.Done()
			displayRez(checkCity(url+strconv.FormatInt(time.Now().Unix(), 10)), city, session, message)
		}(city, url)
	}

	wg.Wait()
}
