package printer

import (
	"fmt"

	"github.com/fatih/color"
)

type Printer interface {
	Green(format string, a ...any)
	HiRed(format string, a ...any)
	Println(a ...any)
	Red(format string, a ...any)
	Yellow(format string, a ...any)
}

type printer struct{}

func New(colorize bool) Printer {
	color.NoColor = !colorize

	return &printer{}
}

func (p *printer) Green(format string, a ...any) {
	color.Green(format, a...)
}

func (p *printer) HiRed(format string, a ...any) {
	color.HiRed(format, a...)
}

func (p *printer) Println(a ...any) {
	fmt.Println(a...)
}

func (p *printer) Red(format string, a ...any) {
	color.Red(format, a...)
}

func (p *printer) Yellow(format string, a ...any) {
	color.Yellow(format, a...)
}
