package main
import "fmt"

func main(){
// declaring of slice through short declaration method	
	sc := [] int{2,45,3,6,7,21}
	fmt.Println(sc)
	fmt.Println("The length of the slice is = ", len(sc))
	fmt.Println("The cap of the slice is = ",cap(sc))
	arr := [5] int{23,44,57,32,2}
	s := arr[2:4]//meaning start at index 2 and end at index 4 of the array
    fmt.Println(s)

	sl := arr[0:4]
	fmt.Println("\nArray before changing the value in the slice\n",arr)
	//slices share the same memomry with the arrary ,any changes made through the slices will effect the array
	sl[2]= 455
	fmt.Println("\nArray after  changing the value in the slice\n",arr)

    numbers := [] int{23,344,4,3,2,76}
	fmt.Println("\nSlice before append : ",numbers)
	numbers=append(numbers,678)
	numbers=append(numbers,10543)
	fmt.Println("\nSlice After append : ",numbers)

	//num := make([]data_type,len,cap) use to create a empty slice 

	num := make([]int,4,8)
	fmt.Println(num)

}