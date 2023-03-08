package constant

// 用户状态
type CUserStatus uint8

const (
	StatusNormal CUserStatus = iota + 1 // 正常
	StatusBanned                        // bannen
)

// 登录方式
type CUserLogin uint8

const (
	AccountLogin CUserLogin = iota // 账号
	PhoneLogin                     // 手机
	WXLogin                        // 微信
	MailLogin                      // 邮箱
)

// 性别
type UserGender uint8

const (
	GenderUnknown UserGender = iota + 1 // 未知
	GenderMan                           // 男
	GenderWoman                         // 女
)

func (ug UserGender) Valid() bool {
	return ug >= GenderUnknown && ug <= GenderWoman
}
