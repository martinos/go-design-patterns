package go_design_patterns

type Subject struct {
	Observers []*Observer
	state     int
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

func (s *Subject) State() int {
	return s.state
}

// Observer

type Observer interface {
	Update()
	SetSubject(*Subject)
	SubjState() int
}
