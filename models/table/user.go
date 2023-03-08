package table

import (
	"pAppServer/constant"

	"gorm.io/gorm"
)

// 用户信息
type CUser struct {
	Basic
	UUID        string               `gorm:"unique_index:idx_uuid;size:100" json:"-"` // uuid
	Avatar      string               `json:"avatar,omitempty"`                        // 头像
	NickName    string               `json:"nickname,omitempty"`                      // 昵称
	Gender      constant.UserGender  `json:"gender,omitempty"`                        // 性别
	BirthDay    string               `json:"birthday,omitempty"`                      // 生日
	Status      constant.CUserStatus `json:"status,omitempty"`                        // 状态
	StatusExtra string               `json:"status_extra,omitempty"`                  // 状态附带信息
}

// 用户身份
type CUserAuth struct {
	gorm.Model
	UUID string
	// 类型和标识一起创建复合索引 避免不同类型但同标识无法创建
	IdentityType constant.CUserLogin `gorm:"uniqueIndex:idx_identify"`
	// NOTICE 字符串作为索引时需要标明长度
	Identifier string `gorm:"uniqueIndex:idx_identify;size:200"`
	Credential string
}

// 用户会员
type UserMember struct {
	gorm.Model
	MemberPoint int // 会员积分
	MemberLv    int // 会员等级
}

// 用户地址
type UserAddress struct {
	gorm.Model
	AddressIds string // 地址
}

// 用户好友
type UserChatFriend struct {
	gorm.Model
	FriendId int    // 好友ID
	NickName string // 好友昵称
	Status   int    //	好友状态
}

// 用户群组
type UserChatGroup struct {
	gorm.Model
	GroupId  string // 群组ID
	NiceName string // 群昵称
	Status   int    // 群状态
}
