package main

import (
	"github.com/signintech/gopdf"
)

type Rectangle struct {
	fillColor   color
	strokeColor color
	tlPoint     gopdf.Point // top left point
	trPoint     gopdf.Point // top right point
	blPoint     gopdf.Point // btm left point
	brPoint     gopdf.Point // btm right point
}

type CurvedRectangle struct {
	rectangle   Rectangle
	curveLength float64
}

func (r *Rectangle) DrawRectangle(pdf *gopdf.GoPdf) {
	pdf.SetFillColor(r.fillColor.r, r.fillColor.g, r.fillColor.b)
	pdf.SetStrokeColor(r.strokeColor.r, r.strokeColor.g, r.strokeColor.b)
	pdf.Polygon([]gopdf.Point{
		{X: r.tlPoint.X, Y: r.tlPoint.Y},
		{X: r.trPoint.X, Y: r.trPoint.Y},
		{X: r.brPoint.X, Y: r.brPoint.Y},
		{X: r.blPoint.X, Y: r.blPoint.Y}}, "DF")
}

func (cr *CurvedRectangle) DrawLeftCurved(pdf *gopdf.GoPdf) {
	pdf.SetLineWidth(1.0)
	cr.rectangle.DrawRectangle(pdf)
	pdf.Curve(
		cr.rectangle.tlPoint.X, cr.rectangle.tlPoint.Y, // start point
		cr.rectangle.tlPoint.X-cr.curveLength, cr.rectangle.tlPoint.Y, // control point 1
		cr.rectangle.blPoint.X-cr.curveLength, cr.rectangle.blPoint.Y, // control point 2
		cr.rectangle.blPoint.X, cr.rectangle.blPoint.Y, // end point
		"DF")
}

func (cr *CurvedRectangle) DrawRightCurved(pdf *gopdf.GoPdf) {
	pdf.SetLineWidth(1.0)
	cr.rectangle.DrawRectangle(pdf)
	pdf.Curve(
		cr.rectangle.tlPoint.X, cr.rectangle.tlPoint.Y, // start point
		cr.rectangle.tlPoint.X+cr.curveLength, cr.rectangle.tlPoint.Y, // control point 1
		cr.rectangle.blPoint.X+cr.curveLength, cr.rectangle.blPoint.Y, // control point 2
		cr.rectangle.blPoint.X, cr.rectangle.blPoint.Y, // end point
		"DF")
}
