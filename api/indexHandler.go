package api

import (
	"encoding/json"
	"net/http"

	"github.com/bohdan-massive/gentest2/config"
	"github.com/gorilla/context"
)

//Example route handler
func IndexHandler(rw http.ResponseWriter, req *http.Request) HttpError {
	encoder := json.NewEncoder(rw)
	resData := make(map[string]string)
	resData["example"] = config.Conf.GetExample()

	val, has := context.GetOk(req, "test")
	if has {
		resData["context"] = val.(string)
	}

	if err := encoder.Encode(resData); err != nil {
		return NewHttpError(err, http.StatusInternalServerError)
	}
	return nil
}
