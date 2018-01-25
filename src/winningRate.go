//根据当前总场次以及胜场计算出期望胜率的总场次以及胜场

package main

import "fmt"


func wingningRate(win, total, rate float64) (float64, float64){
	for win / total < rate {
		win++
		total++
	}
	return win, total
}

func main(){
	a, b := wingningRate(529, 757, 0.70)
	fmt.Printf("期望胜率: %v%%\n", int(a/b * 100))
	fmt.Printf("胜场: %v, 总场次: %v\n", a, b)
}
