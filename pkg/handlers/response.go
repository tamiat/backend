package handlers

type response struct {
	Message string `json:"message"`
	Status string `json:"status"`
}

func response200(msg string) response{
	var res response
	res = response{Message: msg, Status: "200"}
	return res
}

func response400(msg string) response{
	var res response
	res = response{Message: msg, Status: "400"}
	return res
}

func response404(msg string) response{
	var res response
	res = response{Message: msg, Status: "404"}
	return res
}

func response500(msg string) response{
	var res response
	res = response{Message: msg, Status: "500"}
	return res
}

