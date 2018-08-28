package main

import (
	"fmt"
	"strconv"
)

func Dec2Bin(dec int8)[]int{
	bin := []int{0, 0, 0, 0, 0, 0, 0, 0}
	if(dec >= 0){
		bin[0] = 0
	}else{
		bin[0] = 1
		dec = 0 - dec
	}
	i := 0
	for (i<7){
		if(dec % 2) == 1{
			bin[7-i] = 1
			dec = dec / 2
		}else{
			bin[7 - i] = 0
			dec = dec / 2
		}
		i = i + 1
	}
	return bin
}

func Fan(Yuan []int, dec int8)[]int{
	if (dec >= 0){
		return Yuan
	}else{
		i := 1
		for (i < 8){
			if (Yuan[i] == 1){
				Yuan[i] = 0
			}else{
				Yuan[i] = 1
			}
			i = i + 1
		}
		return Yuan
	}
}

func Bu(fan []int, dec int8)[]int{
	if(dec < 0){
		fan[7] = fan[7] + 1
	}
	i := 7
	for (i > 0){
		if (fan[i]) == 2{
			fan[i] = fan[i] - 2
			fan[i - 1] = fan[i - 1] + 1
		}
		i = i - 1
	}
	return fan
}

func test(nums ...int8){
	var a2 []int
	for _, i := range nums{
		tmp := Bu(Fan(Dec2Bin(i), i), i)
		 for _, t :=range tmp{
			 a2 = append(a2, t)
		 }
	}
	var a3 string
	isMinus := false
	if(a2[0] ==1){
		isMinus = true
	}
	for _, item := range a2{
		if(isMinus){
			if (item == 1){
				a3 = a3 + "0"
			}else{
				a3 = a3 + "1"
			}
		}else{
			if (item == 1){
				a3 = a3 + "1"
			}else{
				a3 = a3 + "0"
			}
		}
	}
	fmt.Println(a3)
	if(isMinus){
		result, _ := strconv.ParseInt(a3, 2, 64)
		result = 0 - result
		fmt.Println(result)
	}else{
		result, _ := strconv.ParseInt(a3, 2, 64)
		fmt.Println(result)
	}
}

func main(){
	var a =[]int8{71, -28, -61, 100, -48, 85, -18, 17, -67, -1, 33, 73, 36, 39, 99, 125, -12, 66, 81, 1}
	test(a...)

}