package quantum

import (
	"fmt"
	"math"
	"math/cmplx"
	"strconv"
	"strings"
)

// A Gate is a generic object that modifies a quantum
// computer in some primitive way.
type Gate interface {
	fmt.Stringer
	Apply(c Computer)
	Invert(c Computer)
}

type Circuit []Gate

func (c Circuit) String() string {
	var parts []string
	for _, g := range c {
		parts = append(parts, g.String())
	}
	return strings.Join(parts, " ")
}

func (c Circuit) Apply(comp Computer) {
	for _, g := range c {
		g.Apply(comp)
	}
}

func (c Circuit) Invert(comp Computer) {
	for i := len(c) - 1; i >= 0; i-- {
		c[i].Invert(comp)
	}
}

type HGate struct {
	Bit int
}

func (h *HGate) String() string {
	return "H(" + strconv.Itoa(h.Bit) + ")"
}

func (h *HGate) Apply(c Computer) {
	Hadamard(c, h.Bit)
}

func (h *HGate) Invert(c Computer) {
	Hadamard(c, h.Bit)
}

type TGate struct {
	Bit       int
	Conjugate bool
}

func (t *TGate) String() string {
	conjStr := ""
	if t.Conjugate {
		conjStr = "*"
	}
	return "T" + conjStr + "(" + strconv.Itoa(t.Bit) + ")"
}

func (t *TGate) Apply(c Computer) {
	if t.Conjugate {
		(&TGate{Bit: t.Bit}).Invert(c)
	} else {
		c.Unitary(t.Bit, 1, 0, 0, cmplx.Exp(complex(0, math.Pi/4)))
	}
}

func (t *TGate) Invert(c Computer) {
	if t.Conjugate {
		(&TGate{Bit: t.Bit}).Apply(c)
	} else {
		c.Unitary(t.Bit, 1, 0, 0, cmplx.Exp(complex(0, -math.Pi/4)))
	}
}

type XGate struct {
	Bit int
}

func (x *XGate) String() string {
	return "X(" + strconv.Itoa(x.Bit) + ")"
}

func (x *XGate) Apply(c Computer) {
	X(c, x.Bit)
}

func (x *XGate) Invert(c Computer) {
	X(c, x.Bit)
}

type YGate struct {
	Bit int
}

func (y *YGate) String() string {
	return "Y(" + strconv.Itoa(y.Bit) + ")"
}

func (y *YGate) Apply(c Computer) {
	Y(c, y.Bit)
}

func (y *YGate) Invert(c Computer) {
	for i := 0; i < 4; i++ {
		Y(c, y.Bit)
	}
}

type ZGate struct {
	Bit int
}

func (z *ZGate) String() string {
	return "Z(" + strconv.Itoa(z.Bit) + ")"
}

func (z *ZGate) Apply(c Computer) {
	Z(c, z.Bit)
}

func (z *ZGate) Invert(c Computer) {
	Z(c, z.Bit)
}

type CNotGate struct {
	Control int
	Target  int
}

func (c *CNotGate) String() string {
	return fmt.Sprintf("CNot(%d, %d)", c.Control, c.Target)
}

func (c *CNotGate) Apply(comp Computer) {
	comp.CNot(c.Control, c.Target)
}

func (c *CNotGate) Invert(comp Computer) {
	comp.CNot(c.Control, c.Target)
}
