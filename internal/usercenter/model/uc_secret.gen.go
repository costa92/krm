package model

import (
	"time"
)

const TableNameSecretM = "uc_secret"

// SecretM mapped from table <uc_secret>
type SecretM struct {
	ID          int64     `gorm:"column:id;type:bigint(20) unsigned;primaryKey;autoIncrement:true;comment:主键 ID" json:"id"`                        // 主键 ID
	UserID      string    `gorm:"column:user_id;type:varchar(253);not null;index:idx_user_id,priority:1;comment:用户 ID" json:"user_id"`             // 用户 ID
	Name        string    `gorm:"column:name;type:varchar(253);not null;comment:密钥名称" json:"name"`                                                 // 密钥名称
	SecretID    string    `gorm:"column:secret_id;type:varchar(36);not null;uniqueIndex:uniq_secret_id,priority:1;comment:密钥 ID" json:"secret_id"` // 密钥 ID
	SecretKey   string    `gorm:"column:secret_key;type:varchar(36);not null;comment:密钥 Key" json:"secret_key"`                                    // 密钥 Key
	Status      int32     `gorm:"column:status;type:tinyint(3) unsigned;not null;default:1;comment:密钥状态，0-禁用；1-启用" json:"status"`                  // 密钥状态，0-禁用；1-启用
	Description string    `gorm:"column:description;type:varchar(255);not null;comment:密钥描述" json:"description"`                                   // 密钥描述
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;not null;default:current_timestamp();comment:创建时间" json:"created_at"`             // 创建时间
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;not null;default:current_timestamp();comment:最后修改时间" json:"updated_at"`           // 最后修改时间
	Expires     int64     `gorm:"column:expires;type:bigint(64);not null" json:"expires"`
}

// TableName SecretM's table name
func (*SecretM) TableName() string {
	return TableNameSecretM
}
