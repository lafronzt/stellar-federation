package internal

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"

	"github.com/lafronzt/stellar-federation/model"
	logger "go.lafronz.com/tools/logger/stackdriver"
	"gopkg.in/yaml.v3"
)

var federations model.Federations

func init() {
	yamlFile, err := ioutil.ReadFile("stellar.yaml")
	if err != nil {
		logger.Error("Error reading conf.yaml: %s", err)
	}
	err = yaml.Unmarshal(yamlFile, &federations)
	if err != nil {
		logger.Error("Error unmarshalling conf.yaml: %s", err)
	}

	logger.Info("Loaded %d federations", len(federations))
	logger.Info("Federation Users: %v", federations)

}

func federationHandler(w http.ResponseWriter, r *http.Request) {

	query := r.URL.Query().Get("q")
	queryType := r.URL.Query().Get("type")

	var (
		federation *model.Users
		domain     string
	)

	switch queryType {
	case "name":

		name, err := url.QueryUnescape(query)
		if err != nil {
			logger.Error("Error unescaping query: %s", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		var memo *string

		domain = strings.Split(name, "*")[1]
		name = strings.Split(name, "*")[0]

		if strings.Contains(name, " ") {
			nameMemo := strings.Split(name, " ")
			name = nameMemo[0]
			memo = &nameMemo[1]
			logger.Info("Name: %s, Memo: %s", name, *memo)
		}

		data, ok := searchByUser(name)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		} else {
			if memo != nil && data.Memo == "" {
				data.Memo = *memo
			}

			if data.Memo != "" {
				data.MemoType = "text"
			}

			federation = data
		}
	case "id":
		data, ok := searchByID(query)
		if !ok {
			w.WriteHeader(http.StatusNotFound)
			return
		}

		if data.Memo != "" {
			data.MemoType = "text"
		}

		domain = r.Host

		federation = data
	default:
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	federation.Name = federation.Name + "*" + domain

	data, err := json.Marshal(federation)
	if err != nil {
		logger.Error("Error marshalling data: %s", err)
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func searchByUser(name string) (*model.Users, bool) {
	for _, user := range federations {
		if user.Name == name {
			return &user, true
		}
	}

	return nil, false
}

func searchByID(id string) (*model.Users, bool) {
	for _, user := range federations {
		if user.ID == id {
			return &user, true
		}
	}

	return nil, false
}
