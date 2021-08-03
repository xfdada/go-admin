package uuid

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func GetUuid() string {
	nowTime := time.Now().Format("20060102150405")
	rand.Seed(time.Now().UnixNano())
	times, _ := strconv.Atoi(nowTime)
	return strconv.FormatInt(int64(times), 16) + fmt.Sprintf("%08d", rand.Intn(100000000))
}
