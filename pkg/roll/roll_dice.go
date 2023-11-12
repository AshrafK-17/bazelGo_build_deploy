package roll

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func Roll() string {
	fmt.Println("Rolling the dice")
	return generateNumber()
}

func generateNumber() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	return strconv.Itoa(random.Intn(100))
}
