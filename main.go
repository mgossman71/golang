package main

import (
	"fmt"
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
func getalldeploy(w http.ResponseWriter, r *http.Request) {
	cmd, _ := exec.Command("kubectl", "get", "deploy", "--all-namespaces").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func getallpods(w http.ResponseWriter, r *http.Request) {
	cmd, _ := exec.Command("kubectl", "get", "pods", "--all-namespaces", "-o", "wide").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func getallrs(w http.ResponseWriter, r *http.Request) {
	cmd, _ := exec.Command("kubectl", "get", "replicasets").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func updateimage(w http.ResponseWriter, r *http.Request) {
	// kubectl set image -n default deployment/gozznet-old gozznet-old=mgossman71/gozznet-old:v4
	param := mux.Vars(r)
	nsname := param["nsname"]
	objtype := param["objtype"]
	objname := param["objname"]
	image := param["image"]
	a := objtype + "/" + objname
	b := objname + "=" + image

	cmd, _ := exec.Command("kubectl", "set", "image", "-n", nsname, a, b).Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func describedeploy(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	nsname := param["nsname"]
	deployname := param["deployname"]

	cmd, _ := exec.Command("kubectl", "describe", "deploy", "-n", nsname, deployname).Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func scaleSabnzb(w http.ResponseWriter, r *http.Request) {
	// kubectl scale -n $1 --replicas=$2 deployment/sabnzbd-deployment
	param := mux.Vars(r)
	replicas := param["replicas"]
	nsname := param["nsname"]
	a := "--replicas=" + replicas
	cmd, _ := exec.Command("kubectl", "scale", "-n", nsname, a, "deployment/sabnzbd-deployment").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func scaleCp(w http.ResponseWriter, r *http.Request) {
	// kubectl scale -n $1 --replicas=$2 deployment/couchpotato-deployment
	param := mux.Vars(r)
	replicas := param["replicas"]
	nsname := param["nsname"]
	a := "--replicas=" + replicas
	cmd, _ := exec.Command("kubectl", "scale", "-n", nsname, a, "deployment/couchpotato-deployment").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func scaleSonarr(w http.ResponseWriter, r *http.Request) {
	// kubectl scale -n $1 --replicas=$2 deployment/sonarr-deployment
	param := mux.Vars(r)
	replicas := param["replicas"]
	nsname := param["nsname"]
	a := "--replicas=" + replicas
	cmd, _ := exec.Command("kubectl", "scale", "-n", nsname, a, "deployment/sonarr-deployment").Output()
	w.WriteHeader(http.StatusOK)
	w.Write(cmd)
}
func setupMuxRouter() *mux.Router {
	router := mux.NewRouter()
	api := router.PathPrefix("/api/v1").Subrouter()
	apiScale := router.PathPrefix("/api/v1/scale").Subrouter()
	apiUpdate := router.PathPrefix("/api/v1/update").Subrouter()
	apiDescribe := router.PathPrefix("/api/v1/describe").Subrouter()
	apiGet := router.PathPrefix("/api/v1/get").Subrouter()

	apiGeneric := router.PathPrefix("/api").Subrouter()

	apiGeneric.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("pong"))
	})
	api.HandleFunc("/new", myhandler)
	apiGet.HandleFunc("/allns", getallns)
	apiGet.HandleFunc("/onens", getonens).Queries("name", "{name}")
	apiGet.HandleFunc("/allpods", getallpods)
	apiGet.HandleFunc("/alldeploy", getalldeploy)
	apiGet.HandleFunc("/allnodes", getallnodes)
	apiGet.HandleFunc("/allrs", getallrs)
	apiDescribe.HandleFunc("/rs", describers).Queries("rsname", "{rsname}")
	apiDescribe.HandleFunc("/pod", describepod).Queries("nsname", "{nsname}", "podname", "{podname}")
	apiDescribe.HandleFunc("/deploy", describedeploy).Queries("nsname", "{nsname}", "deployname", "{deployname}")
	apiUpdate.HandleFunc("/image", updateimage).Queries("nsname", "{nsname}", "objtype", "{objtype}", "objname", "{objname}", "image", "{image}")
	apiScale.HandleFunc("/sabnzb", scaleSabnzb).Queries("replicas", "{replicas}", "nsname", "{nsname}")
	apiScale.HandleFunc("/cp", scaleCp).Queries("replicas", "{replicas}", "nsname", "{nsname}")
	apiScale.HandleFunc("/sonarr", scaleSonarr).Queries("replicas", "{replicas}", "nsname", "{nsname}")
	return router
}
func main() {

	muxRouter := setupMuxRouter()
	loggedRouter := handlers.LoggingHandler(os.Stdout, muxRouter)
	err := http.ListenAndServe(":8080", loggedRouter)
	// err := http.ListenAndServe(":32106", loggedRouter)
	if err != nil {
		fmt.Println(err)
	}

}
