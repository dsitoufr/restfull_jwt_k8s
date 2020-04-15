package handlers

import (
    "encoding/json"
    "net/http"
)


const version = "1.0.0"

type VersionResponse struct {
    Version string `json:"version"`
}

type versionHandler struct {
    version string
}

func VersionHandler(response http.ResponseWriter, resquest * http.Request)  {
    r := VersionResponse{
        Version: version,
    }

    json.NewEncoder(response).Encode(r)
    return
}