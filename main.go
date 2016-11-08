package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const (
	NM = iota
	CO
	WY
	MT
	UT
	NV
	ID
	WA
	OR
	CA
	AZ
	AK
	HI
	TX
	LA
	MS
	AR
	OK
	MO
	KS
	NE
	SD
	ND
	MN
	WI
	MI
	IA
	IL
	IN
	KY
	TN
	AL
	FL
	GA
	SC
	NC
	OH
	WV
	DC
	VA
	PA
	MD
	DE
	NJ
	NY
	MA
	CT
	RH
	NH
	VT
	ME
)

func setColor(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello World!")
}

func sendColor(color string) {
	url := "https://api.particle.io/v1/devices/4f0042000b51353335323535/color"

	payload := strings.NewReader("access_token=df21db852a652855a4dcd72f55951879b3a39c09&color=" + color)

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "7b35fd10-c97f-5417-07e3-7a512ce078bf")
	req.Header.Add("content-type", "application/x-www-form-urlencoded")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}

func getColors() {
	url := "https://api.particle.io/v1/devices/4f0042000b51353335323535/colorString"

	payload := strings.NewReader("access_token=df21db852a652855a4dcd72f55951879b3a39c09&color=RGB")

	req, _ := http.NewRequest("GET", url, payload)

	req.Header.Add("authorization", "Bearer df21db852a652855a4dcd72f55951879b3a39c09")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "bbfdc3b2-56b5-00de-1227-f8f48f4ed40f")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(body)
}

type JSON struct {
	id           string
	last_app     string
	connected    bool
	return_value int
}

type Map struct {
	ColorString []byte
}

func (*Map) SetStateColor(state string, color string) {
	Map.ColorString[3] = 'f'
}

func main() {

	fs := http.FileServer(http.Dir("example"))
	http.Handle("/", fs)
	http.HandleFunc("/color", getColors)

	log.Println("Listening.....")
	http.ListenAndServe(":3000", nil)
}
