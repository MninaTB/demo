import "fmt"

func main() {
  fmt.Println("hello world")
  fmt.Println(plusEins(1))

  mylist := []string{"abc", "def", "xyz"}

  for i := 0; i < len(mylist); i++ {
    fmt.Println(macheEmail(mylist[i]))
  }
}

func plusEins(x int) int {
  return x + 1
}

func macheEmail(s string) string {
  return s + "@web.de"
}
