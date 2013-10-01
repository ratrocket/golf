package main

import "fmt"

type Status int

const (NEW Status = iota; CONFIRMED; INVITED)

func (s Status) String() string {
        switch s {
        case NEW: return "new"
        case CONFIRMED: return "confirmed"
        case INVITED: return "invited"
        }
        return "OOPS"
}

func main() {
        fmt.Println(NEW)
}
