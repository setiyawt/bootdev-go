package main

type cost struct {
	day   int
	value float64
}

func getCostsByDay(costs []cost) []float64 {
	totalCost := []float64{}          // Slice ini akan menyimpan total cost untuk tiap hari.
	for i := 0; i < len(costs); i++ { // Looping melalui tiap elemen dalam slice costs.
		cost := costs[i]                 // Mengambil elemen cost pada indeks i.
		for cost.day >= len(totalCost) { // Kalau day dari cost lebih besar atau sama dengan panjang totalCost, tambahin elemen 0.0 ke totalCost.
			totalCost = append(totalCost, 0.0)
		}
		totalCost[cost.day] += cost.value // Menambahkan nilai cost.value ke elemen totalCost yang sesuai dengan day.
	}
	return totalCost
}
