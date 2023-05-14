package message

const (
	LoginMesType    = "LoginMes"
	LoginResMesType = "LoginResMes"
)

type Message struct {
	Type string `json:"type"` //消息类型
	Data string `json:"data"` //消息数据
}

type LoginMes struct {
	UserId   int    `json:"userId"`
	UserPwd  string `json:"userPwd"`
	UserName string `json:"userName"` //name
}

type LoginResMes struct {
	Code  int    `json:"code"`  //返回状态码500表示该用户尚未注册，200表示登录成功
	Error string `json:"error"` //返回的错误信息
}
