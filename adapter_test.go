package go_design_patterns

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAdapter(t *testing.T) {
	Convey("Given an Desk which as a surface of 2", t, func() {
		desk := Desk{4, 2}
		Convey("Should be able to add it to a rectangle shape in a Shapes object", func() {
			rect := Rectangle{4, 3}
			shapes := Shapes{[]Shape{}}
			shapes.shapes = append(shapes.shapes, &rect)
			shapes.shapes = append(shapes.shapes, DeskAdapter{desk})
			So(shapes.Area(), ShouldAlmostEqual, 20)
		})
	})
}
