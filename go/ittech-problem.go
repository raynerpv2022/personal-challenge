package main

// Crear lista de eventos valor de entrada
// ordenar la lista por fecha
// crear una lista  de map-diccionario donde clave sea maquina y valor un slice de los usuarios logueados
// si el usuario esta logeado incluirlo en la lista, si no esta logueado eliminarlo
// imprimir la lista con las maquinas y los usuarios ordenados

import (
	"fmt"
	"math/rand"
	"sort"
	"strconv"
	"time"
)

type Event struct {
	date    time.Time
	user    string
	machine string
	typeL   string
}

type ByDate []Event

func (q ByDate) Len() int           { return len(q) }
func (q ByDate) Swap(i, j int)      { q[i], q[j] = q[j], q[i] }
func (q ByDate) Less(i, j int) bool { return q[i].date.Before(q[j].date) }

// func Getdate(event Event) time.Time {
// 	return event.date
// }

func removeSlice(first []string, i int) []string {
	if i == -1 {
		return first
	}
	copy(first[i:], first[i+1:])
	return first[:len(first)-1]

}

func IndexItem(slice []string, a string) int {
	for i, v := range slice {
		if v == a {
			return i
		}
	}
	return -1

}

func Current_Machine(List_Event []Event) map[string][]string {
	map_machine := map[string][]string{}
	sort.Sort(ByDate(List_Event))

	for _, v := range List_Event {
		userArray := []string{}
		if _, ok := map_machine[v.machine]; !ok {

			map_machine[v.machine] = userArray
		}
		if v.typeL == "LOGUIN" {
			map_machine[v.machine] = append(map_machine[v.machine], v.user)

		} else if v.typeL == "LOGOUT" {
			map_machine[v.machine] = removeSlice(map_machine[v.machine], IndexItem(map_machine[v.machine], v.user))
		}
	}
	return map_machine

}

func CreateEventList() []Event {
	ListEvent := []Event{}
	machine := []string{"Linux Server", "MAC server", "Windows 2020 server", "UBUNTU", "DEBIAN"}
	typel := []string{"LOGUIN", "LOGOUT"}
	for i := 0; i <= 100; i++ {
		event := Event{
			date:    time.Now(),
			user:    "User" + strconv.Itoa((rand.Intn(200))),
			machine: machine[rand.Intn(5)],
			typeL:   typel[rand.Intn(2)],
		}

		ListEvent = append(ListEvent, event)

	}

	return ListEvent
}
func PrintReport(List_Event []Event, Map_Machine map[string][]string) {
	fmt.Println("All Data")
	for _, v := range List_Event {
		fmt.Printf("User %v\t Machine %v \t TYpe %v\t\n", v.user, v.machine, v.typeL)
	}
	fmt.Println("Current User Loguin ")
	for i, v := range Map_Machine {
		fmt.Printf("%v \t %v \n", i, v)
	}

}

func main() {

	List_Event := CreateEventList()

	Map_machine := Current_Machine(List_Event)
	PrintReport(List_Event, Map_machine)

}
