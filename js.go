package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"io/ioutil"

	"github.com/robertkrimen/otto"
)

// JsRun is ...
func JsRun(js string) {

	vm := otto.New()
	vm.Run(js + "; var VV = Hub.config.get('sku');var V2 = '22'; //console.log(VV)")

	hub, _ := vm.Get("VV")

	hubInte, _ := hub.Export()

	fmt.Println(hubInte)

	fmt.Println(reflect.TypeOf(hubInte))
	b, _ := json.Marshal(hubInte)

	// os.Stdout.Write(b)

	fmt.Println(string(b))

	var ws = []byte(string(b))

	ioutil.WriteFile("sku.json", ws, 0666)

}
