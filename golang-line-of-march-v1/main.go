package main

func Sortheight(height []int) []int {
	n := len(height)
	for i := 0; i < n-1; i++ {
		for j := 0; j < n-i-1; j++ {
			if height[j] > height[j+1] {				
				height[j], height[j+1] = height[j+1], height[j]
			}
		}
	}
	return height
}
