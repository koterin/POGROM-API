package controller

import (
    "encoding/json"
    "fmt"
    "io"
    "net/http"

    "pogrom/config"
    "pogrom/internal/entity"
)

func ReadJson(req *http.Request, data *entity.ClientRequest) error {
    resp, err := io.ReadAll(req.Body)
    if err != nil {
        return fmt.Errorf("\nError in ReadJson %w", err)
    }

    err = json.Unmarshal([]byte(resp), data)
    if err != nil {
        return fmt.Errorf("\nError in ReadJson %w", err)
    }

    return nil
}

func addBasicHeaders(w http.ResponseWriter) {
    w.Header().Set("Access-Control-Allow-Origin", config.Args.HOST_URL)
    w.Header().Set("Access-Control-Allow-Credentials", "true")
    w.Header().Set("Access-Control-Allow-Methods", "GET,HEAD,OPTIONS,POST,PUT")
    w.Header().Set("Access-Control-Allow-Headers", "Origin, Accept, X-Requested-With, Content-Type, Access-Control-Request-Method, Access-Control-Request-Headers")
    w.Header().Set("Content-Type", "application/json")
}
