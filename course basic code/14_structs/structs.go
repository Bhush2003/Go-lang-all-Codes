package main

import "fmt"

type me struct {
	id     int
	status string
	name   string
}

func (m *me) chanhgeStatus(status string){
    m.status=status
}
func obj(id int,status string, name string) *me{
    myObject:=me{
        id:id,
        status: status,
        name: name,
    }
    return &myObject
}
func main() {

	var myOrder me = me{
		id:     10,
		status: "paid",
		name:   "protin",
	}

	fmt.Println(myOrder.id)
    fmt.Println(myOrder.status)
	fmt.Println(myOrder.name)
    fmt.Println(myOrder)

    myObj:=obj(10,"cnf","bhushan")

    fmt.Println(myObj)


}