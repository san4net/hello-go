package main

import (
	"bytes"
	"fmt"
	"io"
	"sort"
	"time"
)

func main() {

	dd := [...]int{1, 2, 3}
	for _, i := range dd {
		fmt.Println(i)
	}

	var w io.Writer
	w = new(bytes.Buffer)
	doWrite(w)
	//w = &os.Stdout
	var f1 Foo
	var f2 *Foo = &Foo{}
	doFoo(&f1)
	doFoo(f2)
	var tt = []*Track{
		{"Go", "Delliah", "Water", 2012, time.Duration(23)},
		{"Go", "Delliah", "Blue", 2011, time.Duration(23)},
		{"Go", "Delliah", "sky", 2010, time.Duration(23)},
	}
	sort.Sort(Album{
		track: tt,
		less: func(x, y *Track) bool {
			if x.Title != y.Title {
				return x.Title < y.Title
			}
			if x.Year != y.Year {
				return x.Year < y.Year
			}
			if x.Length != y.Length {
				return x.Length < y.Length
			}
			return false
		},
	})
	fmt.Printf("sorted album %v \n", tt)

	for _, i := range tt {
		fmt.Printf("%s \n", i)
	}
}

type Fooer interface {
	Dummy()
}

type Foo struct{}

func (f *Foo) Dummy() {
}

func doFoo(f Fooer) {
	fmt.Printf("%v and type %T and %v \n", f, f, f == nil)
}

func doWrite(w io.Writer) {
	d := time.Duration(3)

	fmt.Println(d)
	fmt.Printf("%v and type %T \n", w, w)
	if w != nil {
		w.Write([]byte("hellow\n"))
		fmt.Println("done")
	}
}

type Track struct {
	Title  string
	Artist string
	Album  string
	Year   int
	Length time.Duration
}

type Album struct {
	track []*Track
	less  func(x, y *Track) bool
}

func (a Album) Len() int {
	return len(a.track)
}

func (a Album) Less(i, j int) bool {
	return a.less(a.track[i], a.track[j])
}

func (a Album) Swap(i, j int) {
	a.track[i], a.track[j] = a.track[j], a.track[i]
}
