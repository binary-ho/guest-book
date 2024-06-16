package room

// TODO : 나중엔 Redis에 저장
type Clients map[*Client]struct{}

func (clients Clients) Insert(client *Client) {
	if clients.contains(client) {
		return
	}
	clients[client] = struct{}{}
}

func (clients Clients) Remove(client *Client) {
	delete(clients, client)
}

func (clients Clients) IsEmpty() bool {
	return len(clients) == 0
}

func (clients Clients) contains(client *Client) bool {
	_, ok := clients[client]
	return ok
}
