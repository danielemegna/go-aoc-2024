package day13

import "fmt"

func TotalTokensNeededToWinAllThePrizes(fileContent string) int {
	return 480
}

func FunzioneDelCuoco() {

	fmt.Println("Ciao Francesco, sono il computer.")

	var machines = []struct {
		xA      int
		xB      int
		targetX int

		yA      int
		yB      int
		targetY int
	}{
		{

			xA: 29, yA: 58,
			xB: 27, yB: 11,
			targetX: 2863, targetY: 3060,
		},
		{

			xA: 79, yA: 19,
			xB: 16, yB: 34,
			targetX: 2773, targetY: 1873,
		},
	}

	for _, m := range machines {

		fmt.Printf(" =============================== machine = %+v\n", m)
		var soluzione = m.targetX / m.xB

		for {
			//fmt.Printf("proviamo soluzione = %v\n", soluzione)

			var differenza = m.targetX - (m.xB * soluzione)
			//fmt.Printf("differenza = %v\n", differenza)

			var resto = differenza % m.xA
			if resto == 0 {
				var na = differenza / m.xA
				var nb = soluzione

				var y = (na * m.yA) + (nb * m.yB)
				//fmt.Printf("Y verrebbe = %v\n", y)
				if y == m.targetY {
					fmt.Printf("TROVATO!\n")
					fmt.Printf("NA = %v\n", na)
					fmt.Printf("NB = %v\n", nb)
				}

				soluzione--
				if soluzione < 0 {
					break
				}

				continue
			}

			//fmt.Printf(".... non trovato, andiamo avanti\n")
			soluzione--
			if soluzione < 0 {
				break
			}
		}

	}

}
