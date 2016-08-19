package policy

type PolicyManager struct {
	policies map[string]Policy
}

var manager *PolicyManager

func init() {
	manager = &PolicyManager{make(map[string]Policy)}
}

func GetManager() *PolicyManager {
	return manager
}

func (m *PolicyManager) GetPolicy(serviceName string) (Policy, bool) {
	p, ok := m.policies[serviceName]
	return p, ok
}

func (m *PolicyManager) GetAllPolicies() []Policy {

	allPolicies := make([]Policy, len(m.policies))

	i := 0
	for _, p := range m.policies {
		allPolicies[i] = p
		i++
	}

	return allPolicies
}

func (m *PolicyManager) AddPolicy(p Policy) {
	m.policies[p.Service] = p
	GetEvaluator().AddEvaluation(p)
}

func (m *PolicyManager) AddPolicies(policies []Policy) {
	for _, p := range policies {
		m.AddPolicy(p)
	}
}

func (m *PolicyManager) DeletePolicy(serviceName string) {
	GetEvaluator().RemoveEvaluation(serviceName)
	delete(m.policies, serviceName)
}

func (m *PolicyManager) DeleteAllPolicy() {
	for s := range m.policies {
		m.DeletePolicy(s)
	}
}

func validatePolicy(p Policy) bool {
	return true
}
