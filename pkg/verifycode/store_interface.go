package verifycode

type Store interface {
	Set(id string, value string) bool
	Get(id string, clear bool) string
	// Verify 方法、参数和返回值都要相同
	Verify(id, answer string, clear bool) bool
}
