package main

type Vogel struct {
	name string 
	alter int16
    
}

func (v Vogel) Name() string {
	return v.name
}
func (v Vogel) Alter() int16 {
	return v.alter
}
