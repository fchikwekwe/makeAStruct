package main

import "fmt"
type Thief struct {
    id int
    name string
    tools []string
    experience int
    skills []string
    hideoutCoord []int
}

func stealTreasure (tools []string, skills []string, exp int) bool {
    fmt.Printf("skills: %v ,tools: %v , years experience: %v.", skills, tools, exp)
    if exp > 2 {
        return true
    }
    return false
}

func main () {
    var tools[]string
    tools[0] = "lockpick"
    tools[1] = "knife"

    var skills[]string
    skills[0] = "lockpicking"

    var exp int = 3

    var hideout[]int
    hideout[0] = 15
    hideout[1] = 20

    // t := Thief{1, "bob", tools, exp, skills, hideout}

    stealTreasure(tools, skills, exp)
}
