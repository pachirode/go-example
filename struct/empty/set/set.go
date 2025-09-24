package main

type Set map[string]struct{}

func (s Set) Add(ele string) {
	s[ele] = struct{}{}
}

func (s Set) Remove(ele string) {
	delete(s, ele)
}

func (s Set) Contain(ele string) bool {
	_, exists := s[ele]
	return exists
}

func (s Set) Size() int {
	return len(s)
}

func (s Set) String() string {
	res := ""
	for e := range s {
		res += e + ", "
	}
	return res
}
