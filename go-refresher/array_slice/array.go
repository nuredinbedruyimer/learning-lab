package arrayslice

import "fmt"

//  to create array in golang we gonna use this rule
//  var variable_name [size]type
//  where  the size is the number of elements in the array and type is the data type of the elements in the array

func ArraySample() {

	//  let me concatinate them to the one word

	arr := []int{1, 2, 3}

	for i := 0; i < 6; i++ {
		arr = append(arr, arr[len(arr)-1]+arr[len(arr)-2])

	}

	fmt.Println("Total", arr)

	fmt.Println("Length:", len(arr))
	fmt.Println("Capacity:", cap(arr))

}
