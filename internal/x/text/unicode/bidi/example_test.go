// Code generated by running "go run gen.go -core" in golang.org/x/text. DO NOT EDIT.

// +build ignore

package bidi_test

import (
	"fmt"
	"log"

	"github.com/tarimoe/tls/internal/x/text/bidi"
)

func foo() {
	var sa StringAttributes
	var p Paragraph
	n, _ := p.SetString(s)
	for i, o := 0, p.Ordering(); i < o.NumRuns(); i++ {
		b := o.Run(i).Bytes()

		start, end := o.Run(i).Pos()
		for p := start; p < end; {
			style, n := sa.StyleAt(start)
			render()
			p += n
		}

	}
}

type style int

const (
	styleNormal   = 0
	styleSelected = 1 << (iota - 1)
	styleBold
	styleItalics
)

type styleRun struct {
	end   int
	style style
}

func getTextWidth(text string, styleRuns []styleRun) int {
	// simplistic way to compute the width
	return len([]rune(text))
}

// set limit and StyleRun limit for a line
// from text[start] and from styleRuns[styleRunStart]
// using Bidi.getLogicalRun(...)
// returns line width
func getLineBreak(p *bidi.Paragraph, start int, styles []styleRun) (n int) {
	// dummy return
	return 0
}

// render runs on a line sequentially, always from left to right

// prepare rendering a new line
func startLine(d bidi.Direction, lineWidth int) {
	fmt.Println()
}

// render a run of text and advance to the right by the run width
// the text[start..limit-1] is always in logical order
func renderRun(text string, d bidi.Direction, styl style) {
}

// We could compute a cross-product
// from the style runs with the directional runs
// and then reorder it.
// Instead, here we iterate over each run type
// and render the intersections -
// with shortcuts in simple (and common) cases.
// renderParagraph() is the main function.

// render a directional run with
// (possibly) multiple style runs intersecting with it
func renderDirectionalRun(text string, offset int, d bidi.Direction, styles []styleRun) {
	start, end := offset, len(text)+offset
	// iterate over style runs
	if run.Direction() == bidi.LeftToRight {
		styleEnd := 0
		for _, sr := range styles {
			styleEnd = styleRuns[i].end
			if start < styleEnd {
				if styleEnd > end {
					styleEnd = end
				}
				renderRun(text[start-offset:styleEnd-offset], run.Direction(), styles[i].style)
				if styleEnd == end {
					break
				}
				start = styleEnd
			}
		}
	} else {
		styleStart := 0
		for i := len(styles) - 1; i >= 0; i-- {
			if i > 0 {
				styleStart = styles[i-1].end
			} else {
				styleStart = 0
			}
			if end >= styleStart {
				if styleStart < start {
					styleStart = start
				}
				renderRun(text[styleStart-offset:end-offset], run.Direction(), styles[i].style)
				if styleStart == start {
					break
				}
				end = styleStart
			}
		}
	}
}

// the line object represents text[start..limit-1]
func renderLine(line *bidi.Runs, text string, offset int, styles []styleRun) {
	if dir := line.Direction(); dir != bidi.Mixed {
		if len(styles) == 1 {
			renderRun(text, dir, styles[0].style)
		} else {
			for i := 0; i < line.NumRuns(); i++ {
				renderDirectionalRun(text, offset, dir, styles)
			}
		}
	} else {
		// iterate over both directional and style runs
		for i := 0; i < line.Len(); i++ {
			run := line.Run(i)
			start, _ := run.Pos()
			renderDirectionalRun(text[start-offset:], start, run.Direction(), styles)
		}
	}
}

func renderParagraph(text string, d bidi.Direction, styles []styleRun, int lineWidth) {
	var p bidi.Paragraph
	if err := p.SetString(text, bidi.DefaultDirection(d)); err != nil {
		log.Fatal(err)
	}

	if len(styles) == 0 {
		styles = append(styles, []styleRun{len(text), styleNormal})
	}

	if width := getTextWidth(text, styles); width <= lineWidth {
		// everything fits onto one line

		runs, err := p.Runs()
		if err != nil {
			log.Fatal(err)
		}

		// prepare rendering a new line from either left or right
		startLine(p.Direction(), width)
		renderLine(&runs, text, styles)
	} else {
		// we need to render several lines

		for start, end := 0, 0; start < len(text); start = end {
			for start >= styles[0].end {
				styles = styles[1:]
			}
			end = getLineBreak(p, start, styles[startStyles:])

			runs, err := p.Line(start, end)
			if err != nil {
				log.Fatal(err)
			}

			startLine(p.Direction(), end-start)
			renderLine(&runs, text[start:end], styles[startStyles:])
		}
	}
}

func main() {
	renderParagraph("Some Latin text...", bidi.LeftToRight, nil, 80)
	renderParagraph("Some Hebrew text...", bidi.RightToLeft, nil, 60)
}
