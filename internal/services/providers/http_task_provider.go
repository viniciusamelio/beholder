package providers

import (
	"beholder-api/internal/dtos"
	"beholder-api/internal/services"
	"log"
	"net/http"

	"github.com/vicanso/go-axios"
)

type HttpTaskProvider struct{}

func NewHttpTasksProvider() services.TaskService {
	return &HttpTaskProvider{}
}

func (p *HttpTaskProvider) Execute(i *dtos.GetCallsFromSessionResponseDto) error {
	instance := axios.NewInstance(&axios.InstanceConfig{
		BaseURL: i.BaseURL,
		Headers: http.Header{
			"Content-Type": []string{"application/json"},
		},
	})
	for _, v := range *i.Calls {
		call := v
		go func() {
			instance.Request(&axios.Config{
				Method: call.Method,
				URL:    *call.Path,
				Body:   *call.Body,
				OnError: func(err error, config *axios.Config) (newErr error) {
					log.Default().Printf("Call %s responded with error %s\n", call.Name, err.Error())
					return err
				},
				OnDone: func(config *axios.Config, resp *axios.Response, err error) {
					log.Default().Printf("Call %s responded with status %d\n", call.Name, resp.Status)
				},
			})
		}()
	}
	return nil
}
