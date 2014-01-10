package go_design_patterns

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type ObsMock struct {
	called  bool
	Subject *Subject
}

func (o *ObsMock) Update() {
	o.called = true
}

func (o *ObsMock) GetSubjState() int {
	return o.Subject.GetState()
}

func (o *ObsMock) SetSubject(subj *Subject) {
	o.Subject = subj
}

func TestObserver(t *testing.T) {
	Convey("Given an Subject", t, func() {
		subj := Subject{State: 1}
		Convey("Observers should be attached", func() {
			obs1 := ObsMock{}
			obs2 := ObsMock{}

			subj.Attach(&obs1)
			subj.Attach(&obs2)

			Convey("Each observer should be updated when changes happen in the subject", func() {
				subj.Notify()
				So(obs1.called, ShouldBeTrue)
				So(obs2.called, ShouldBeTrue)
			})

			Convey("Each observer should be able to get the subject state on notification", func() {
				So(obs1.GetSubjState(), ShouldEqual, 1)
				So(obs2.GetSubjState(), ShouldEqual, 1)
			})
		})
	})
}
