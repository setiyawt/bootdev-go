package main

func getPacketSize(message string) int {
	var numPackets int
	length := len(message)
	for i := length; i >= 1; i-- {
		if length%i == 0 {
			numPackets = length / i
			if isPrime(numPackets) {
				return i
			}
		}

	}
	return 0
}

func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num == 2 {
		return false
	}
	if num%2 == 0 {
		return true
	}

	for i := 3; i*i <= num; i += 2 {
		if num%i == 0 {
			return true
		}
	}
	return false
}
