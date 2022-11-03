
package k8s

import (
	"fmt"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts"
	"net/http"
	"github.com/dome9/dome9-sdk-go/services/cloudaccounts/k8s"
)

type Service k8s.Service

func (service *Service) EnableImageAssurance(body k8s.ImageAssuranceEnableRequest) (*http.Response, error) {
	relativeURL := fmt.Sprintf("%s/%s/%s", cloudaccounts.RESTfulPathK8S, cloudaccounts.RESTfulPathK8SImageAssurance, cloudaccounts.RESTfulPathK8sEnable)
	resp, err := service.Client.NewRequestDo("POST", relativeURL, nil, body, nil)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

