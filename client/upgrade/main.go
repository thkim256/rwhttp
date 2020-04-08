package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	server := "10.20.200.111:30000"
	// server := "10.96.1.14:80"
	// server := "localhost:8080"
	dialer := &net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}
	conn, err := dialer.Dial("tcp", server)
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	reader := bufio.NewReader(conn)
	request, _ := http.NewRequest(
		"GET",
		fmt.Sprintf("https://%s/terminal.sh.hira.nginx-598c9d76dc-djgjb.app/357/ozlwrtfg/websocket", server),
		nil)
	request.Header.Set("Connection", "upgrade")
	request.Header.Set("Upgrade", "websocket")
	request.Header.Set("Sec-WebSocket-Accept", "60PcGwcfLoJd184+eNSFiOFmR8A=")
	request.Header.Set("Sec-WebSocket-Extensions", "permessage-deflate")
	request.Header.Set("Set-Cookie", "ACCSESSIONID=0B16B84848BD19C997C2926689B48DBD; Path=/; HttpOnly")
	request.Header.Set("Authorization", " Bearer eyJhbGciOiJSUzI1NiIsInR5cCIgOiAiSldUIiwia2lkIiA6ICIwMU5GdUNfdjVSRG5yelNjTjQxaVZnMXhUX2tKQS10M3pZek9YMHp5UEJVIn0.eyJqdGkiOiIwYTI3NGRlOS0yOTFiLTQ5OTEtYmNhZC02NDk1NmViNTIwNGYiLCJleHAiOjE1ODUwMTU3NDksIm5iZiI6MCwiaWF0IjoxNTg1MDE1NDQ5LCJpc3MiOiJodHRwczovLzEwLjIwLjIwMC4xMTE6MzAwMDMvYXV0aC9yZWFsbXMvYWNjb3JkaW9uIiwiYXVkIjoiYWNjb3VudCIsInN1YiI6IjM5ZDRkNGVmLWJiN2ItNGQ3NS04MmRlLTgyN2VjNGY1MjczNCIsInR5cCI6IkJlYXJlciIsImF6cCI6ImFjY29yZGlvbiIsImF1dGhfdGltZSI6MCwic2Vzc2lvbl9zdGF0ZSI6IjE5YWMwMjhhLTg4YzQtNDQ2Mi1iZDhiLTlmMzRiN2ZhYjUxOSIsImFjciI6IjEiLCJyZWFsbV9hY2Nlc3MiOnsicm9sZXMiOlsib2ZmbGluZV9hY2Nlc3MiLCJ1bWFfYXV0aG9yaXphdGlvbiJdfSwicmVzb3VyY2VfYWNjZXNzIjp7ImFjY291bnQiOnsicm9sZXMiOlsibWFuYWdlLWFjY291bnQiLCJtYW5hZ2UtYWNjb3VudC1saW5rcyIsInZpZXctcHJvZmlsZSJdfX0sInNjb3BlIjoiYWNjb3JkaW9uX2dyb3VwcyB1c2VyX25hbWUgcHJvZmlsZSBlbWFpbCIsImVtYWlsX3ZlcmlmaWVkIjpmYWxzZSwidXNlcl9uYW1lIjoiYWRtaW4iLCJncm91cHMiOlsiYWRtaW5zIl0sInByZWZlcnJlZF91c2VybmFtZSI6ImFkbWluIiwiYWNjb3JkaW9uX2dyb3VwcyI6WyJhZG1pbnMiXSwiZ2l2ZW5fbmFtZSI6IiIsImZhbWlseV9uYW1lIjoiIiwiZW1haWwiOiJvc3MtdGVjaEBtYW50ZWNoLmNvLmtyIn0.mjhbJNhYjV5FeKRmH16IukzZgZDByVgB3ECbtGwJ_QWodU-_0MTXVyNZuHAst8aZEu7mXowcv-Fgi0bL8N_Vj9ZO5AbLRF6DvoYUrRycXqytOkQxu1ydS0M_ZMRCU694fzuLviuPc2h9x5SwplVKl-CiRrS-xgiyqdH53g9DjI_SPcUNc44F_ZzZhdy2ddOSSbnRRxF81rlhhWWUJ2PtcJjT2k1AFotceQlqr2hR0sHYsGrigzcx9tq88IUGABJhhFDh907rQe2lmVIu1SgCfbHveN6TR8RCDJkjK2gHpu291qsPb4NZ1IrFPUBxYJbH5rPXCA-tG1xj3fsuQNU_1A")
	err = request.Write(conn)
	if err != nil {
		panic(err)
	}

	resp, err := http.ReadResponse(reader, request)
	if err != nil {
		panic(err)
	}

	log.Println("Status:", resp.Status)
	log.Println("Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println(string(body))

}
