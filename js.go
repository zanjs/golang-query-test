package main

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/robertkrimen/otto"
)

// JsRun is ...
func JsRun(js string) {

	vm := otto.New()
	vm.Run(js + "; var VV = Hub.config.get('sku');var V2 = '22'; //console.log(VV)")

	hub, _ := vm.Get("VV")

	hubInte, _ := hub.Export()

	// _hub, _ := json.Marshal(hubstr)
	fmt.Println(hubInte)
	// fmt.Println(_hub)

	fmt.Println(reflect.TypeOf(hubInte))
	b, _ := json.Marshal(hubInte)

	// os.Stdout.Write(b)

	fmt.Println(string(b))

}
