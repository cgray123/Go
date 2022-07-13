var (
	printNow chan bool
	i        int
)

func main() {
	printNow = make(chan bool)
	go printer()
	go sender()

	for {
	}
}
func printer() {
	for {
		if _, ok := <-printNow; ok {
			if i%2 == 0 {
				fmt.Print(" is even\n")
			} else {
				fmt.Print(" is odd\n")
			}
		}

	}
}
func sender() {
	for {
		for i = 0; i <= 100; i++ {
			fmt.Print(i)

			printNow <- true

			time.Sleep(2 * time.Millisecond)
		}
		os.Exit(0)
	}
}
