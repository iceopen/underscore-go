package underscorego

// 初始化结构体
type U struct {
}

const (
	Version = "0.0.1.dev"
)

// 返回当前版本号
func (u *U) GetVersion() string {
	return Version
}
