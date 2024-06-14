package tiering

import (
	"testing"
)

func TestTierListsContainsEachSpecialWeapon(t *testing.T) {
	allWeaponSet := make(map[Weapon]struct{}, len(AllWeapons))
	for _, weapon := range AllWeapons {
		allWeaponSet[weapon] = struct{}{}
	}

	for situation, tierList := range TierLists() {
		missingWeapons := make([]Weapon, 0, 8)
		containedWeaponSet := make(map[Weapon]struct{}, len(AllWeapons))

		for _, tier := range [][]Weapon{tierList.Best, tierList.Good, tierList.Decent, tierList.Niche, tierList.DoNot} {
			for _, weapon := range tier {
				containedWeaponSet[weapon] = struct{}{}
			}
		}

		for weapon, _ := range allWeaponSet {
			if _, ok := containedWeaponSet[weapon]; !ok {
				missingWeapons = append(missingWeapons, weapon)
			}
		}

		if len(missingWeapons) > 0 {
			t.Fatalf("Missing weapons in situation '%v': %v", situation, missingWeapons)
		}
	}
}
