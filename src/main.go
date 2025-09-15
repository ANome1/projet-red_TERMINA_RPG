package main

import RED "RED/structure"

func main() {
	perso := RED.InitCharacter("Doby", "Elfe", 10, 100, 40, []string{"Arc", "Flèches", "Potion de soin"})
	RED.DisplayInfo(perso)
	RED.AccessInventory(perso)
	RED.TakePotion(&perso, "Potion de soin")
	RED.DisplayInfo(perso)
	RED.AccessInventory(perso)
}
