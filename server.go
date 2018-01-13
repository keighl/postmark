package postmark

// GetThisServer gets details for the server associated
// with the currently in-use server API Key
func (client *Client) GetThisServer() (Server, error) {
	res := Server{}
	err := client.doRequest(parameters{
		Method: "GET",
		Path: "server",
		TokenType: server_token,
	}, &res)

	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// EditThisServer updates details for the server associated
// with the currently in-use server API Key
func (client *Client) EditThisServer(server Server) (Server, error) {
	res := Server{}
	err := client.doRequest(parameters{
		Method:    "PUT",
		Path:      "server",
		TokenType: server_token,
	}, &res)
	return res, err
}
