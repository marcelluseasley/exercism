package robotname

import (
	"math/rand"
	"time"
)



// Define the Robot type here.
type Robot struct {
	name string
	nameRecord 
}

func (r *Robot) Name() (string, error) {
	if r.name == "" {
		//return generateName()
	}
	return r.name, nil
}

func (r *Robot) Reset() {
	r.name = "";
}
