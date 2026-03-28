package palindrome
import (
    "fmt"
    "sort"
)

// Define Product type here.
type Product struct{
    value int
    Factorizations [][2]int
}

func isPalindrome (num int) bool {
    original := num
	reversed := 0
    for num > 0{
        reversed = reversed*10 + num %10
        num /= 10
    }
    return original == reversed
}
func Products(fmin, fmax int) (Product, Product, error) {
    if fmin > fmax {
        return Product{}, Product{}, fmt.Errorf("fmin > fmax")
    }
    // var list  []int
    // for i := fmin ;i <= fmax ;i ++ {
    //     for j := i;j <= fmax;j ++{
    //         res := i * j
    //         list = append(list,res)
    //     }
    // }
    var palindromes []Product

    for i := fmin;i <= fmax;i++{
        for j := i;j <= fmax;j++{
            product := i * j
            if isPalindrome(product){
                found := false
                for index := range palindromes{
                    if palindromes[index].value == product{
                        palindromes[index].Factorizations = append(palindromes[index].Factorizations,[2]int{i,j})
                        found = true
                        break
                    }
                }
                if !found{
                    palindromes = append(palindromes,Product{
                        value : product,
                        Factorizations : [][2]int{{i,j}},
                    })
                }
            }
        }
    }
   
    if len(palindromes) == 0 {
        return Product{}, Product{}, fmt.Errorf("no palindromes")
    }

	sort.Slice(palindromes, func(i, j int) bool {
        return palindromes[i].value < palindromes[j].value
    })

	return palindromes[0], palindromes[len(palindromes)-1], nil
}

