package service

import (
	"bytes"
	"errors"
	"io"
	"net/http"

	"github.com/goccy/go-json"

	"github.com/XxThunderBlastxX/thunder-cli/internal/model"
)

type IProject interface {
	AddProject(project model.Project) error
}

type ProjectService struct {
	http    *http.Client
	baseUrl string
}

func NewProjectService(http *http.Client, baseUrl string) IProject {
	return &ProjectService{
		http:    http,
		baseUrl: baseUrl,
	}
}

func (p *ProjectService) AddProject(project model.Project) error {
	projectJson, err := json.Marshal(project)
	if err != nil {
		return err
	}

	res, err := p.http.Post(p.baseUrl+"/projects/add", "application/json", bytes.NewBuffer(projectJson))
	if err != nil {
		return err
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
