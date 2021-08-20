package main
// Reversort from Google's 2021 Code Jam Qualification Round
/* I attempted to figure out how to reverse a slice in Go
 after realizing I made a bubble sort style swap and was throwing
 the duress factor out of whack. The swap is returning some odd output
		[4/4]0xc000018100
		2
		[4/4]0xc000018100
		2
		[4/4]0xc000018100
		2
		[4/4]0xc000018100
		[1 1 1 3]
	Doesn't help that it nuked two of the values. it turns the 7 6 5 4 3 2 1
	into all ones, this isn't an issue with my input portion, just the sort.
	I decided to stop it here as I have other things to do.
	It may be disheartening to get hit out by the first question on the preliminary,
	but I learned a lot about golang in the course of a couple of hours.
	-Sam : Aug 20 2021

	Update: on the tmp passoff I mixed up the order, it still failed the test set but
	this was a good start.
*/

import (
    "fmt"
		_ "errors"
    )

//func populate_array (input int) []int {
//	var myArray [input]int
//	for i := 0; i < input; i++ {
//		fmt.Scanf("%d ", &myArray[i])
//	}
//	return myArray
//}

func main () {
    // first line is the amount of cases
		var cases int
		fmt.Scanf("%d\n", &cases)
		//fmt.Printf("Cases: %d\n", cases)
    // each case has 2 lines
		// Array must be flexible, a slice is to be defined first
		// slice capacity is flexible unlike arrays... okay?
		var caseLen int = 4
    for i := 0; i < cases; i++ {
			duress := 0
			// first is array length
			fmt.Scanf("%d\n", &caseLen)
			// second is array itself
			theArray := make([]int, caseLen, caseLen)
			for j := 0 ; j < caseLen ; j++ {
				fmt.Scanf("%d", &theArray[j])
			}
			// we have an array now! ( a slice but whatever )
			// compute the operation cost per case
			// reversort(L):
			for a := 0; a < caseLen - 1; a ++ {
				bvar := a
				for b := a+1; b < caseLen ; b++ {
					if (theArray[bvar] > theArray[b]){
						//fmt.Println("before", bvar)
						bvar = b
						//fmt.Println("after", bvar)
					}
				}
				duress = duress + 1 + bvar - a
				// reverse
				// if else here, if the subarray is even (bvar-a is odd) then bump up by 1
				mid := (bvar - a + 1)/2 + a
				//if((bvar-a)%2!=0){
				//	mid = mid + 1
				//}
				// set sub for the reverse routine
				var tmp int = -1
				for sub := 0; sub + a < mid + 1; sub++ {
					tmp = theArray[bvar-sub]
					theArray[bvar-sub] = theArray[a + sub]
					theArray[a + sub] = tmp
				}
				// bubble sort
				// this works,
				//tmp := theArray[bvar]
				//theArray[bvar] = theArray[a]
				//theArray[a] = tmp
				//fmt.Println(theArray)
			}
			//fmt.Println(theArray)
			fmt.Printf("Case #%d: %d\n", i+1, duress)
		}

}
