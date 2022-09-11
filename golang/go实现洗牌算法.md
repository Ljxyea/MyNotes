## go实现洗牌算法

### 使用perm函数实现

```golang
使用perm函数实现随机算法
func Shuffle(vals []int) []int {
    r := make([]int, len(vals))
    prem := r.Perm(len(vals))
    for i, randIndex := range perm {
        ret[i] = vals[randIndex]
    }
    return ret
}
```

### 使用lntn

```
func Shuffle(vals []int) {
	r := rand.New(rand.NewSource(time.Now().Unix()))
	for len(vals) >0 {
		n := len(vals)
		randIndex:= r.lntn(n)
		vals[n-1], val[randIndex] = vals[randIndex], vals[n-1]
		vals = vals[:n-1]
	}
}
```

使用官方提供的rand.shuffle方法

```
package main
import (
	"fmt"
	"math/rand"
	"time"
)
func main(){
	var arr = []unit64{1,2,3,4,5,6,7}
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(arr), func(i, j int)) {
		arr[i], arr[j] = arr[j], arr[i]
	}
	fmt.Println(arr)
	
}
```

