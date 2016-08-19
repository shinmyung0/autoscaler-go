package policy

import (
	"encoding/json"
	"testing"
)

func TestPolicyManager(t *testing.T) {

	m := GetManager()
	p, ok := m.GetPolicy("name")
	if ok || p.Service != "" {
		t.Error("Ok returned true")
	}

	data := []byte(policySample)
	var policies []Policy
	err := json.Unmarshal(data, &policies)
	if err != nil {
		t.Error("Marshalling error")
	}

	p = policies[0]
	m.AddPolicy(p)
	p, ok = m.GetPolicy("frontend-fib")
	if !ok || p.Service != "frontend-fib" {
		t.Error("Policy not added correctly")
	}

	m.DeletePolicy("frontend-fib")
	p, ok = m.GetPolicy("frontend-fib")

	if ok || p.Service != "" {
		t.Error("Policy not deleted")
	}

}
