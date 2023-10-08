package models

type User struct {
	Model
	Name        string `json:"name" form:"name"`               // 姓名
	Password    string `json:"password" form:"password"`       // 密码
	Avatar      string `json:"avatar" form:"avatar"`           // 头像
	Sex         int    `json:"sex" form:"sex"`                 // 性别
	Phone       string `json:"phone" form:"phone"`             // 手机号
	Email       string `json:"email" form:"email"`             // 邮箱
	Status      uint8  `json:"status" form:"status"`           // 状态
	AccountName string `json:"accountName" form:"accountName"` // 账号名
	Role        int    `json:"role" form:"role"`               // 角色
	Post        int    `json:"post" form:"post"`               // 岗位
	Salt        string `json:"salt" form:"salt"`               // 随机字符串解密
	LoginIp     string `json:"loginIp" form:"loginIp"`         // 最后登录ip
	LoginTime   int64  `json:"loginTime" form:"loginTime"`     // 最后登录时间
}

func (u *User) TableName() string {
	return "user"
}
