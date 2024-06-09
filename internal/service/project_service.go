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
	RemoveProject(name string) error
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
	cfg, err := config.NewAppConfig()
	if err != nil {
		return err
	}

	accessToken := cfg.Viper.GetString("access_token")

	// Creating new http request
	req, err := http.NewRequest("POST", p.config.BaseApiUrl+"/projects/add", bytes.NewBuffer(projectJson))
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

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
		return errors.New("ops! could not add project. Please try again : " + res.Status)
	}

	return nil
}

func (p *ProjectService) RemoveProject(name string) error {
	cfg, err := config.NewAppConfig()
	if err != nil {
		return err
	}

	accessToken := cfg.Viper.GetString("access_token")

	// Creating new http request
	req, err := http.NewRequest("DELETE", p.config.BaseApiUrl+"/projects/remove?name="+name, nil)
	if err != nil {
		log.Fatal("Error creating request: ", err)
	}

	// Set headers
	req.Header.Set("Content-Type", "application/json")
	req.Header.Add("Authorization", "Bearer "+accessToken)

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
		return errors.New("ops! could not remove project. Please try again : " + res.Status)
	}

	return nil
}
