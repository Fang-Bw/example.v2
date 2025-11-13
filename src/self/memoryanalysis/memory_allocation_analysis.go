package memoryanalysis

type User struct {
	ID   int
	Name string
}

// 编译优化参数
// go build -gcflags="-m -l -N" memory_allocation_analysis.go

// 例子1：栈分配
func createLocal() User {
	var user User // 在栈上分配（未逃逸）
	user.ID = 1
	return user // 值拷贝，原对象仍在栈上
}

// 例子2：堆分配
func createEscaped() *User {
	user := &User{ID: 1} // 逃逸到堆（返回指针）
	return user
}
