package postmark

import "context"

// GetCurrentServer calls GetCurrentServerWithContext with empty context
func (client *Client) GetCurrentServer() (Server, error) {
	return client.GetCurrentServerWithContext(context.Background())
}

// GetCurrentServer gets details for the server associated
// with the currently in-use server API Key
func (client *Client) GetCurrentServerWithContext(ctx context.Context) (Server, error) {
	res := Server{}
	err := client.doRequest(ctx, parameters{
		Method:    "GET",
		Path:      "server",
		TokenType: server_token,
	}, &res)

	return res, err
}

///////////////////////////////////////
///////////////////////////////////////

// EditCurrentServer calls EditCurrentServerWithContext with empty context
func (client *Client) EditCurrentServer(server Server) (Server, error) {
	return client.EditCurrentServerWithContext(context.Background(), server)
}

// EditCurrentServerWithContext updates details for the server associated
// with the currently in-use server API Key
func (client *Client) EditCurrentServerWithContext(ctx context.Context, server Server) (Server, error) {
	res := Server{}
	err := client.doRequest(ctx, parameters{
		Method:    "PUT",
		Path:      "server",
		TokenType: server_token,
	}, &res)
	return res, err
}
