package intf

import (
	"bytes"
	"fmt"
)

type Duck interface {
	Quack() string
}

type duck struct{}

// Duck interfaceを通じてdockのメモリを確保
func NewDuck() Duck {
	return &duck{}
}

func (d *duck) Quack() string {
	return "QUUAAAACCCCKK!!!"
}

type Farmer interface {
	Breed() string
}

type farmer struct {
	ducks []Duck
}

func NewFarmer(ducks ...Duck) Farmer {
	return &farmer{ducks}
}

func (f *farmer) Breed() string {
	var b bytes.Buffer
	for i, d := range f.ducks {
		fmt.Fprintf(&b, "#%d %s\n", i, d.Quack())
	}
	return b.String()
}

type selfCallingDuck struct {
	name string
}

func NewNamedDuck(name string) Duck {
	return &selfCallingDuck{name}
}

func (d *selfCallingDuck) Quack() string {
	return d.name
}
