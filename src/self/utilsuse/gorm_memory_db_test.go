package utilsuse

import (
	"fmt"
	"os"
	"testing"
)

// setupTestDB 创建测试用的数据库实例
func setupTestDB(t *testing.T) *MemoryDB {
	db, err := NewMemoryDB()
	if err != nil {
		t.Fatalf("创建测试数据库失败: %v", err)
	}
	return db
}

// teardownTestDB 清理测试数据库
func teardownTestDB(t *testing.T, db *MemoryDB) {
	if db != nil {
		if err := db.Close(); err != nil {
			t.Errorf("关闭测试数据库失败: %v", err)
		}
	}
}

// TestMain 测试主函数，用于统一管理测试数据库实例
func TestMain(m *testing.M) {
	// 运行所有测试前的准备工作
	fmt.Println("=== 开始运行 GORM 内存数据库测试 ===")

	// 可以在这里创建全局测试数据库实例（如果需要）
	// 对于内存数据库，每个测试使用独立实例更合适

	// 运行测试
	code := m.Run()

	// 测试后的清理工作
	fmt.Println("=== 测试完成 ===")

	// 退出
	os.Exit(code)
}

// TestMemoryDB 测试内存数据库基本功能
func TestMemoryDB(t *testing.T) {
	// 使用辅助函数创建测试数据库
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 测试创建
	user := &User{
		Name:  "测试用户",
		Email: "test@example.com",
		Age:   25,
	}
	if err := db.CreateUser(user); err != nil {
		t.Errorf("创建用户失败: %v", err)
	}
	if user.ID == 0 {
		t.Error("用户 ID 应该被自动生成")
	}

	// 测试查询
	foundUser, err := db.GetUserByID(user.ID)
	if err != nil {
		t.Errorf("查询用户失败: %v", err)
	}
	if foundUser.Name != user.Name {
		t.Errorf("期望名称 %s, 实际 %s", user.Name, foundUser.Name)
	}

	// 测试更新
	user.Name = "更新后的名称"
	if err := db.UpdateUser(user); err != nil {
		t.Errorf("更新用户失败: %v", err)
	}

	updatedUser, _ := db.GetUserByID(user.ID)
	if updatedUser.Name != "更新后的名称" {
		t.Errorf("更新失败，期望 %s, 实际 %s", "更新后的名称", updatedUser.Name)
	}

	// 测试删除
	if err := db.DeleteUser(user.ID); err != nil {
		t.Errorf("删除用户失败: %v", err)
	}

	_, err = db.GetUserByID(user.ID)
	if err == nil {
		t.Error("用户应该已被删除")
	}
}

// TestBatchOperations 测试批量操作
func TestBatchOperations(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 批量创建
	users := []*User{
		{Name: "用户1", Email: "user1@example.com", Age: 20},
		{Name: "用户2", Email: "user2@example.com", Age: 25},
		{Name: "用户3", Email: "user3@example.com", Age: 30},
	}

	if err := db.CreateUsers(users); err != nil {
		t.Errorf("批量创建失败: %v", err)
	}

	// 验证数量
	count, err := db.CountUsers()
	if err != nil {
		t.Errorf("统计失败: %v", err)
	}
	if count != 3 {
		t.Errorf("期望 3 个用户，实际 %d", count)
	}

	// 测试年龄范围查询
	usersByAge, err := db.GetUsersByAgeRange(20, 25)
	if err != nil {
		t.Errorf("年龄范围查询失败: %v", err)
	}
	if len(usersByAge) != 2 {
		t.Errorf("期望 2 个用户，实际 %d", len(usersByAge))
	}
}

// TestSearch 测试搜索功能
func TestSearch(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	users := []*User{
		{Name: "张三", Email: "zhangsan@example.com", Age: 25},
		{Name: "李四", Email: "lisi@example.com", Age: 30},
	}
	db.CreateUsers(users)

	// 搜索测试
	results, err := db.SearchUsers("张")
	if err != nil {
		t.Errorf("搜索失败: %v", err)
	}
	if len(results) != 1 {
		t.Errorf("期望找到 1 个结果，实际 %d", len(results))
	}
}

// TestPagination 测试分页功能
func TestPagination(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	// 创建测试数据
	users := make([]*User, 10)
	for i := 0; i < 10; i++ {
		users[i] = &User{
			Name:  fmt.Sprintf("用户%d", i+1),
			Email: fmt.Sprintf("user%d@example.com", i+1),
			Age:   20 + i,
		}
	}
	db.CreateUsers(users)

	// 测试分页
	pageUsers, total, err := db.GetUsersWithPagination(1, 3)
	if err != nil {
		t.Errorf("分页查询失败: %v", err)
	}
	if total != 10 {
		t.Errorf("期望总数 10，实际 %d", total)
	}
	if len(pageUsers) != 3 {
		t.Errorf("期望每页 3 条，实际 %d", len(pageUsers))
	}
}

// TestTransaction 测试事务
func TestTransaction(t *testing.T) {
	db := setupTestDB(t)
	defer teardownTestDB(t, db)

	user := &User{
		Name:  "事务用户",
		Email: "transaction@example.com",
		Age:   25,
	}

	if err := db.CreateUserWithTransaction(user); err != nil {
		t.Errorf("事务创建失败: %v", err)
	}

	// 验证用户已创建
	foundUser, err := db.GetUserByID(user.ID)
	if err != nil {
		t.Errorf("查询事务创建的用户失败: %v", err)
	}
	if foundUser.Email != user.Email {
		t.Errorf("事务创建的用户数据不正确")
	}
}
