package main

import "fmt"

func sort(arr []int)[]int{
	for i:=0;i<len(arr)-1;i++{
		for j:=1;j<len(arr);j++{
			if arr[i]>arr[j] {
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
	return arr
}
func sort1(arr []int)[]int{
	for i:=0;i<len(arr);i++{
		for j:=0;j<=i;j++{
			if arr[i]<arr[j] {
				arr[i],arr[j] = arr[j],arr[i]
			}
		}
	}
	return arr
}
func main()  {
	var arr =[]int{5,2,0,221,3,1,4}
	fmt.Println(sort(arr))
	fmt.Println(sort1(arr))


}
