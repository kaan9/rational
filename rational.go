package main

import "fmt"

type Rat struct {
	factors map[uint] int
	neg bool
}

func newRat(num int) *Rat {
	rat := new(Rat)
	rat.factors = make(map[uint] int)
	if num < 0 {
		rat.neg = true
		num *= -1
	}
	factor := 2
	for ; num > 1; factor++ {
		for ; num % factor == 0; num /= factor {
			rat.factors[uint(factor)]++
		}
	}
	return rat
}

// numerator of a, always returns positive sign
func (a *Rat) numer() *Rat {
	c := new(Rat)
	c.factors = make(map[uint] int)
	for k, v := range a.factors {
		if v > 0 {
			c.factors[k] = v
		}
	}
	return c
}

// denominator of a, always returns positive sign
func (a *Rat) denom() *Rat {
	c := new(Rat)
	c.factors = make(map[uint] int)
	for k, v := range a.factors {
		if v < 0 {
			c.factors[k] = v
		}
	}
	return c
}


func (a *Rat) add(b *Rat) *Rat {
	a_num := a.numer()
	a_num.neg = a.neg
	a_den := a.denom()
	b_num := b.numer()
	b_num.neg = b.neg
	b_den := b.denom()
	c_num := newRat(a_num.times(b_den).to_int() + b_num.times(a_den).to_int())
	c_den := b_den.times(a_den)
	return c_num.times(c_den.inv())
}

func (a *Rat) times(b *Rat) *Rat {
	c := new(Rat)
	c.factors = make(map[uint] int)
	c.neg = (a.neg || b.neg) && !(a.neg && b.neg)
	for k, v := range a.factors {
		c.factors[k] += v
	}
	for k, v := range b.factors {
		c.factors[k] += v
	}
	return c
}

func (a *Rat) div(b *Rat) *Rat {
	c := new(Rat)
	c.factors = make(map[uint] int)
	c.neg = (a.neg || b.neg) && !(a.neg && b.neg)
	for k, v := range a.factors {
		c.factors[k] += v
	}
	for k, v := range b.factors {
		c.factors[k] -= v
	}
	return c
}


func (r *Rat) inv() *Rat {
	c := new(Rat)
	c.factors = make(map[uint] int)
	c.neg = r.neg
	for k, v := range r.factors {
		c.factors[k] = -v
	}
	return c
}

func (r *Rat) pow(n int) *Rat {
	p := new(Rat)
	p.factors = make(map[uint] int)
	p.neg = r.neg && n%2 == 1
	for k, v := range r.factors {
		p.factors[k] = v * n
	}
	return p
}

func (r *Rat) to_int() int {
	num := 1
	if r.neg {
		num = -1
	}
	for k, v := range r.factors {
		for ; v > 0; v-- {
			num *= int(k)
		}
	}
	return num
}

func main() {
	var num, num2, p int
	fmt.Printf("num1 num2 pow: ")
	fmt.Scanf("%d %d %d", &num, &num2, &p)
	r := newRat(num)
	r2 := newRat(num2)
	fmt.Printf("times: %v\n", r.times(r2))
	fmt.Printf("div: %v\n", r.div(r2))
	fmt.Printf("inv: %v\n", r.inv())
	fmt.Printf("pow: %v\n", r.pow(p))
	fmt.Printf("add: %v\n", r.add(r2))
}
