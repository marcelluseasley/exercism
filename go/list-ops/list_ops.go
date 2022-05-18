package listops

type IntList []int 

type binFunc func(int, int) int 

type predFunc func(int) bool 

type unaryFunc func(x int) int

func (l IntList) Length() int {
	cont :=0

	for cont = range l {
		cont++
	}
	return cont
}

func (l IntList) Reverse() IntList {
	rl := make(IntList, len(l))
	last := len(l) - 1
	for i := last; i >= 0; i-- {
		rl[last-i] = l[i]
	}
	return rl
}

func (l IntList) Append(l2 IntList) IntList {
	return append(l, l2...)
}

func (l IntList) Concat(l2 []IntList) IntList {
	for _, v := range l2 {
		l = l.Append(v)
	}
	return l
}

func (l IntList) Filter(f predFunc) IntList {
	rl := IntList{}
	for _, v := range l {
		if f(v) {
			rl = append(rl, v)
		}
	}
	return rl
}

func (l IntList) Map(f unaryFunc) IntList {
	rl := make(IntList, len(l))
	for i, v := range l {
		rl[i] = f(v)
	}
	return rl
}

func (l IntList) Foldr(f binFunc, acc int) int {
	for i := len(l) - 1; i >= 0; i-- {
		acc = f(l[i], acc)
	}
	return acc
}

func (l IntList) Foldl(f binFunc, acc int) int {
	for _, v := range l {
		acc = f(acc, v)
	}
	return acc
}