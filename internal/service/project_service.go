package service

import (
	"bytes"
	"errors"
	"io"
	"log"
	"net/http"

	"github.com/goccy/go-json"

	"github.com/XxThunderBlastxX/thunder-cli/internal/config"
	"github.com/XxThunderBlastxX/thunder-cli/internal/model"
)

type IProject interface {
	AddProject(project model.Project) error
}

type ProjectService struct {
	config *config.Config
}

func NewProjectService(config *config.Config) IProject {
	return &ProjectService{
		config: config,
	}
}

func (p *ProjectService) AddProject(project model.Project) error {
	projectJson, err := json.Marshal(project)
	if err != nil {
		return err
	}
	k, err := NewKeyRingService()
	if err != nil {
		return err
	}

	accessToken, err := k.Get("AUTH_ACCESS_TOKEN")
	if err != nil {
		return err
	}

	// Creating new http request
	req, err := http.NewRequest("POST", p.config.BaseApiUrl+"/projects/add", bytes.NewBuffer(projectJson))
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+string(accessToken.Data))

	// Create a new HTTP client and send the request
	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		log.Fatal("Error sending request: ", err)
	}

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {

		}
	}(res.Body)

	if res.StatusCode != 200 {
		return errors.New("ops! could not add project. Please try again")
	}

	return nil
}
