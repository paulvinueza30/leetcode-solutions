package strings

import "fmt"

func addStrings(num1, num2 string) string {
	num1Idx, num2Idx := len(num1)-1, len(num2)-1
	res := make([]byte, max(len(num1), len(num2)+1))
	carryOver := 0
	for num1Idx >= 0 && num2Idx >= 0 {
		num1Num, num2Num := byte('0'), byte('0')
		if num1Idx >= 0 {
			num1Num = num1[num1Idx]
		}
		if num2Idx >= 0 {
			num2Num = num2[num2Idx]
		}
		num1Int := int(num1Num - '0')
		num2Int := int(num2Num - '0')
		sum := num1Int + num2Int + carryOver
		sumString := fmt.Sprintf("%v", sum)
		fmt.Println(sumString)
		var resNum byte
		if sum > 9 {
			carryOver = 1
			resNum = sumString[1]
		} else {
			carryOver = 0
			resNum = sumString[0]
		}
		res[max(num2Idx, num1Idx)+1] = resNum - '0'
		num1Idx--
		num2Idx--
	}
	fmt.Println(res)
	if res[0] == '0' {
		return string(res[1:])
	} else {
		return string(res)
	}
}
