/////////////////////////////////////////////////// Exercise: Slices

package main

import "golang.org/x/tour/pic"

func Pic(dx, dy int) [][]uint8 {

	//a slice of length dy
	p := make([][]uint8, dy)

	//each element of which is a slice of dx 8-bit unsigned integers
	for i := 0; i < dy; i++ {
		p[i] = make([]uint8, dx)

		for j := 0; j < dx; j++ {
			p[i][j] = uint8(i*i + j*j)
			// p[i][j] = uint8(i ^ j)
			// p[i][j] = uint8((i + j*i) / 2)
			// p[i][j] = uint8((i + j*i) / 2)
			// p[i][j] = uint8((i + j) / 2)
		}
	}
	return p
}

func main() {
	pic.Show(Pic)
}

/////////////////////////////////////////////////// Exercise: Maps

package main

import (
	"fmt"
)

type Vertex struct {
	Lat, Long float64
}

var map_locations = map[string]Vertex{
	"Ayoub's": {40.2345, 23.23456},
	"Chris's": {12.23456, -14.234},
}

func main() {
	fmt.Println(map_locations)

	map_cooler := make(map[string]string)

	map_cooler["Answer"] = "Ayoub is cooler"
	fmt.Println("The value: ", map_cooler["Answer"])

	map_cooler["Answer"] = "Chris is cooler"
	fmt.Println("The value: ", map_cooler["Answer"])

	delete(map_cooler, "Answer")
	fmt.Println("The value: ", map_cooler["Answer"])

	value, check := map_cooler["Answer"]
	fmt.Println("The Value:", value, "Present?", check)
	if !check {
		fmt.Println("Value deleted")
	}

	map_cooler["Answer"] = "Mia is cooler"
	fmt.Println("The Value:", map_cooler["Answer"])

	value2, check2 := map_cooler["Answer"]
	fmt.Println("The Value:", value2, "Is present?", check2)

	delete(map_cooler, "Answer")

	map_testMap := make(map[string]int)

	map_testMap["One"] = 1
	map_testMap["Two"] = 2

	fmt.Println(map_testMap)

	for key, value := range map_testMap {
		fmt.Println("Key: ", key, "Value: ", value)
	}

	map_testMap["Two"] += 10
	fmt.Println(map_testMap)
}

package main

import (
	"fmt"
	"strings"
)

//Returns a MAP of ints (count)
//Count == every word in string s
func WordCount(s string) map[string]int {

	count := make(map[string]int)
	words := strings.Fields(s)

	fmt.Println("INPUT: ", words)
	fmt.Println("Counting words...")

	for index, value := range words {
		count[value] += 1
		fmt.Println("Index: ", index, " Value: ", value)
	}

	fmt.Println("OUTPUT: ", count)
	return count
}

func main() {
	s := "Hey im so cool but are you cool"
	WordCount(s)

}

package main

import (
	"fmt"
	"strings"

	"golang.org/x/tour/wc"
)

//Returns a MAP of ints (count)
//Count == every word in string s
func WordCount(s string) map[string]int {

	count := make(map[string]int)
	words := strings.Fields(s)

	fmt.Println("INPUT: ", words)
	fmt.Println("Counting words...")

	for index, value := range words {
		count[value] += 1
		fmt.Println("Index: ", index, " Value: ", value)
	}

	fmt.Println("OUTPUT: ", count)
	return count
}
func main() {
	wc.Test(WordCount)
}

package main

import (
	"fmt"
)

func add(function func(int, int) int) int {
	return function(2, 2)
}

func main() {
	test := func(x, y int) int {
		return x + y
	}
	fmt.Println(test(8, 8))
	fmt.Println(add(test))

}

///////////////////////////////////////////////// Exercise: Fibonacci closure

package main

import (
	"fmt"
)

func closure_func() func(string) {
	fmt.Println("In closure function")
	return func(msg string) {
		fmt.Println(msg)
	}
}

func outer_function() func() int {
	i := 0
	fmt.Println("i: ", i)
	return func() int {
		i++
		fmt.Println("i: ", i)
		return i
	}
}

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci( {
	x := 0
	y := 1

	fmt.Println("X BEFORE: ", x, "\n", "Y BEFORE: ", y)

	for i := 0; i <= 21; i++ {
		z := x + y
		x++
		y++
		fmt.Println("X AFTER: ", x, "\n", "Y AFTER: ", y)
		fmt.Println("FIB NUMBER: ", z)
	}
}

func fibonacci(n int) int {
	if n == 0 {
		return 0
	} else if n == 1 {
		return 1
	} else {
		return fibonacci(n-1) + fibonacci(n-2)
	}
}

func fibonacci() func() int {
	f0 := 0
	f1 := 1
	return func() int {
		fn := f0
		f0 = f1
		f1 = f1 + fn
		return fn
	}
}

func main() {

	f := fibonacci()

	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}

	// for i := 1; i <= 5; i++ {
	// 	fmt.Println(fibonacci(i), " ")
	// }

	// fibonacci(22)
	// fmt.Println("In main function")
	// returned_func := closure_func()
	// returned_func("In returned func")

	// func(msg string) {
	// 	fmt.Println("In anon function")
	// 	fmt.Println(msg)
	// }("Im the Message")

	// nextInt := outer_function()
	// nextInt()

	// nextInt2 := outer_function()
	// nextInt2()

	// for i := 0; i <= 10; i++ {
	// 	fmt.Println(i)
	// 	nextInt()
	// 	nextInt2()
	// }

}

/////////////////////////////////////////////////// Exercise: Stringers

package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type square struct {
	width  float64
	height float64
}

func (c *circle) area() float64 {
	area := math.Pi * c.radius * c.radius
	fmt.Println("Circle Area:", area)
	return area
}

func (s *square) area() float64 {
	area := s.width * s.height
	fmt.Println("Square Area:", area)
	return area
}

func area(s shape) float64 {
	return s.area()
}

func main() {
	c := circle{
		5.00,
	}
	s := square{
		10.50,
		10.50,
	}

	shapes := []shape{&c, &s}

	fmt.Println(c.radius)
	fmt.Println(c.area())
	fmt.Println(s.height, " & ", s.width)
	fmt.Println(s.area())

	for _, shape := range shapes {
		shape.area()
	}

	area(&c)

}

package main

import "fmt"

type IPAddr [4]byte

// TODO: Add a "String() string" method to IPAddr.
func (address IPAddr) String() string {
	// output := ""
	// for i := 0; i < 4; i++ {
	// 	output += fmt.Sprintf("%d.", address[i])
	// }
	// fmt.Println(output)
	// return output
	return fmt.Sprintf("%d.%d.%d.%d", address[0], address[1], address[2], address[3])
}
func main() {
	hosts := map[string]IPAddr{
		"loopback":  {127, 0, 0, 1},
		"googleDNS": {8, 8, 8, 8},
	}
	for name, ip := range hosts {
		fmt.Printf("%v: %v\n", name, ip)
	}
}

/////////////////////////////////////////////////// Exercise: Errors
package main

import (
	"fmt"
	"time"
)

type MyError struct {
	When time.Time
	What string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s",
		e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}

package main

import "fmt"

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %v", float64(e))
}

func Sqrt(x float64) (float64, error) {
	z := 1.0
	if x >= 0 {
		for i := 1; i <= 10; i++ {
			z -= (z*z - x) / (2 * z)
			fmt.Println("Z Value: ", z)
		}
	} else {
		return 0, ErrNegativeSqrt(x)
	}
	return z, nil
}

func main() {
	fmt.Println(Sqrt(2))
	fmt.Println(Sqrt(-2))
}

/////////////////////////////////////////////////// Exercise: Readers

Implement a Reader type that emits an infinite stream of the ASCII character 'A'.
package main

import (
	// "fmt"

	"golang.org/x/tour/reader"
)

type MyReader struct{}

// TODO: Add a Read([]byte) (int, error) method to MyReader.
func (r MyReader) Read(b []byte) (int, error) {
	for i := range b {
		// fmt.Println("-------------------------------------------")
		// fmt.Printf("b[%v] = %s", i, b[i])
		// fmt.Println("-------------------------------------------")
		b[i] = 65
		// fmt.Println("///////////////////////////////////////////")
		// fmt.Printf("b[%v] = %s", i, b[i])
		// fmt.Println("///////////////////////////////////////////")

	}
	return len(b), nil
}

func main() {
	reader.Validate(MyReader{})
}

/////////////////////////////////////////////////// Exercise: Images

package main

import (
	"image"
	"image/color"

	"golang.org/x/tour/pic"
)

type Image struct {
	x1, y1 int
}

func (i Image) ColorModel() color.Model {
	return color.RGBAModel
}

func (i Image) Bounds() image.Rectangle {
	return image.Rect(0, 0, i.x1, i.y1)
}

func (i Image) At(x, y int) color.Color {
	output := uint8(x * y)
	return color.RGBA{output, output, 255, 255}
}

func main() {
	m := Image{256, 65}
	pic.ShowImage(m)
}

/////////////////////////////////////////////////// Exercise: Equivalent Binary Trees

package main

import (
	"fmt"
	"time"
)

func greetings(quit <-chan bool) <-chan string {
	out := make(chan string, 10)

	for i := 0; i < 10; i++ {
		select {
		case out <- "Hello World":
		case out <- "Howdy World":
		case out <- "Greetings World":
		case <-quit:
			out <- "No greetings for yo today!"
		default:
			fmt.Println("Default")
		}
	}

	close(out)

	return out
}

func main() {
	quit := make(chan bool, 10)
	quit <- true
	c := greetings(quit)
	for v := range c {
		fmt.Println(v)
	}
}

func compute(value int) {
	for i := 1; i < 5; i++ {
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func main() {
	go compute(5)
	go compute(5)
	time.Sleep(time.Second * 10)
// }

package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	} else if t.Left == nil {
		ch <- t.Value
		if t.Right != nil {
			Walk(t.Right, ch)
		}
		return
	} else {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	chan1 := make(chan int)
	chan2 := make(chan int)
	exit := make(chan bool)

	go run(t1, chan1)
	go run(t2, chan2)

	var output bool

	go func() {
		for i := range chan1 {
			if i == <-chan2 {
				output = true
			} else {
				output = false
			}
		}
		exit <- true
	}()
	<-exit
	return output
}

func run(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

func main() {
	// fmt.Println("Tour of Go Solution - Equivalent Binary Trees\n")
	// t1 := tree.New(5)
	// t2 := tree.New(5)
	// fmt.Println("Tree 1:", t1)
	// fmt.Println("Tree 2:", t2)
	// fmt.Println("Are they same? - ", Same(t1, t2))
	// fmt.Println("------")
	// t3 := tree.New(2)
	// t4 := tree.New(3)
	// fmt.Println("Tree 3:", t3)
	// fmt.Println("Tree 4:", t4)
	// fmt.Println("Are they same? - ", Same(t3, t4))
	// fmt.Println("")

	fmt.Println("Generating trees . . .")
	t1 := tree.New(10)
	t2 := tree.New(5)
	fmt.Println("-------TEST ONE------")
	fmt.Println("Tree 1: ", t1)
	fmt.Println("Tree 2: ", t2)
	fmt.Println("Comparing Tree 1 & Tree 2 . . .")
	fmt.Println("Test Result = ", Same(t1, t2))
	fmt.Println("---------END----------")

	t3 := tree.New(5)
	t4 := tree.New(5)
	fmt.Println("-------TEST ONE------")
	fmt.Println("Tree 3: ", t3)
	fmt.Println("Tree 4: ", t4)
	fmt.Println("Comparing Tree 3 & Tree 4 . . .")
	fmt.Println("Test Result = ", Same(t3, t4))
	fmt.Println("---------END----------")

}

/////////////////////////////////////////////////// Exercise: Web Crawler

package main

import (
	"fmt"
	"sync"
)

var (
	mutex   sync.Mutex
	balance int
)

func deposit(amount int, waitG *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Depositing %d to Account with current balance %d\n", amount, balance)
	balance += amount
	mutex.Unlock()
	waitG.Done()
}

func withdraw(amount int, waitG *sync.WaitGroup) {
	mutex.Lock()
	fmt.Printf("Withdrawing %d from Accoint with current balance %d\n", amount, balance)
	balance -= amount
	mutex.Unlock()
	waitG.Done()
}

func main() {
	fmt.Println("Hello World")

	balance = 1000

	var waitG sync.WaitGroup
	waitG.Add(3)
	go withdraw(700, &waitG)
	go withdraw(100000, &waitG)
	go deposit(500, &waitG)
	waitG.Wait()

	// time.Sleep(time.Second * 5)
	fmt.Printf("Balance %d\n", balance)
}

package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

func recursiveCrawl(url string, depth int, fetcher Fetcher, gt map[string]bool) {
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

package main

import (
	"fmt"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		Crawl(u, depth-1, fetcher)
	}
	return
}

func main() {
	Crawl("https://golang.org/", 4, fetcher)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{

		"The Go Programming Language",

		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

package main
import (
  "fmt"
  "sync"
)
type Fetcher interface {
  // Fetch returns the body of URL and
  // a slice of URLs found on that page.
  Fetch(url string) (body string, urls []string, err error)
}
var cache = make(Cache)
var wg sync.WaitGroup
var mux sync.Mutex
// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
  defer wg.Done()
  if cache.get(url) {
    fmt.Printf("xx Skipping: %s\n", url)
    return
  }
  fmt.Printf("** Crawling: %s\n", url)
  cache.set(url, true)

  if depth <= 0 {
    return
  }
  body, urls, err := fetcher.Fetch(url)
  if err != nil {
    fmt.Println(err)
    return
  }
  fmt.Printf("found: %s %q\n", url, body)
  for _, u := range urls {
    wg.Add(1)
    go Crawl(u, depth-1, fetcher)
  }
  return
}
func main() {
  wg.Add(1)
  Crawl("https://golang.org/", 4, fetcher)
  wg.Wait()
}
type Cache map[string]bool
func (ch Cache) get(key string) bool {
  mux.Lock()
  defer mux.Unlock()
  return cache[key]
}
func (ch Cache) set(key string, value bool) {
  mux.Lock()
  defer mux.Unlock()
  cache[key] = value
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var (
	mutex sync.Mutex
	waitG sync.WaitGroup
	cache := make(map[string]bool)
)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	mutex.Lock()
	if cache[url] == nil{
		fmt.Println("Duplicate: ", url)
		return
	}
	mutex.Unlock()
	fmt.Println("None duplicate: ", url)

	mutex.Lock()
	cache[url] = true
	mutex.Unlock()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		waitG.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
	return

	waitG.Done()
}

func main() {
	wg.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	wg.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

package main

import (
	"fmt"
	"sync"
	//"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var (
	mutex sync.Mutex
	waitG sync.WaitGroup
	cache = make(map[string]bool)
)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer waitG.Done()
	if cache[url] {
		mutex.Lock()
		fmt.Println("Duplicate: ", url)
		mutex.Unlock()
		return
	}
	mutex.Lock()
	fmt.Println("None-duplicate: ", url)
	cache[url] = true
	mutex.Unlock()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		waitG.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
	//waitG.Done()
	return
}

func main() {
	waitG.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	//time.Sleep(time.Second * 5)
	waitG.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

fatal error: all goroutines are asleep - deadlock!

goroutine 1 [semacquire]:
sync.runtime_Semacquire(0x0)
	/usr/local/go-faketime/src/runtime/sema.go:56 +0x25
sync.(*WaitGroup).Wait(0x49a702)
	/usr/local/go-faketime/src/sync/waitgroup.go:130 +0x71
main.main()
	/tmp/sandbox2165778371/prog.go:58 +0x55

package main

import (
	"fmt"
	"sync"
	//"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

var (
	mutex sync.Mutex
	waitG sync.WaitGroup
	cache = make(map[string]bool)
)

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer waitG.Done()
	fmt.Println("----------------START-----------------")
	for u, b := range cache {
		fmt.Printf("Cache: %s, %b\n", u, b)
	}
	fmt.Println("-----------------END------------------")
	if cache[url] {
		mutex.Lock()
		fmt.Println("Already crawled: ", url)
		mutex.Unlock()
		return
	}
	fmt.Println("Crawling: ", url)
	mutex.Lock()
	cache[url] = true
	mutex.Unlock()

	if depth <= 0 {
		return
	}
	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	for _, u := range urls {
		waitG.Add(1)
		go Crawl(u, depth-1, fetcher)
	}
	//waitG.Done()
	return
}

func main() {
	waitG.Add(1)
	Crawl("https://golang.org/", 4, fetcher)
	//time.Sleep(time.Second * 5)
	waitG.Wait()
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

// fatal error: all goroutines are asleep - deadlock!

// goroutine 1 [semacquire]:
// sync.runtime_Semacquire(0x0)
// 	/usr/local/go-faketime/src/runtime/sema.go:56 +0x25
// sync.(*WaitGroup).Wait(0x49a702)
// 	/usr/local/go-faketime/src/sync/waitgroup.go:130 +0x71
// main.main()
// 	/tmp/sandbox2165778371/prog.go:58 +0x55

found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/
not found: https://golang.org/cmd/
found: https://golang.org/pkg/fmt/ "Package fmt"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/pkg/os/ "Package os"
found: https://golang.org/ "The Go Programming Language"
found: https://golang.org/pkg/ "Packages"
not found: https://golang.org/cmd/

None-duplicate:  https://golang.org/pkg/os/
found: https://golang.org/pkg/os/ "Package os"
Duplicate:  https://golang.org/cmd/
Duplicate:  https://golang.org/
None-duplicate:  https://golang.org/pkg/fmt/
found: https://golang.org/pkg/fmt/ "Package fmt"
Duplicate:  https://golang.org/pkg/
Duplicate:  https://golang.org/
Duplicate:  https://golang.org/pkg/
Duplicate:  https://golang.org/

** Crawling: https://golang.org/
found: https://golang.org/ "The Go Programming Language"
** Crawling: https://golang.org/cmd/
not found: https://golang.org/cmd/
** Crawling: https://golang.org/pkg/
found: https://golang.org/pkg/ "Packages"
** Crawling: https://golang.org/pkg/os/
found: https://golang.org/pkg/os/ "Package os"
xx Skipping: https://golang.org/pkg/
** Crawling: https://golang.org/pkg/fmt/
found: https://golang.org/pkg/fmt/ "Package fmt"
xx Skipping: https://golang.org/
xx Skipping: https://golang.org/
xx Skipping: https://golang.org/pkg/
xx Skipping: https://golang.org/cmd/
xx Skipping: https://golang.org/

found: http://golang.org/ "The Go Programming Language"
not found: http://golang.org/cmd/
found: http://golang.org/pkg/ "Packages"
found: http://golang.org/pkg/os/ "Package os"
not found: http://golang.org/cmd/
found: http://golang.org/pkg/fmt/ "Package fmt"

found: https://golang.org/ "The Go Programming Language"
not found: https://golang.org/cmd/
found: https://golang.org/pkg/ "Packages"
found: https://golang.org/pkg/os/ "Package os"
found: https://golang.org/pkg/fmt/ "Package fmt"

----------------START-----------------
-----------------END------------------
Crawling:  https://golang.org/
found: https://golang.org/ "The Go Programming Language"
----------------START-----------------
Cache: https://golang.org/, %!b(bool=true)
-----------------END------------------
Crawling:  https://golang.org/cmd/
not found: https://golang.org/cmd/
----------------START-----------------
Cache: https://golang.org/, %!b(bool=true)
Cache: https://golang.org/cmd/, %!b(bool=true)
-----------------END------------------
Crawling:  https://golang.org/pkg/
found: https://golang.org/pkg/ "Packages"
----------------START-----------------
----------------START-----------------
Cache: https://golang.org/, %!b(bool=true)
Cache: https://golang.org/cmd/, %!b(bool=true)
Cache: https://golang.org/pkg/, %!b(bool=true)
-----------------END------------------
Already crawled:  https://golang.org/cmd/
----------------START-----------------
Cache: https://golang.org/, %!b(bool=true)
Cache: https://golang.org/cmd/, %!b(bool=true)
Cache: https://golang.org/pkg/, %!b(bool=true)
-----------------END------------------
Already crawled:  https://golang.org/
----------------START-----------------
Cache: https://golang.org/, %!b(bool=true)
Cache: https://golang.org/cmd/, %!b(bool=true)
Cache: https://golang.org/pkg/, %!b(bool=true)
-----------------END------------------
Crawling:  https://golang.org/pkg/fmt/
Cache: https://golang.org/, %!b(bool=true)
Cache: https://golang.org/cmd/, %!b(bool=true)
Cache: https://golang.org/pkg/, %!b(bool=true)
Cache: https://golang.org/pkg/fmt/, %!b(bool=true)
-----------------END------------------
found: https://golang.org/pkg/fmt/ "Package fmt"
----------------START-----------------
----------------START-----------------
Cache: https://golang.org/, %!b(bool=true)
Cache: https://golang.org/cmd/, %!b(bool=true)
Cache: https://golang.org/pkg/, %!b(bool=true)
Crawling:  https://golang.org/pkg/os/
found: https://golang.org/pkg/os/ "Package os"
----------------START-----------------
Cache: https://golang.org/, %!b(bool=true)
----------------START-----------------
Cache: https://golang.org/, %!b(bool=true)
Cache: https://golang.org/cmd/, %!b(bool=true)
Cache: https://golang.org/pkg/, %!b(bool=true)
Cache: https://golang.org/pkg/fmt/, %!b(bool=true)
Cache: https://golang.org/pkg/os/, %!b(bool=true)
-----------------END------------------
Already crawled:  https://golang.org/
Cache: https://golang.org/, %!b(bool=true)
Cache: https://golang.org/cmd/, %!b(bool=true)
Cache: https://golang.org/pkg/, %!b(bool=true)
Cache: https://golang.org/pkg/fmt/, %!b(bool=true)
-----------------END------------------
Already crawled:  https://golang.org/pkg/
Cache: https://golang.org/cmd/, %!b(bool=true)
Cache: https://golang.org/pkg/, %!b(bool=true)
Cache: https://golang.org/pkg/fmt/, %!b(bool=true)
Cache: https://golang.org/pkg/os/, %!b(bool=true)
-----------------END------------------
Already crawled:  https://golang.org/pkg/
Cache: https://golang.org/pkg/fmt/, %!b(bool=true)
Cache: https://golang.org/pkg/os/, %!b(bool=true)
-----------------END------------------
Already crawled:  https://golang.org/

package main

import (
	"fmt"
	"sync"
	"time"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}
// SafeUrlMap is safe to use concurrently.
type SafeUrlMap struct {
    v   map[string]string
    mux sync.Mutex
}

func (c *SafeUrlMap) Set(key string, body string) {
    c.mux.Lock()
    // Lock so only one goroutine at a time can access the map c.v.
    c.v[key] = body
    c.mux.Unlock()
}

// Value returns mapped value for the given key.
func (c *SafeUrlMap) Value(key string) (string, bool) {
    c.mux.Lock()
    // Lock so only one goroutine at a time can access the map c.v.
    defer c.mux.Unlock()
    val, ok := c.v[key]
    return val, ok
}
func Crawl(url string, depth int, fetcher Fetcher) {
    mux.Lock()
    defer mux.Unlock()
    if depth <= 0 || IsVisited(url) {
        return
    }
    visit[url] = true
    body, urls, err := fetcher.Fetch(url)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Printf("found: %s %q\n", url, body)
    for _, u := range urls {
        //
        go Crawl(u, depth-1, fetcher)
    }
    return
}

func IsVisited(s string) bool {
    _, ok := visit[s]
    return ok
}

var mux sync.Mutex

var visit = make(map[string]bool)

func main() {
    Crawl("https://golang.org/", 4, fetcher)
    time.Sleep(time.Second)
}

// fakeFetcher is Fetcher that returns canned results.
type fakeFetcher map[string]*fakeResult

type fakeResult struct {
	body string
	urls []string
}

func (f fakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = fakeFetcher{
	"https://golang.org/": &fakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &fakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &fakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &fakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}
