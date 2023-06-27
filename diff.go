package main

import (
	"fmt"
	"math"
)

type editOperation int

const (
    Insert editOperation = iota
    Replace
    Delete
)

type Diff struct {
    operation editOperation
    index int
    tochar string
}

type EditScript []Diff

func main() {
	before := "horse"
	after := "ros"

    _, script := diff(before, after);
    fmt.Println(script)


}

func min(nums ...int) int {
	min := math.MaxInt64
	for _, num := range nums {
		if num < min {
			min = num
		}
	}

	return min
}

func diff(before string, after string) (int, EditScript) {

    script := make([]Diff, 0);

	bsze := len(before) + 1
	asze := len(after) + 1

	dp := make([][]int, bsze)

	for i := range dp {
		dp[i] = make([]int, asze)
	}

	for i := 0; i < bsze; i++ {
		dp[i][0] = i
	}
	for i := 0; i < asze; i++ {
		dp[0][i] = i
	}

	for bi := 1; bi < bsze; bi++ {
		for ai := 1; ai < asze; ai++ {
			if before[bi-1] == after[ai-1] {
				dp[bi][ai] = dp[bi-1][ai-1]
			} else {
				r := dp[bi-1][ai-1] + 1
				i := dp[bi][ai-1] + 1
				d := dp[bi-1][ai] + 1

				dp[bi][ai] = min(r, i, d)

			}
		}
	}

	//Generating edit script
	i, j := bsze-1, asze-1

	for i >= 1 || j >= 1 {
		if i > 0 && j > 0 && before[i-1] == after[j-1] {
			i--
			j--
			continue
		}
		if i > 0 && dp[i][j] == dp[i-1][j]+1 {
            script = append(script, Diff{operation: Delete, index: i-1});
			i--
		} else if j > 0 && dp[i][j] == dp[i][j-1]+1 {
            script = append(script, Diff{operation: Replace, index: i-1, tochar: string(after[j-1])})
			j--
		} else if i > 0 && j > 0 {
            script = append(script, Diff{operation: Replace, index: i-1, tochar: string(after[j-1])});
			i--
			j--
		}
	}

    //The gen list is in reverse order
    for i, j := 0, len(script)-1; i < j; i, j = i+1, j-1 {
		script[i], script[j] = script[j], script[i]
	}

	return dp[bsze-1][asze-1], script

}
