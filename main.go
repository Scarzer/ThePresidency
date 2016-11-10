package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/gorilla/mux"
)

const ()

type Map struct {
	states []string
	colors []int
}

func MakeNewStateMap() Map {

	foo := []string{
		"NM",
		"CO",
		"WY",
		"MT",
		"UT",
		"NV",
		"ID",
		"WA",
		"OR",
		"CA",
		"AZ",
		"AK",
		"HI",
		"TX",
		"LA",
		"MS",
		"AR",
		"OK",
		"MO",
		"KS",
		"NE",
		"SD",
		"ND",
		"MN",
		"WI",
		"MI",
		"IA",
		"IL",
		"IN",
		"KY",
		"TN",
		"AL",
		"FL",
		"GA",
		"SC",
		"NC",
		"OH",
		"WV",
		"DC",
		"VA",
		"PA",
		"MD",
		"DE",
		"NJ",
		"NY",
		"MA",
		"CT",
		"RH",
		"NH",
		"VT",
		"ME",
	}

	var colors []int
	for v := range foo {
		println(v)
		colors = append(colors, 0)
	}
	fmt.Println(len(colors))
	fmt.Println(len(foo))
	return Map{states: foo, colors: colors}
}

func (m *Map) GenerateColorString() string {
	cs := ""
	for ind := range m.states {
		switch m.colors[ind] {
		case 0:
			cs = cs + " "
			break
		case 1:
			cs = cs + "R"
			break
		case 2:
			cs = cs + "B"
			break
		case 3:
			cs = cs + "G"
			break
		}
	}
	return cs
}
func (m *Map) GetColorOfState(state string) int {
	c := 0
	for ind, st := range m.states {
		if st == state {
			c = m.colors[ind]
		}
	}
	return c
}

func (m *Map) ChangeColorOfState(state string, color int) int {
	c := 0
	for ind, st := range m.states {
		if st == state {
			m.colors[ind] = color
			c = m.colors[ind]
		}
	}
	return c
}

var Colors Map

func getColorFromState(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r)
	state := vars["state"]

	c := Colors.GetColorOfState(state)
	bs := []byte(strconv.Itoa(c))
	w.Write(bs)

	w.Write([]byte("Foo"))
}
func postColorFromState(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	state := vars["state"]
	party := vars["party"]
	log.Println(state)
	number, _ := strconv.Atoi(party)
	log.Println(party)
	log.Println(number)
	Colors.ChangeColorOfState(state, number)

	sendColor(Colors.GenerateColorString())
}

func clearMap(w http.ResponseWriter, r *http.Request) {

}

func main() {
	Colors = MakeNewStateMap()
	r := mux.NewRouter()
	r.HandleFunc("/clearMap", clearMap).Methods("GET")
	r.HandleFunc("/color/{state}", getColorFromState).Methods("GET")
	r.HandleFunc("/color/{state}/{party}", postColorFromState).Methods("POST")
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("example")))

	http.Handle("/", r)

	log.Println("Listening on port 3000")

	go func() {
		for {
			timer := time.NewTimer(time.Second * 2)
			<-timer.C

		}
	}()

	http.ListenAndServe(":8000", nil)

}

type NYT struct {
	President struct {
		Current struct {
			Timestamp             time.Time `json:"timestamp"`
			ElectoralVotesCounted struct {
				Clintonh  int `json:"clintonh"`
				Johnsong  int `json:"johnsong"`
				Mcmulline int `json:"mcmulline"`
				Other     int `json:"other"`
				Steinj    int `json:"steinj"`
				Trumpd    int `json:"trumpd"`
			} `json:"electoral_votes_counted"`
			ElectoralVotesDistribution struct {
				Clintonh struct {
					Num108 float64 `json:"108"`
					Num128 float64 `json:"128"`
					Num129 float64 `json:"129"`
					Num131 float64 `json:"131"`
					Num132 float64 `json:"132"`
					Num135 float64 `json:"135"`
					Num136 float64 `json:"136"`
					Num138 float64 `json:"138"`
					Num139 float64 `json:"139"`
					Num140 float64 `json:"140"`
					Num142 float64 `json:"142"`
					Num145 float64 `json:"145"`
					Num148 float64 `json:"148"`
					Num152 float64 `json:"152"`
					Num153 float64 `json:"153"`
					Num154 float64 `json:"154"`
					Num156 float64 `json:"156"`
					Num158 float64 `json:"158"`
					Num160 float64 `json:"160"`
					Num161 float64 `json:"161"`
					Num162 float64 `json:"162"`
					Num163 float64 `json:"163"`
					Num164 float64 `json:"164"`
					Num165 float64 `json:"165"`
					Num166 float64 `json:"166"`
					Num167 float64 `json:"167"`
					Num168 float64 `json:"168"`
					Num169 float64 `json:"169"`
					Num171 float64 `json:"171"`
					Num172 float64 `json:"172"`
					Num173 float64 `json:"173"`
					Num174 float64 `json:"174"`
					Num175 float64 `json:"175"`
					Num176 float64 `json:"176"`
					Num177 float64 `json:"177"`
					Num178 float64 `json:"178"`
					Num179 float64 `json:"179"`
					Num180 float64 `json:"180"`
					Num181 float64 `json:"181"`
					Num182 float64 `json:"182"`
					Num183 float64 `json:"183"`
					Num184 float64 `json:"184"`
					Num185 float64 `json:"185"`
					Num186 float64 `json:"186"`
					Num187 float64 `json:"187"`
					Num188 float64 `json:"188"`
					Num189 float64 `json:"189"`
					Num190 float64 `json:"190"`
					Num191 float64 `json:"191"`
					Num192 float64 `json:"192"`
					Num193 float64 `json:"193"`
					Num194 float64 `json:"194"`
					Num195 float64 `json:"195"`
					Num196 float64 `json:"196"`
					Num197 float64 `json:"197"`
					Num198 float64 `json:"198"`
					Num199 float64 `json:"199"`
					Num200 float64 `json:"200"`
					Num201 float64 `json:"201"`
					Num202 float64 `json:"202"`
					Num203 float64 `json:"203"`
					Num204 float64 `json:"204"`
					Num205 float64 `json:"205"`
					Num206 float64 `json:"206"`
					Num207 float64 `json:"207"`
					Num208 float64 `json:"208"`
					Num209 float64 `json:"209"`
					Num210 float64 `json:"210"`
					Num211 float64 `json:"211"`
					Num212 float64 `json:"212"`
					Num213 float64 `json:"213"`
					Num214 float64 `json:"214"`
					Num215 float64 `json:"215"`
					Num216 float64 `json:"216"`
					Num217 float64 `json:"217"`
					Num218 float64 `json:"218"`
					Num219 float64 `json:"219"`
					Num220 float64 `json:"220"`
					Num221 float64 `json:"221"`
					Num222 float64 `json:"222"`
					Num223 float64 `json:"223"`
					Num224 float64 `json:"224"`
					Num225 float64 `json:"225"`
					Num226 float64 `json:"226"`
					Num227 float64 `json:"227"`
					Num228 float64 `json:"228"`
					Num229 float64 `json:"229"`
					Num230 float64 `json:"230"`
					Num231 float64 `json:"231"`
					Num232 float64 `json:"232"`
					Num233 float64 `json:"233"`
					Num234 float64 `json:"234"`
					Num235 float64 `json:"235"`
					Num236 float64 `json:"236"`
					Num237 float64 `json:"237"`
					Num238 float64 `json:"238"`
					Num239 float64 `json:"239"`
					Num240 float64 `json:"240"`
					Num241 float64 `json:"241"`
					Num242 float64 `json:"242"`
					Num243 float64 `json:"243"`
					Num244 float64 `json:"244"`
					Num245 float64 `json:"245"`
					Num246 float64 `json:"246"`
					Num247 float64 `json:"247"`
					Num248 float64 `json:"248"`
					Num249 float64 `json:"249"`
					Num250 float64 `json:"250"`
					Num251 float64 `json:"251"`
					Num252 float64 `json:"252"`
					Num253 float64 `json:"253"`
					Num254 float64 `json:"254"`
					Num255 float64 `json:"255"`
					Num256 float64 `json:"256"`
					Num257 float64 `json:"257"`
					Num258 float64 `json:"258"`
					Num259 float64 `json:"259"`
					Num260 float64 `json:"260"`
					Num261 float64 `json:"261"`
					Num262 float64 `json:"262"`
					Num263 float64 `json:"263"`
					Num264 float64 `json:"264"`
					Num265 float64 `json:"265"`
					Num266 float64 `json:"266"`
					Num267 float64 `json:"267"`
					Num268 float64 `json:"268"`
					Num269 float64 `json:"269"`
					Num270 float64 `json:"270"`
					Num271 float64 `json:"271"`
					Num272 float64 `json:"272"`
					Num273 float64 `json:"273"`
					Num274 float64 `json:"274"`
					Num275 float64 `json:"275"`
					Num276 float64 `json:"276"`
					Num277 float64 `json:"277"`
					Num278 float64 `json:"278"`
					Num279 float64 `json:"279"`
					Num280 float64 `json:"280"`
					Num281 float64 `json:"281"`
					Num282 float64 `json:"282"`
					Num283 float64 `json:"283"`
					Num284 float64 `json:"284"`
					Num285 float64 `json:"285"`
					Num286 float64 `json:"286"`
					Num287 float64 `json:"287"`
					Num288 float64 `json:"288"`
					Num289 float64 `json:"289"`
					Num290 float64 `json:"290"`
					Num291 float64 `json:"291"`
					Num292 float64 `json:"292"`
					Num293 float64 `json:"293"`
					Num294 float64 `json:"294"`
					Num295 float64 `json:"295"`
					Num296 float64 `json:"296"`
					Num297 float64 `json:"297"`
					Num298 float64 `json:"298"`
					Num299 float64 `json:"299"`
					Num300 float64 `json:"300"`
					Num301 float64 `json:"301"`
					Num302 float64 `json:"302"`
					Num303 float64 `json:"303"`
					Num304 float64 `json:"304"`
					Num305 float64 `json:"305"`
					Num306 float64 `json:"306"`
					Num307 float64 `json:"307"`
					Num308 float64 `json:"308"`
					Num309 float64 `json:"309"`
					Num310 float64 `json:"310"`
					Num311 float64 `json:"311"`
					Num312 float64 `json:"312"`
					Num313 float64 `json:"313"`
					Num314 float64 `json:"314"`
					Num315 float64 `json:"315"`
					Num316 float64 `json:"316"`
					Num317 float64 `json:"317"`
					Num318 float64 `json:"318"`
					Num319 float64 `json:"319"`
					Num320 float64 `json:"320"`
					Num321 float64 `json:"321"`
					Num322 float64 `json:"322"`
					Num323 float64 `json:"323"`
					Num324 float64 `json:"324"`
					Num325 float64 `json:"325"`
					Num326 float64 `json:"326"`
					Num327 float64 `json:"327"`
					Num328 float64 `json:"328"`
					Num329 float64 `json:"329"`
					Num330 float64 `json:"330"`
					Num331 float64 `json:"331"`
					Num332 float64 `json:"332"`
					Num333 float64 `json:"333"`
					Num334 float64 `json:"334"`
					Num335 float64 `json:"335"`
					Num336 float64 `json:"336"`
					Num337 float64 `json:"337"`
					Num338 float64 `json:"338"`
					Num339 float64 `json:"339"`
					Num340 float64 `json:"340"`
					Num341 float64 `json:"341"`
					Num342 float64 `json:"342"`
					Num343 float64 `json:"343"`
					Num344 float64 `json:"344"`
					Num345 float64 `json:"345"`
					Num346 float64 `json:"346"`
					Num347 float64 `json:"347"`
					Num348 float64 `json:"348"`
					Num349 float64 `json:"349"`
					Num350 float64 `json:"350"`
					Num351 float64 `json:"351"`
					Num352 float64 `json:"352"`
					Num353 float64 `json:"353"`
					Num354 float64 `json:"354"`
					Num355 float64 `json:"355"`
					Num356 float64 `json:"356"`
					Num357 float64 `json:"357"`
					Num358 float64 `json:"358"`
					Num359 float64 `json:"359"`
					Num360 float64 `json:"360"`
					Num361 float64 `json:"361"`
					Num362 float64 `json:"362"`
					Num363 float64 `json:"363"`
					Num364 float64 `json:"364"`
					Num365 float64 `json:"365"`
					Num366 float64 `json:"366"`
					Num367 float64 `json:"367"`
					Num368 float64 `json:"368"`
					Num369 float64 `json:"369"`
					Num370 float64 `json:"370"`
					Num371 float64 `json:"371"`
					Num372 float64 `json:"372"`
					Num373 float64 `json:"373"`
					Num374 float64 `json:"374"`
					Num375 float64 `json:"375"`
					Num376 float64 `json:"376"`
					Num377 float64 `json:"377"`
					Num378 float64 `json:"378"`
					Num379 float64 `json:"379"`
					Num380 float64 `json:"380"`
					Num381 float64 `json:"381"`
					Num382 float64 `json:"382"`
					Num383 float64 `json:"383"`
					Num384 float64 `json:"384"`
					Num385 float64 `json:"385"`
					Num386 float64 `json:"386"`
					Num387 float64 `json:"387"`
					Num388 float64 `json:"388"`
					Num389 float64 `json:"389"`
					Num390 float64 `json:"390"`
					Num391 float64 `json:"391"`
					Num392 float64 `json:"392"`
					Num393 float64 `json:"393"`
					Num394 float64 `json:"394"`
					Num395 float64 `json:"395"`
					Num396 float64 `json:"396"`
					Num397 float64 `json:"397"`
					Num398 float64 `json:"398"`
					Num399 float64 `json:"399"`
					Num400 float64 `json:"400"`
					Num401 float64 `json:"401"`
					Num402 float64 `json:"402"`
					Num403 float64 `json:"403"`
					Num404 float64 `json:"404"`
					Num405 float64 `json:"405"`
					Num406 float64 `json:"406"`
					Num407 float64 `json:"407"`
					Num408 float64 `json:"408"`
					Num409 float64 `json:"409"`
					Num410 float64 `json:"410"`
					Num411 float64 `json:"411"`
					Num412 float64 `json:"412"`
					Num413 float64 `json:"413"`
					Num414 float64 `json:"414"`
					Num415 float64 `json:"415"`
					Num416 float64 `json:"416"`
					Num417 float64 `json:"417"`
					Num418 float64 `json:"418"`
					Num419 float64 `json:"419"`
					Num420 float64 `json:"420"`
					Num421 float64 `json:"421"`
					Num422 float64 `json:"422"`
					Num423 float64 `json:"423"`
					Num424 float64 `json:"424"`
					Num425 float64 `json:"425"`
					Num426 float64 `json:"426"`
					Num427 float64 `json:"427"`
					Num428 float64 `json:"428"`
					Num429 float64 `json:"429"`
					Num430 float64 `json:"430"`
					Num431 float64 `json:"431"`
					Num432 float64 `json:"432"`
					Num433 float64 `json:"433"`
					Num434 float64 `json:"434"`
					Num435 float64 `json:"435"`
					Num436 float64 `json:"436"`
					Num437 float64 `json:"437"`
					Num438 float64 `json:"438"`
					Num439 float64 `json:"439"`
					Num440 float64 `json:"440"`
					Num441 float64 `json:"441"`
					Num442 float64 `json:"442"`
					Num443 float64 `json:"443"`
					Num444 float64 `json:"444"`
					Num445 float64 `json:"445"`
					Num446 float64 `json:"446"`
					Num447 float64 `json:"447"`
					Num448 float64 `json:"448"`
					Num449 float64 `json:"449"`
					Num450 float64 `json:"450"`
					Num452 float64 `json:"452"`
					Num457 float64 `json:"457"`
					Num458 float64 `json:"458"`
					Num460 float64 `json:"460"`
					Num461 float64 `json:"461"`
					Num464 float64 `json:"464"`
					Num466 float64 `json:"466"`
					Num471 float64 `json:"471"`
					Num481 float64 `json:"481"`
					Num482 float64 `json:"482"`
				} `json:"clintonh"`
				Trumpd struct {
					Num56  float64 `json:"56"`
					Num57  float64 `json:"57"`
					Num67  float64 `json:"67"`
					Num71  float64 `json:"71"`
					Num72  float64 `json:"72"`
					Num74  float64 `json:"74"`
					Num78  float64 `json:"78"`
					Num80  float64 `json:"80"`
					Num81  float64 `json:"81"`
					Num88  float64 `json:"88"`
					Num89  float64 `json:"89"`
					Num90  float64 `json:"90"`
					Num91  float64 `json:"91"`
					Num92  float64 `json:"92"`
					Num93  float64 `json:"93"`
					Num94  float64 `json:"94"`
					Num95  float64 `json:"95"`
					Num96  float64 `json:"96"`
					Num97  float64 `json:"97"`
					Num98  float64 `json:"98"`
					Num99  float64 `json:"99"`
					Num100 float64 `json:"100"`
					Num101 float64 `json:"101"`
					Num102 float64 `json:"102"`
					Num103 float64 `json:"103"`
					Num104 float64 `json:"104"`
					Num105 float64 `json:"105"`
					Num106 float64 `json:"106"`
					Num107 float64 `json:"107"`
					Num108 float64 `json:"108"`
					Num109 float64 `json:"109"`
					Num110 float64 `json:"110"`
					Num111 float64 `json:"111"`
					Num112 float64 `json:"112"`
					Num113 float64 `json:"113"`
					Num114 float64 `json:"114"`
					Num115 float64 `json:"115"`
					Num116 float64 `json:"116"`
					Num117 float64 `json:"117"`
					Num118 float64 `json:"118"`
					Num119 float64 `json:"119"`
					Num120 float64 `json:"120"`
					Num121 float64 `json:"121"`
					Num122 float64 `json:"122"`
					Num123 float64 `json:"123"`
					Num124 float64 `json:"124"`
					Num125 float64 `json:"125"`
					Num126 float64 `json:"126"`
					Num127 float64 `json:"127"`
					Num128 float64 `json:"128"`
					Num129 float64 `json:"129"`
					Num130 float64 `json:"130"`
					Num131 float64 `json:"131"`
					Num132 float64 `json:"132"`
					Num133 float64 `json:"133"`
					Num134 float64 `json:"134"`
					Num135 float64 `json:"135"`
					Num136 float64 `json:"136"`
					Num137 float64 `json:"137"`
					Num138 float64 `json:"138"`
					Num139 float64 `json:"139"`
					Num140 float64 `json:"140"`
					Num141 float64 `json:"141"`
					Num142 float64 `json:"142"`
					Num143 float64 `json:"143"`
					Num144 float64 `json:"144"`
					Num145 float64 `json:"145"`
					Num146 float64 `json:"146"`
					Num147 float64 `json:"147"`
					Num148 float64 `json:"148"`
					Num149 float64 `json:"149"`
					Num150 float64 `json:"150"`
					Num151 float64 `json:"151"`
					Num152 float64 `json:"152"`
					Num153 float64 `json:"153"`
					Num154 float64 `json:"154"`
					Num155 float64 `json:"155"`
					Num156 float64 `json:"156"`
					Num157 float64 `json:"157"`
					Num158 float64 `json:"158"`
					Num159 float64 `json:"159"`
					Num160 float64 `json:"160"`
					Num161 float64 `json:"161"`
					Num162 float64 `json:"162"`
					Num163 float64 `json:"163"`
					Num164 float64 `json:"164"`
					Num165 float64 `json:"165"`
					Num166 float64 `json:"166"`
					Num167 float64 `json:"167"`
					Num168 float64 `json:"168"`
					Num169 float64 `json:"169"`
					Num170 float64 `json:"170"`
					Num171 float64 `json:"171"`
					Num172 float64 `json:"172"`
					Num173 float64 `json:"173"`
					Num174 float64 `json:"174"`
					Num175 float64 `json:"175"`
					Num176 float64 `json:"176"`
					Num177 float64 `json:"177"`
					Num178 float64 `json:"178"`
					Num179 float64 `json:"179"`
					Num180 float64 `json:"180"`
					Num181 float64 `json:"181"`
					Num182 float64 `json:"182"`
					Num183 float64 `json:"183"`
					Num184 float64 `json:"184"`
					Num185 float64 `json:"185"`
					Num186 float64 `json:"186"`
					Num187 float64 `json:"187"`
					Num188 float64 `json:"188"`
					Num189 float64 `json:"189"`
					Num190 float64 `json:"190"`
					Num191 float64 `json:"191"`
					Num192 float64 `json:"192"`
					Num193 float64 `json:"193"`
					Num194 float64 `json:"194"`
					Num195 float64 `json:"195"`
					Num196 float64 `json:"196"`
					Num197 float64 `json:"197"`
					Num198 float64 `json:"198"`
					Num199 float64 `json:"199"`
					Num200 float64 `json:"200"`
					Num201 float64 `json:"201"`
					Num202 float64 `json:"202"`
					Num203 float64 `json:"203"`
					Num204 float64 `json:"204"`
					Num205 float64 `json:"205"`
					Num206 float64 `json:"206"`
					Num207 float64 `json:"207"`
					Num208 float64 `json:"208"`
					Num209 float64 `json:"209"`
					Num210 float64 `json:"210"`
					Num211 float64 `json:"211"`
					Num212 float64 `json:"212"`
					Num213 float64 `json:"213"`
					Num214 float64 `json:"214"`
					Num215 float64 `json:"215"`
					Num216 float64 `json:"216"`
					Num217 float64 `json:"217"`
					Num218 float64 `json:"218"`
					Num219 float64 `json:"219"`
					Num220 float64 `json:"220"`
					Num221 float64 `json:"221"`
					Num222 float64 `json:"222"`
					Num223 float64 `json:"223"`
					Num224 float64 `json:"224"`
					Num225 float64 `json:"225"`
					Num226 float64 `json:"226"`
					Num227 float64 `json:"227"`
					Num228 float64 `json:"228"`
					Num229 float64 `json:"229"`
					Num230 float64 `json:"230"`
					Num231 float64 `json:"231"`
					Num232 float64 `json:"232"`
					Num233 float64 `json:"233"`
					Num234 float64 `json:"234"`
					Num235 float64 `json:"235"`
					Num236 float64 `json:"236"`
					Num237 float64 `json:"237"`
					Num238 float64 `json:"238"`
					Num239 float64 `json:"239"`
					Num240 float64 `json:"240"`
					Num241 float64 `json:"241"`
					Num242 float64 `json:"242"`
					Num243 float64 `json:"243"`
					Num244 float64 `json:"244"`
					Num245 float64 `json:"245"`
					Num246 float64 `json:"246"`
					Num247 float64 `json:"247"`
					Num248 float64 `json:"248"`
					Num249 float64 `json:"249"`
					Num250 float64 `json:"250"`
					Num251 float64 `json:"251"`
					Num252 float64 `json:"252"`
					Num253 float64 `json:"253"`
					Num254 float64 `json:"254"`
					Num255 float64 `json:"255"`
					Num256 float64 `json:"256"`
					Num257 float64 `json:"257"`
					Num258 float64 `json:"258"`
					Num259 float64 `json:"259"`
					Num260 float64 `json:"260"`
					Num261 float64 `json:"261"`
					Num262 float64 `json:"262"`
					Num263 float64 `json:"263"`
					Num264 float64 `json:"264"`
					Num265 float64 `json:"265"`
					Num266 float64 `json:"266"`
					Num267 float64 `json:"267"`
					Num268 float64 `json:"268"`
					Num269 float64 `json:"269"`
					Num270 float64 `json:"270"`
					Num271 float64 `json:"271"`
					Num272 float64 `json:"272"`
					Num273 float64 `json:"273"`
					Num274 float64 `json:"274"`
					Num275 float64 `json:"275"`
					Num276 float64 `json:"276"`
					Num277 float64 `json:"277"`
					Num278 float64 `json:"278"`
					Num279 float64 `json:"279"`
					Num280 float64 `json:"280"`
					Num281 float64 `json:"281"`
					Num282 float64 `json:"282"`
					Num283 float64 `json:"283"`
					Num284 float64 `json:"284"`
					Num285 float64 `json:"285"`
					Num286 float64 `json:"286"`
					Num287 float64 `json:"287"`
					Num288 float64 `json:"288"`
					Num289 float64 `json:"289"`
					Num290 float64 `json:"290"`
					Num291 float64 `json:"291"`
					Num292 float64 `json:"292"`
					Num293 float64 `json:"293"`
					Num294 float64 `json:"294"`
					Num295 float64 `json:"295"`
					Num296 float64 `json:"296"`
					Num297 float64 `json:"297"`
					Num298 float64 `json:"298"`
					Num299 float64 `json:"299"`
					Num300 float64 `json:"300"`
					Num301 float64 `json:"301"`
					Num302 float64 `json:"302"`
					Num303 float64 `json:"303"`
					Num304 float64 `json:"304"`
					Num305 float64 `json:"305"`
					Num306 float64 `json:"306"`
					Num307 float64 `json:"307"`
					Num308 float64 `json:"308"`
					Num309 float64 `json:"309"`
					Num310 float64 `json:"310"`
					Num311 float64 `json:"311"`
					Num312 float64 `json:"312"`
					Num313 float64 `json:"313"`
					Num314 float64 `json:"314"`
					Num315 float64 `json:"315"`
					Num316 float64 `json:"316"`
					Num317 float64 `json:"317"`
					Num318 float64 `json:"318"`
					Num319 float64 `json:"319"`
					Num320 float64 `json:"320"`
					Num321 float64 `json:"321"`
					Num322 float64 `json:"322"`
					Num323 float64 `json:"323"`
					Num324 float64 `json:"324"`
					Num325 float64 `json:"325"`
					Num326 float64 `json:"326"`
					Num327 float64 `json:"327"`
					Num328 float64 `json:"328"`
					Num329 float64 `json:"329"`
					Num330 float64 `json:"330"`
					Num331 float64 `json:"331"`
					Num332 float64 `json:"332"`
					Num333 float64 `json:"333"`
					Num334 float64 `json:"334"`
					Num335 float64 `json:"335"`
					Num336 float64 `json:"336"`
					Num337 float64 `json:"337"`
					Num338 float64 `json:"338"`
					Num339 float64 `json:"339"`
					Num340 float64 `json:"340"`
					Num341 float64 `json:"341"`
					Num342 float64 `json:"342"`
					Num343 float64 `json:"343"`
					Num344 float64 `json:"344"`
					Num345 float64 `json:"345"`
					Num346 float64 `json:"346"`
					Num347 float64 `json:"347"`
					Num348 float64 `json:"348"`
					Num349 float64 `json:"349"`
					Num350 float64 `json:"350"`
					Num351 float64 `json:"351"`
					Num352 float64 `json:"352"`
					Num353 float64 `json:"353"`
					Num354 float64 `json:"354"`
					Num355 float64 `json:"355"`
					Num356 float64 `json:"356"`
					Num357 float64 `json:"357"`
					Num358 float64 `json:"358"`
					Num359 float64 `json:"359"`
					Num360 float64 `json:"360"`
					Num361 float64 `json:"361"`
					Num362 float64 `json:"362"`
					Num363 float64 `json:"363"`
					Num364 float64 `json:"364"`
					Num365 float64 `json:"365"`
					Num366 float64 `json:"366"`
					Num367 float64 `json:"367"`
					Num369 float64 `json:"369"`
					Num370 float64 `json:"370"`
					Num371 float64 `json:"371"`
					Num372 float64 `json:"372"`
					Num373 float64 `json:"373"`
					Num374 float64 `json:"374"`
					Num375 float64 `json:"375"`
					Num376 float64 `json:"376"`
					Num377 float64 `json:"377"`
					Num378 float64 `json:"378"`
					Num380 float64 `json:"380"`
					Num382 float64 `json:"382"`
					Num384 float64 `json:"384"`
					Num385 float64 `json:"385"`
					Num386 float64 `json:"386"`
					Num390 float64 `json:"390"`
					Num393 float64 `json:"393"`
					Num396 float64 `json:"396"`
					Num398 float64 `json:"398"`
					Num399 float64 `json:"399"`
					Num400 float64 `json:"400"`
					Num402 float64 `json:"402"`
					Num403 float64 `json:"403"`
					Num406 float64 `json:"406"`
					Num407 float64 `json:"407"`
					Num409 float64 `json:"409"`
					Num410 float64 `json:"410"`
					Num430 float64 `json:"430"`
				} `json:"trumpd"`
				Johnsong struct {
					Num0 float64 `json:"0"`
					Num4 float64 `json:"4"`
				} `json:"johnsong"`
				Mcmulline struct {
					Num0  float64 `json:"0"`
					Num4  float64 `json:"4"`
					Num6  float64 `json:"6"`
					Num10 float64 `json:"10"`
				} `json:"mcmulline"`
				Steinj struct {
					Num0 int `json:"0"`
				} `json:"steinj"`
				Other struct {
					Num0 int `json:"0"`
				} `json:"other"`
			} `json:"electoral_votes_distribution"`
			ElectoralVotesExpected struct {
				Clintonh  []int `json:"clintonh"`
				Trumpd    []int `json:"trumpd"`
				Johnsong  []int `json:"johnsong"`
				Mcmulline []int `json:"mcmulline"`
				Steinj    []int `json:"steinj"`
				Other     []int `json:"other"`
			} `json:"electoral_votes_expected"`
			VoteShareCounted struct {
				Clintonh  float64 `json:"clintonh"`
				Trumpd    float64 `json:"trumpd"`
				Johnsong  float64 `json:"johnsong"`
				Mcmulline float64 `json:"mcmulline"`
				Steinj    float64 `json:"steinj"`
				Other     float64 `json:"other"`
			} `json:"vote_share_counted"`
			VoteShareExpected struct {
				Clintonh  []float64 `json:"clintonh"`
				Trumpd    []float64 `json:"trumpd"`
				Johnsong  []float64 `json:"johnsong"`
				Mcmulline []float64 `json:"mcmulline"`
				Steinj    []float64 `json:"steinj"`
				Other     []float64 `json:"other"`
			} `json:"vote_share_expected"`
			WinProb struct {
				Clintonh  float64 `json:"clintonh"`
				Trumpd    float64 `json:"trumpd"`
				Johnsong  int     `json:"johnsong"`
				Mcmulline int     `json:"mcmulline"`
				Steinj    int     `json:"steinj"`
				Other     int     `json:"other"`
			} `json:"win_prob"`
		} `json:"current"`
		Races []struct {
			Type     string `json:"type"`
			State    string `json:"state"`
			District int    `json:"district"`
			Current  struct {
				Timestamp          time.Time `json:"timestamp"`
				PrecinctsReporting int       `json:"precincts_reporting"`
				PrecinctsTotal     int       `json:"precincts_total"`
				PercentCounted     int       `json:"percent_counted"`
				VoteShareCounted   struct {
					Castled     string `json:"castled"`
					Clintonh    string `json:"clintonh"`
					DeLaFuenter string `json:"de_la_fuenter"`
					Johnsong    string `json:"johnsong"`
					Other       string `json:"other"`
					Steinj      string `json:"steinj"`
					Trumpd      string `json:"trumpd"`
				} `json:"vote_share_counted"`
				VoteShareExpected struct {
					Clintonh  []float64 `json:"clintonh"`
					Trumpd    []float64 `json:"trumpd"`
					Johnsong  []float64 `json:"johnsong"`
					Mcmulline []int     `json:"mcmulline"`
					Steinj    []float64 `json:"steinj"`
					Other     []int     `json:"other"`
				} `json:"vote_share_expected"`
				WinProb struct {
					Clintonh  float64 `json:"clintonh"`
					Trumpd    float64 `json:"trumpd"`
					Johnsong  int     `json:"johnsong"`
					Mcmulline int     `json:"mcmulline"`
					Steinj    int     `json:"steinj"`
					Other     int     `json:"other"`
				} `json:"win_prob"`
				Winner interface{} `json:"winner"`
			} `json:"current"`
		} `json:"races"`
		Timeseries []struct {
			Timestamp             time.Time `json:"timestamp"`
			ElectoralVotesCounted struct {
				Clintonh  int `json:"clintonh"`
				Johnsong  int `json:"johnsong"`
				Mcmulline int `json:"mcmulline"`
				Other     int `json:"other"`
				Steinj    int `json:"steinj"`
				Trumpd    int `json:"trumpd"`
			} `json:"electoral_votes_counted"`
			ElectoralVotesExpected struct {
				Clintonh  []int `json:"clintonh"`
				Trumpd    []int `json:"trumpd"`
				Johnsong  []int `json:"johnsong"`
				Mcmulline []int `json:"mcmulline"`
				Steinj    []int `json:"steinj"`
				Other     []int `json:"other"`
			} `json:"electoral_votes_expected"`
			VoteShareCounted struct {
				Clintonh  float64 `json:"clintonh"`
				Trumpd    float64 `json:"trumpd"`
				Johnsong  float64 `json:"johnsong"`
				Mcmulline float64 `json:"mcmulline"`
				Steinj    float64 `json:"steinj"`
				Other     float64 `json:"other"`
			} `json:"vote_share_counted"`
			VoteShareExpected struct {
				Clintonh  []float64 `json:"clintonh"`
				Trumpd    []float64 `json:"trumpd"`
				Johnsong  []float64 `json:"johnsong"`
				Mcmulline []float64 `json:"mcmulline"`
				Steinj    []float64 `json:"steinj"`
				Other     []float64 `json:"other"`
			} `json:"vote_share_expected"`
			WinProb struct {
				Clintonh  float64 `json:"clintonh"`
				Trumpd    float64 `json:"trumpd"`
				Johnsong  int     `json:"johnsong"`
				Mcmulline int     `json:"mcmulline"`
				Steinj    int     `json:"steinj"`
				Other     int     `json:"other"`
			} `json:"win_prob"`
		} `json:"timeseries"`
	} `json:"president"`
	Senate struct {
		Current struct {
			Timestamp    time.Time `json:"timestamp"`
			SeatsCounted struct {
				Democrat   int `json:"democrat"`
				Republican int `json:"republican"`
			} `json:"seats_counted"`
			SeatsExpected struct {
				Democrat   []int `json:"democrat"`
				Republican []int `json:"republican"`
			} `json:"seats_expected"`
			SeatsDistribution struct {
				Democrat struct {
					Num43 float64 `json:"43"`
					Num44 float64 `json:"44"`
					Num45 float64 `json:"45"`
					Num46 float64 `json:"46"`
					Num47 float64 `json:"47"`
					Num48 float64 `json:"48"`
					Num49 float64 `json:"49"`
					Num50 float64 `json:"50"`
					Num51 float64 `json:"51"`
					Num52 float64 `json:"52"`
					Num53 float64 `json:"53"`
					Num54 float64 `json:"54"`
					Num55 float64 `json:"55"`
					Num56 float64 `json:"56"`
					Num58 float64 `json:"58"`
				} `json:"democrat"`
				Republican struct {
					Num42 float64 `json:"42"`
					Num44 float64 `json:"44"`
					Num45 float64 `json:"45"`
					Num46 float64 `json:"46"`
					Num47 float64 `json:"47"`
					Num48 float64 `json:"48"`
					Num49 float64 `json:"49"`
					Num50 float64 `json:"50"`
					Num51 float64 `json:"51"`
					Num52 float64 `json:"52"`
					Num53 float64 `json:"53"`
					Num54 float64 `json:"54"`
					Num55 float64 `json:"55"`
					Num56 float64 `json:"56"`
					Num57 float64 `json:"57"`
				} `json:"republican"`
			} `json:"seats_distribution"`
			SeatsMean struct {
				Democrat   float64 `json:"democrat"`
				Republican float64 `json:"republican"`
			} `json:"seats_mean"`
			WinProb struct {
				Democrat   float64 `json:"democrat"`
				Republican float64 `json:"republican"`
			} `json:"win_prob"`
		} `json:"current"`
	} `json:"senate"`
	House struct {
		Current struct {
			Timestamp    time.Time `json:"timestamp"`
			SeatsCounted struct {
				Democrat   int `json:"democrat"`
				Republican int `json:"republican"`
			} `json:"seats_counted"`
			SeatsDistribution struct {
				Democrat struct {
					Num142 float64 `json:"142"`
					Num147 float64 `json:"147"`
					Num148 float64 `json:"148"`
					Num153 float64 `json:"153"`
					Num155 float64 `json:"155"`
					Num157 float64 `json:"157"`
					Num158 float64 `json:"158"`
					Num160 float64 `json:"160"`
					Num161 float64 `json:"161"`
					Num162 float64 `json:"162"`
					Num163 float64 `json:"163"`
					Num164 float64 `json:"164"`
					Num165 float64 `json:"165"`
					Num166 float64 `json:"166"`
					Num167 float64 `json:"167"`
					Num168 float64 `json:"168"`
					Num169 float64 `json:"169"`
					Num170 float64 `json:"170"`
					Num171 float64 `json:"171"`
					Num172 float64 `json:"172"`
					Num173 float64 `json:"173"`
					Num174 float64 `json:"174"`
					Num175 float64 `json:"175"`
					Num176 float64 `json:"176"`
					Num177 float64 `json:"177"`
					Num178 float64 `json:"178"`
					Num179 float64 `json:"179"`
					Num180 float64 `json:"180"`
					Num181 float64 `json:"181"`
					Num182 float64 `json:"182"`
					Num183 float64 `json:"183"`
					Num184 float64 `json:"184"`
					Num185 float64 `json:"185"`
					Num186 float64 `json:"186"`
					Num187 float64 `json:"187"`
					Num188 float64 `json:"188"`
					Num189 float64 `json:"189"`
					Num190 float64 `json:"190"`
					Num191 float64 `json:"191"`
					Num192 float64 `json:"192"`
					Num193 float64 `json:"193"`
					Num194 float64 `json:"194"`
					Num195 float64 `json:"195"`
					Num196 float64 `json:"196"`
					Num197 float64 `json:"197"`
					Num198 float64 `json:"198"`
					Num199 float64 `json:"199"`
					Num200 float64 `json:"200"`
					Num201 float64 `json:"201"`
					Num202 float64 `json:"202"`
					Num203 float64 `json:"203"`
					Num204 float64 `json:"204"`
					Num205 float64 `json:"205"`
					Num206 float64 `json:"206"`
					Num207 float64 `json:"207"`
					Num208 float64 `json:"208"`
					Num209 float64 `json:"209"`
					Num210 float64 `json:"210"`
					Num211 float64 `json:"211"`
					Num212 float64 `json:"212"`
					Num213 float64 `json:"213"`
					Num214 float64 `json:"214"`
					Num215 float64 `json:"215"`
					Num216 float64 `json:"216"`
					Num217 float64 `json:"217"`
					Num218 float64 `json:"218"`
					Num219 float64 `json:"219"`
					Num220 float64 `json:"220"`
					Num221 float64 `json:"221"`
					Num222 float64 `json:"222"`
					Num223 float64 `json:"223"`
					Num224 float64 `json:"224"`
					Num225 float64 `json:"225"`
					Num226 float64 `json:"226"`
					Num227 float64 `json:"227"`
					Num228 float64 `json:"228"`
					Num229 float64 `json:"229"`
					Num230 float64 `json:"230"`
					Num231 float64 `json:"231"`
					Num232 float64 `json:"232"`
					Num233 float64 `json:"233"`
					Num234 float64 `json:"234"`
					Num235 float64 `json:"235"`
					Num236 float64 `json:"236"`
					Num237 float64 `json:"237"`
					Num238 float64 `json:"238"`
					Num239 float64 `json:"239"`
					Num240 float64 `json:"240"`
					Num241 float64 `json:"241"`
					Num242 float64 `json:"242"`
					Num243 float64 `json:"243"`
					Num244 float64 `json:"244"`
					Num245 float64 `json:"245"`
					Num246 float64 `json:"246"`
					Num247 float64 `json:"247"`
					Num250 float64 `json:"250"`
					Num253 float64 `json:"253"`
					Num254 float64 `json:"254"`
					Num258 float64 `json:"258"`
					Num276 float64 `json:"276"`
				} `json:"democrat"`
				Republican struct {
					Num159 float64 `json:"159"`
					Num177 float64 `json:"177"`
					Num181 float64 `json:"181"`
					Num182 float64 `json:"182"`
					Num185 float64 `json:"185"`
					Num188 float64 `json:"188"`
					Num189 float64 `json:"189"`
					Num190 float64 `json:"190"`
					Num191 float64 `json:"191"`
					Num192 float64 `json:"192"`
					Num193 float64 `json:"193"`
					Num194 float64 `json:"194"`
					Num195 float64 `json:"195"`
					Num196 float64 `json:"196"`
					Num197 float64 `json:"197"`
					Num198 float64 `json:"198"`
					Num199 float64 `json:"199"`
					Num200 float64 `json:"200"`
					Num201 float64 `json:"201"`
					Num202 float64 `json:"202"`
					Num203 float64 `json:"203"`
					Num204 float64 `json:"204"`
					Num205 float64 `json:"205"`
					Num206 float64 `json:"206"`
					Num207 float64 `json:"207"`
					Num208 float64 `json:"208"`
					Num209 float64 `json:"209"`
					Num210 float64 `json:"210"`
					Num211 float64 `json:"211"`
					Num212 float64 `json:"212"`
					Num213 float64 `json:"213"`
					Num214 float64 `json:"214"`
					Num215 float64 `json:"215"`
					Num216 float64 `json:"216"`
					Num217 float64 `json:"217"`
					Num218 float64 `json:"218"`
					Num219 float64 `json:"219"`
					Num220 float64 `json:"220"`
					Num221 float64 `json:"221"`
					Num222 float64 `json:"222"`
					Num223 float64 `json:"223"`
					Num224 float64 `json:"224"`
					Num225 float64 `json:"225"`
					Num226 float64 `json:"226"`
					Num227 float64 `json:"227"`
					Num228 float64 `json:"228"`
					Num229 float64 `json:"229"`
					Num230 float64 `json:"230"`
					Num231 float64 `json:"231"`
					Num232 float64 `json:"232"`
					Num233 float64 `json:"233"`
					Num234 float64 `json:"234"`
					Num235 float64 `json:"235"`
					Num236 float64 `json:"236"`
					Num237 float64 `json:"237"`
					Num238 float64 `json:"238"`
					Num239 float64 `json:"239"`
					Num240 float64 `json:"240"`
					Num241 float64 `json:"241"`
					Num242 float64 `json:"242"`
					Num243 float64 `json:"243"`
					Num244 float64 `json:"244"`
					Num245 float64 `json:"245"`
					Num246 float64 `json:"246"`
					Num247 float64 `json:"247"`
					Num248 float64 `json:"248"`
					Num249 float64 `json:"249"`
					Num250 float64 `json:"250"`
					Num251 float64 `json:"251"`
					Num252 float64 `json:"252"`
					Num253 float64 `json:"253"`
					Num254 float64 `json:"254"`
					Num255 float64 `json:"255"`
					Num256 float64 `json:"256"`
					Num257 float64 `json:"257"`
					Num258 float64 `json:"258"`
					Num259 float64 `json:"259"`
					Num260 float64 `json:"260"`
					Num261 float64 `json:"261"`
					Num262 float64 `json:"262"`
					Num263 float64 `json:"263"`
					Num264 float64 `json:"264"`
					Num265 float64 `json:"265"`
					Num266 float64 `json:"266"`
					Num267 float64 `json:"267"`
					Num268 float64 `json:"268"`
					Num269 float64 `json:"269"`
					Num270 float64 `json:"270"`
					Num271 float64 `json:"271"`
					Num272 float64 `json:"272"`
					Num273 float64 `json:"273"`
					Num274 float64 `json:"274"`
					Num275 float64 `json:"275"`
					Num277 float64 `json:"277"`
					Num278 float64 `json:"278"`
					Num280 float64 `json:"280"`
					Num282 float64 `json:"282"`
					Num287 float64 `json:"287"`
					Num288 float64 `json:"288"`
					Num293 float64 `json:"293"`
				} `json:"republican"`
			} `json:"seats_distribution"`
			SeatsExpected struct {
				Democrat   []int `json:"democrat"`
				Republican []int `json:"republican"`
			} `json:"seats_expected"`
			WinProb struct {
				Democrat   float64 `json:"democrat"`
				Republican float64 `json:"republican"`
			} `json:"win_prob"`
		} `json:"current"`
	} `json:"house"`
}

func NYTApi() {

	url := "https://intf.nyt.com/newsgraphics/2016/11-08-election-forecast/president.json"

	payload := strings.NewReader("access_token=df21db852a652855a4dcd72f55951879b3a39c09&color=RGB")

	req, _ := http.NewRequest("GET", url, payload)

	req.Header.Add("authorization", "Bearer df21db852a652855a4dcd72f55951879b3a39c09")
	req.Header.Add("cache-control", "no-cache")
	req.Header.Add("postman-token", "2f545f2f-8d93-2896-ec41-98eea2e33881")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)
	fmt.Println(body)

}

/*
func main() {
,
	r := mux.NewRouter()
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("example")))
	r.HandleFunc("/color", setColor)
	http.Handle("/", r)

	log.Println("Listening.....")
	http.ListenAndServe(":3000", nil)
}

func setColor(w http.ResponseWriter, r *http.Request) {
	reqBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Println(err)
	}
	io.WriteString(w, string(reqBody))
}
*/

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

	fmt.Println(string(body))

}

/*
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
}
*/
