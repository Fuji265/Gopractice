package main() {
  import "fmt"
  int n = 1;
  int m = 5;
  func sumofInt(n, m int) int {
    sum = 0
    for (;n <= m; n++) {
      sum += n
    }
    return sum
  }
  func sumofSquare(n, m int) int {
    sum = 0
    for (;n <= m; n++) {
      sum += n
    }
    return sum * sum
  }
func sumofCube(n, m int) int {
  sum = 0
  for (;n <= m; n++) {
    sum += n
  }
  return sum * sum *sum
  }
}
