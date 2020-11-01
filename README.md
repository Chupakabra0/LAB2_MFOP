# Лабораторная работа №2
## Отчёт по лабораторной работе<br>по дисциплине "Математические основы программирования"<br>студента группы ПА-18-2<br>Сафиюлина Александра Александровича

### Постановка задачи

Тема: "Неформальное доказательство правильности программ, которые имеют разветвление"

Разделы лабораторной работы:
1. Написать программу в соответсвии с индивидуальным вариантом заданий (16 вариант);
2. Составить спецификацию и доказать правильность программы в ручном режиме по её спецификации с помощью фиксирования состояния программы после выполнения каждой команды.

#### Вариант 16

Имеем на входе множество {A, B, C}, причем элементы множества могут быть одинаковы. Вычислить сумму подмножества элементов ближайших к заданому X, если таких элементов больше, чем самых дальних от того же X, иначе - сумму подмножества элементов самых дальних от X.

### Описание решения:
    
1. Констатируем, что существует предусловие, код и постусловие `{Q} S {R}`, которые соответствуют поставленной задаче.
2. Идея: Составим множество {A', B' C'} такое что A' = abs(A -X), B' = abs(B - X), C' = abs(C - X). Найдём среди элементов этого множества минимум и максимум, обозначив их Mn и Mx соответсвенно.
Далее всё будет зависить от того, сколько элементов множества находится на расстоянии *(подразумеваем под расстоянием модуль разности между двумя элементами)* Mn и Mx. Если элементов на расстроянии Mn от X окажется больше, тогда считаем сумму таких элементов, иначе -- сумму элементов множаства, находящихся на расстоянии Mx от X.

3. Напишем код (использовал язык программирования Go):
  
  ```Go
package main

import (
	"fmt"
	"math"
)

type ThreeElementInt64Set struct {
	a, b, c int64
}

func CountDifference(set ThreeElementInt64Set, number int64) ThreeElementInt64Set {
	return ThreeElementInt64Set{ int64(math.Abs(float64(set.a - number))), int64(math.Abs(float64(set.b - number))), int64(math.Abs(float64(set.c - number))) }
}

func FindNum(dif ThreeElementInt64Set, pred func(int64, int64) bool) int64 {
	if pred(dif.a, dif.b) {
		if pred(dif.a, dif.c) {
			return dif.a
		} else {
			return dif.c
		}
	} else if pred(dif.b, dif.c) {
		return dif.b
	} else {
		return dif.c
	}
}

func CountNum(first int64, number int64) int64 {
	if first == number {
		return 1
	}
	return 0
}

func CountNumbers(set ThreeElementInt64Set, distance int64, number int64) (counter int64) {
	counter += CountNum(set.a, number + distance)
	counter += CountNum(set.b, number + distance)
	counter += CountNum(set.c, number + distance)
	if distance != 0 {
		counter += CountNum(set.a, number - distance)
		counter += CountNum(set.b, number - distance)
		counter += CountNum(set.c, number - distance)
	}
	return
}

func ElByDistance(first int64, distance int64, number int64) int64 {
	if first - distance == number || first + distance == number {
		return first
	}
	return 0
}

func SumElementsByDistance(set ThreeElementInt64Set, distance int64, number int64) (sum int64) {
	sum += ElByDistance(set.a, distance, number)
	sum += ElByDistance(set.b, distance, number)
	sum += ElByDistance(set.c, distance, number)
	return
}

func main() {
	var (
		abc ThreeElementInt64Set
		x   int64
	)

	_, _ = fmt.Print("Введите множество ABC: ")
	_, _ = fmt.Scan(&abc.a, &abc.b, &abc.c)

	_, _ = fmt.Print("Введите x: ")
	_, _ = fmt.Scan(&x)

	var (
		dif      = CountDifference(abc, x)
		minDIf   = FindNum(dif, func(first int64, second int64) bool { return first <= second })
		maxDif   = FindNum(dif, func(first int64, second int64) bool { return first >= second })
		countMin = CountNumbers(abc, minDIf, x)
		countMax = CountNumbers(abc, maxDif, x)
		result   int64
	)

	if countMin > countMax {
		result = SumElementsByDistance(abc, minDIf, x)
	} else {
		result = SumElementsByDistance(abc, maxDif, x)
	}

	fmt.Println(result)
}
  ```
  
4. Составим предусловие:
  
  `Q: {∀a, b, c, x: multiset1(A = a, B = b, C = c) ∧ X = x}`
  
5. Составим постусловие:
  
  `R : {(P1 ∧ C) ∨ (P2 ∧ D)}`
  
6. Составим спецификацию данной программы:
  
```Go
func main() {
	var (
		abc ThreeElementInt64Set
		x   int64
	)

	_, _ = fmt.Print("Введите множество ABC: ")
	_, _ = fmt.Scan(&abc.a, &abc.b, &abc.c)

	_, _ = fmt.Print("Введите x: ")
	_, _ = fmt.Scan(&x)
  
  // Начало выполнения программы
	var (
		dif      = CountDifference(abc, x)
                // {∀a, b, c, x: multiset1(A = a, B = b, C = c) ∧ X = x ∧ multiset2(A' = abs(A - X), B' = abs(B - X), C' = abs(C- X) }
		// где abs(A) = {∀a: A = a ∧ ((A < 0 ∧ result = -A) ∨ (A >= 0 ∧ result = A))}
    
		minDIf   = FindNum(dif, func(first int64, second int64) bool { return first <= second })
		maxDif   = FindNum(dif, func(first int64, second int64) bool { return first >= second })
                // {∀a, b, c, x: multiset1(A = a, B = b, C = c) ∧ multiset2(A' = abs(A - X), B' = abs(B - X), C' = abs(C - X)) ∧ (∀i = 0, 1, 2 ∧ ∀j = 0, 1, 2: multiset2[i] <= multiset2[j]) ∧ (∀i = 0, 1, 2 ∧ ∀j = 0, 1, 2: multiset2[i] <= multiset2[j]) ∧ X} = {multiset1((A = a, B = b, C = c) ∧ multiset2(A' = abs(A - X), B' = abs(B - X), C' = abs(C - X)) ∧ Mn ∧ Mx ∧ X}
    
		countMin = CountNumbers(abc, minDIf, x)
		countMax = CountNumbers(abc, maxDif, x)
                // {∀a, b, c, x: multiset1(A = a, B = b, C = c) ∧ multiset2(A' = abs(A - X), B' = abs(B - X), C' = abs(C - X)) ∧ Mn ∧ Mx ∧ {∀i = 0, 1, 2: multiset1[i] = X + Mn} ∧ {∀i = 0, 1, 2: multiset1[i] = X + Mx} ∧ X} = {multiset1(A, B, C) ∧ multiset2(A', B', C') ∧ Mn ∧ Mx ∧ Cnx ∧ Cmx ∧ X}
    
		result   int64
	)

	if countMin > countMax {
        // {Cmn > Cnx} = P1
		result = SumElementsByDistance(abc, minDIf, x)
                // P1 ∧ (C01 ∨ C02) ∧ (C11 ∨ C12) ∧ (C21 ∨ C22), где ∀j = 0, 1: Cj1 = {∃i = 0, 1, 2: X - Mn = multiset1[i]}, Cj2 = {∃i = 0, 1, 2: X + Mn = multiset1[i]}
                // (C01 ∨ C02) ∧ (C11 ∨ C12) ∧ (C21 ∨ C22) = C
                // P1 ∧ C
	} else {
        // {Cmn <= Cnx} = P2
		result = SumElementsByDistance(abc, maxDif, x)
                // P2 ∧ (D01 ∨ D02) ∧ (D11 ∨ D12) ∧ (D21 ∨ D22), где ∀j = 0, 1: Dj1 = {∃i = 0, 1, 2: X - Mx = multiset1[i]}, Cj2 = {∃i = 0, 1, 2: X + Mx = multiset1[i]}
                // (D01 ∨ D02) ∧ (D11 ∨ D12) ∧ (D21 ∨ D22) = D
                // P2 ∧ D
	}

	fmt.Println(result)
        // {(P1 ∧ C) ∨ (P2 ∧ D)}
}
```

### Тестовые примеры

![Первый скриншот](Screenshots/Screenshot1.png)
