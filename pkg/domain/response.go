package domain

type response struct {
	Message string `json:"message"`
	Status  string `json:"status"`
}

func Response200(msg string) response {
	var res response
	res = response{Message: msg, Status: "200"}
	return res
}

func Response400(msg string) response {
	var res response
	res = response{Message: msg, Status: "400"}
	return res
}

func Response404(msg string) response {
	var res response
	res = response{Message: msg, Status: "404"}
	return res
}

func Response500(msg string) response {
	var res response
	res = response{Message: msg, Status: "500"}
	return res
}

func Response401(msg string) response {
	var res response
	res = response{Message: msg, Status: "401"}
	return res
}