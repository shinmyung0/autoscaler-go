package policy

type Policy struct {
	Service             string
	EvaluationInterval  int
	ObservationInterval int
	MinInstances        int
	MaxInstances        int

	DesiredState struct {
		SpringBoot struct {
			Max       int
			Endpoints []string
		}
		CAdvisor struct {
			AvgCpu struct {
				Min int
				Max int
			}
			AvgMemory struct {
				Min int
				Max int
			}
		}
	}
}
