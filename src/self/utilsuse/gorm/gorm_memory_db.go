package gorm

import (
	"fmt"
	"time"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
===========================================
基于 GORM 的内存数据库 CRUD 操作
===========================================

使用 SQLite 内存数据库（:memory:）作为 H2 的等价方案
在 Go 生态中，SQLite 内存模式是最常用的内存数据库方案

===========================================
*/

// User 用户模型示例
type User struct {
	ID        uint      `gorm:"primaryKey" json:"id"`
	Name      string    `gorm:"not null" json:"name"`
	Email     string    `gorm:"uniqueIndex;not null" json:"email"`
	Age       int       `json:"age"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// TableName 指定表名
func (User) TableName() string {
	return "users"
}

// MemoryDB 内存数据库管理器
type MemoryDB struct {
	db *gorm.DB
}

// NewMemoryDB 创建新的内存数据库连接
func NewMemoryDB() (*MemoryDB, error) {
	// 使用 SQLite 内存数据库（:memory:）
	// 这相当于 Java 中的 H2 内存数据库
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info), // 显示 SQL 日志
	})
	if err != nil {
		return nil, fmt.Errorf("连接内存数据库失败: %w", err)
	}

	memoryDB := &MemoryDB{db: db}

	// 自动迁移表结构
	if err := memoryDB.AutoMigrate(); err != nil {
		return nil, fmt.Errorf("自动迁移失败: %w", err)
	}

	return memoryDB, nil
}

// AutoMigrate 自动迁移数据库表结构
func (m *MemoryDB) AutoMigrate() error {
	return m.db.AutoMigrate(&User{})
}

// Close 关闭数据库连接（SQLite 内存数据库在程序结束时自动关闭）
func (m *MemoryDB) Close() error {
	sqlDB, err := m.db.DB()
	if err != nil {
		return err
	}
	return sqlDB.Close()
}

// ===========================================
// CRUD 操作 - Create (创建)
// ===========================================

// CreateUser 创建用户
func (m *MemoryDB) CreateUser(user *User) error {
	result := m.db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("创建用户失败: %w", result.Error)
	}
	return nil
}

// CreateUsers 批量创建用户
func (m *MemoryDB) CreateUsers(users []*User) error {
	result := m.db.Create(users)
	if result.Error != nil {
		return fmt.Errorf("批量创建用户失败: %w", result.Error)
	}
	return nil
}

// ===========================================
// CRUD 操作 - Read (查询)
// ===========================================

// GetUserByID 根据 ID 查询用户
func (m *MemoryDB) GetUserByID(id uint) (*User, error) {
	var user User
	result := m.db.First(&user, id)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("用户不存在: ID=%d", id)
		}
		return nil, fmt.Errorf("查询用户失败: %w", result.Error)
	}
	return &user, nil
}

// GetUserByEmail 根据邮箱查询用户
func (m *MemoryDB) GetUserByEmail(email string) (*User, error) {
	var user User
	result := m.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		if result.Error == gorm.ErrRecordNotFound {
			return nil, fmt.Errorf("用户不存在: Email=%s", email)
		}
		return nil, fmt.Errorf("查询用户失败: %w", result.Error)
	}
	return &user, nil
}

// GetAllUsers 查询所有用户
func (m *MemoryDB) GetAllUsers() ([]User, error) {
	var users []User
	result := m.db.Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("查询所有用户失败: %w", result.Error)
	}
	return users, nil
}

// GetUsersByAge 根据年龄查询用户
func (m *MemoryDB) GetUsersByAge(age int) ([]User, error) {
	var users []User
	result := m.db.Where("age = ?", age).Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("根据年龄查询用户失败: %w", result.Error)
	}
	return users, nil
}

// GetUsersByAgeRange 根据年龄范围查询用户
func (m *MemoryDB) GetUsersByAgeRange(minAge, maxAge int) ([]User, error) {
	var users []User
	result := m.db.Where("age >= ? AND age <= ?", minAge, maxAge).Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("根据年龄范围查询用户失败: %w", result.Error)
	}
	return users, nil
}

// CountUsers 统计用户数量
func (m *MemoryDB) CountUsers() (int64, error) {
	var count int64
	result := m.db.Model(&User{}).Count(&count)
	if result.Error != nil {
		return 0, fmt.Errorf("统计用户数量失败: %w", result.Error)
	}
	return count, nil
}

// ===========================================
// CRUD 操作 - Update (更新)
// ===========================================

// UpdateUser 更新用户（根据 ID）
func (m *MemoryDB) UpdateUser(user *User) error {
	result := m.db.Save(user)
	if result.Error != nil {
		return fmt.Errorf("更新用户失败: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("用户不存在: ID=%d", user.ID)
	}
	return nil
}

// UpdateUserByID 根据 ID 更新用户指定字段
func (m *MemoryDB) UpdateUserByID(id uint, updates map[string]interface{}) error {
	result := m.db.Model(&User{}).Where("id = ?", id).Updates(updates)
	if result.Error != nil {
		return fmt.Errorf("更新用户失败: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("用户不存在: ID=%d", id)
	}
	return nil
}

// UpdateUserEmail 更新用户邮箱
func (m *MemoryDB) UpdateUserEmail(id uint, email string) error {
	return m.UpdateUserByID(id, map[string]interface{}{
		"email": email,
	})
}

// UpdateUserName 更新用户名称
func (m *MemoryDB) UpdateUserName(id uint, name string) error {
	return m.UpdateUserByID(id, map[string]interface{}{
		"name": name,
	})
}

// ===========================================
// CRUD 操作 - Delete (删除)
// ===========================================

// DeleteUser 根据 ID 删除用户
func (m *MemoryDB) DeleteUser(id uint) error {
	result := m.db.Delete(&User{}, id)
	if result.Error != nil {
		return fmt.Errorf("删除用户失败: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("用户不存在: ID=%d", id)
	}
	return nil
}

// DeleteUserByEmail 根据邮箱删除用户
func (m *MemoryDB) DeleteUserByEmail(email string) error {
	result := m.db.Where("email = ?", email).Delete(&User{})
	if result.Error != nil {
		return fmt.Errorf("删除用户失败: %w", result.Error)
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("用户不存在: Email=%s", email)
	}
	return nil
}

// DeleteAllUsers 删除所有用户
func (m *MemoryDB) DeleteAllUsers() error {
	result := m.db.Where("1 = 1").Delete(&User{})
	if result.Error != nil {
		return fmt.Errorf("删除所有用户失败: %w", result.Error)
	}
	return nil
}

// ===========================================
// 事务操作示例
// ===========================================

// CreateUserWithTransaction 使用事务创建用户
func (m *MemoryDB) CreateUserWithTransaction(user *User) error {
	return m.db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行多个操作
		if err := tx.Create(user).Error; err != nil {
			return err
		}
		// 可以继续执行其他操作
		// 如果返回错误，事务会自动回滚
		return nil
	})
}

// ===========================================
// 高级查询示例
// ===========================================

// GetUsersWithPagination 分页查询用户
func (m *MemoryDB) GetUsersWithPagination(page, pageSize int) ([]User, int64, error) {
	var users []User
	var total int64

	// 统计总数
	if err := m.db.Model(&User{}).Count(&total).Error; err != nil {
		return nil, 0, fmt.Errorf("统计总数失败: %w", err)
	}

	// 分页查询
	offset := (page - 1) * pageSize
	result := m.db.Offset(offset).Limit(pageSize).Find(&users)
	if result.Error != nil {
		return nil, 0, fmt.Errorf("分页查询失败: %w", result.Error)
	}

	return users, total, nil
}

// SearchUsers 搜索用户（根据名称或邮箱）
func (m *MemoryDB) SearchUsers(keyword string) ([]User, error) {
	var users []User
	result := m.db.Where("name LIKE ? OR email LIKE ?",
		"%"+keyword+"%", "%"+keyword+"%").Find(&users)
	if result.Error != nil {
		return nil, fmt.Errorf("搜索用户失败: %w", result.Error)
	}
	return users, nil
}

// GetDB 获取底层 GORM DB 实例（用于高级操作）
func (m *MemoryDB) GetDB() *gorm.DB {
	return m.db
}
