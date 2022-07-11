package main

import (
	"fmt"
	"regexp"
	"strings"
  "math/rand"
  "strconv"
  "math"
)
// Data Collections mod6
func main() {
	//act 1
	arr := RanArr(100)
	fmt.Println(arr)
	//act 2
	fmt.Println("Max num",MaxArr(arr))
	fmt.Println("Max num index",MaxArrIn(arr))
	fmt.Println("Min num",MinArr(arr))
	fmt.Println("Min num index",MinArrIn(arr))
	fmt.Println(SortArrD(arr))
	fmt.Println(SortArrA(arr))
	fmt.Println("mean",MeanArr(arr))
	fmt.Println("med",MedArr(arr))
	fmt.Println(PosArr(arr))
	fmt.Println(NegArr(arr))
	fmt.Println(SeqArr(arr))
	fmt.Println(RemoveDupArr(arr))
	//act 3 inputs 1 2 3 4 5  y all on new lines
	sub.Act3()
	//act 4
	sub.Para()
	//act 5
	arr1 := RanSlice(100)
	arr2 := RanSlice(100)
	fmt.Println("Slice 1: ", arr1)
	fmt.Println("Slice 2: ", arr2)
	fmt.Println("Sorted desccending Slice 1: ", SliceSort(arr1, 1))
	fmt.Println("Sorted ascending Slice 2: ", SliceSort(arr2, 0))
	fmt.Println("Combined and sorted slice 1 and 2", SliceCom(arr1, arr2))
	//act 6
	s := new(Sphere)
	s.Radius = 3
	fmt.Println(s.Volume())
	c := new((Cube))
	c.Length = 5
	fmt.Println(c.Volume())
	a := new(Box)
	a.Height = 3
	a.Length = 4
	a.Width = 2
	fmt.Println(a.Volume())
	b := new(Cylinder)
	b.Height = 2
	b.Radius = 5
	fmt.Println(b.Volume())
}
func RanArr(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(201)-100)
	}
	return arr
}

func MaxArr(arr []int) int {
	var max = arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if max < arr[i] {
			max = arr[i]
		}
	}
	return max
}
func MaxArrIn(arr []int) int {
	var max, I = arr[0], 0
	for i := 1; i < len(arr)-1; i++ {
		if max < arr[i] {
			max = arr[i]
			I = i
		}
	}
	return I
}
func MinArr(arr []int) int {
	var min = arr[0]
	for i := 1; i < len(arr)-1; i++ {
		if min > arr[i] {
			min = arr[i]
		}
	}
	return min
}
func MinArrIn(arr []int) int {
	var min, I = arr[0], 0
	for i := 1; i < len(arr)-1; i++ {
		if min > arr[i] {
			min = arr[i]
			I = i
		}
	}
	return I
}
func SortArrD(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		x := i
		for x > 0 {
			if arr[x-1] < arr[x] {
				arr[x-1], arr[x] = arr[x], arr[x-1]
			}
			x = x - 1
		}
	}
	return arr
}
func SortArrA(arr []int) []int {
	for i := 1; i < len(arr); i++ {
		x := i
		for x > 0 {
			if arr[x-1] > arr[x] {
				arr[x-1], arr[x] = arr[x], arr[x-1]
			}
			x = x - 1
		}
	}
	return arr
}
func MeanArr(arr []int) float64 {
	var tot int
	for i := 0; i < len(arr); i++ {
		tot += arr[i]
	}
	return float64(tot) / float64(len(arr))
}
func MedArr(arr []int) int {
	arr = SortArrD(arr)
	return arr[49]
}
func PosArr(arr []int) []int {
	var pArr []int
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] >= 0 {
			pArr = append(pArr, arr[i])
		}
	}
	return pArr
}
func NegArr(arr []int) []int {
	var nArr []int
	for i := 1; i < len(arr)-1; i++ {
		if arr[i] < 0 {
			nArr = append(nArr, arr[i])
		}
	}
	return nArr
}
func SeqArr(arr []int) []int {
	var long, tLong []int
	for i := 1; i < len(arr); i++ {

		if len(tLong) == 0 {
			if arr[i-1] < arr[i] {
				tLong = append(tLong, arr[i-1])
				tLong = append(tLong, arr[i])
			}
		} else if tLong[len(tLong)-1] < arr[i] {
			tLong = append(tLong, arr[i])

		} else if len(long) < len(tLong) {
			long = append([]int(nil), tLong...)
			tLong = nil
		}
	}
	return long
}
func RemoveDupArr(arr []int) []int {
	var noDup []int
	for i := 0; i < len(arr); i++ {
	dup:
		for x := 0; x < len(noDup); x++ {
			if arr[i] == noDup[x] {
				i++
				goto dup
			}
		}
		noDup = append(noDup, arr[i])
	}
	return noDup
}


func Act3() {
	fmt.Println("Input 5 integers")
	var input string
	var numList []int
	for i := 0; i < 5; i++ {
		fmt.Scanln(&input)
		x, _ := strconv.Atoi(input)
		numList = append(numList, x)
	}
	fmt.Println("Input up to 5 more integers, input 'y' to exit")
	for i := 0; i < 5; i++ {
		fmt.Scanln(&input)
		if input == "y" {
			break
		}
		x, _ := strconv.Atoi(input)
		numList = append(numList, x)
	}
	var sum int
	product := 1
	for i := 0; i < len(numList); i++ {
		sum += numList[i]
		product *= numList[i]
	}
	fmt.Println("you inputted:", numList, "Sum:", sum, "Product:", product)
}


func Para() {
	p := "Book One is geared toward teaching students how to write an instructional paragraph using the FNTF formula: First, Next, Then, Finally. The book first describes the How-to paragraph, gives examples, and has students dissect paragraphs using multiple choice and short answer questions. Once students are familiar with the purpose and style of a How-to paragraph, they make their own using provided prompts. Students then learn how to edit and format paragraphs with instruction given on margins, editing marks, and more. Students then edit sample paragraphs. The book then provides various lessons for using correct capitalization, punctuation, subjects, and verbs. Thought content of the paragraphs is then covered and students gain practice in reviewing if paragraphs make sense, correcting ones that do not, and making more of their own paragraphs. Once students have become familiar with reading, writing, and editing paragraphs, the last chapter of the book instructs students on how to write a whole How-to essay using the same step-by-step process they have learned for paragraphs."
	p = strings.ToLower(p)
	re, _ := regexp.Compile(`[^\w]`)
	p = re.ReplaceAllString(p, " ")
	s := strings.Fields(p)
	count := make(map[string]int)
	Keys := make(map[string]bool)

	for _, w := range s {
		if _, check := Keys[w]; !check {
			Keys[w] = true
			count[w] = strings.Count(p, w+" ")
		}
	}
	fmt.Println(count)
}


func RanSlice(n int) []int {
	var arr []int
	for i := 0; i < n; i++ {
		arr = append(arr, rand.Intn(101))
	}
	return arr
}
func SliceSort(arr []int, x int) []int {
	if x > 0 {
		for i := 1; i < len(arr); i++ {
			x := i
			for x > 0 {
				if arr[x-1] < arr[x] {
					arr[x-1], arr[x] = arr[x], arr[x-1]
				}
				x = x - 1
			}
		}
	} else if x <= 0 {
		for i := 1; i < len(arr); i++ {
			x := i
			for x > 0 {
				if arr[x-1] > arr[x] {
					arr[x-1], arr[x] = arr[x], arr[x-1]
				}
				x = x - 1
			}
		}

	}
	return arr
}
func SliceCom(arr1, arr2 []int) []int {
	arr1 = append(arr1, arr2...)
	return SliceSort(arr1, 1)
}

type Cube struct {
	Length float64
}
type Box struct {
	Length float64
	Height float64
	Width  float64
}
type Sphere struct {
	Radius float64
}
type Cylinder struct {
	Radius float64
	Height float64
}

func (c Cube) Volume() float64 {
	return math.Pow(c.Length, 3)
}
func (b Box) Volume() float64 {
	return b.Height * b.Length * b.Width
}
func (s Sphere) Volume() float64 {
	return (4 / 3) * math.Pi * math.Pow(s.Radius, 3)
}
func (s Cylinder) Volume() float64 {
	return math.Pi * math.Pow(s.Radius, 2) * s.Height
}
