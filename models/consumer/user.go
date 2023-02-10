package model_consumer

import "gorm.io/gorm"

// 用户基础
type User struct {
	gorm.Model
	Avatar   string // 头像
	NickName string // 昵称
	Sex      int    // 性别
	BirthDay string // 生日
	Status   int    // 状态
}

// 用户账户
type UserAccount struct {
	gorm.Model
	OpenId   string // 微信openid
	UserName string // 用户名
	Email    string // 邮箱
	Password string // 密码
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
