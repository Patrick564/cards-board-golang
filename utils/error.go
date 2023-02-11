package utils

type CustomError struct {
	Message string `json:"error"`
}

func (c *CustomError) Error() string {
	return c.Message
}
