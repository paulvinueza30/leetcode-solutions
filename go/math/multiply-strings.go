package main

// wont work for big numbers
// func multiply(num1 string, num2 string) string{
//  numA := helper(num1)
//  numB := helper(num2)
//  product := numA * numB
//  res := []byte{}
//  if product == 0 {
// 	return "0"
//  }
//  for product != 0{
// 	rightMost := byte('0' + (product % 10))
// 	product /= 10
// 	res = append(res, rightMost)
//  }
//  for i, j := 0, len(res) - 1 ; i < j ; i, j = i + 1, j - 1{
// 	res[i] , res[j] = res[j] , res[i]
//  }
//  return string(res)
// }

// func helper(s string) int{
// 	res := 0

// 	for _ , c := range s{
// 		num := c - '0'
// 		res *= 10
// 		res += int(num)
// 	}
// 	return res
// }

func multiply(num1 string , num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}
	m, n := len(num1) , len(num2)
	res := make([]int, m + n)
	for i := m -1 ; i >= 0 ; i--{
		for j := n - 1 ; j >=0 ; j--{
			a := int('0' + num1[i])
			b := int('0' + num2[j])
			prod := a * b
			pos1 , pos2 := i + j , i + j + 1
			total := res[pos2] + prod
			res[pos1] += total / 10
			res[pos2] = total % 10
		}
	}
	leading := true
	bytes := []byte{}
	for _, n := range res{
		if leading && n == 0 {
			continue
		}
		leading = false
		bytes = append(bytes, byte(n) + '0')
	}
	return string(bytes)
}