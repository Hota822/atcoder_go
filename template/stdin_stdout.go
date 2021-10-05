// ========================  default print  =============================
// yes or no
// slice output by type include new line by each values
// normal output, like int, string, etc...
import (
	"fmt"
)

func print(ans interface{}) {
	if v, ok := ans.(bool); ok {
		if v {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
		return
	}
	if sli, ok := ans.([]int); ok {
		for _, v := range sli {
			fmt.Println(v)
		}
		return
	}
	if sli, ok := ans.([]int64); ok {
		for _, v := range sli {
			fmt.Println(v)
		}
		return
	}
	if sli, ok := ans.([]string); ok {
		for _, v := range sli {
			fmt.Println(v)
		}
		return
	}
	fmt.Println(ans)
}


//============================ print array to one line ===============================
import (
	"fmt"
)

func print(ans interface{}) {
	if v, ok := ans.([]int); ok {
		for _, v := range ans {
			fmt.Print(v)
		}
	}
	if v, ok := ans.([]string); ok {
		for _, v := range ans {
			fmt.Print(v)
		}
	}
	fmt.Println("")
}

//============================ printf ===============================
// Printf => PRINT args formatted by supplied string
// Sprintf => RETURN args formatted by supplied string
// %v => usable all of types, parse by default format
// %g => float64, e.g. 67.333333333
