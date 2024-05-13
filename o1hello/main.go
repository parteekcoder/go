package main

import (
	"fmt"
	"sync"
	"time"
)

var a = "h" // package level variable
var A = "j" // Global Variable
// statictally types: tell go type while declaraing
// type inference

// structure struct , like ligtweight class
type UserData struct {
	name   string
	rollNo int
}

var u = UserData{
	name:   "hi",
	rollNo: 233,
}

var signals = []string{"test"}
var wg sync.WaitGroup //pointer
var mut sync.Mutex    //pointer

// u.name

func main() {
	// var conferenceName = "Go Conference"
	conferenceName := "Go Conference" // local variable
	const tickets int = 20
	fmt.Print("Welcome to ", conferenceName, " booking appp \n")
	fmt.Printf("Welcome to %v", conferenceName)
	// const conferenceTickets = 20
	fmt.Println("hello")
	fmt.Printf("type %T", conferenceName)
	var userName string
	fmt.Scan(&userName)
	//fsgsg
	fmt.Print(userName)

	// array   fixed size
	// var users = [50]string{}
	// var users [50] string
	// print(len(users))
	// print(users)

	//slices
	// abstraction of array
	// dynamic size

	// var users = []string{}
	var users []string
	users = append(users, "hi")
	fmt.Println(users)

	//loops
	// only for loop
	//  not for-each, while, do-while
	// for i := 0; i < 10; i++ {

	// }

	// for-each
	// range iterates over elements for diff DS not specific to arrays/slice
	// for arrray it provides index and value
	for index, user := range users {
		fmt.Print(index, user)
		if index == 2 {
			break
		}
	}

	// for condition{

	// }
	// if index not using we can use blank identifier(_)
	// to ignore a variable

	// maps
	// var mp = make(map[string]string)
	// var mp map[string]string
	// var listMp []map[string]string

	wg.Add(1)
	go sendTicket()

	wg.Wait()
}

func sendTicket() {
	time.Sleep(2 * time.Second)
	mut.Lock()
	signals = append(signals, "d")
	mut.Unlock()
	defer wg.Done()
}
func sayHi(str string) string {
	return str + "hi"
}
func sayHis(str string) (string, string) {
	return "hi", str
}

// Go have thread abstraction "Green Thread"
// we interact with goroutine managed by go runtimes
// it can create millions of thread without any degradation in performence
// it is easy , cheaper
// go have channels for communication between the threads

// in other languages it is complex
// it requires to interact with direct OS level threads
// here there is difficult to communicate
