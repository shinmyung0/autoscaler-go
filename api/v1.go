package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/shinmyung0/autoscaler/policy"
)

func v1ApiControllerInit(r *mux.Router) {

	s := r.PathPrefix("/v1").Subrouter()

	s.Methods("GET").Path("/policies").HandlerFunc(getAllPolicies)
	s.Methods("GET").Path("/services").HandlerFunc(getAllServices)
	s.Methods("GET").Path("/instances/{serviceName}").HandlerFunc(getAllInstances)
	s.Methods("GET").Path("/policy/{serviceName}").HandlerFunc(getPolicy)
	s.Methods("GET").Path("/status/{serviceName}").HandlerFunc(getStatus)
	s.Methods("POST", "PUT").Path("/policies").HandlerFunc(addPolicies)
	s.Methods("POST", "PUT").Path("/policy/{serviceName}").HandlerFunc(addPolicy)
	s.Methods("DELETE").Path("/policies").HandlerFunc(deleteAllPolicies)
	s.Methods("DELETE").Path("/policy/{serviceName}").HandlerFunc(deletePolicy)

}

func getAllPolicies(w http.ResponseWriter, req *http.Request) {
	m := policy.GetManager()
	policies := m.GetAllPolicies()
	json.NewEncoder(w).Encode(policies)
}

func getAllServices(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Not implemented", 404)
}

func getAllInstances(w http.ResponseWriter, req *http.Request) {
	//serviceName := mux.Vars(req)["serviceName"]
	http.Error(w, "Not implemented", 404)
}

func getPolicy(w http.ResponseWriter, req *http.Request) {
	serviceName := mux.Vars(req)["serviceName"]
	m := policy.GetManager()
	p, ok := m.GetPolicy(serviceName)
	if ok {
		json.NewEncoder(w).Encode(p)
	} else {
		http.Error(w, "Resource Unavailable", 404)
	}
}

func getStatus(w http.ResponseWriter, req *http.Request) {
	http.Error(w, "Not implemented", 404)
}

func addPolicies(w http.ResponseWriter, req *http.Request) {

	if req.Body == nil {
		http.Error(w, "Request body was empty", 400)
		return
	}

	var policies []policy.Policy
	err := json.NewDecoder(req.Body).Decode(&policies)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	m := policy.GetManager()
	m.AddPolicies(policies)

	fmt.Fprintf(w, "Added %d policies", len(policies))

}

func addPolicy(w http.ResponseWriter, req *http.Request) {

	if req.Body == nil {
		http.Error(w, "Please send policy", 400)
		return

	}
	var p policy.Policy
	err := json.NewDecoder(req.Body).Decode(&p)
	if err != nil {
		http.Error(w, err.Error(), 400)
		return
	}
	m := policy.GetManager()
	m.AddPolicy(p)

	fmt.Fprintf(w, "Added policy for service %s\n", p.Service)
}

func deleteAllPolicies(w http.ResponseWriter, req *http.Request) {
	m := policy.GetManager()
	m.DeleteAllPolicy()
	return
}

func deletePolicy(w http.ResponseWriter, req *http.Request) {
	serviceName := mux.Vars(req)["serviceName"]
	m := policy.GetManager()
	m.DeletePolicy(serviceName)
	return

}
