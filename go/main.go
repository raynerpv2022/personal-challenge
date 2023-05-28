package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// type persona struct {
// 	name int
// }

// func ChangeToMe(p *persona) {
// 	p.name = 999999 //(*p).name too is ok

// }

// **********************************************************
// func fibo(n int) int {
// 	// if n == 0 || n == 1 {
// 	// 	return n
// 	// }

// 	// return fibo(n -1) + fibo(n -2)

// 	switch n {
// 	case 0:
// 		return 0
// 	case 1:
// 		return 1
// 	default:
// 		return fibo(n-1) + fibo(n-2)
// 	}

// }

// ***********************************************************************
// Json

//	type fitness struct {
//		Name    string
//		Country string
//		Repe    int
//	}
//
// ****************************
// Sort personalized
// type Sport struct {
// 	Name      string
// 	Player    int
// 	Countries []string
// }

// type byName []Sport
// type porjugadores []Sport

// func (a byName) Len() int           { return len(a) }
// func (a byName) Less(w, e int) bool { return a[w].Name < a[e].Name }
// func (a byName) Swap(w, e int)      { a[w], a[e] = a[e], a[w] }

// func (s porjugadores) Len() int           { return len(s) }
// func (s porjugadores) Less(w, e int) bool { return s[w].Player < s[e].Player }
// func (s porjugadores) Swap(w, e int)      { s[w], s[e] = s[e], s[w] }

// **********************

func main() {
	// *********************************************
	// concurrencia race.... acceso a variables compartidas
	var wg sync.WaitGroup
	// var mu sync.Mutex //con Atomic no lo necesito
	var counter int64
	goRoutine := 20
	fmt.Println("CUPs NUmber", runtime.NumCPU())
	fmt.Println("Go Routine A", runtime.NumGoroutine())
	wg.Add(goRoutine)
	for i := 0; i < goRoutine; i++ {
		go func() {
			// mu.Lock()
			// shareVar := counter  con atoic no lo necesito
			atomic.AddInt64(&counter, 1)
			runtime.Gosched()
			//shareVar++
			fmt.Println("Counter in For :", atomic.LoadInt64(&counter))
			// counter = shareVar
			// mu.Unlock()    con atomic no lo necesito
			fmt.Println("Go Routine B", runtime.NumGoroutine())
			wg.Done()
		}()
		fmt.Println("Go Routine C", runtime.NumGoroutine())

	}
	wg.Wait()
	fmt.Println("Counter is", counter)

	// *******************************************
	// runtime ....concurrencia
	// var wg sync.WaitGroup
	// runtimeFunc := func() {
	// 	fmt.Printf("ARchitecture\t%v\n", runtime.GOARCH)
	// 	fmt.Printf("Num CPU\t%v\n", runtime.NumCPU())
	// 	fmt.Printf("NUm Go Routine\t%v\n", runtime.NumGoroutine())

	// }
	// s := func(mess string) {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(mess, i)
	// 	}
	// 	wg.Done()
	// }
	// g := func(mess string) {
	// 	for i := 0; i < 10; i++ {
	// 		fmt.Println(mess, i)
	// 	}
	// 	//wg.Done() // Elimina el proceso que add puso... y n se ejecuta mas ninguna concurrencia
	// }
	// runtimeFunc()
	// wg.Add(1)
	// go s("With GO")
	// g("Wkukuschima")
	// runtimeFunc()
	// wg.Wait()

	// ***********************************************************
	//video 129 Sort personalized

	// Sport2 := Sport{"Baseball", 29, []string{"Cuba", "USA", "Japon"}}
	// Sport3 := Sport{"FootBall", 11, []string{"Spain", "Argentine"}}
	// Sport1 := Sport{"Jokey   ", 16, []string{"Rusia", "Germany"}}
	// AllSport := []Sport{Sport1, Sport2, Sport3}

	// printSport := func(AllSport []Sport) {
	// 	fmt.Println("--------------------------Sports-----------------------------------")
	// 	fmt.Printf("\tSport\t\tPlayer\t\tCountry\n")
	// 	for i, v := range AllSport {
	// 		fmt.Printf(" %d\t%s\t%d", i+1, v.Name, v.Player)
	// 		for _, c := range v.Countries {
	// 			fmt.Printf("\t\t%s", c)
	// 		}
	// 		fmt.Println()
	// 	}

	// }

	// printSport(AllSport)
	// sort.Sort(byName(AllSport))
	// printSport(AllSport)
	// sort.Sort(porjugadores(AllSport))
	// printSport(AllSport)

	// *************************************************************
	// Sort Video 128
	// intArray := []int{3, 9, 1, 5, 9, 88, 77, 33, 1, 0}
	// stringArray := []string{"zzz", "dd", "wq", "aaa"}
	// sort.Ints(intArray)
	// sort.Strings(stringArray)
	// fmt.Println(intArray, stringArray)

	// **********************************************************
	// Video 127
	// type Sport struct {
	// 	Name      string
	// 	Player    int
	// 	Countries []string
	// }
	// Sport1 := Sport{"Baseball", 9, []string{"Cuba", "USA", "Japon"}}
	// Sport2 := Sport{"FootBall", 11, []string{"Spain", "Argentine"}}
	// Sport3 := Sport{"Jokey", 6, []string{"Rusia", "Germany"}}
	// AllSport := []Sport{Sport1, Sport2, Sport3}

	// error := json.NewEncoder(os.Stdout).Encode(AllSport) // esta funcion envia a stdout algo... el enredo.. como devuelve un puntero a Encoder
	// if error != nil {                                    // entonces puedeo usar el metodo encoder que recibe una interface(cualquier valor)
	// 	fmt.Fprintln("Error: ", error)
	// }

	// *******************************************************
	//Video 126 unmarshal
	// type TSfromJsontoGo struct {
	// 	Name string
	// 	Page int
	// }
	// jsonToGo := `[{"Name":"Black swan","Page":300},{"Name":"Invicto","Page":249},{"Name":"Antifragil","Page":500}]`

	// var fromJsontoGo []TSfromJsontoGo
	// error := json.Unmarshal([]byte(jsonToGo), &fromJsontoGo)
	// if error != nil {
	// 	fmt.Println("Error from Json:", error)
	// }
	// fmt.Println(fromJsontoGo)
	// for i, v := range fromJsontoGo {
	// 	fmt.Println("Book :", i+1)
	// 	fmt.Printf("\tName : %s\n \tPage : %d\n", v.Name, v.Page)
	// }

	// ********************************************
	// video 125 ex Marshall JSON

	// type Book struct {
	// 	Name string
	// 	Page int
	// }

	// reader1 := Book{"Black swan", 300}
	// reader2 := Book{"Invicto", 249}
	// reader3 := Book{"Antifragil", 500}
	// allReader := []Book{reader1, reader2, reader3}
	// fmt.Println(allReader)
	// ToJson, error := json.Marshal(allReader)
	// if error != nil {
	// 	fmt.Println("Error insider", error)
	// }
	// fmt.Println("To Json:", string(ToJson))

	// ***********************************
	// Sort
	// type TfitNess struct {
	// 	Tnametype string
	// 	repe      int
	// }

	// ex1 := TfitNess{"Muscle-up", 5}
	// ex2 := TfitNess{"Push up", 45}
	// ex3 := TfitNess{"Pull up", 10}
	// ex4 := TfitNess{"Abs Leg", 30}

	// fmt.Println(ex1)
	// fmt.Println(ex2)
	// fmt.Println(ex3)
	// AllInOne := []TfitNess{ex1, ex2, ex3, ex4}
	// sort.Sort(ByAge(AllInOne))
	// for i, v := range AllInOne {
	// 	fmt.Println("Index : Value :", i, v)
	// }
	// **************************************************
	// JSon
	// unmarshall
	// type Scountry struct {
	// 	Name    string //`json:"Name"`
	// 	Country string //`json: "Country`
	// 	City    string //`json: "City"`
	// }

	// var country []Scountry

	// json1 := `[{
	// 	"Name":    "Web",
	// 	"Country": "Nz",
	// 	"City":    "Wellingston"
	// 	},{
	// 	"Name":    "Web",
	// 	"Country": "Nz",
	// 	"City":    "Wellingston"
	// 	}]`
	// bytjson := []byte(json1)
	// err := json.Unmarshal(bytjson, &country)
	// if err != nil {
	// 	fmt.Println("Error != nil", err)
	// }
	// fmt.Printf("%T%T \n", json1, bytjson)
	// for i, v := range country {
	// 	fmt.Printf("Index %d\n", i+1)
	// 	fmt.Printf("Name: %v, Country :%v, City : %v\n", v.Name, v.Country, v.City)
	// }

	// marshal
	// tryner := fitness{
	// 	Name:    "marcos",
	// 	Country: "spain",
	// 	Repe:    10,
	// }
	// tryner2 := fitness{
	// 	Name:    "Robb",
	// 	Country: "USA",
	// 	Repe:    5,
	// }

	// GeneralTryner := []fitness{
	// 	tryner,
	// 	tryner2,
	// }

	// fmt.Println(GeneralTryner)
	// byt, error := json.Marshal(GeneralTryner)
	// if error != nil {
	// 	fmt.Println("Error :", error)
	// }
	// fmt.Println(string(byt))
	// fmt.Println(byt)

	// **************************************************
	//Fibonacci
	// func (pr *rectangule) perimeter()int{
	// 	return pr .x+pr.y
	// }
	// func (pc *circle)perimeter()int{
	// 	return 2*Math.PI*pc.r
	// }
	// type shape interface{
	// 	area()int
	// 	perimeter() int
	// }

	// ***************************************
	// remove item
	// d := func(item int, array []int) {
	// 	array = append(array[:item], array[item+1:]...)
	// 	fmt.Println(array)
	// }
	// d(6, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9})

	//***************************************
	//Add element to array
	// d := func(item int, array []int) []int {
	// 	array = append(array, item)
	// 	//fmt.Println(array)
	// 	return array

	// }

	// fmt.Println(d(10, []int{1, 2, 3, 4}))

	// ************************************
	// print reverse array
	// a := func(array string) string {
	// 	reverse := ""
	// 	//for i := 0; i < len(array); i++ {
	// 	for _, v := range array {
	// 		reverse = string(v) + reverse
	// 	}
	// 	return reverse
	// }
	// fmt.Println(a("hola cabeza hueca"))

	// *****************************************************
	// video := func(array []int) (int, int) {
	// 	max := array[0]
	// 	min := array[0]
	// 	for i := 0; i < len(array); i++ {
	// 		if min > array[i] {
	// 			min = array[i]
	// 		}
	// 		if max < array[i] {
	// 			max = array[i]
	// 		}
	// 	}
	// 	return max, min
	// }
	// //aa := []int{-1, -2, -3}
	// max1, min1 := video([]int{-1, -2, -3})
	// fmt.Printf("Max number is : %v,\n Min number is : %v\n", max1, min1)

	// ********************************************
	// Is someone obese
	// BIM := func(w, h float64) string {
	// 	re := w / (h * h)
	// 	fmt.Println(re)
	// 	if re <= 18.5 {
	// 		return "UnderWeight"
	// 	}
	// 	if re > 18.5 && re <= 24.9 {
	// 		return "Normal"
	// 	} else {
	// 		return "Obese"
	// 	}

	// }
	// fmt.Println(BIM(60, 1.9))

	// ******************************************************
	// Print rectangule area
	// aa := func(l, w int) int {
	// 	return l * w
	// }(12, 13)
	// fmt.Print(aa)

	// ******************************************************
	// Print sum of two number
	// number := func(x1, x2 int) int {
	// 	return x1 + x2
	// }
	// fmt.Println(number(23, 45))

	//************************************************************

	// ******************************************************
	// Print name and Last name
	// f := func(name, last string) string {
	// 	return name + " " + last
	// }
	// fmt.Println(f("fulanito", "Menganito"))

	//************************************************************
	//video 117 ex
	// Per := persona{name: 111111}
	// fmt.Println(Per)
	// ChangeToMe(&Per)
	// fmt.Println(Per)

	// *******************************************************
	// video116 ex
	// value := 23
	// fmt.Printf("valeu %v, Type %T, Adress %v", value, value, &value)

	// **************************************************
	// pointer and Deference

	// a := 11
	// b := &a
	// *b = 123
	// fmt.Println("&a, a ", &a, a)
	// fmt.Printf(" type a %T  type &a %T\n", a, &a)
	// fmt.Printf(" b %v , &b %v, *b %v\n", b, &b, *b)
	// **************************************************************
	// s := "ksdjksjk"
	// fmt.Println("main before &s", s, &s)
	// aa := str_l(&s)
	// fmt.Println("main  after &s", s, &s)
	// fmt.Println("aa  &aa", aa, &aa)

	//	func str_l(s *string) *string {
	//		fmt.Println("inside first", &s, *s, s)
	//		*s = "Changed"
	//		fmt.Println("insider afters", &s, *s, s)
	//		return s
	//	}
	// *********************************************************
	// Sum NUmber
	// 	SumNumber := func(array ...int) int {
	// 		sum := 0
	// 		for _, v := range array {
	// 			sum += v
	// 		}
	// 		return sum
	// 	}
	// 	fmt.Println(" Sum is :", SumNumber(1, 2, 3, 4, 5, 6, 7, 8, 9))
	// }
	// **********************************************************?
	// summ odd and even number
	// 	sumNumber := func(array []int) (int, int) {
	// 		sumO := 0
	// 		sumE := 0
	// 		for _, v := range array {
	// 			if v%2 == 0 {
	// 				sumE += v
	// 			} else {
	// 				sumO += v
	// 			}
	// 		}
	// 		return sumE, sumO
	// 	}
	// 	e, o := sumNumber([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10})
	// 	fmt.Printf("Sum even number %d, Sum Odd number %d\n", e, o)
	//
}
