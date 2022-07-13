package main

import (
	"fmt"
	"math/rand"
	"strconv"
)

type Entity struct {
	Health  int
	Attack  int
	Defense int
}
type Hero struct {
	Entity
	MaxHealth int
	Potions   [5]int
	Gold      int
}

func main() {
	PlayGame()
}

//resets the Hero but mantains gold
func (h *Hero) MakeHero() int {
	tHealth := rand.Intn(11) + 20
	h.MaxHealth = tHealth
	h.Health = tHealth
	h.Defense = rand.Intn(2) + 1 //I change the defense  so there is no a chance for the game to run infinitely ex: (hero has 3 or more defense would make all possible goblins do 0 dmg)
	h.Attack = rand.Intn(2) + 3  //changed for same resaon above
	h.Potions = [5]int{}
	return tHealth
}

//make a goblin
func (g *Entity) MakeGoblin() {
	g.Health = rand.Intn(6) + 5
	g.Attack = rand.Intn(3) + 2
	g.Defense = rand.Intn(2) + 1
}

//main game loop
func PlayGame() {
	h := new(Hero)
replay:
	var gobKills, lvl, steps int
	h.MakeHero()
	for h.Health != 0 {
		kill, lCheck, alive, pStep := h.OutOfComOp(steps)
		steps = pStep
		if !alive {
		renter:
			fmt.Println("You have died, would you like to play again?(y or n)")
			var input string
			fmt.Scanln(&input)
			if input == "y" {
				goto replay
			} else if input == "n" {
				fmt.Println("You reached level", lvl, "You killed", gobKills, "Goblins")
				break
			} else {
				fmt.Println("Please input y or n")
				goto renter
			}
		}
		if kill {
			gobKills++
		}
		if lCheck {
			lvl++
		}

	}

}

//prints out info, handles auto walking and single walking
func (h *Hero) OutOfComOp(steps int) (bool, bool, bool, int) {
top:
	var lvl bool
	alive := true
	var hasPot, canShop bool
	var pCount int
	for i := 0; i < len(h.Potions); i++ {
		if h.Potions[i] == 2 {
			pCount++
		}
	}
	fmt.Println("Step:", steps, "Health", h.Health, "Potions:", pCount, "Gold:", h.Gold)
	fmt.Println("Input 1 to Take one step")
	fmt.Println("Input 2 to take steps until you find a goblin or find a potion shop")
	if h.Potions[0] == 2 {
		fmt.Println("Input 3 to drink a potion")
		hasPot = true
	}

	if steps%10 == 0 && steps != 0 {
		fmt.Println("Input 4 to go to the potion shop")
		canShop = true
		lvl = true
	}
	var input string
	fmt.Scanln(&input)
	x, _ := strconv.Atoi(input)
	if x == 1 {
		if rand.Intn(100) >= 60 {
			steps++
			fmt.Println("You found a Goblin, You are on step", steps)
			h.InComOp()
			if h.Health < 1 {
				alive = false
			}
		} else {
			steps++
			fmt.Println("No Goblin found, You are on step", steps)
			return false, lvl, alive, steps
		}

	} else if x == 2 {
		var flag, stop bool
		for !stop {
			if rand.Intn(100) >= 60 {
				steps++
				fmt.Println("You found a Goblin, You are on step", steps)
				h.InComOp()
				if h.Health < 1 {
					alive = false
				}
				stop = true
			} else {
				steps++
				fmt.Println("No Goblin found, You are on step", steps)
				if steps%10 == 0 {
					flag = true
					break
				}
			}
		}
		if flag {
			return false, lvl, alive, steps
		}
	} else if x == 3 && hasPot {
		if h.Potions[0] == 0 {
			fmt.Println("You have no Potions, you can by more at the potion shop after every tenth step")
		} else if h.Health == h.MaxHealth {
			fmt.Println("You are already at max Health")
		} else {
			for i := len(h.Potions) - 1; i >= 0; i-- {
				if h.Potions[i] == 2 {
					h.Potions[i] = 0
					h.Health += 2
					if h.Health > h.MaxHealth {
						h.Health = h.MaxHealth
					}
					fmt.Println("Current health is", h.Health)
					break
				}
			}
		}
		goto top

	} else if x == 4 && canShop {
		h.PotionShop(steps)
		goto top
	} else {
		fmt.Println("Please input one of the options")
		goto top
	}
	return true, lvl, alive, steps
}

//handle selling potions
func (h *Hero) PotionShop(steps int) {
top:
	if h.Potions[len(h.Potions)-1] == 2 {
		fmt.Println("you already have the max amount of potions")
	} else {
		var missingPots int
		for i := len(h.Potions) - 1; i >= 0; i-- {
			if h.Potions[i] == 0 {
				missingPots++
			} else {
				break
			}
		}
		fmt.Println("how many potions would you like to buy? You can buy up to", missingPots)
		var input string
		fmt.Scanln(&input)
		x, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please input an int")
			goto top
		} else if x > missingPots {
			fmt.Println("You can not have more than 5 potions")
			goto top
		} else if h.Gold >= x*4 {
			var count int
			for i := 0; i < len(h.Potions); i++ {
				if h.Potions[i] == 0 {
					h.Potions[i] = 2
					h.Gold -= 4
					count++
					if x == count {
						break
					}
				}

			}
			fmt.Println(h.Potions)

		} else {
			fmt.Println("You do not have enough gold to by that many potions")
			goto top
		}

	}

}

//handles fighting
func (h *Hero) InComOp() {
	g := new(Entity)
	g.MakeGoblin()

turn:

	fmt.Println("Input 1 to Attack once")
	fmt.Println("Input 2 to Auto Attack (you will attack untill either you or the goblin is dead)")

	var input string
	fmt.Scanln(&input)
	x, err := strconv.Atoi(input)
	if err != nil {
		fmt.Println("Input must be an int")
	} else if x == 1 {
		h.Health, g.Health = h.Health-(g.Attack-h.Defense), g.Health-(h.Attack-g.Defense)
		fmt.Println("You have ", h.Health, "health and the Goblin has", g.Health, "health")
		goto turn

	} else if x == 2 {
		for h.Health > 0 {
			h.Health, g.Health = h.Health-(g.Attack-h.Defense), g.Health-(h.Attack-g.Defense)
			fmt.Println("You have ", h.Health, "health and the Goblin has", g.Health, "health")
			if g.Health < 1 {
				break
			}
		}
		if g.Health <= 0 {
			fmt.Println("Congrats you kill the goblin and found 2 Gold")
			h.Gold += 2
		}
	} else {
		fmt.Println("Please Input one of the options")
		goto turn
	}

}
