package malsin  
  
import (  
    "fmt"  
    "sort"  
)  
  
type Person struct {  
    name string  
    age  int  
}  
  
type persons []Person  
  
func (ps persons) Len() int {  
    return len(ps)  
}  
  
func (ps persons) Less(i, j int) bool {  
    return ps[i].age < ps[j].age  
}  
  
func (ps persons) Swap(i, j int) {  
    ps[i], ps[j] = ps[j], ps[i]  
}  
  
func main() {  
    ps := Persons{  
        {"larry", 19},  
        {"jackey", 18},  
        {"lucy", 20},  
    }  
    sort.Sort(ps)  
    fmt.Println(ps)  
} 