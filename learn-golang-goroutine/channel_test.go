package learn_golang_goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestChannel(t *testing.T) {
	channel := make(chan string)

	channel <- "New Channel"

	data := <-channel

	fmt.Println(data)
	fmt.Println(<-channel)

	// untuk close channel
	//close(channel)
	// disarankan pakai defer
	defer close(channel)
}

func TestCreateChannel(t *testing.T) {
	channel := make(chan string)
	defer close(channel)

	go func() {
		//time.Sleep(2 * time.Second)
		channel <- "Hello, Channel For Babayo"
		fmt.Println("Sudah Selesai Dikirim")
	}()

	dataFromChannel := <-channel
	fmt.Printf("Data %s \n", dataFromChannel)
	time.Sleep(5 * time.Second)
}

func CreateResponse(channel chan string) {
	time.Sleep(2 * time.Second)
	channel <- "Data From Channel As Param"
}

func TestChannelAsParameter(t *testing.T) {
	channel := make(chan string)

	go CreateResponse(channel)

	data := <-channel
	fmt.Printf("Data From Channel : %s \n", data)

	close(channel)
	time.Sleep(5 * time.Second)
}

func ChannelOnlyOut(outChannel chan<- string) {
	time.Sleep(2 * time.Second)
	outChannel <- "OutChannel, Hello"
}

func ChannelOnlyIn(inChannel <-chan string) {
	fmt.Println(<-inChannel)
}

func TestChannelInAndOut(t *testing.T) {
	channel := make(chan string)

	go ChannelOnlyIn(channel)
	go ChannelOnlyOut(channel)

	time.Sleep(5 * time.Second)
}

func TestBufferedChannel(t *testing.T) {
	bufferedChannel := make(chan string, 3)
	defer close(bufferedChannel)

	go func() {
		time.Sleep(2 * time.Second)
		bufferedChannel <- "Bayu"
		bufferedChannel <- "Seno"
		bufferedChannel <- "Ariefyanto"
	}()

	go func() {
		fmt.Println(<-bufferedChannel)
		fmt.Println(<-bufferedChannel)
		fmt.Println(<-bufferedChannel)
		time.Sleep(3 * time.Second)
	}()

	time.Sleep(7 * time.Second)
}
func TestRangeChannel(t *testing.T) {
	rangeChannel := make(chan string, 3)

	go func() {
		for i := 0; i < 10; i++ {
			rangeChannel <- fmt.Sprintf("Data Ke-%d", i)
			//time.Sleep(2 * time.Second)
		}
		close(rangeChannel)
	}()

	// Iterate through channel until channel closed
	for data := range rangeChannel {
		fmt.Println(data)
	}

	fmt.Println("Selesai")
}

func TestSelectChannel(t *testing.T) {
	channelOne := make(chan string)
	channelTwo := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channelOne <- "Hello, Gaes"
	}()
	go CreateResponse(channelTwo)

	go func() {
		counter := 0
		for {
			// Select based on channel
			select {
			case data := <-channelOne:
				fmt.Printf("Data Channel One : %s\n", data)
				counter++
			case data := <-channelTwo:
				fmt.Printf("Data Channel Two : %s\n", data)
				counter++
			}
			if counter == 2 {
				break
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}

func TestDefaultSelectChannel(t *testing.T) {
	channelOne := make(chan string)
	channelTwo := make(chan string)
	channelThree := make(chan string)

	go func() {
		time.Sleep(2 * time.Second)
		channelOne <- "Hello, Gaes"
	}()
	go CreateResponse(channelTwo)
	go CreateResponse(channelThree)

	go func() {
		counter := 0
		for {
			select {
			case data := <-channelOne:
				fmt.Printf("Data Channel One : %s\n", data)
				counter++
			case data := <-channelTwo:
				fmt.Printf("Data Channel Two : %s\n", data)
				counter++

			case data := <-channelThree:
				fmt.Printf("Data Channel Three : %s\n", data)
				counter++

			default:
				fmt.Println("Menunggu Data")
			}

			if counter == 3 {
				break
			}
		}
	}()

	time.Sleep(5 * time.Second)
	fmt.Println("Selesai")
}
