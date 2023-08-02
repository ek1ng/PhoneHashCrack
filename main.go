package main

import (
	"crypto/sha256"
	"encoding/hex"
	"flag"
	"fmt"
	"runtime"
	"time"
)

var (
	logo        = " ____  _                      _   _           _      ____                _    \n|  _ \\| |__   ___  _ __   ___| | | | __ _ ___| |__  / ___|_ __ __ _  ___| | __\n| |_) | '_ \\ / _ \\| '_ \\ / _ \\ |_| |/ _` / __| '_ \\| |   | '__/ _` |/ __| |/ /\n|  __/| | | | (_) | | | |  __/  _  | (_| \\__ \\ | | | |___| | | (_| | (__|   < \n|_|   |_| |_|\\___/|_| |_|\\___|_| |_|\\__,_|___/_| |_|\\____|_|  \\__,_|\\___|_|\\_\\\n                                                                              \n"
	targetStart string //"11cbd"
	targetEnd   string //"893d4"

	areacodelist = []int{130, 131, 132, 133, 134, 135, 136,
		137, 138, 139, 145, 147, 149, 150,
		151, 152, 153, 155, 156, 157, 158,
		159, 162, 165, 166, 167, 170, 171,
		172, 173, 174, 175, 176, 177, 178,
		180, 181, 182, 183, 184, 185, 186,
		187, 188, 189, 190, 191, 192, 193,
		195, 196, 197, 198, 199}
)

func main() {
	fmt.Println(logo)
	threads := runtime.NumCPU() // 获取cpu逻辑核心数（包括超线程）
	c := make(chan struct{}, threads)
	start := time.Now()
	flag.StringVar(&targetStart, "start", "", "input start phone hash")
	flag.StringVar(&targetEnd, "end", "", "input end phone hash")
	flag.Parse()

	for _, areacode := range areacodelist {
		c <- struct{}{}
		go func(areacode int) {
			fmt.Println("[-] start to crack areacode :", areacode)
			defer func() { <-c }()
			for i := 0; i < 100000000; i++ {
				phone := fmt.Sprintf("86%d%08d", areacode, i)
				targetTest := sha256.Sum256([]byte(phone))
				var hexTest = make([]byte, hex.EncodedLen(len(targetTest)))
				hex.Encode(hexTest, targetTest[:])
				if targetStart == string(hexTest[0:5]) && targetEnd == string(hexTest[len(hexTest)-5:]) {
					fmt.Println("[+] find phone number : " + phone)
				}
			}
		}(areacode)
	}
	end := time.Since(start)
	fmt.Println(end)
}

