package main

import (
	"fmt"
	//"log"
	"net/http"
	"os"
	"os/exec"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func describepod(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	nsname := param["nsname"]
	podname := param["podname"]
	fmt.Println(nsname)
	cmd, _ := exec.Command("kubectl", "describe", "pod", "-n", nsname, podname).Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}

func myhandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Hey there!!"))
}

func getallns(w http.ResponseWriter, r *http.Request) {
	cmd, _ := exec.Command("kubectl", "get", "ns").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}

func getonens(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	name := param["name"]
	cmd, _ := exec.Command("kubectl", "get", "ns", name).Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func describers(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	rsname := param["rsname"]
	cmd, _ := exec.Command("kubectl", "describe", "rs", rsname).Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func getallnodes(w http.ResponseWriter, r *http.Request) {
	cmd, _ := exec.Command("kubectl", "get", "nodes").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func getallpods(w http.ResponseWriter, r *http.Request) {
	cmd, _ := exec.Command("kubectl", "get", "pods", "--all-namespaces").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func getallrs(w http.ResponseWriter, r *http.Request) {
	cmd, _ := exec.Command("kubectl", "get", "replicasets").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}

func setupMuxRouter() *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	apiGeneric := router.PathPrefix("/api").Subrouter()

	apiGeneric.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	api.HandleFunc("/new", myhandler)
	api.HandleFunc("/getallns", getallns)
	api.HandleFunc("/getonens", getonens).Queries("name", "{name}")
	api.HandleFunc("/getallpods", getallpods)
	api.HandleFunc("/getallnodes", getallnodes)
	api.HandleFunc("/getallrs", getallrs)
	api.HandleFunc("/describers", describers).Queries("rsname", "{rsname}")
	api.HandleFunc("/describepod", describepod).Queries("nsname", "{nsname}", "podname", "{podname}")
	return router
}

func main() {

	muxRouter := setupMuxRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, muxRouter)
	err := http.ListenAndServe(":8080", loggedRouter)
	if err != nil {
		fmt.Println(err)
	}

}
