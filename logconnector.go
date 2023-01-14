package main

/*


params:

- hostname (node0, for example)
- service name (mtrader, for example)
- priority
- timestamp
- logstring
- version (both internal and provided to storage)




*/

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	//- hostname (node0, for example)
	//- service name (mtrader, for example)
	//-/ priority
	//- timestamp
	//- logstring
	//- version (both internal and provided to storage)
	//- table

	hostPtr := flag.String("host", "", "Hostname")
	servicePtr := flag.String("app", "", "App name")
	versionPtr := flag.String("version", "1.0", "version (only v1 supported so far)")
	tableNamePtr := flag.String("table", "", "tablename")

	flag.Parse()

	fmt.Printf("Parameters: %v %v %v %v \n", *hostPtr, *servicePtr, *versionPtr, *tableNamePtr)

	if hostPtr == nil || strings.Compare(*hostPtr, "") == 0 {
		fmt.Printf("correct usage: %v -host <host> -app <appname> [-version <version>] -table <tablename>\n", os.Args[0])
		os.Exit(-1)
	}

	if servicePtr == nil || strings.Compare(*servicePtr, "") == 0 {
		fmt.Printf("correct usage: %v -host <host> -app <appname> [-version <version>] -table <tablename>\n", os.Args[0])
		os.Exit(-1)
	}

	if tableNamePtr == nil || strings.Compare(*tableNamePtr, "") == 0 {
		fmt.Printf("correct usage: %v -host <host> -app <appname> [-version <version>] -table <tablename>\n", os.Args[0])
		os.Exit(-1)
	}

	url := fmt.Sprintf("https://api.us-east.tinybird.co/v0/events?name=%v", *tableNamePtr)

	for {

		text, err := reader.ReadString('\n')

		if err == nil {

			text = strings.Trim(text, "\n")
			text = strconv.Quote(text)

			timestamp := time.Now()

			logString := fmt.Sprintf(`{"timestamp":"%v","host":"%v","app":"%v","version":"%v","priority":"default","log":"%v"}`, timestamp.Format(time.RFC3339Nano), *hostPtr, *servicePtr, *versionPtr, text)
			fmt.Println(logString)

			var jsonStr = []byte(logString)
			//var jsonStr = []byte(`{"timestamp":"2022-10-27T11:43:02.099Z","transaction_id":"8d1e1533-6071-4b10-9cda-b8429c1c7a67","priority_boarding":false,"meal_choice":"vegetarian","seat_number":"15D","airline":"Red Balloon"}`)

			req, err2 := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
			req.Header.Set("Authorization", "Bearer p.eyJ1IjogIjY4ODdjM2E4LTdjMDMtNDJlMC04NzY0LTljMjI0MDcyZTFkOSIsICJpZCI6ICJkZGQ1MmFlNC05ODA2LTQ4NTMtYjAyYS1kMzY0OWFjMGM5MjkifQ.4TAVT2d3UtJvERwroWbsswA9SZX6DbjvpTZLPlLBfYw")
			req.Header.Set("Content-Type", "application/json")

			client := &http.Client{Timeout: time.Second * 10}
			resp, err2 := client.Do(req)
			if err2 != nil {
				panic(err2)
			}
			defer resp.Body.Close()

			body, _ := ioutil.ReadAll(resp.Body)
			fmt.Println(string(body))
		}

	}

}
