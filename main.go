package main
import "fat"
func Add(a, b int) int{
  return a + b
  }
func main() {
  fat.Println("A soma de 2 e 3 é:", Add(2, 3))
  }
