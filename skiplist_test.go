package skiplist

import (
	"fmt"
	"math/rand"
	"testing"
)

type Int int

func (i Int) Less(other interface{}) bool {
	return i < other.(Int)
}

func TestInt(t *testing.T) {
	sl := New()
	if sl.Len() != 0 || sl.Front() != nil && sl.Back() != nil {
		t.Fatal()
	}

	testData := []Int{Int(1), Int(2), Int(3)}

	sl.Insert(testData[0])
	if sl.Len() != 1 || sl.Front().Value.(Int) != testData[0] || sl.Back().Value.(Int) != testData[0] {
		t.Fatal()
	}

	sl.Insert(testData[2])
	if sl.Len() != 2 || sl.Front().Value.(Int) != testData[0] || sl.Back().Value.(Int) != testData[2] {
		t.Fatal()
	}

	sl.Insert(testData[1])
	if sl.Len() != 3 || sl.Front().Value.(Int) != testData[0] || sl.Back().Value.(Int) != testData[2] {
		t.Fatal()
	}

	sl.Insert(Int(-999))
	sl.Insert(Int(-888))
	sl.Insert(Int(888))
	sl.Insert(Int(999))
	sl.Insert(Int(1000))

	expect := []Int{Int(-999), Int(-888), Int(1), Int(2), Int(3), Int(888), Int(999), Int(1000)}
	ret := make([]Int, 0)

	for e := sl.Front(); e != nil; e = e.Next() {
		ret = append(ret, e.Value.(Int))
	}
	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[i] {
			t.Fatal()
		}
	}

	e := sl.Find(Int(2))
	if e == nil || e.Value.(Int) != 2 {
		t.Fatal()
	}

	ret = make([]Int, 0)
	for ; e != nil; e = e.Next() {
		ret = append(ret, e.Value.(Int))
	}
	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[i+3] {
			t.Fatal()
		}
	}

	sl.Remove(sl.Find(Int(2)))
	sl.Delete(Int(888))
	sl.Delete(Int(1000))

	expect = []Int{Int(-999), Int(-888), Int(1), Int(3), Int(999)}
	ret = make([]Int, 0)

	for e := sl.Back(); e != nil; e = e.Prev() {
		ret = append(ret, e.Value.(Int))
	}

	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[len(ret)-i-1] {
			t.Fatal()
		}
	}

	if sl.Front().Value.(Int) != -999 {
		t.Fatal()
	}

	sl.Remove(sl.Front())
	if sl.Front().Value.(Int) != -888 || sl.Back().Value.(Int) != 999 {
		t.Fatal()
	}

	sl.Remove(sl.Back())
	if sl.Front().Value.(Int) != -888 || sl.Back().Value.(Int) != 3 {
		t.Fatal()
	}

	if e = sl.Insert(Int(2)); e.Value.(Int) != 2 {
		t.Fatal()
	}
	sl.Delete(Int(-888))

	if r := sl.Delete(Int(123)); r != nil {
		t.Fatal()
	}

	if sl.Len() != 3 {
		t.Fatal()
	}

	sl.Insert(Int(2))
	sl.Insert(Int(2))
	sl.Insert(Int(1))

	if e = sl.Find(Int(2)); e == nil {
		t.Fatal()
	}

	expect = []Int{Int(2), Int(2), Int(2), Int(3)}
	ret = make([]Int, 0)
	for ; e != nil; e = e.Next() {
		ret = append(ret, e.Value.(Int))
	}
	for i := 0; i < len(ret); i++ {
		if ret[i] != expect[i] {
			t.Fatal()
		}
	}

	sl2 := sl.Init()
	if sl2.Len() != 0 || sl.Len() != 0 || sl2.Front() != nil || sl2.Back() != nil || sl.Front() != nil || sl.Back() != nil {
		t.Fatal()
	}

	// for i := 0; i < 100; i++ {
	// 	sl.Insert(Int(rand.Intn(200)))
	// }
	// output(sl)
}

func BenchmarkIntInsertOrder(b *testing.B) {
	b.StopTimer()
	sl := New()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Insert(Int(i))
	}
}

func BenchmarkIntInsertRandom(b *testing.B) {
	b.StopTimer()
	sl := New()
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Insert(Int(rand.Int()))
	}
}

func BenchmarkIntDeleteOrder(b *testing.B) {
	b.StopTimer()
	sl := New()
	for i := 0; i < 1000000; i++ {
		sl.Insert(Int(i))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Delete(Int(i))
	}
}

func BenchmarkIntDeleteRandome(b *testing.B) {
	b.StopTimer()
	sl := New()
	for i := 0; i < 1000000; i++ {
		sl.Insert(Int(rand.Int()))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Delete(Int(rand.Int()))
	}
}

func BenchmarkIntFindOrder(b *testing.B) {
	b.StopTimer()
	sl := New()
	for i := 0; i < 1000000; i++ {
		sl.Insert(Int(i))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Find(Int(i))
	}
}

func BenchmarkIntFindRandom(b *testing.B) {
	b.StopTimer()
	sl := New()
	for i := 0; i < 1000000; i++ {
		sl.Insert(Int(rand.Int()))
	}
	b.StartTimer()

	for i := 0; i < b.N; i++ {
		sl.Find(Int(rand.Int()))
	}
}

func output(sl *SkipList) {
	var node *Element
	for i := 0; i < SKIPLIST_MAXLEVEL; i++ {
		fmt.Printf("LEVEL[%v]: ", i)
		node = sl.header.level[i].forward
		for node != nil {
			fmt.Printf("%v -> ", node.Value)
			node = node.level[i].forward
		}
		fmt.Println("NULL")
	}
}
