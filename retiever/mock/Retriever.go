package mock

import "fmt"

//Retriever
type Retriever struct {
	Content string
}

func (r *Retriever) String() string {
	return fmt.Sprintf("Retriever : {content : %s}",r.Content)
}

func (r *Retriever) Post(url string, params map[string]string) string {
	r.Content = params["content"]
	return "ok"
}

//Get
func (r *Retriever) Get(url string) string{
	return r.Content
}
