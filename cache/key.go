package cache

import (
	"fmt"
	"strconv"
)

const (
	Rankkey = "rank"
)

func ProductViewKey(id uint) string {
	return fmt.Sprintf("View:Product :%s", strconv.Itoa(int(id)))
}
