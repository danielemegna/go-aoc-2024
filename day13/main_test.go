package day13

import "fmt"

import (
	"testing"
)

func TestDelCuoco(t *testing.T) {

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
			xA:      76,
			xB:      13,
			targetX: 4349,
			yA:      95,
			yB:      94,
			targetY: 12356,
		},
		{
			xA:      32,
			yA:      21,
			xB:      18,
			yB:      94,
			targetX: 2008,
			targetY: 4934,
		},
		{
			xA: 12, yA: 40,
			xB: 76, yB: 46,
			targetX: 18260, targetY: 11116,
		},
		{

			xA: 15, yA: 64,
			xB: 84, yB: 48,
			targetX: 7728, targetY: 9072,
		},
		{

			xA: 18, yA: 42,
			xB: 49, yB: 25,
			targetX: 4558, targetY: 3310,
		},
		{

			xA: 63, yA: 23,
			xB: 16, yB: 56,
			targetX: 16903, targetY: 19663,
		},
		{

			xA: 49, yA: 83,
			xB: 59, yB: 20,
			targetX: 5806, targetY: 8076,
		},
		{

			xA: 56, yA: 14,
			xB: 19, yB: 68,
			targetX: 1983, targetY: 7212,
		},
		{

			xA: 74, yA: 26,
			xB: 24, yB: 62,
			targetX: 3980, targetY: 6648,
		},
		{

			xA: 18, yA: 44,
			xB: 45, yB: 15,
			targetX: 16172, targetY: 13856,
		},
		{
			xA: 54, yA: 84,
			xB: 42, yB: 11,
			targetX: 17114, targetY: 9700,
		},
		{
			xA: 81, yA: 57,
			xB: 15, yB: 64,
			targetX: 5295, targetY: 6719,
		},
		{

			xA: 16, yA: 46,
			xB: 73, yB: 39,
			targetX: 8604, targetY: 15288,
		},
		{

			xA: 30, yA: 64,
			xB: 30, yB: 15,
			targetX: 3890, targetY: 11425,
		},
		{

			xA: 11, yA: 55,
			xB: 60, yB: 28,
			targetX: 18461, targetY: 16257,
		},
		{

			xA: 93, yA: 47,
			xB: 31, yB: 77,
			targetX: 1333, targetY: 919,
		},
		{

			xA: 51, yA: 68,
			xB: 77, yB: 15,
			targetX: 8370, targetY: 5637,
		},
		{

			xA: 11, yA: 47,
			xB: 63, yB: 30,
			targetX: 3441, targetY: 12927,
		},
		{

			xA: 14, yA: 25,
			xB: 42, yB: 14,
			targetX: 15418, targetY: 15509,
		},
		{

			xA: 36, yA: 55,
			xB: 43, yB: 19,
			targetX: 9673, targetY: 17061,
		},
		{

			xA: 15, yA: 51,
			xB: 32, yB: 14,
			targetX: 15661, targetY: 13339,
		},
		{

			xA: 90, yA: 59,
			xB: 19, yB: 73,
			targetX: 7414, targetY: 5829,
		},
		{

			xA: 89, yA: 35,
			xB: 47, yB: 65,
			targetX: 2337, targetY: 2175,
		},
		{

			xA: 55, yA: 19,
			xB: 13, yB: 33,
			targetX: 16671, targetY: 1691,
		},
		{

			xA: 41, yA: 11,
			xB: 13, yB: 57,
			targetX: 13290, targetY: 18772,
		},
		{

			xA: 73, yA: 30,
			xB: 15, yB: 49,
			targetX: 3462, targetY: 7552,
		},
		{

			xA: 64, yA: 20,
			xB: 15, yB: 59,
			targetX: 10670, targetY: 12078,
		},
		{

			xA: 20, yA: 43,
			xB: 62, yB: 41,
			targetX: 18332, targetY: 16063,
		},
		{

			xA: 61, yA: 18,
			xB: 18, yB: 72,
			targetX: 3251, targetY: 18170,
		},
		{

			xA: 25, yA: 48,
			xB: 56, yB: 12,
			targetX: 5352, targetY: 3876,
		},
		{

			xA: 19, yA: 79,
			xB: 62, yB: 13,
			targetX: 15824, targetY: 688,
		},
		{

			xA: 60, yA: 93,
			xB: 66, yB: 18,
			targetX: 8340, targetY: 9555,
		},
		{

			xA: 59, yA: 47,
			xB: 12, yB: 57,
			targetX: 3331, targetY: 3460,
		},
		{

			xA: 20, yA: 99,
			xB: 72, yB: 20,
			targetX: 868, targetY: 2951,
		},
		{

			xA: 67, yA: 30,
			xB: 22, yB: 59,
			targetX: 3077, targetY: 2115,
		},
		{

			xA: 41, yA: 23,
			xB: 27, yB: 54,
			targetX: 7406, targetY: 18395,
		},
		{

			xA: 19, yA: 54,
			xB: 61, yB: 30,
			targetX: 10935, targetY: 9338,
		},
		{

			xA: 76, yA: 53,
			xB: 14, yB: 32,
			targetX: 15124, targetY: 12677,
		},
		{

			xA: 50, yA: 60,
			xB: 92, yB: 13,
			targetX: 7050, targetY: 6025,
		},
		{

			xA: 11, yA: 91,
			xB: 53, yB: 30,
			targetX: 2304, targetY: 4356,
		},
		{

			xA: 18, yA: 52,
			xB: 48, yB: 16,
			targetX: 18176, targetY: 9008,
		},
		{

			xA: 37, yA: 14,
			xB: 32, yB: 53,
			targetX: 6571, targetY: 14640,
		},
		{

			xA: 25, yA: 12,
			xB: 50, yB: 85,
			targetX: 3850, targetY: 6423,
		},
		{

			xA: 51, yA: 43,
			xB: 78, yB: 13,
			targetX: 7602, targetY: 3613,
		},
		{

			xA: 34, yA: 29,
			xB: 85, yB: 13,
			targetX: 6970, targetY: 2732,
		},
		{

			xA: 44, yA: 96,
			xB: 76, yB: 52,
			targetX: 4696, targetY: 7628,
		},
		{

			xA: 17, yA: 62,
			xB: 63, yB: 52,
			targetX: 2237, targetY: 2470,
		},
		{

			xA: 74, yA: 22,
			xB: 13, yB: 64,
			targetX: 8262, targetY: 9536,
		},
		{

			xA: 26, yA: 69,
			xB: 37, yB: 11,
			targetX: 3668, targetY: 17503,
		},
		{

			xA: 57, yA: 12,
			xB: 13, yB: 62,
			targetX: 15518, targetY: 18554,
		},
		{

			xA: 45, yA: 17,
			xB: 33, yB: 64,
			targetX: 2672, targetY: 278,
		},
		{

			xA: 12, yA: 53,
			xB: 61, yB: 12,
			targetX: 2995, targetY: 5248,
		},
		{

			xA: 17, yA: 36,
			xB: 34, yB: 23,
			targetX: 14864, targetY: 7398,
		},
		{

			xA: 17, yA: 44,
			xB: 19, yB: 12,
			targetX: 15714, targetY: 19608,
		},
		{

			xA: 46, yA: 21,
			xB: 11, yB: 59,
			targetX: 6782, targetY: 3934,
		},
		{

			xA: 15, yA: 60,
			xB: 51, yB: 35,
			targetX: 5397, targetY: 8575,
		},
		{

			xA: 16, yA: 38,
			xB: 26, yB: 18,
			targetX: 14938, targetY: 5584,
		},
		{

			xA: 21, yA: 97,
			xB: 71, yB: 73,
			targetX: 7044, targetY: 14180,
		},
		{

			xA: 49, yA: 12,
			xB: 35, yB: 68,
			targetX: 7690, targetY: 3224,
		},
		{

			xA: 22, yA: 90,
			xB: 98, yB: 92,
			targetX: 8404, targetY: 13992,
		},
		{

			xA: 62, yA: 89,
			xB: 88, yB: 12,
			targetX: 11736, targetY: 8044,
		},
		{

			xA: 11, yA: 24,
			xB: 56, yB: 28,
			targetX: 2155, targetY: 2724,
		},
		{

			xA: 23, yA: 59,
			xB: 63, yB: 23,
			targetX: 11601, targetY: 11213,
		},
		{

			xA: 16, yA: 44,
			xB: 55, yB: 36,
			targetX: 12689, targetY: 9588,
		},
		{

			xA: 82, yA: 22,
			xB: 57, yB: 65,
			targetX: 4317, targetY: 1705,
		},
		{

			xA: 70, yA: 21,
			xB: 12, yB: 59,
			targetX: 13144, targetY: 16036,
		},
		{

			xA: 19, yA: 60,
			xB: 63, yB: 18,
			targetX: 13511, targetY: 12536,
		},
		{

			xA: 39, yA: 14,
			xB: 11, yB: 23,
			targetX: 11630, targetY: 7945,
		},
		{

			xA: 83, yA: 14,
			xB: 41, yB: 73,
			targetX: 9182, targetY: 5646,
		},
		{

			xA: 13, yA: 32,
			xB: 48, yB: 21,
			targetX: 2143, targetY: 17376,
		},
		{

			xA: 81, yA: 14,
			xB: 23, yB: 99,
			targetX: 4079, targetY: 4506,
		},
		{

			xA: 35, yA: 57,
			xB: 91, yB: 26,
			targetX: 3724, targetY: 3132,
		},
		{

			xA: 36, yA: 29,
			xB: 11, yB: 80,
			targetX: 3487, targetY: 4872,
		},
		{

			xA: 17, yA: 92,
			xB: 29, yB: 25,
			targetX: 2446, targetY: 2418,
		},
		{

			xA: 69, yA: 13,
			xB: 13, yB: 44,
			targetX: 2886, targetY: 10466,
		},
		{

			xA: 50, yA: 17,
			xB: 12, yB: 51,
			targetX: 6442, targetY: 7078,
		},
		{

			xA: 16, yA: 60,
			xB: 36, yB: 11,
			targetX: 1716, targetY: 16279,
		},
		{

			xA: 65, yA: 11,
			xB: 11, yB: 60,
			targetX: 18930, targetY: 16410,
		},
		{

			xA: 96, yA: 23,
			xB: 60, yB: 95,
			targetX: 7932, targetY: 2626,
		},
		{

			xA: 11, yA: 19,
			xB: 28, yB: 15,
			targetX: 4553, targetY: 13508,
		},
		{

			xA: 86, yA: 48,
			xB: 21, yB: 36,
			targetX: 4013, targetY: 2604,
		},
		{

			xA: 13, yA: 21,
			xB: 98, yB: 12,
			targetX: 9396, targetY: 2742,
		},
		{

			xA: 52, yA: 21,
			xB: 15, yB: 38,
			targetX: 875, targetY: 3596,
		},
		{

			xA: 18, yA: 75,
			xB: 44, yB: 12,
			targetX: 19734, targetY: 7307,
		},
		{

			xA: 12, yA: 45,
			xB: 85, yB: 46,
			targetX: 13467, targetY: 4878,
		},
		{

			xA: 33, yA: 81,
			xB: 80, yB: 34,
			targetX: 6083, targetY: 6001,
		},
		{

			xA: 79, yA: 15,
			xB: 12, yB: 51,
			targetX: 12333, targetY: 7055,
		},
		{

			xA: 33, yA: 88,
			xB: 64, yB: 43,
			targetX: 8196, targetY: 9983,
		},
		{

			xA: 43, yA: 11,
			xB: 26, yB: 57,
			targetX: 17360, targetY: 16125,
		},
		{

			xA: 28, yA: 15,
			xB: 11, yB: 29,
			targetX: 2650, targetY: 8295,
		},
		{

			xA: 13, yA: 32,
			xB: 31, yB: 11,
			targetX: 5400, targetY: 12700,
		},
		{

			xA: 51, yA: 25,
			xB: 21, yB: 36,
			targetX: 7382, targetY: 7507,
		},
		{

			xA: 13, yA: 39,
			xB: 65, yB: 15,
			targetX: 5642, targetY: 3426,
		},
		{

			xA: 47, yA: 34,
			xB: 17, yB: 53,
			targetX: 3067, targetY: 4498,
		},
		{

			xA: 65, yA: 32,
			xB: 33, yB: 62,
			targetX: 4457, targetY: 2606,
		},
		{

			xA: 49, yA: 70,
			xB: 35, yB: 15,
			targetX: 6059, targetY: 3295,
		},
		{

			xA: 59, yA: 11,
			xB: 14, yB: 68,
			targetX: 18274, targetY: 1936,
		},
		{

			xA: 40, yA: 20,
			xB: 47, yB: 93,
			targetX: 5317, targetY: 8983,
		},
		{

			xA: 19, yA: 78,
			xB: 97, yB: 55,
			targetX: 7544, targetY: 7975,
		},
		{

			xA: 66, yA: 16,
			xB: 56, yB: 80,
			targetX: 5012, targetY: 6064,
		},
		{

			xA: 34, yA: 71,
			xB: 88, yB: 36,
			targetX: 4302, targetY: 5585,
		},
		{

			xA: 12, yA: 53,
			xB: 86, yB: 50,
			targetX: 4978, targetY: 6484,
		},
		{

			xA: 63, yA: 90,
			xB: 82, yB: 33,
			targetX: 8537, targetY: 6474,
		},
		{

			xA: 43, yA: 13,
			xB: 29, yB: 46,
			targetX: 19005, targetY: 18165,
		},
		{

			xA: 12, yA: 45,
			xB: 64, yB: 39,
			targetX: 5760, targetY: 2909,
		},
		{

			xA: 19, yA: 42,
			xB: 58, yB: 32,
			targetX: 8338, targetY: 18796,
		},
		{

			xA: 42, yA: 14,
			xB: 26, yB: 97,
			targetX: 3304, targetY: 9758,
		},
		{

			xA: 26, yA: 54,
			xB: 35, yB: 19,
			targetX: 5612, targetY: 304,
		},
		{

			xA: 22, yA: 14,
			xB: 15, yB: 37,
			targetX: 15626, targetY: 614,
		},
		{

			xA: 38, yA: 17,
			xB: 14, yB: 55,
			targetX: 6140, targetY: 7830,
		},
		{

			xA: 61, yA: 16,
			xB: 12, yB: 28,
			targetX: 4088, targetY: 18840,
		},
		{

			xA: 64, yA: 36,
			xB: 12, yB: 31,
			targetX: 16000, targetY: 512,
		},
		{

			xA: 56, yA: 23,
			xB: 17, yB: 35,
			targetX: 9656, targetY: 16934,
		},
		{

			xA: 53, yA: 83,
			xB: 40, yB: 14,
			targetX: 10714, targetY: 15926,
		},
		{

			xA: 21, yA: 54,
			xB: 53, yB: 17,
			targetX: 19178, targetY: 17717,
		},
		{

			xA: 50, yA: 20,
			xB: 11, yB: 45,
			targetX: 10433, targetY: 7075,
		},
		{

			xA: 38, yA: 14,
			xB: 41, yB: 87,
			targetX: 2833, targetY: 4279,
		},
		{

			xA: 14, yA: 29,
			xB: 28, yB: 15,
			targetX: 16790, targetY: 13621,
		},
		{

			xA: 80, yA: 54,
			xB: 29, yB: 83,
			targetX: 7175, targetY: 7063,
		},
		{

			xA: 58, yA: 36,
			xB: 26, yB: 71,
			targetX: 5610, targetY: 5896,
		},
		{

			xA: 67, yA: 35,
			xB: 15, yB: 55,
			targetX: 205, targetY: 16765,
		},
		{

			xA: 11, yA: 83,
			xB: 88, yB: 13,
			targetX: 6425, targetY: 10493,
		},
		{

			xA: 81, yA: 60,
			xB: 13, yB: 32,
			targetX: 15205, targetY: 19068,
		},
		{

			xA: 15, yA: 54,
			xB: 21, yB: 11,
			targetX: 16940, targetY: 6123,
		},
		{

			xA: 23, yA: 65,
			xB: 45, yB: 19,
			targetX: 3830, targetY: 7586,
		},
		{

			xA: 24, yA: 37,
			xB: 41, yB: 11,
			targetX: 1351, targetY: 2435,
		},
		{

			xA: 77, yA: 19,
			xB: 13, yB: 62,
			targetX: 14693, targetY: 1415,
		},
		{

			xA: 72, yA: 16,
			xB: 23, yB: 73,
			targetX: 8394, targetY: 4238,
		},
		{

			xA: 75, yA: 25,
			xB: 18, yB: 71,
			targetX: 6747, targetY: 5759,
		},
		{

			xA: 71, yA: 16,
			xB: 12, yB: 59,
			targetX: 18562, targetY: 2449,
		},
		{

			xA: 75, yA: 25,
			xB: 16, yB: 63,
			targetX: 823, targetY: 12489,
		},
		{

			xA: 81, yA: 15,
			xB: 14, yB: 83,
			targetX: 6930, targetY: 8520,
		},
		{

			xA: 21, yA: 30,
			xB: 80, yB: 33,
			targetX: 2143, targetY: 948,
		},
		{

			xA: 89, yA: 35,
			xB: 58, yB: 79,
			targetX: 4208, targetY: 4352,
		},
		{

			xA: 32, yA: 61,
			xB: 61, yB: 28,
			targetX: 5177, targetY: 13721,
		},
		{

			xA: 16, yA: 53,
			xB: 39, yB: 27,
			targetX: 1206, targetY: 1338,
		},
		{

			xA: 54, yA: 30,
			xB: 31, yB: 56,
			targetX: 19166, targetY: 6314,
		},
		{

			xA: 24, yA: 58,
			xB: 49, yB: 19,
			targetX: 18105, targetY: 17665,
		},
		{

			xA: 44, yA: 26,
			xB: 12, yB: 41,
			targetX: 9008, targetY: 19382,
		},
		{

			xA: 55, yA: 21,
			xB: 18, yB: 36,
			targetX: 12563, targetY: 2591,
		},
		{

			xA: 17, yA: 66,
			xB: 64, yB: 26,
			targetX: 7891, targetY: 7118,
		},
		{

			xA: 21, yA: 38,
			xB: 31, yB: 12,
			targetX: 19443, targetY: 19350,
		},
		{

			xA: 32, yA: 15,
			xB: 18, yB: 32,
			targetX: 7472, targetY: 6202,
		},
		{

			xA: 65, yA: 11,
			xB: 12, yB: 63,
			targetX: 12845, targetY: 2099,
		},
		{

			xA: 68, yA: 75,
			xB: 81, yB: 20,
			targetX: 2821, targetY: 2210,
		},
		{

			xA: 31, yA: 96,
			xB: 83, yB: 34,
			targetX: 10041, targetY: 12360,
		},
		{

			xA: 37, yA: 15,
			xB: 31, yB: 67,
			targetX: 4451, targetY: 3383,
		},
		{

			xA: 11, yA: 49,
			xB: 37, yB: 14,
			targetX: 19422, targetY: 7060,
		},
		{

			xA: 66, yA: 20,
			xB: 68, yB: 92,
			targetX: 7894, targetY: 8532,
		},
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
