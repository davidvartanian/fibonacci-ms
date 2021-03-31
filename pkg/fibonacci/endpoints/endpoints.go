package endpoints

import (
	"encoding/json"
	"fibonacci-ms/pkg/fibonacci"
	"io"
	"net/http"
	"strconv"
)

func GetRequestHandler(svc fibonacci.Service) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		posVal := req.FormValue("pos")
		if len(req.FormValue("pos")) == 0 {
			res.WriteHeader(http.StatusBadRequest)
			_, _ = io.WriteString(res, "Error: missing pos parameter")
			return
		}
		pos, err := strconv.ParseInt(posVal, 10, 32)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			_, _ = io.WriteString(res, "Error: "+err.Error())
			return
		}

		result, err := svc.Get(pos)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(res, "Error: "+err.Error())
			return
		}
		_, err = io.WriteString(res, strconv.Itoa(int(result)))
	}
}

func ListRequestHandler(svc fibonacci.Service) func(res http.ResponseWriter, req *http.Request) {
	return func(res http.ResponseWriter, req *http.Request) {
		minVal := req.FormValue("min")
		maxVal := req.FormValue("max")
		if len(minVal) == 0 || len(maxVal) == 0 {
			res.WriteHeader(http.StatusBadRequest)
			_, _ = io.WriteString(res, "Error: missing min or max parameters")
			return
		}
		min, err := strconv.ParseInt(minVal, 10, 32)
		max, err := strconv.ParseInt(maxVal, 10, 32)
		if err != nil {
			res.WriteHeader(http.StatusBadRequest)
			_, _ = io.WriteString(res, "Request Error: "+err.Error())
			return
		}

		result, err := svc.List(min, max)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(res, "Service Error: "+err.Error())
			return
		}
		serialized, err := json.Marshal(result)
		if err != nil {
			res.WriteHeader(http.StatusInternalServerError)
			_, _ = io.WriteString(res, "Serialization Error: "+err.Error())
			return
		}
		_, err = io.WriteString(res, string(serialized))
	}
}
