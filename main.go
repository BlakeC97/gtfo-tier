package main

import (
	"cmp"
	"fmt"
	"gtfo-tier/tiering"
	"slices"
	"strings"
)

func weaponsToStrings(weapons []tiering.Weapon) []string {
	weaponStrings := make([]string, 0, len(weapons))
	for _, w := range weapons {
		weaponStrings = append(weaponStrings, string(w))
	}
	return weaponStrings
}

func main() {
	tierLists := tiering.TierLists()

	type pair struct {
		situation tiering.Situation
		tierList  tiering.TierList
	}

	tierListsSorted := make([]pair, 0, len(tierLists))
	for situation, tierList := range tierLists {
		tierListsSorted = append(tierListsSorted, pair{situation, tierList})
	}
	slices.SortStableFunc(tierListsSorted, func(a, b pair) int {
		return cmp.Compare(a.situation, b.situation)
	})

	for _, tl := range tierListsSorted {
		situation := tl.situation
		tierList := tl.tierList
		fmt.Println(situation)
		fmt.Printf("\tBest: %s\n", strings.Join(weaponsToStrings(tierList.Best), ", "))
		fmt.Printf("\tGood: %s\n", strings.Join(weaponsToStrings(tierList.Good), ", "))
		fmt.Printf("\tDecent: %s\n", strings.Join(weaponsToStrings(tierList.Decent), ", "))
		fmt.Printf("\tNiche: %s\n", strings.Join(weaponsToStrings(tierList.Niche), ", "))
		fmt.Printf("\tDon't: %s\n", strings.Join(weaponsToStrings(tierList.DoNot), ", "))
		fmt.Println()
	}
}
