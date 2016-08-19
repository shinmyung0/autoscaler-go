package policy

import (
	"encoding/json"
	"testing"
)

const policySample = `
[

    {
        "Service" : "frontend-fib",
        "EvaluationInterval" : 5,
        "ObservationInterval" : 30,
        "MinInstances" : 5,
        "MaxInstances" : 10,
        "DesiredState" : {

            "Springboot" : {
                "Max" : 200,
                "Endpoints" : [
                    "/fib"
                ]
            },
            "Cadvisor" : {
                "AvgCpu" : {
                    "Min" : 10,
                    "Max" : 80
                },
                "AvgMemory" : {
                    "Min" : 10,
                    "Max" : 80
                }
            }
        }
    }

]
`

func TestPolicyParsing(t *testing.T) {

	data := []byte(policySample)

	var policies []Policy
	err := json.Unmarshal(data, &policies)
	if err != nil {
		t.Error("Unmarshalling returned an error")
	}

	p := policies[0]
	if p.Service != "frontend-fib" {
		t.Error("Value not bound correctly")
	}

	if p.EvaluationInterval != 5 {
		t.Error("Value not bound correctly")
	}

	if p.ObservationInterval != 30 {
		t.Error("Value not bound correctly")
	}

	if p.MinInstances != 5 {
		t.Error("Value not bound correctly")
	}

	if p.MaxInstances != 10 {
		t.Error("Value not bound correctly")
	}

	sb := p.DesiredState.SpringBoot
	endpoints := sb.Endpoints

	// Parsed as float64
	if sb.Max != 200.00 {
		t.Error("Value not bound correctly")
	}

	if endpoints[0] != "/fib" {
		t.Error("Value not bound correctly")
	}

}
