package main

func main() {

}

// 算法核心： 两两对比把最大（小）的数放到一端， 经过n-1次外层循环就能得出正确数组
func MaoPao(arr []int) {
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr)-1-i; j++ {
			arr[j], arr[j+1] = arr[j+1], arr[j]
		}
	}
}

func KuaiPai(arr []int) {

}
