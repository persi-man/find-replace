package main

import (
	"fmt"
	"strings"
)

/**

- ProcessLine searches for old in line and replace it by new
- Return found = true if the pattern was found , res for the result of string and
occ for the number if occurence of old
*/
func ProcessLine(line, old, new string) (found bool, res string, occ int) {
	//Traitment of lower go
	oldLower := strings.ToLower(old)
	newLower := strings.ToLower(new)
	oldUpper := strings.ToUpper(old)
	newUpper := strings.ToUpper(new)
	res = line

	if strings.Contains(line, oldLower) || strings.Contains(line, old) || strings.Contains(line, oldUpper) {
		found = true
		//Number of occurences
		occ += strings.Count(line, oldLower)
		occ += strings.Count(line, old)
		occ += strings.Count(line, oldUpper)
		//Result
		res = strings.Replace(line, oldLower, newLower, -1)
		res = strings.Replace(line, old, new, -1)
		res = strings.Replace(line, oldUpper, newUpper, -1)
	}

	return found, res, occ

}

func main() {
	_, res, occ := ProcessLine("Go was publicly announced in November 2009,[23] and version 1.0 was released in March 2012.[24][25] Go is widely used in production at Google[26] and in many other organizations and open-source projects .In April 2018, the original logo (Gopher mascot) was replaced with a stylized GO slanting right with trailing streamlines. However, the mascot remained the same.[27]", "Go", "Python")
	fmt.Printf("Result:%s\n Number of occurences:%d\n", res, occ)
}
