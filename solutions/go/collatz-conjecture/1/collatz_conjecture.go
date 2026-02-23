package collatzconjecture
import "fmt"
func CollatzConjecture(n int) (int, error) {
	// panic("Please implement the CollatzConjecture function")

    if n == 1 {
        return 0,nil	
    } else if n <= 0 {
        return 0, fmt.Errorf("error")
    }
    step := 0
	calc := n
    for calc > 1 {
        if calc % 2 > 0 {
            calc = calc * 3 + 1
        } else {
            calc = calc / 2
        }
    	step++
    }
    return step, nil
}
