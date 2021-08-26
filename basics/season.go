package main

import "fmt"

func main() {
	fmt.Println(Season(3))     // calling function to find the season
}

func Season(month int) string {
	switch month {  // switch on the basis of value of the months(1-12)
		case 12,1,2:  return "Winter"  // Jan, Feb and Dec have winter
		case 3,4,5:	  return "Spring"  // March, Apr and May have spring
		case 6,7,8:   return "Summer"  // June, July and Aug have summer
		case 9,10,11: return "Autumn"  // Sept, Oct and Nov have autumn

		default: return "Season unknown"         //value outside [1,12], then season is unkown
	}
}
