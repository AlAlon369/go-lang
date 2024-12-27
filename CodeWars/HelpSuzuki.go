package CodeWars

import "sort"

func ChoreAssignment(chores []int) []int {
	// Сортируем массив работ
	sort.Ints(chores)

	// Формируем пары минимального и максимального элементов
	workloads := make([]int, len(chores)/2)
	for i := 0; i < len(chores)/2; i++ {
		workloads[i] = chores[i] + chores[len(chores)-1-i]
	}

	// Сортируем итоговые рабочие нагрузки
	sort.Ints(workloads)
	return workloads
}
