//gens number then checks if it is even or odd in parallel mod7
var (
	printNow chan bool
	i        int
)

func main() {
	printNow = make(chan bool)

	go maker()
	go checker()
	for {
	}
}
func maker() {
	for {
		if _, ok := <-printNow; ok {
			if i%2 == 0 {
				fmt.Println("even:", i)
			} else {
				fmt.Println("odd:", i)
			}
		}

	}
}
func checker() {
	for {
		for i = 0; i <= 10; i++ {
			fmt.Println("Call", i)

			printNow <- true

			time.Sleep(1 * time.Millisecond)
		}
		os.Exit(0)
	}
}
