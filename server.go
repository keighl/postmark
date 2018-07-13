package postmark

// GetCurrentServer gets details for the server associated
// with the currently in-use server API Key
func (client *Client) GetCurrentServer() (Server, error) {
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

// EditCurrentServer updates details for the server associated
// with the currently in-use server API Key
func (client *Client) EditCurrentServer(server Server) (Server, error) {
	res := Server{}
	err := client.doRequest(parameters{
		Method:    "PUT",
		Path:      "server",
		TokenType: server_token,
	}, &res)
	return res, err
}
