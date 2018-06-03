package chat

type Object map[string]interface{}

type Request struct {
	ID     int    `json:"id"`
	Method string `json:"method"`
	Params Object `json:"params"`
}

type Response struct {
	ID     int    `json:"id"`
	Result Object `json:"result"`
}

type Error struct {
	ID    int    `json:"id"`
	Error Object `json:"error"`
}
