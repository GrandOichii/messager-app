
type Chat = {
	id: string
    participants: Array<string>
	// Messages           []*Message `json:"messages"`
}

type Message = {
    uhandle: string,
    text: string
}
