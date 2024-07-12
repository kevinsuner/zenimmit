package stepper

type StepType int8

const (
	CommitType StepType = iota
	BreakingChange
	CommitScope
)

type Stepper struct {
	Type  StepType
	Steps map[StepType]Step
}

type Step struct {
	Question string
}

func New() Stepper {
	return Stepper{
		Type: CommitType,
		Steps: map[StepType]Step{
			CommitType:     Step{},
			BreakingChange: Step{},
			CommitScope:    Step{},
		},
	}
}

func (s *Stepper) Next() {
	s.Type++
	switch s.Type {
	case CommitType:
		// TODO: Do A
	case BreakingChange:
		// TODO: Do B
	case CommitScope:
		// TODO: Do C
	default:
		// TODO: Do D
	}
}

func (s *Stepper) Current() Step {
	return s.Steps[s.Type]
}
