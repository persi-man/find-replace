package main

import (
	"bufio"
	"fmt"
	"os"
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
		occ += strings.Count(res, oldLower)
		occ += strings.Count(res, old)
		occ += strings.Count(res, oldUpper)
		//Result
		res = strings.Replace(line, oldLower, newLower, -1)
		res = strings.Replace(res, old, new, -1)
		res = strings.Replace(res, oldUpper, newUpper, -1)
	}

	return found, res, occ

}

/**
-FindReplaceFile searchrd for old and replace by new using ProcessLine in the file
-Return number of occurences,number of lines on the file with "Go" and an error
*/
func FindReplaceFile(src, old, new string) (occ int, lines []int, err error) {
	srcFile, err := os.Open(src)
	if err != nil {
		return occ, lines, err
	}
	defer srcFile.Close()
	old = old + " "
	new = new + " "

	indexLine := 1
	scanner := bufio.NewScanner(srcFile)

	for scanner.Scan() {
		found, res, oc := ProcessLine(scanner.Text(), old, new)
		if found {
			occ += oc
			lines = append(lines, indexLine)

		}
		fmt.Println(res)
		indexLine++

	}
	return occ, lines, nil
}

func main() {
	old := "Go"
	new := "Python"
	occ, lines, err := FindReplaceFile("wikigo.txt", old, new)
	if err != nil {
		fmt.Printf("Error while executing find replace  %v\n", err)
	}
	fmt.Println("---- Start replace stats ----")
	defer fmt.Println("---- End of Replace stats ----")
	fmt.Printf("Number of occurences of \"%v\":%v\n", old, occ)
	fmt.Printf("Number of lines :%d\n", len(lines))
	//Lines with "Go"
	fmt.Print("In lines [ ")
	len := len(lines)
	for i, j := range lines {
		fmt.Printf("%v", j)
		if i < len-1 {
			fmt.Printf(" - ")
		}
	}
	fmt.Println(" ]")

}
