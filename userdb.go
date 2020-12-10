package main

// Append contacts from dstar.su database and TG list to TYT MD380/390 userdb.bin
// (c) 2020, EU1ADI

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
)

type mobileCode struct {
	Name string
	ISO  string
}

var wg sync.WaitGroup
var mobileCodes map[int]mobileCode
var userDB map[int]string
var userDStar map[int]string
var bmGroups map[int]string
var prefixes []string
var skipTG bool

func checkPrefix(ID string) bool {
	for _, prefix := range prefixes {
		if strings.HasPrefix(ID, prefix) {
			return true
		}
	}
	return false
}

func processContacts() {
	count := 0
	for dstarID := range userDStar {
		_, ok := userDB[dstarID]
		if ok {
			continue
		}

		data := strings.Split(userDStar[dstarID], "\t")

		countryName := ""
		iso := ""

		strID := data[0]
		if len(strID) >= 7 {
			countryID, _ := strconv.Atoi(strID[0:3])
			country, ok := mobileCodes[countryID]
			if ok {
				countryName = country.Name
				iso = country.ISO
			}
		}

		userDB[dstarID] = fmt.Sprintf("%d,%s,%s,%s,,,%s", dstarID, data[1], data[2], countryName, iso)
		count++
	}

	if !skipTG {
		for groupID := range bmGroups {
			_, ok := userDB[groupID]
			if ok {
				continue
			}

			countryName := ""
			iso := ""

			strID := strconv.Itoa(groupID)
			if len(strID) >= 3 {
				countryID, _ := strconv.Atoi(strID[0:3])
				country, ok := mobileCodes[countryID]
				if ok {
					countryName = country.Name
					iso = country.ISO
				}
			}

			userDB[groupID] = fmt.Sprintf("%d,TG%d,%s,%s,,,%s", groupID, groupID, strings.TrimSpace(bmGroups[groupID]), countryName, iso)
			count++
		}
	}

	println("total", len(userDB), "lines, added", count)

	keys := make([]int, len(userDB))
	i := 0
	total := 0
	for k := range userDB {
		total = total + len(userDB[k]) + 1
		keys[i] = k
		i++
	}
	sort.Ints(keys)

	f, err := os.Create("userdb.bin")
	if err != nil {
		return
	}
	defer f.Close()

	f.WriteString(fmt.Sprintf("%d\n", total))
	for _, k := range keys {
		f.WriteString(userDB[k] + "\n")
	}

}

func loadCodes() {
	defer wg.Done()

	mobileCodes = make(map[int]mobileCode, 0)
	for _, code := range strings.Split(codes, "\n") {
		data := strings.Split(code, "\t")
		if len(data) != 3 {
			continue
		}

		code, _ := strconv.Atoi(strings.TrimSpace(data[0]))
		mobileCodes[code] = mobileCode{
			Name: data[1],
			ISO:  data[2],
		}
	}

	println("Codes list done")
}

func loadUserDB() {
	defer wg.Done()

	resp, err := http.Get("https://raw.githubusercontent.com/DMR-Database/database/master/user.bin")
	if err != nil {
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		println("UserDB load error")
		return
	}

	userDB = make(map[int]string, 0)
	for _, line := range strings.Split(string(body), "\n") {
		data := strings.Split(line, ",")
		if len(data) != 7 {
			continue
		}

		if checkPrefix(data[0]) {
			continue
		}

		id, _ := strconv.Atoi(strings.TrimSpace(data[0]))
		userDB[id] = line
	}
	println("UserDB done")
}

func loadDStarSU() {
	defer wg.Done()

	resp, err := http.Get("http://registry.dstar.su/dmr/DMRIds2.php")
	if err != nil {
		println("DStar.su load error")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	userDStar = make(map[int]string)
	for _, line := range strings.Split(string(body), "\n") {
		if len(line) > 0 && line[:1] == "#" { // skip commented line
			continue
		}

		data := strings.Split(line, "\t")
		if len(data) != 3 {
			continue
		}

		if checkPrefix(data[0]) {
			continue
		}

		id, _ := strconv.Atoi(strings.TrimSpace(data[0]))
		userDStar[id] = line
	}
	println("DStar.su done")
}

func loadTGList() {
	defer wg.Done()

	resp, err := http.Get("https://api.brandmeister.network/v1.0/groups/")
	if err != nil {
		println("TG list load error")
		return
	}
	defer resp.Body.Close()

	bmGroups = make(map[int]string)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}

	err = json.Unmarshal([]byte(body), &bmGroups)
	if err != nil {
		return
	}
	println("TG list done")
}

func main() {
	exclude := flag.String("exclude", "", "exclude country prefixes, example: 310,311,312")
	skip := flag.Bool("skiptg", false, "skip adding TGs")
	flag.Parse()

	if len(*exclude) > 0 {
		prefixes = strings.Split(*exclude, ",")
	}

	skipTG = *skip
	println("exclude:", *exclude)
	println("skipTG:", skipTG)

	wg.Add(3)
	go loadCodes()
	go loadUserDB()
	go loadDStarSU()

	if !skipTG {
		wg.Add(1)
		go loadTGList()
	}

	wg.Wait()
	processContacts()
}
