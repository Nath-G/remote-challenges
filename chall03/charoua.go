package main

import "time"
import "net/http"
import "fmt"
import "io/ioutil"
import "strings"
import "strconv"

 func main() {
	start := time.Now()
	fmt.Printf("%dms - GET http://0.0.0.0:8080/\n", time.Since(start).Milliseconds())
	resp, err := http.Get("http://0.0.0.0:8080/")
	if err != nil {
		panic("failure of the GET request to http://0.0.0.0:8080/")
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic("failure of the reading of body")
	}
	fmt.Printf("%dms - GET http://0.0.0.0:8080/\n\tanswer:%s\n", time.Since(start).Milliseconds(), body)
	split := strings.Split(string(body), " ")
	split = strings.Split(string(split[0]), ",")
	id, err := strconv.Atoi(strings.Split(split[0], "=")[1])
	if err != nil {
		panic("failure of atoi : id is not an int")
	}
	r, err := strconv.Atoi(strings.Split(split[1], "=")[1])
	if err != nil {
		panic("failure of atoi : r is not an int")
	}
	g, err := strconv.Atoi(strings.Split(split[2], "=")[1])
	if err != nil {
		panic("failure of atoi : g is not an int")
	}
	b, err := strconv.Atoi(strings.Split(split[3], "=")[1])
	if err != nil {
		panic("failure of atoi : b is not an int")
	}
	resp, err = http.Get(fmt.Sprintf("http://0.0.0.0:8080/?id=%d&resp=%02x%02x%02x", id, r, g, b))
	fmt.Printf("%dms - GET http://0.0.0.0:8080/?id=%d&resp=%02x%02x%02x\n",time.Since(start).Milliseconds(), id, r, g, b)
	if err != nil {
		panic("failure of the Gnswer back to http://0.0.0.0:8080/")
	}
	if resp.StatusCode == http.StatusOK {
		fmt.Printf("%dms - GET http://0.0.0.0:8080/?id=%d&resp=%02x%02x%02x\n\tanswer: ok!\n",time.Since(start).Milliseconds(), id, r, g, b)
	} else {
		fmt.Printf("%dms - GET http://0.0.0.0:8080/?id=%d&resp=%02x%02x%02x\n\tanswer: ko :(\n",time.Since(start).Milliseconds(), id, r, g, b)
	}
	fmt.Printf("%dms - DONE", time.Since(start).Milliseconds())
}
