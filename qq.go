// package main

// import (
//     "container/list"
//     "fmt"
// )

// func main() {
//     // Create a new list and put some numbers in it.
//     lst := list.New()
//     // Last Element
//     e4 := lst.PushBack("Welcome")
//     // First Element
//     e1 := lst.PushFront("One")
//     // Insert before e4
//     lst.InsertBefore("Testing", e4)
//     // Insert after e1
//     lst.InsertAfter("Slow", e1)

//     // Iterate through list and print its contents.
//     for e := lst.Front(); e != nil; e = e.Next() {
//         fmt.Println(e.Value)
//     }
// }

/////////////////

// package main

// import (
//     "container/list"
//     "fmt"
// )

// func displayList(lst *list.List) {
//     fmt.Println("~~~~~~~~~~~~~~~~~~")
//     // Iterate through list and print its contents.
//     for e := lst.Front(); e != nil; e = e.Next() {
//         fmt.Println(e.Value)
//     }
// }
// func main() {
//     // Step 1: Create a new lisl
//     lst := list.New()
//     /* Step 2: Populating the list */
//     e3 := lst.PushFront("Three")
//     e2 := lst.PushFront("Two")
//     e1 := lst.PushFront("One")
//     e0 := lst.PushFront("Zero")
//     displayList(lst)
//     lst.MoveAfter(e1, e2)
//     displayList(lst)
//     lst.MoveBefore(e0, e3)
//     displayList(lst)
//     lst.MoveToFront(e0)
//     lst.MoveToBack(e3)
//     displayList(lst)
// }

////////////
// package main

// import (
//     "flag"
//     "fmt"
// )

// func main() {
//     var username string
//     flag.StringVar(&username, "user", "rakesh", "user name")

//     heightPtr := flag.Int("height", 68, "height")
//     metricFlagPtr := flag.Bool("metric", true, "metric / imperial")

//     flag.Parse()

//     fmt.Println("username:", username)
//     fmt.Println("height:", *heightPtr)
//     fmt.Println("fork:", *metricFlagPtr)
//     fmt.Println("Remaining:", flag.Args())
// }

////////////////////
package main

import (
    "fmt"
    "path"
)

func main() {
    fmt.Println(path.Base("/etc/gdm"))
    fmt.Println(path.Base("/"))
    fmt.Println(path.Base("/home/mayank/go"))
    fmt.Println(path.Base("C://Windows/System32"))
    fmt.Println(path.Base("C:\\Windows\\System32"))
}