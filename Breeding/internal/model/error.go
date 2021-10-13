package model

type Error struct {
	HTTPCode        int    `json:"httpCode"`
	Message         string `json:"message"`
	InternalMessage string `json:"internalMessage"`
}

type ErrorResp struct {
	Error string `json:"error"`
}

func (e Error) ErrorResp() ErrorResp {
	return ErrorResp{
		Error: e.Message,
	}
}
