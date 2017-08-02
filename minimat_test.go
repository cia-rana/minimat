package minimat

import (
	"os"
	"testing"
)

var a, b *Mat

func TestMain(m *testing.M) {
	a, _ = NewMat(3, 3)
	a.Set(0, 0, float64(-4))
	a.Set(0, 1, float64(-3))
	a.Set(0, 2, float64(-2))
	a.Set(1, 0, float64(-1))
	a.Set(1, 1, float64(0))
	a.Set(1, 2, float64(1))
	a.Set(2, 0, float64(2))
	a.Set(2, 1, float64(3))
	a.Set(2, 2, float64(4))

	b, _ = NewMat(3, 3)
	b.Set(0, 0, float64(4))
	b.Set(0, 1, float64(3))
	b.Set(0, 2, float64(3))
	b.Set(1, 0, float64(4))
	b.Set(1, 1, float64(0))
	b.Set(1, 2, float64(1))
	b.Set(2, 0, float64(2))
	b.Set(2, 1, float64(3))
	b.Set(2, 2, float64(4))

	code := m.Run()
	os.Exit(code)
}

func TestRow(t *testing.T) {
	expect := 3
	if actual := a.Row(); expect != actual {
		t.Errorf("expect: %d, actual: %d", expect, actual)
	}
}

func TestCol(t *testing.T) {
	expect := 3
	if actual := a.Col(); expect != actual {
		t.Errorf("expect: %d, actual: %d", expect, actual)
	}
}

func TestGet(t *testing.T) {
	testCase := []struct {
		i      int
		j      int
		expect float64
	}{
		{0, 0, float64(-4)},
		{1, 1, float64(0)},
		{2, 2, float64(4)},
	}
	for _, tc := range testCase {
		actual, _ := a.Get(tc.i, tc.j)
		if actual != tc.expect {
			t.Errorf("expect: %f, actual: %f", tc.expect, actual)
		}
	}
}

func TestAdd(t *testing.T) {
	expect, _ := NewMat(3, 3)
	expect.Set(0, 0, float64(0))
	expect.Set(0, 1, float64(0))
	expect.Set(0, 2, float64(1))
	expect.Set(1, 0, float64(3))
	expect.Set(1, 1, float64(0))
	expect.Set(1, 2, float64(2))
	expect.Set(2, 0, float64(4))
	expect.Set(2, 1, float64(6))
	expect.Set(2, 2, float64(8))

	actual, _ := a.Add(b)
	if !expect.Equal(actual) {
		t.Errorf("expect: #%v, actual: #%v\n", expect, actual)
	}
}

func TestSub(t *testing.T) {
	expect, _ := NewMat(3, 3)
	expect.Set(0, 0, float64(-8))
	expect.Set(0, 1, float64(-6))
	expect.Set(0, 2, float64(-5))
	expect.Set(1, 0, float64(-5))
	expect.Set(1, 1, float64(0))
	expect.Set(1, 2, float64(0))
	expect.Set(2, 0, float64(0))
	expect.Set(2, 1, float64(0))
	expect.Set(2, 2, float64(0))

	actual, _ := a.Sub(b)
	if !expect.Equal(actual) {
		t.Errorf("expect: #%v, actual: #%v\n", expect, actual)
	}
}

func TestDot(t *testing.T) {
	expect, _ := NewMat(3, 3)
	expect.Set(0, 0, float64(-32))
	expect.Set(0, 1, float64(-18))
	expect.Set(0, 2, float64(-23))
	expect.Set(1, 0, float64(-2))
	expect.Set(1, 1, float64(0))
	expect.Set(1, 2, float64(1))
	expect.Set(2, 0, float64(28))
	expect.Set(2, 1, float64(18))
	expect.Set(2, 2, float64(25))

	actual, _ := a.Dot(b)
	if !expect.Equal(actual) {
		t.Errorf("expect: #%v, actual: #%v\n", expect, actual)
	}
}

func TestPow(t *testing.T) {
	expect, _ := NewMat(3, 3)
	expect.Set(0, 0, float64(770166310))
	expect.Set(0, 1, float64(510791457))
	expect.Set(0, 2, float64(649957935))
	expect.Set(1, 0, float64(449111146))
	expect.Set(1, 1, float64(297860751))
	expect.Set(1, 2, float64(379013536))
	expect.Set(2, 0, float64(665249420))
	expect.Set(2, 1, float64(441208218))
	expect.Set(2, 2, float64(561416593))

	actual, _ := b.Pow(10)
	if !expect.Equal(actual) {
		t.Errorf("expect: #%v, actual: #%v\n", expect, actual)
	}
}
