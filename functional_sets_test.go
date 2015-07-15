package functional_sets


import (
	"testing"

)

func TestSingleton(t *testing.T) {
	var s Set
	s = SingletonSet(34)

	if !Contains(s, 34) {
		t.Error("Singleton not created")
	}

	if Contains(s, 1) {
		t.Error("1 cannot be in a singleton of set 34")
	}

}

func TestUnion(t *testing.T) {
	s1 := SingletonSet(34)
	s2 := SingletonSet(73)

	if !Contains(s1, 34) || ! Contains(s2, 73) {
		t.Error("Singleton not created correctly")
	}

	un := Union(s1, s2)

	if !Contains(un, 34) || ! Contains(un, 73) {
		t.Error("Union not created correctly")
	}

	if Contains(un, 1) {
		t.Error("1 cannot be in the union")
	}

	s3 := SingletonSet(23)
	s4 := SingletonSet(25)

	un2 := Union(s3, s4)

	un3:= Union(un, un2)

	if Contains(un3, 1) {
		t.Error("1 cannot be in the union")
	}

	if !Contains(un3, 34) ||  !Contains(un3, 73) || !Contains(un3, 23) || !Contains(un3, 25) {
		t.Error("Union of Unions failed")
	}

}

func TestIntersection(t *testing.T) {
	s1 := SingletonSet(34)
	s2 := SingletonSet(73)

	if !Contains(s1, 34) || ! Contains(s2, 73) {
		t.Error("Singleton not created correctly")
	}

	in := Intersection(s1, s2)

	if Contains(in, 34) ||  Contains(in, 73) {
		t.Error("Intersection not created correctly")
	}

	if Contains(in, 1) {
		t.Error("1 cannot be in the union")
	}

    un := Union(s1, s2)
	s3 := SingletonSet(34)

	in2 := Intersection(un, s3)


	if Contains(in2, 73) {
		t.Error("73 cannot be in the union")
	}

	if !Contains(in2, 34)  {
		t.Error("Intersection failed")
	}

}

func TestDiff(t *testing.T) {
	s1 := SingletonSet(23)
	s2 := SingletonSet(34)
	s3 := SingletonSet(73)
	un := Union(s1, s2)
	un = Union(un, s3)
	df := Diff(un, s1)

	if Contains(df, 23) {
		t.Error("Diff failed. ")
	}

	if !Contains(df, 73) {
		t.Error("Diff failed. Expected 73")
	}

}

func TestFilter(t *testing.T) {
	//type pred func(int) bool
	//var p Pred
	p := func (x int) bool { return x > 30 }
    s1 := SingletonSet(23)
	s2 := SingletonSet(34)
	s3 := SingletonSet(73)
	un := Union(s1, s2)
	un = Union(un, s3)
	df := Filter(un, p)

	if Contains(df, 23) {
		t.Error("Filter failed. ")
	}

	if !Contains(df, 73) {
		t.Error("Filter failed. Expected 73")
	}

}
