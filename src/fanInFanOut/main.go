package main

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

func main() {
	f, err := os.Open("src/fanInFanOut/worldcities.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	rows := genRows(f)
	up1 := upperCityNames(filterPopulation(45000, 50000,rows))
	up2 := upperCityNames(filterPopulation(45000, 50000, rows))
	upperRows := fanIn(up1, up2)
	for c := range upperRows {
		log.Println(c)
	}
}

type City struct {
	Name       string
	Population int
}

func genRows(r io.Reader) <-chan City {
	out := make(chan City)
	go func() {
		reader := csv.NewReader(r)
		for {
			row, err := reader.Read()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			population, err := strconv.Atoi(row[9])
			if err != nil {
				continue
			}
			out <- City{Name: row[1], Population: population}
		}
		close(out)
	}()
	return out
}

func upperCityNames(cities <-chan City) <-chan City {
	out := make(chan City)
	go func() {
		for c := range cities {
			out <- City{Name: strings.ToUpper(c.Name), Population: c.Population}
		}
		close(out)
	}()
	return out
}

func filterPopulation(min, max int, cities <-chan City) <-chan City {
	out := make(chan City)
	go func() {
		for c := range cities {
			if c.Population > min && c.Population < max {
				out <- City{Name: strings.ToUpper(c.Name), Population: c.Population}
			}	
		}
		close(out)
	}()
	return out
}

func fanIn(cities ...<-chan City) <-chan City {
	result := make(chan City)
	wg := &sync.WaitGroup{}
	wg.Add(len(cities))
	for _, c:= range cities {
		go func(city <-chan City) {
			for r := range city {
				result <- r
			}
			wg.Done()
		}(c)
	}
	go func() {
		wg.Wait()
		close(result)
	}()
	return result
}