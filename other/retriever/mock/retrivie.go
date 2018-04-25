package mock


//使用者去实现Get方法
type Retriever struct {
	Contents string
}
//实现了接口Retriver的get方法，也就是实现了Retriver接口
func (r Retriever) Get(url string) string {
	return r.Contents
}