package go_design_patterns

type Subject struct {
	Observers []*Observer
	State     int
}

func (s *Subject) Attach(o Observer) {
	o.SetSubject(s)
	s.Observers = append(s.Observers, &o)
}

func (s *Subject) Notify() {
	for _, obs := range s.Observers {
		(*obs).Update()
	}
}

func (s *Subject) GetState() int {
	return s.State
}

// Observer

type Observer interface {
	Update()
	SetSubject(*Subject)
	GetSubjState() int
}
