package go_design_patterns

type Desk struct {
	x, y float64
}

func (d Desk) Surface() float64 {
	return d.x * d.y
}

type DeskAdapter struct {
	desk Desk
}

func (d DeskAdapter) Area() float64 {
	return d.desk.Surface()
}
