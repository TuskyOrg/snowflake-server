package main

import (
	"encoding/json"
	"fmt"
	"github.com/AmreeshTyagi/goldflake"
	"net/http"
	"strconv"
)

var gf *goldflake.Goldflake

func getMachineID() (uint16, error) {
	return 1234, nil
}

func init() {
	var settings goldflake.Settings
	settings.MachineID = getMachineID
	gf = goldflake.NewGoldflake(settings)
	if gf == nil {
		panic("goldflake not created")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	id, err := gf.NextID()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	fmt.Printf("%v\n", id)
	body, err := json.Marshal(
		map[string]string{
			"id": strconv.FormatUint(id, 10),
		})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}
	_, _ = w.Write(body)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
