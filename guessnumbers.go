package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano()) // VSCode已默认优化,不需要再次设置
	secertNumber := rand.Intn(maxNum)
	// fmt.Println(("The secret number is: "), secertNumber)

	fmt.Println("Please input your guess number: ")
	reader := bufio.NewReader(os.Stdin) // 读取用户输入
	for {
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("An error occurred while reading input. Please try again.", err)
			continue
		}
		input = strings.TrimSuffix(input, "\n") // 去掉输入中的换行符

		guess, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input. Please enter an integer value.")
			continue
		}
		fmt.Println("You guess is: ", guess)
		if guess > secertNumber {
			fmt.Println("Your guess is bigger than the secret number. Please try again.")
		} else if guess < secertNumber {
			fmt.Println("Your guess is smaller than the secret number. Please try again.")
		} else {
			fmt.Println("Correct, you Legend!")
		}
	}
}
