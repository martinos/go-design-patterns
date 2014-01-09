package go_design_patterns

import (
	// "fmt"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type Point struct {
	x, y float64
}

func (p *Point) Scale(scale float64) {
	p.x = p.x * scale
	p.y = p.y * scale
}
func TestMethods(t *testing.T) {
	Convey("Given a Point", t, func() {
		pt := Point{2, 3}
		pt.Scale(3)
		Convey("It should rescale to point", func() {
			So(pt.x, ShouldAlmostEqual, 6)
			So(pt.y, ShouldAlmostEqual, 9)
		})
	})
}
