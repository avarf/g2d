package graphics

import (
	"github.com/lucasepe/g2d/object"
	"github.com/lucasepe/g2d/typing"
)

// Arc draws a circular arc centered at `x, y` with a radius of `r`.
// The path starts at `angle1`, ends at `angle2`, and travels in the direction given by anticlockwise.
func Arc(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check("arc", args, typing.ExactArgs(5)); err != nil {
		return object.NewError(err.Error())
	}

	x, err := typing.ToFloat(args[0])
	if err != nil {
		return object.NewError("TypeError: arc() argument #1 `x` %s", err.Error())
	}

	y, err := typing.ToFloat(args[1])
	if err != nil {
		return object.NewError("TypeError: arc() argument #2 `y` %s", err.Error())
	}

	r, err := typing.ToFloat(args[2])
	if err != nil {
		return object.NewError("TypeError: arc() argument #3 `r` %s", err.Error())
	}

	rad1, err := typing.ToFloat(args[3])
	if err != nil {
		return object.NewError("TypeError: arc() argument #4 `degrees1` %s", err.Error())
	}

	rad2, err := typing.ToFloat(args[4])
	if err != nil {
		return object.NewError("TypeError: arc() argument #4 `degrees2` %s", err.Error())
	}

	env.Canvas().Value.DrawArc(x, y, r, rad1, rad2)
	return &object.Null{}
}

// Point draws a point at specified coordinates.
func Point(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check("point", args, typing.RangeOfArgs(2, 3)); err != nil {
		return object.NewError(err.Error())
	}

	x, err := typing.ToFloat(args[0])
	if err != nil {
		return object.NewError("TypeError: point() argument #1 %s", err.Error())
	}

	y, err := typing.ToFloat(args[1])
	if err != nil {
		return object.NewError("TypeError: point() argument #1 %s", err.Error())
	}

	// default point size
	r := env.Canvas().Value.LineWidth()
	if len(args) == 3 {
		r, _ = typing.ToFloat(args[2])
	}

	env.Canvas().Value.DrawPoint(x, y, r)
	return &object.Null{}
}

// Circle draws a circle centered at [x, y] coordinates and with the radius `r`.
func Circle(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check("circle", args, typing.ExactArgs(3)); err != nil {
		return object.NewError(err.Error())
	}

	x, err := typing.ToFloat(args[0])
	if err != nil {
		return object.NewError("TypeError: circle() argument #1 `x` %s", err.Error())
	}

	y, err := typing.ToFloat(args[1])
	if err != nil {
		return object.NewError("TypeError: circle() argument #2 `y` %s", err.Error())
	}

	r, err := typing.ToFloat(args[2])
	if err != nil {
		return object.NewError("TypeError: circle() argument #3 `r` %s", err.Error())
	}

	env.Canvas().Value.DrawCircle(x, y, r)

	return &object.Null{}
}

// Ellipse draws an ellipse centered at [x, y] coordinates and with the radii `rx` and `ry`.
func Ellipse(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check("ellipse", args, typing.RangeOfArgs(3, 4)); err != nil {
		return object.NewError(err.Error())
	}

	x, err := typing.ToFloat(args[0])
	if err != nil {
		return object.NewError("TypeError: ellipse() argument #1 `x` %s", err.Error())
	}

	y, err := typing.ToFloat(args[1])
	if err != nil {
		return object.NewError("TypeError: ellipse() argument #2 `y` %s", err.Error())
	}

	rx, err := typing.ToFloat(args[2])
	if err != nil {
		return object.NewError("TypeError: ellipse() argument #3 `rx` %s", err.Error())
	}

	if len(args) == 3 {
		env.Canvas().Value.DrawCircle(x, y, rx)
		return &object.Null{}
	}

	ry, err := typing.ToFloat(args[3])
	if err != nil {
		return object.NewError("TypeError: ellipse() argument #4 `ry` %s", err.Error())
	}

	env.Canvas().Value.DrawEllipse(x, y, rx, ry)
	return &object.Null{}
}

// Quad draws a quadrilateral, a four sided polygon.
func Quad(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check("quad", args, typing.ExactArgs(8)); err != nil {
		return object.NewError(err.Error())
	}

	x1, err := typing.ToFloat(args[0])
	if err != nil {
		return object.NewError("TypeError: quad() argument #1 `x1` %s", err.Error())
	}

	y1, err := typing.ToFloat(args[1])
	if err != nil {
		return object.NewError("TypeError: quad() argument #2 `y1` %s", err.Error())
	}

	x2, err := typing.ToFloat(args[2])
	if err != nil {
		return object.NewError("TypeError: quad() argument #1 `x2` %s", err.Error())
	}

	y2, err := typing.ToFloat(args[3])
	if err != nil {
		return object.NewError("TypeError: quad() argument #2 `y2` %s", err.Error())
	}

	x3, err := typing.ToFloat(args[4])
	if err != nil {
		return object.NewError("TypeError: quad() argument #1 `x3` %s", err.Error())
	}

	y3, err := typing.ToFloat(args[5])
	if err != nil {
		return object.NewError("TypeError: quad() argument #2 `y3` %s", err.Error())
	}

	x4, err := typing.ToFloat(args[6])
	if err != nil {
		return object.NewError("TypeError: quad() argument #1 `x4` %s", err.Error())
	}

	y4, err := typing.ToFloat(args[7])
	if err != nil {
		return object.NewError("TypeError: quad() argument #2 `y4` %s", err.Error())
	}

	env.Canvas().Value.DrawQuadrilateral(x1, y1, x2, y2, x3, y3, x4, y4)
	return &object.Null{}
}

// Triangle draws a triangle.
// The first two arguments specify the first point, the middle
// two arguments specify the second point, and the last two
// arguments specify the third point.
func Triangle(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check("triangle", args, typing.ExactArgs(6)); err != nil {
		return object.NewError(err.Error())
	}

	x1, err := typing.ToFloat(args[0])
	if err != nil {
		return object.NewError("TypeError: triangle() argument #1 `x1` %s", err.Error())
	}

	y1, err := typing.ToFloat(args[1])
	if err != nil {
		return object.NewError("TypeError: triangle() argument #2 `y1` %s", err.Error())
	}

	x2, err := typing.ToFloat(args[2])
	if err != nil {
		return object.NewError("TypeError: triangle() argument #1 `x2` %s", err.Error())
	}

	y2, err := typing.ToFloat(args[3])
	if err != nil {
		return object.NewError("TypeError: triangle() argument #2 `y2` %s", err.Error())
	}

	x3, err := typing.ToFloat(args[4])
	if err != nil {
		return object.NewError("TypeError: triangle() argument #1 `x3` %s", err.Error())
	}

	y3, err := typing.ToFloat(args[5])
	if err != nil {
		return object.NewError("TypeError: triangle() argument #2 `y3` %s", err.Error())
	}

	dc := env.Canvas().Value
	dc.NewSubPath()
	dc.MoveTo(x1, y1)
	dc.LineTo(x2, y2)
	dc.LineTo(x3, y3)
	dc.ClosePath()

	return &object.Null{}
}

// Line draws a line from point (x1, y1) to point (x2, y2)
func Line(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check("line", args, typing.ExactArgs(4)); err != nil {
		return object.NewError(err.Error())
	}

	x1, err := typing.ToFloat(args[0])
	if err != nil {
		return object.NewError("TypeError: line() argument #1 `x1` %s", err.Error())
	}

	y1, err := typing.ToFloat(args[1])
	if err != nil {
		return object.NewError("TypeError: line() argument #2 `y1` %s", err.Error())
	}

	x2, err := typing.ToFloat(args[2])
	if err != nil {
		return object.NewError("TypeError: line() argument #3 `x2` %s", err.Error())
	}

	y2, err := typing.ToFloat(args[3])
	if err != nil {
		return object.NewError("TypeError: line() argument #4 `y2` %s", err.Error())
	}

	env.Canvas().Value.DrawLine(x1, y1, x2, y2)
	return &object.Null{}
}

// Rect draws a (w x h) rectangle with upper left corner located at (x, y).
// rect(x, y, w, h, [r]) if radius `r` is specified, the rectangle will have rounded corners.
func Rect(env *object.Environment, args ...object.Object) object.Object {
	if err := typing.Check("rect", args); err != nil {
		return object.NewError(err.Error())
	}

	if len(args) < 4 {
		return object.NewError("TypeError: rect() takes at least 4 arguments, given: %d", len(args))
	}

	x, err := typing.ToFloat(args[0])
	if err != nil {
		return object.NewError("TypeError: rect() argument #1 %s", err.Error())
	}

	y, err := typing.ToFloat(args[1])
	if err != nil {
		return object.NewError("TypeError: rect() argument #2 %s", err.Error())
	}

	w, err := typing.ToFloat(args[2])
	if err != nil {
		return object.NewError("TypeError: rect() argument #3 %s", err.Error())
	}

	h, err := typing.ToFloat(args[3])
	if err != nil {
		return object.NewError("TypeError: rect() argument #4 %s", err.Error())
	}

	if len(args) == 5 {
		r, err := typing.ToFloat(args[4])
		if err != nil {
			return object.NewError("TypeError: rect() argument #5 %s", err.Error())
		}

		env.Canvas().Value.DrawRoundedRectangle(x, y, w, h, r)
		return &object.Null{}
	}

	// Let's dance!
	if len(args) == 8 {
		tl, err := typing.ToFloat(args[4])
		if err != nil {
			return object.NewError("TypeError: rect() argument #5 %s", err.Error())
		}
		tr, err := typing.ToFloat(args[5])
		if err != nil {
			return object.NewError("TypeError: rect() argument #6 %s", err.Error())
		}
		br, err := typing.ToFloat(args[6])
		if err != nil {
			return object.NewError("TypeError: rect() argument #7 %s", err.Error())
		}
		bl, err := typing.ToFloat(args[7])
		if err != nil {
			return object.NewError("TypeError: rect() argument #8 %s", err.Error())
		}

		env.Canvas().Value.DrawRoundedRectangleExtended(x, y, w, h, tl, tr, br, bl)
		return &object.Null{}
	}

	env.Canvas().Value.DrawRectangle(x, y, w, h)
	return &object.Null{}
}
