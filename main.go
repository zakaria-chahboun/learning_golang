package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"example.com/bill"
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

	// an implementation of setTimeOut (js)
	// example6()

	// playing with context
	example7()
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
func example6() {

	// duration
	duration := time.Duration(time.Second * 2)

	// we have just 1 timeout function (aka go routines)
	var wg sync.WaitGroup
	wg.Add(1)

	// our anonymous time out func
	go func() {
		time.Sleep(duration)
		fmt.Println("i'm runing after ", duration)
		wg.Done()
	}()

	// our main code here
	fmt.Println("---- i don't wait here ----")

	// block go routines until finished
	wg.Wait()
}

// --------------------------- //
func example7() {

	/*
		You can use context as a store (like context/store in svelte)
		The context store the data in a "map", so yu can use what ever type of data!
	*/
	myStore := context.Background() // scoped context: means my parent is "this" func: example7()

	myStore = addValue(myStore) // return a new store (immutable)
	readValue(myStore)

	addValue2(&myStore) // mutable store
	readValue(myStore)

	//timeout, cancel := context.WithTimeout(context.Background(), time.Second*2)

}

func addValue(ctx context.Context) context.Context {
	// context, key, value!
	newContext := context.WithValue(ctx, "name", "zaki")
	return newContext
	/*
		store the "key,value" in the "context"
	*/
}
func addValue2(ctx *context.Context) {
	// edit the current context! pointer
	*ctx = context.WithValue(*ctx, "name", "karim")
}
func readValue(ctx context.Context) {
	// retrieve the key "name"
	val := ctx.Value("name")
	fmt.Println(val)
}
