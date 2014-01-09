package go_design_patterns

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestAreaCalculation(t *testing.T) {
	Convey("Given a 2 x 3 rectangle", t, func() {
		rect := Rectangle{2, 3}

		Convey("its area should be 9", func() {
			So(rect.Area(), ShouldAlmostEqual, 6)
		})
	})

	Convey("Given a circle with a diameter of 2", t, func() {
		circle := Circle{2}

		Convey("its area should be 6.28318", func() {
			So(circle.Area(), ShouldAlmostEqual, 6.28318)
		})
	})

	Convey("Given a square with a side of 3", t, func() {
		square := Square{3}

		Convey("its area should be 9", func() {
			So(square.Area(), ShouldAlmostEqual, 9)
		})
	})

	Convey("Given that I have a 2 x 3 Rectangle, a circle with a diameter of 2 and a 3 x 3 Square", t, func() {
		rect := Rectangle{2, 3}
		circle := Circle{2}
		square := Square{3}

		shapes := Shapes{make([]Shape, 0)}
		shapes.shapes = append(shapes.shapes, &rect)
		shapes.shapes = append(shapes.shapes, &circle)
		shapes.shapes = append(shapes.shapes, &square)
		Convey("The the area of all the shapes should be 21.28318", func() {
			So(shapes.Area(), ShouldAlmostEqual, 21.28318)
		})
	})
}
