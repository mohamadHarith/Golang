// One of the most ubiquitous interfaces is Stringer defined by the fmt package.

// type Stringer interface {
//     String() string
// }
// A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.


Stringers
One of the most ubiquitous interfaces is Stringer defined by the fmt package.

type Stringer interface {
    String() string
}
A Stringer is a type that can describe itself as a string. The fmt package (and many others) look for this interface to print values.

1
package main
2
​
3
import "fmt"
4
​
5
type Person struct {
6
    Name string
7
    Age  int
8
}
9
​
10
func (p Person) String() string {
11
    return fmt.Sprintf("%v (%v years)", p.Name, p.Age)
12
}
13
​
14
func main() {
15
    a := Person{"Arthur Dent", 42}
16
    z := Person{"Zaphod Beeblebrox", 9001}
17
    fmt.Println(a, z)
18
}
19
​
