package main

import RED "RED/structure"

func main() {
	perso := RED.InitCharacter("Aragorn", "Ranger", 10, 100, 100, []string{"Arc", "Flèches", "Potion"})
	RED.DisplayInfo(perso)
	RED.AccessInventory(perso)
}
