package main

import (
	"bufio"
	"context"
	"crypto/md5"
	"crypto/sha256"
	"embed"
	"encoding/base64"
	"encoding/json"
	"example.com/bill"
	"fmt"
	"io"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"
	"sync"
	"sync/atomic"
	"text/template"
	"time"
)

func main() {
	// playing with strcuts and their functions
	// example1()

	// playign with arrays & slices
	// example2()

	// playing with interfaces
	// example3()

	// playing with maps
	// example4()

	// playing with concurency
	// example5()

	// implementation of setTimeout (js)
	// ---- Way 1 ----
	// setTimeout1()
	// ---- Way 2 ----
	// setTimeout2()

	// Go Tickers: implementation of setInterval (js) + (Sleep Vs Ticker)
	//example6()

	// playing with context
	// example7()

	// playing with context (advanced) start the server fisrt (look inside it)
	//example8()

	// playing with the singleton pattern
	// example9()

	// playing with atomic counters
	// example10()

	// playing with mutex
	// example11()

	// playing with sorting
	// example12()

	// playing with panic() & recover(): a.k.a try catch
	// example13()

	// playing with strings
	// example14()

	// playing with templates
	// example15()

	// playing with regular expressions
	// example16()

	// playing with json
	// example17()

	// playing with url parsing
	// example18()

	// playing with encryption
	// example19()

	// playing with base64 encroding
	// example20()

	// playing with reading files
	// example21()

	// playing with writing files
	// example22()

	// playing with line filters (aka pipe program)
	// echo "Hi, how are you?" | go run main.go
	// cat files/write1.txt | go run main.go
	// example23()

	// playing with File Path (crossplatform)
	// example24()

	// playing with temporary files and dirs
	// example25()

	// playing with Embed Directive
	example26()
}

// --------------------------- //
func example1() {

	perfum := bill.Product{ID: 1, Title: "Raghba Lattafa", Price: 120.00}
	invoice := bill.Invoice{ID: 1, Product: perfum, Date: time.Now()}
	fmt.Println(invoice)
	fmt.Println(invoice.IsOutdated())
	invoice.Date = invoice.Date.Add(time.Hour * 24 * 7)
	fmt.Println(invoice)
	fmt.Println(invoice.IsOutdated())
}

// --------------------------- //
func example2() {
	array := [3]int{1, 2, 3}       // this is an array with length = 3
	array1 := [...]int{1, 2, 3, 4} // this is an array with length = 4 generated by the compiler
	slice := array[1:2]            // this i a slice of array, [index, index not included]
	slice1 := []int{1, 2, 3, 4, 5} // this is a slice! the compiler create an array then return a slice
	fmt.Println("array:", array, "array1:", array1, "slice:", slice, "slice1:", slice1)
	/*
		The length of an array is matter: "array != array1"
		The function "readSlice(..)" accept only a slice not an array!
			but you can pass the array to it as a slice: "array[:]"
	*/
	readSlice(slice1)
	readSlice(array[:])
	fmt.Println("length of array:", len(array))
	fmt.Println("capacity of array:", cap(array)) // cap is how mush free space?

	// -- Dynamic Arrays --
	var table = make([]int, 3) // (array type, init capatity or size)
	table[0] = 5
	table[1] = 6
	table[2] = 7
	/*
		If we want to add table[3] we have to expand the array by "append(..)" function
	*/
	table = append(table, 8) // (table, the element to add)
	fmt.Println("dynamic table", table)
}

func readSlice(s []int) {
	// visit the example12()

	fmt.Println("readSlice:")
	for _, v := range s {
		fmt.Println(v)
	}
}

// --------------------------- //
func example3() {
	jamal := Freelancer{ID: 1, workingHours: 75, pricePeerHour: 100}
	amal := Salary{ID: 2, monthlyPrice: 3500, months: 2, bonus: 100}
	total := totalExpense([]Worker{jamal, amal})
	fmt.Println("Total Expense:", total)

	// trying the empty interface
	anything(45)
	anything("hello")
	anything([]int{1, 2, 3})
}

type Worker interface {
	CalculateSalary() float64
}
type Freelancer struct {
	ID            int
	workingHours  float64
	pricePeerHour float64
}
type Salary struct {
	ID           int
	monthlyPrice float64
	months       int
	bonus        int
}

// The "Freelancer" implements the "Worker" interface
func (this Freelancer) CalculateSalary() float64 {
	return this.pricePeerHour * this.workingHours
}

// The "Salary" implements the "Woker" interface
func (this Salary) CalculateSalary() float64 {
	return (this.monthlyPrice * float64(this.months)) + float64(this.bonus)
}

// Becuase "Freelancer" and "Salary" are both "Worker"s
func totalExpense(w []Worker) float64 {
	var total float64 = 0
	for _, v := range w {
		total += v.CalculateSalary()
	}
	return total
}

// empty interface is like the "any" type of Typecript!
func anything(t interface{}) {
	fmt.Printf("my value: %v, my type: %T\n", t, t)
}

// --------------------------- //
func example4() {
	results := map[string]int{} // [key]value
	results["kamal"] = 12
	results["said"] = 15
	results["mohsin"] = 13
	fmt.Println("results:", results)
	// to delete a value in map
	delete(results, "kamal")
	fmt.Println("results:", results)
	// kamal now is deleted but if we call it it give us a "0" value, cuze of go initialization
	fmt.Println("kamal value after deleting:", results["kamal"])
	// how we can make the deference between "0" of null and the real "0"?
	// the map return two things, "value, existing"
	value, ok := results["kamal"]
	fmt.Println("kamal value:", value, ",but is kanmal exist?", ok)
	fmt.Println("how many results there?", len(results))

	// trying anonymous function
	isExist := func(m map[string]int, key string) string {
		_, ok := m[key]
		if ok {
			return "yes"
		}
		return "no"
	}
	fmt.Println("kamal value:", value, ",but is kanmal exist?", isExist(results, "kamal"))
}

// --------------------------- //
func example5() {
	// ---- First way ----

	// Desctiption: now both of the process 'A'&'B' run in the same time concurentlry

	var wg sync.WaitGroup // think of it as a counter
	wg.Add(2)             // how many go routines we have? in this example (2 go functions)

	// async anonymous function run immediately
	go func() {
		process("A") // run process "A"
		wg.Done()    // emit the 'done' signal after finishing the process 'A'
	}()

	go func() {
		process("B")
		wg.Done()
	}()

	wg.Wait() // important for wating for the go routines: (blocks code here until all go routines above are finished)

	// ---- Second way ---- (communicate with process using channels)
	/*
	   - A Go channel is a communication mechanism that allows Goroutines to exchange data.
	   When developers have numerous Goroutines running at the same time,
	   channels are the most convenient way to communicate with each other.
	   - Developers often use these channels for notifications and managing concurrency in applications.
	*/

	out := make(chan string) // create a unbaffred channel
	go process1("C", out)

	out1 := <-out // the code is blocked until someone write into the pipe :) (this is the job of the process1 func)
	println("Receiver side x:", out1)

	// the 'go function' and the 'for loop', runs in concurently
	// the 'range chan', is reading from the channel until it closed!
	for msg := range out {
		println("Receiver side:", msg)
	}

	// ----- Third way ----- (baffred channels)
	/*
		the unbaffred channels blocks until some 'Go routines' read or write to it
		but we can avoid that by creating a `buffred channles` with a `capacity`!
	*/
	cha := make(chan int, 1) // buffred channel with the capatity of 1, that's means we can write&read from it (1 time) without a go routine!

	cha <- 5
	data, open := <-cha
	fmt.Println(data, open) // '5 true' why? 'true' means is open
	//data, open = <-cha // error? cuz we have only capacity == 1, so to avoid that increase the capacity or write to the channel again, it's a read write sequence

	cha <- 42
	data = <-cha
	close(cha)
	_, open = <-cha
	fmt.Println(data, open) // 42 false

}

func process(name string) {
	for i := 1; i <= 5; i++ {
		println(name, i)
		time.Sleep(time.Second)
	}
}
func process1(name string, out chan string) {
	for i := 1; i <= 5; i++ {
		out <- fmt.Sprint(name, i)
		time.Sleep(time.Second)
	}
	println("-- process1 end --")
	close(out) // manually close channel, never doing that in the receiver side!
	/*
		- cuz if we send data to a closed channel, the runtime crashed!
		- to avoid blocking, always close channel after the writing process
		- if we didn't that we got this error message:
		"fatal error: all goroutines are asleep - deadlock!"
	*/
	//	defer close(out)
	/*
		[ defer = fach tssali ]
			we can do that instead of closing manually the channel
			the 'defer' run the code after the end of the scope!
			so you can put it in the top if you want!
	*/
}

// --------------------------- //
func setTimeout1() {
	// set a duration
	duration := time.Duration(time.Second * 2)

	// we have just 1 timeout function (a.k.a go routines)
	var wg sync.WaitGroup
	wg.Add(1)

	// our anonymous setTimeout function
	go func() {
		defer wg.Done() // will work in any case, even if the func broken
		time.Sleep(duration)
		fmt.Println("i'm runing after ", duration)
	}()

	// ---- Our main code here ----
	println("---- I don't wait ----")

	// block the main until our (go routines) finished
	wg.Wait()
}
func setTimeout2() {
	// to let the "main" wating for us
	done := make(chan int)
	// block or wait for our setTimeout
	defer func() { <-done }()

	// setTimeout :) after 3 secs
	time.AfterFunc(time.Second*3, func() {
		fmt.Println("i'm running after 3 s")
		done <- 0 // finish or exit
	})

	// ---- Our main code here ----
	println("---- I don't wait ----")
}

// --------------------------- //
func example6() {
	/*
		Tickers are for when you want to do something repeatedly at regular intervals
	*/

	// ticker with Every 1 secs
	tk := time.NewTicker(time.Second)

	fmt.Println("Start..")

	// loop over "tk.C" channel, until "tk.Stop()"
	counter := 1
	for range tk.C {
		log.Println("Calling", counter)

		// stop after 4 iteration
		if counter == 4 {
			tk.Stop()
			break
		}
		counter++
	}

	/*
		of course if you want a non-blocking alternative to setInterval of js
		you have to work with tickers + go routines
		----------------
		Sleep Vs Ticker
		----------------
		- time.Sleep just waits for the provided time and continues with the program. There is no adjustment if the rest of the code takes longer.

		- The ticker takes the execution time of the provided block into account and skips an interval, if necessary.

		Imagine this scenario: You provide an interval of one minute and your code takes 10 seconds to execute.

		- In your first version your program executes your code for ten seconds and then sleeps for 60 seconds. Practically it gets called every 70 seconds.

		- In your second version your code gets executed for 10 seconds, then the ticker adjusts the wait time to 50 seconds. Your code gets executed exactly each minute.
	*/
}

// --------------------------- //
func example7() {

	/*
		We can use context as a store (like context in svelte)
		The context store the data in a "map", so you can use what ever type of data!
	*/
	myStore := context.Background() // scoped context: means my parent is "this" func: example7()
	myStore = storeValue(myStore)   // return a new store (immutable)
	readValue(myStore)

	storeValue2(&myStore) // mutable store
	readValue(myStore)

	/*
		we can also make a deadline or timeout to our functions that subscibe to the context
		by calling "cancel()" manually or by a duration
	*/
	parent := context.Background()                            // like "this" in other langs
	ctx, cancel := context.WithTimeout(parent, time.Second*3) // rule: after 3 secs, cancel the job
	defer cancel()                                            // don't forget this, cuz the timer is eating some resources if the main is done before the time is done

	// our func is subscribed to the context, so
	// the context terminate in a 3 secs, the job is cancelled!
	sleepAndTalk(ctx)

	/*
		- Why we need the context? we can do that by just a normal timer!
		the answer is, Nope!

		- The power of the context is, if any reason the "root or parent" context gets cancelled
		that cancellation will propagate to all the children contexts
		and all of those operations will stop!
	*/
}

func storeValue(ctx context.Context) context.Context {
	// context, key, value!
	newContext := context.WithValue(ctx, "name", "zaki")
	return newContext
	/*
		store the "key,value" in the "context"
	*/
}
func storeValue2(ctx *context.Context) {
	// edit the current context! pointer
	*ctx = context.WithValue(*ctx, "name", "karim")
}
func readValue(ctx context.Context) {
	// retrieve the key "name"
	val := ctx.Value("name")
	fmt.Println(val)
}
func sleepAndTalk(ctx context.Context) {
	// sleep a 5 sec then talk, but if the context is canceled, then print an error
	// it's a deadline bro!

	select {
	case <-time.After(time.Second * 5):
		fmt.Println("i'm talking after 5 secs")
	case <-ctx.Done():
		println(ctx.Err().Error())
	}
}

// --------------------------- //
func example8() {
	/*
		- Note:
			You have to start the server by:
			go run server/server.go

			You can use the browser or the "client1()" func!
	*/

	fmt.Println("loading ...")
	// trying a normal http call to the server
	//client1()

	/*
	 Calling the request with a context
	 The client abort the connection after 3 secs (timeout),
	 So the server now reveive this signal and save some resources! Cool!
	*/
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		// this will aborted afer 3 secs
		client2("http://localhost:8080", time.Second*2)
		wg.Done()
	}()

	wg.Add(1)
	func() {
		// this will recieve the "Hello Go! [2]" body msg, cuz the time out is too big!
		client2("http://localhost:8080/2", time.Second*7)
		wg.Done()
	}()
	wg.Wait()
}

func client1() {
	// call the server
	res, err := http.Get("http://localhost:8080")

	// error handling
	if err != nil {
		log.Fatal(err.Error())
	}
	if res.StatusCode != http.StatusOK {
		log.Fatal(res.Status)
	}
	// important
	defer res.Body.Close()

	// just print what the server gives us! :)
	bytes, _ := io.ReadAll(res.Body)
	fmt.Println("Client received: ", string(bytes))
}

func client2(url string, timeout time.Duration) {
	/*
		Like the client1() func, but this time we add a timeout to the context in the request
		So with that, We have to separate the request and the responce funcs

		Now the request context has two things:
		(1) the abort calling (by user hand) "the default req context"
		(2) + our new rule: the time out!
		Cool!, we don't mess with the server or the handlers,
		We just did that here by the context inheritence!
	*/

	ctx := context.Background()
	// abort the request if we passed 3 secs
	ctx, cancel := context.WithTimeout(ctx, timeout)
	defer cancel()

	// create a new request with our context
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	// error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	// make the call of the request
	res, err := http.DefaultClient.Do(req)

	// error handling
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if res.StatusCode != http.StatusOK {
		fmt.Println(res.Status)
		return
	}
	// important
	defer res.Body.Close()

	// just print what the server gives us! :)
	bytes, _ := io.ReadAll(res.Body)
	fmt.Println("Client received: ", string(bytes))
}

// --------------------------- //
func example9() {
	/*
		The singleton pattern:
		Sometimes we need our code to be executed only once in the entire lifetime!
		So we can do this with mutext.lock/unlock, but to be more crleary and 100% safe
		We use the "Sync" package with the "..Do()" funtion!
	*/

	var Once sync.Once

	// Here is a sample program that shows how even if you call it multiple times, it gets executed only once.
	for range []int{1, 2, 3} {
		Once.Do(func() {
			fmt.Println("I executed jut once!")
		})
	}

	/*
		The "singleton pattern" is useful when
			we one to instantiate a database connection one time
			even if our app is calling the constractor function many times!
	*/
}

// --------------------------- //
func example10() {

	/*
		Here we’ll look at using the sync/atomic package for atomic counters
		accessed by multiple goroutines.
	*/

	var counter uint64 //  always-positive
	var wg sync.WaitGroup

	// We’ll start 50 goroutines that each increment the counter exactly 1000 times.
	for i := 0; i < 50; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			/*
				To atomically increment the counter we use AddUint64,
				giving it the memory address of our 'counter' with the '&' syntax.
			*/
			for j := 0; j < 1000; j++ {
				// add +1 to the counter atomically
				atomic.AddUint64(&counter, 1)
			}
		}()
	}

	wg.Wait()

	// we expected counter=50,000
	fmt.Println("We have now", counter, "operations")

	/*
		It’s safe to access 'counter' now because we know no other goroutine is writing to it.
		 Reading atomics safely while they are being updated is also possible
		 using functions like 'atomic.LoadUint64'.

		 NOTE:
		 We expect to get exactly 50,000 operations.
		 Had we used the non-atomic ops++ to increment the counter,
		 we’d likely get a different number, changing between runs,
		 because the goroutines would interfere with each other.
		 Moreover, we’d get data race failures when running with the -race flag.
	*/
}

// --------------------------- //
func example11() {
	// alibaba is our container :p
	// we don't need to fill the 'mu' mutex here
	alibaba := container{
		counters: map[string]int{"a": 0, "b": 0},
	}
	var wg sync.WaitGroup

	// This function increments a named counter in a loop.
	counting := func(name string, max int) {
		defer wg.Done()
		for i := 0; i < max; i++ {
			alibaba.increment(name)
		}
	}

	/*
		Run several goroutines concurrently;
		note that they all access the same 'container',
		and two of them access the same 'counter'.
	*/
	wg.Add(3)
	go counting("a", 1000)
	go counting("b", 1000)
	go counting("a", 1000)
	wg.Wait()

	// map[a:2000 b:1000]
	fmt.Println("Alibaba counters:", alibaba.counters)
}

type container struct {
	counters map[string]int // map, every name has a counter
	mu       sync.Mutex     // mutex to organize the go routines access
}

func (this *container) increment(name string) {
	// lock the go routines, just 1 at time
	this.mu.Lock()
	this.counters[name]++  // now we have like the atomic counter, thanks to mutex!
	defer this.mu.Unlock() // finally, it's the others go routines turn :p

	/*
		if we don't use mutex we get:
		"fatal error: concurrent map writes"
	*/
}

// --------------------------- //
func example12() {
	names := []string{"zakaria", "mona", "ayoub", "ilyas"}
	numbers := []int{8, 5, 2, 6, 3, 7, 1}

	/*
		everythings in Go are bassed by value by default
		the slices also passed by value, BUT it points to the original array
		so, yes it create a copy but it refers to the same array!

		A slice is a descriptor of an array segment:
		[ptr *array, len int, cap int]

		Sort is ascending order by default
	*/
	sort.Strings(names) // attention!
	sort.Ints(numbers)  // attention!

	fmt.Println("my names are sorted now:", names)
	fmt.Println("my numbers are sorted now:", numbers)

	isSorted := sort.StringsAreSorted(names) // check if our 'names' are sorted
	fmt.Println("is our names sorted?", isSorted)

	// sort in decreasing order ---- way 1 ----
	sort.Sort(sort.Reverse(sort.StringSlice(names))) // see "custom sorting" bellow to understand
	fmt.Println("names in decreasing order(1):", names)

	// sort in decreasing order ---- way 2 ----
	sort.Slice(names, func(i, j int) bool {
		return names[i] < names[i]
	})
	fmt.Println("names in decreasing order(2):", names)

	/*
		Custom Sorting!
		we can implement out custom sort algoritm to sort for example by "lenght"
		instead of alphabetically sorting

		use vscode go plugin to create sort functions for you :p
	*/

	sort.Sort(sortByLenght(names)) // sort names by lenght
	fmt.Println("names sorted by lenght", names)
}

type sortByLenght []string

func (a sortByLenght) Len() int      { return len(a) }
func (a sortByLenght) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a sortByLenght) Less(i, j int) bool {
	// here is out implementation
	return len(a[i]) < len(a[j])
}

// --------------------------- //
func example13() {
	/*
	   A panic typically means something went unexpectedly wrong.
	   Mostly we use it to fail fast on errors that shouldn’t occur during normal operation,
	   or that we aren’t prepared to handle gracefully
	*/

	fmt.Println("---- Doing something .. ----")

	/*
		Like catch
		recover must be called inside a deferred function
	*/
	defer func() {
		if catch := recover(); catch != nil {
			fmt.Println("Recovered - The panic error is:", catch)
		}
	}()

	// it's panic! like throw
	justPanic()

	fmt.Println("---- I can't show up! Cuz of the panic ----")
}

func justPanic() {
	panic(`"we can't continue baby!"`) // abort the whole app!
}

// --------------------------- //
func example14() {
	var p = fmt.Println

	/*
		The standard library’s strings package provides many useful string-related functions.
	*/

	p("Contains:  ", strings.Contains("test", "es"))
	p("Count:     ", strings.Count("test", "t"))
	p("HasPrefix: ", strings.HasPrefix("test", "te"))
	p("HasSuffix: ", strings.HasSuffix("test", "st"))
	p("Index:     ", strings.Index("test", "e"))
	p("Join:      ", strings.Join([]string{"a", "b"}, "-"))
	p("Repeat:    ", strings.Repeat("a", 5))
	p("Replace:   ", strings.Replace("foo", "o", "0", -1))
	p("Replace:   ", strings.Replace("foo", "o", "0", 1))
	p("Split:     ", strings.Split("a-b-c-d-e", "-"))
	p("ToLower:   ", strings.ToLower("TEST"))
	p("ToUpper:   ", strings.ToUpper("test"))
}

// --------------------------- //
func example15() {
	t := template.New("template1")

	// ---- parse from string ----
	t, err := t.Parse("Hi, i'm {{.}}\n")
	if err != nil {
		log.Fatalln(err)
	}
	// fill the template
	fmt.Println("---- t ----")
	err = t.Execute(os.Stdout, "zaki") // Hi, i'm zaki
	if err != nil {
		log.Fatalln(err)
	}

	// ---- parse from file ----
	path := "templates/t1.txt"
	t1 := template.Must(template.ParseFiles(path)) // 'Must' instead of cheking 'err' manually: it's panic()
	fmt.Println("---- t1 ----")
	t1.Execute(os.Stderr, "zaki")

	// ---- named filled data ----
	path = "templates/t2.txt"
	t2 := template.Must(template.ParseFiles(path))
	// you can use structs
	fmt.Println("---- t2 ----")
	t2.Execute(os.Stderr, struct {
		Name    string
		Age     int
		Hobbies []string
	}{
		Name:    "zaki",
		Age:     26,
		Hobbies: []string{"voice over", "singing", "trip", "cooking"},
	})

	/*
		Note:
		- instead of calling `ParseFiles(..)` each time
		you can use `ParseGlob("templates/*")` one time,
		- then you can excute each file like that:
		.ExecuteTemplate(os.Stdout, "t1.txt", ..)
		.ExecuteTemplate(os.Stdout, "t2.txt", ..)
	*/

	// ---- Declaring variables inside templates ----
	path = "templates/t3.txt"
	t3 := template.Must(template.ParseFiles(path))
	fmt.Println("---- t3 ----")
	// Note: '-' is for trim whitespaces
	t3.Execute(os.Stderr, nil)

	// ---- Conditional Statements ----
	path = "templates/t4.txt"
	t4 := template.Must(template.ParseFiles(path))
	fmt.Println("---- t4 ----")
	// Generate random number between 100 and 300
	rand.Seed(time.Now().UnixNano())
	min := 100
	max := 300
	num := rand.Intn((max-min)+1) + min
	t4.Execute(os.Stderr, num)

	// ---- Using functions in templates ----
	/*
		Note:
		- functions must have only 1 return value, or 1 return value and an error
		- bind it before the parsing
	*/
	// define our func
	decorator := func(text string) string {
		return "**" + text + "**"
	}
	// put it into func map
	myFunctions := template.FuncMap{
		"toUpper":   strings.ToUpper,
		"decorator": decorator,
	}
	path = "templates/t5.txt"
	/*
		Since the templates created by ParseFiles
		are named by the base name of the argument files,
		the base name here is "t5.txt"
	*/
	t5 := template.Must(template.New("t5.txt").Funcs(myFunctions).ParseFiles(path))
	// just call it inside the template as a pipeline "|" or argument
	fmt.Println("---- t5 ----")
	err = t5.Execute(os.Stderr, "i'm a beautiful text")
	if err != nil {
		log.Fatalln(err)
	}

	/*
		You can do the same in HTML by "html/template" :)
	*/
}

// --------------------------- //
func example16() {
	re, err := regexp.Compile("^z.*a$")
	if err != nil {
		panic(err)
		// or regexp.MustCompile()
	}

	// check if strings mutch the regular expression
	ok := re.Match([]byte("zakaria"))
	fmt.Println(ok) // true
	ok = re.MatchString("zakia")
	fmt.Println(ok) // true

	// ----  Find ----
	re = regexp.MustCompile(`z[a-z]+a`)
	text := "my name is zakaria, my wife is zakia"

	// get the 1st match string
	match := re.FindString(text)
	fmt.Println(match) // 'zakaria'

	// get the 1st match index
	location := re.FindStringIndex(text)
	fmt.Println(location) // slice of 2: [start, end not included]: [11, 18]

	// get all matches strings
	matches := re.FindAllString(text, -1) // -1 to match all strings, or you can set a max number
	fmt.Println(matches)                  // ['zakaria', 'zakia']

	// get all matches indexs
	locations := re.FindAllStringIndex(text, -1)
	fmt.Println(locations) // slice or slices: [[11 18] [31 36]]

	// ---- Replace ----
	replaced := re.ReplaceAllString(text, "anonymous") // "my name is anonymous, my wife is anonymous"
	fmt.Println(replaced)

	// ---- Replace function ----
	out := re.ReplaceAllStringFunc(text, strings.ToUpper) // "my name is ZAKARIA, my wife is ZAKIA"
	fmt.Println(out)

	decorator := func(match string) string {
		return "**" + match + "**"
	}
	out = re.ReplaceAllStringFunc(text, decorator) // "my name is **zakaria**, my wife is **zakia**"
	fmt.Println(out)

	// ---- Split ----
	text = "akram[1]asaad[2]mona[3]zakaria[4]"
	re = regexp.MustCompile(`\[\d\]`)
	splited := re.Split(text, -1)
	fmt.Println(splited) // [akram asaad mona zakaria ]
}

// --------------------------- //
func example17() {

	/*
		Form Go data types to Json data types
		a.k.a "JSON.stringify"
	*/

	// basic types
	jsonBool, _ := json.Marshal(true)
	jsonInt, _ := json.Marshal(3)
	jsonString, _ := json.Marshal("hello")

	fmt.Println(true, "to", string(jsonBool))
	fmt.Println(3, "to", string(jsonInt))
	fmt.Println("hello", "to", string(jsonString))
	fmt.Println("-------------------")

	// convert slices and maps to JSON arrays and objects
	goSlice := []string{"zakaria", "mona"}
	goMap := map[string]int{"code": 35, "label": 12}
	jsonArray, _ := json.Marshal(goSlice)
	jonObject, _ := json.Marshal(goMap)

	fmt.Println(goSlice, "to", string(jsonArray))
	fmt.Println(goMap, "to", string(jonObject))
	fmt.Println("-------------------")

	// ---- Custom data types ----
	res1 := &responce_1{
		ID:   104,
		Data: []int{10, 20, 30, 40, 50, 60},
	}
	jsonRes1, _ := json.Marshal(res1)
	fmt.Println(res1, "to", string(jsonRes1))
	fmt.Println("-------------------")

	// with custom json keys
	res2 := &responce_2{
		ID:   104,
		Data: goSlice,
	}
	jsonRes2, _ := json.Marshal(res2)
	fmt.Println(res1, "to", string(jsonRes2))
	fmt.Println("-------------------")

	/*
		Form Json data types to Go data types
		a.k.a "JSON.parse"
	*/
	json_string := `{"id": 99, "names": ["zaki", "mohammed", "imane"]}`
	var obj responce_2
	err := json.Unmarshal([]byte(json_string), &obj)
	if err != nil {
		panic(err)
	}
	fmt.Println(obj.Data)

}

type responce_1 struct {
	// Note: JSON will only include exported fields
	ID   int
	Data []int
}
type responce_2 struct {
	// Note: JSON will only include exported fields
	ID   int      `json:"id"`    // tag: to custom JSON key names
	Data []string `json:"names"` // tag
}

// --------------------------- //
func example18() {
	stringUrl := "postgres://zaki:123456@host.com:5432/api?version=1.0#f"
	res, err := url.Parse(stringUrl)
	if err != nil {
		panic(err)
	}

	// Scheme --> postgres
	fmt.Println(res.Scheme)

	// Username --> zaki
	fmt.Println(res.User.Username())

	// Password --> 123456, true
	fmt.Println(res.User.Password())

	// Host --> host.com:5432
	fmt.Println(res.Host)

	// Hostname --> host.com
	fmt.Println(res.Hostname())

	// Port --> 5432
	fmt.Println(res.Port())

	// ----- or ------

	host, port, _ := net.SplitHostPort(res.Host)
	// Host, Port --> host.com, 5432
	fmt.Println(host, port)

	// Path --> /api
	fmt.Println(res.Path)

	// Fragment --> f
	fmt.Println(res.Fragment)

	// Query string --> version=1.0
	fmt.Println(res.RawQuery)

	// Query map --> map[version:[1.0]]
	fmt.Println(res.Query())
	// Get query --> 1.0
	fmt.Println(res.Query().Get("version"))
	// Check query --> true
	fmt.Println(res.Query().Has("version"))
}

// --------------------------- //
func example19() {
	// text to encrypt
	text := "my text"

	// encryption type: sha256
	hash := sha256.New()

	// encrypt text to sha256 in a byte stream
	hash.Write([]byte(text))
	bStream := hash.Sum(nil)

	// convert []byte to string
	hash_text := fmt.Sprintf("%x", bStream)

	// output: 7330d2b39ca35eaf4cb95fc846c21ee6a39af698154a83a586ee270a0d372104
	fmt.Println(hash_text)

	// enctyption type: md5
	hash = md5.New()
	hash.Write([]byte(text))
	bStream = hash.Sum(nil)
	hash_text = fmt.Sprintf("%x", bStream)

	// output: d3b96ce8c9fb4e9bd0198d03ba6852c7
	fmt.Println(hash_text)
}

// ------------------------------ //
func example20() {
	/*
		Go supports both standard and URL-compatible base64.
	*/

	// my data to encode, can be binary too
	data := "Hi! i have 26 years old!"

	// ---- Standard encoding: like btoa() in js ----
	encodedData := base64.StdEncoding.EncodeToString([]byte(data))
	// SGkhIGkgaGF2ZSAyNiB5ZWFycyBvbGQh
	fmt.Println(encodedData)

	// ---- Standard decoding: like atob() in js ----
	originalData, err := base64.StdEncoding.DecodeString(encodedData)
	if err != nil {
		panic(err)
	}
	// Hi! i have 26 years old!
	fmt.Println(string(originalData))

	// ---- URL encoding: like btoa() in js ----
	encodedData = base64.URLEncoding.EncodeToString([]byte(data))
	// SGkhIGkgaGF2ZSAyNiB5ZWFycyBvbGQh
	fmt.Println(encodedData)

	// ---- URL decoding: like atob() in js ----
	originalData, err = base64.URLEncoding.DecodeString(encodedData)
	if err != nil {
		panic(err)
	}
	// Hi! i have 26 years old!
	fmt.Println(string(originalData))

}

// ------------------------------ //
func example21() {
	// basic file reading is to grab entire content into memory
	data, err := os.ReadFile("./templates/t1.txt")
	check(err)
	fmt.Println(string(data))

	// ---- Advanced more details ----
	file, err := os.Open("./templates/t1.txt")
	check(err)
	// important
	defer file.Close()

	// define how many bytes we want to read: ex 5 bytes
	store := make([]byte, 5)
	//return the number of bytes read (in this example will be 5)
	length, err := file.Read(store)
	check(err)
	// 5 bytes --> Hi i'
	fmt.Printf("%d bytes --> %s\n", length, string(store[:length]))

	/*
		We can also `Seek` to a known location in the file
		and Read from there.

		Seek() takes 2 arguments
		------------------------
		1st: set the offset (read from this position)
		2nd: a number (0, 1 or 3) = positions
			0= from the begining of file
			1= from the offset
			2= from the ned of file

		it return the new offset
	*/

	// new offset = 13
	_, err = file.Seek(13, 0)
	check(err)
	store = make([]byte, 4)
	length, err = file.Read(store)
	check(err)
	// Yes!
	fmt.Println(string(store))

	// ---- More useful things ----
	/*
		we can also use `io` package to get some helpful functions
	*/

	// reset
	file.Seek(0, 0)
	store = make([]byte, 20)
	// from, into, min
	length, err = io.ReadAtLeast(file, store, 6) // read until store is fulfilled, but at least 6 bytes
	check(err)
	fmt.Println(`------- io pacakge -------`)
	fmt.Println(string(store))
}
func example22() {
	// basic file writing
	myText := []byte("Hi i'm zaki how are you?")

	/*
		Params: (file path, btyes to write, file mode)

		File Mode (Unix):
		0000     no permissions
		0700     read, write, & execute only for owner
		0770     read, write, & execute for owner and group
		0777     read, write, & execute for owner, group and others
		0111     execute
		0222     write
		0333     write & execute
		0444     read
		0555     read & execute
		0666     read & write
		0740     owner can read, write, & execute; group can only read; others have no permissions

		------------------
		for example if we create a file with '02222' and we want to access to it
		we will get: "cat: files/write1.txt: Permission denied"
		cuz we give to it just the 'write' permissions
	*/
	err := os.WriteFile("./files/write1.txt", myText, 0666) // 0666 or fs.FileMode(os.O_RDWR))
	check(err)

	// ----- More details -----
	// create a file, then open itm with file mode 0666
	file, err := os.Create("./files/write2.txt")
	defer file.Close()
	check(err)

	// write bytes to file
	_, err = file.Write([]byte("Hi!\n"))
	check(err)
	// write string to file (append to "Hi\n")
	_, err = file.WriteString("السلام عليكم")
	check(err)

	/*
		This function not only commits the current contents of the file to persistent storage
		but also flushes the file system’s in-memory copy of recently written data to the persistent store.
		Taking care of the memory footprints is always a good idea in programming.
	*/
	file.Sync()

	// ---- io package ----
	_, err = io.WriteString(file, "\nHoa! ")
	check(err)

	// ---- bufio package (buffered writers) ----
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString("como estas\n")
	check(err)

	// commit!
	writer.Flush()
}
func check(err error) {
	if err != nil {
		panic(err)
	}
}

// ------------------------------ //
func example23() {
	/*
		A line filter is a common type of program that reads input on 'stdin', processes it,
		and then prints some derived result to 'stdout'.
		'grep' is common line filters.

		like pipe:
		echo "Hi, how are you?" | go run main.go

		--------------
		Wrapping the unbuffered 'os.Stdin' with a buffered scanner
		gives us a convenient Scan method that advances the scanner to the next token,
		which is the next line in the default scanner.
	*/
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		upper := strings.ToUpper(line)
		fmt.Println(upper)
	}

	// check errors
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}
}

// ------------------------------ //
func example24() {
	path := filepath.Join("dir1", "dir2", `filename`)
	// -> dir1/dir2/filename
	fmt.Println(path)

	/*
		Join will also normalize paths
		by removing superfluous separators and directory changes.
	*/
	path = filepath.Join("dir1//", `filename`)
	// -> dir1/filename
	fmt.Println(path)

	path = filepath.Join("dir1/../dir1", `filename`)
	// -> dir1/filename
	fmt.Println(path)

	// ---- get directory from path ----
	path = "/home/zaki/Documents/book.pdf"
	dir := filepath.Dir(path)
	// -> /home/zaki/Documents
	fmt.Println(dir)
	// -> book.pdf
	base := filepath.Base(path)
	fmt.Println(base)

	// ---- Absolute path checkng ----
	// -> false
	fmt.Println(filepath.IsAbs("desktop/file.txt"))
	// -> true
	fmt.Println(filepath.IsAbs("/desktop/file.txt"))

	// ---- File Extension ----
	filename := "Readme.md"
	ext := filepath.Ext(filename)
	// -> .md
	fmt.Println(ext)
	// -> Readme
	filename = strings.TrimSuffix(filename, ext)
	fmt.Println(filename)

	/*
		`Rel` finds a relative path between a `base` and a `target`.
		It returns an error if the target cannot be made relative to base.
	*/

	rel, err := filepath.Rel("a/b", "a/b/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)
	rel, err = filepath.Rel("a/b", "a/c/t/file")
	if err != nil {
		panic(err)
	}
	fmt.Println(rel)

}

// ------------------------------ //
func example25() {
	/*
		Throughout program execution,
		we often want to create data that isn’t needed after the program exits.
		Temporary files and directories are useful for this purpose since they don’t pollute the file system over time.
	*/

	// create a temp file, open it with read/write modes, in the default location in the os (linux /tmp)
	tempFile, err := os.CreateTemp("", "sample") // "" -> default lcoation, "sample" is the pattern, the file name will ba like that: "/tmp/sample1722126074"
	check(err)
	// remove file after closing
	defer os.Remove(tempFile.Name())
	// display the file name (the path) -> "/tmp/sample1722126074"
	fmt.Println("temporary file name:", tempFile.Name())

	// ---- Custom Path with Custom name ----
	name := "myfile-*" // with be like that -> "myfile-154965564", the func replace the '*'
	directory := "./files"
	tempFile2, err := os.CreateTemp(directory, name)
	check(err)
	defer os.Remove(tempFile2.Name())
	// -> ./files/myfile-3811349410
	fmt.Println("temporary file2 name:", tempFile2.Name())

	// ---- temporary directories ----
	tempDir, err := os.MkdirTemp("", "mydir")
	check(err)
	// display dir name (path)
	fmt.Println("temporary dir name:", tempDir)
	// remove the dir with all content
	defer os.RemoveAll(tempDir)

	// add a file to the temporary directory
	newFilePath := filepath.Join(tempDir, "newFile.txt")
	err = os.WriteFile(newFilePath, []byte("hello!\n"), 0666)
	check(err)

	// just to wait to see the files
	//time.Sleep(time.Second * 10)
}

// ------------------------------ //
//go:embed templates/t1.txt
var fileString string

//go:embed templates/t2.txt
var fileByte []byte

//go:embed templates/*
var folder embed.FS

func example26() {
	/*
		<go:embed> is a compiler directive that allows programs to include
		arbitrary files and folders in the Go binary at build time.
	*/

	// Print the embed "templates/t1.txt" content
	fmt.Println(fileString)
	fmt.Println("----------------")
	// Print the embed "templates/t2.txt" content
	fmt.Println(string(fileByte))
	fmt.Println("----------------")

	// Print anyting from the embed templates folder
	content, err := folder.ReadFile("templates/t3.txt")
	check(err)
	fmt.Println(string(content))
}
