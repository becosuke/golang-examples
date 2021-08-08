package main

import (
	"encoding/json"
	helloService "example/hello/domain/service/hello"
	"fmt"
	"net/http"
)

func GetHello(w http.ResponseWriter, _ *http.Request) {
	res := helloService.GetHello()
	json, _ := json.Marshal(res)
	fmt.Fprint(w, string(json))
}
