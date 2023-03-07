package main

func main() {
	// 오버플로우
	// maxuint 들은 최대값을 출력시키는데 , 그 최대값에 1을 더하니 오버플로우됨.
	// var n1 uint8 = math.MaxUint8 + 1
	// var n2 uint8 = math.MaxUint8 + 1
	// var n3 uint8 = math.MaxUint16 + 1
	// var n4 uint8 = math.MaxUint32 + 1
	// var n5 uint8 = math.MaxUint64 + 1

	//오버플로우 ( 범위 미만 )
	// 양수저장해야되는데 음수 저장해서 범위 미만되어 오버플로우 발생
	// var n5 uint8 = -1
	// var n5 uint16 = -1
	// var n5 uint32 = -1
	// var n5 uint64 = -1

}
