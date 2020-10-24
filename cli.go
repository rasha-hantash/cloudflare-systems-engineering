package main
import (
	"log"
	"github.com/spf13/cobra"
	"fmt"
	"io"
	"crypto/tls"
	"strings"
	"time"
	
)

var (
	profileFlag int
	rootCmd = &cobra.Command{
		Use: "go run cli.go <website you want to visit> ", //name of the command
		Short: "An example cobra program", // short description of the command
		Long: `This is a simple example of a cobra progam.`,
		Args: cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args [] string) {
			if profileFlag > 0 { 
				byteSizes := make([]int, profileFlag)
				time := make([]time.Duration, profileFlag)
				errorCodes := make([]string, profileFlag)
				successPercent := make([]int, profileFlag)
				for i := 0; i < profileFlag; i++ {
					byteSizes[i], time[i],successPercent[i], errorCodes[i] = URLRequest(args)
				}
				percentage :=  float64(0);
				for i := 0; i < profileFlag; i++  {
					
					percentage += float64(successPercent[i])
				}

				fastestTime := time[0]
				slowestTime := time[0]
				meanTime := float64(time[0])
				smallestByte := byteSizes[0]
				largestByte := byteSizes[0]
				errorMessages := errorCodes[0]
				for i := 1; i < profileFlag; i++ {
					if time[i] < fastestTime {
						fastestTime = time[i]
					}

					if time[i] > slowestTime {
						slowestTime = time[i]
					}

					if byteSizes[i] > largestByte {
						largestByte = byteSizes[i]
					}
					if byteSizes[i] < smallestByte {
						smallestByte = byteSizes[i]
					}
					errorMessages += " " + errorCodes[i]
					
					meanTime += float64(time[i])
				}
				fmt.Println()
				fmt.Println("Number of requests: ", profileFlag)
				fmt.Println("Fastest time: ", fastestTime)
				fmt.Println("Slowest time:  ", slowestTime)
				fmt.Printf("Mean and median times: %0.6fms \n", (meanTime / float64(profileFlag)) / 1000000)
				fmt.Printf("Percentage of succeeded requests: %0.2f %% \n",(percentage / float64(profileFlag)) * 100)
				fmt.Println("Error codes: ", errorMessages)
				fmt.Println("Size in bytes of the smallest response: ", smallestByte)
				fmt.Println("size in bytes of the largest response: ", largestByte)
			} else {
				URLRequest(args)
			}

		},
	}
)

func init() {
	rootCmd.PersistentFlags().IntVarP(&profileFlag, "profile", "p", 0, "the profile flag")
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}

}

// URLRequest grabs the JSON data from the http request
func URLRequest(url[] string) (int, time.Duration, int, string) {
	startTime := time.Now()
	stringURL := strings.SplitN(url[0], "/",2)
	conn, err := tls.Dial("tcp", stringURL[0] + ":443", nil)
    if err != nil {
		log.Fatal("dial error:", err)
		endTime := time.Now()
		timeTaken := endTime.Sub(startTime)
	
		return 0, timeTaken, 0, err.Error()
	}
	defer conn.Close()
	buf := make([]byte, 0, 4096) // big buffer
	fmt.Fprintf(conn, "GET /" + stringURL[1] + " HTTP/1.1\r\nHost: "+ stringURL[0]+ "\r\nConnection: Close\r\n\r\n")
	endTime := time.Now()
	timeTaken := endTime.Sub(startTime)
	for {
		tmp := make([]byte, 256) 
		n, err := conn.Read(tmp)
		if err != nil {
			if err != io.EOF {
				fmt.Println("read error:", err)
				return 0, timeTaken, 0, err.Error()
			}
			break
		}
		buf = append(buf, tmp[:n]...)

	}
	s := string(buf)
	result := strings.Split(s,"\n")
	fmt.Println(result[len(result)-1])

	
	return len(buf), timeTaken, 1, ""
}