package sonargo

import "net/http"

type HotspotsService struct {
	client *Client
}

type HotspotsSearchOption struct {
	ProjectKey      string `url:"projectKey,omitempty"` // Description:"Comma-separated list of component keys. Retrieve issues associated to a specific list of components (and all its descendants). A component can be a portfolio, project, module, directory or file.",ExampleValue:"my_project"
	P               int    `url:"p,omitempty"`          // Description:"1-based page number",ExampleValue:"42"
	Ps              int    `url:"ps,omitempty"`         // Description:"Page size. Must be greater than 0 and less or equal than 500",ExampleValue:"20"
	Status          string `url:"status,omitempty"`     // Description:"Status of the hotspot",Example:"TO_REVIEW"
	OnlyMine        bool   `url:"onlyMine,omitempty"`   // Description:"Only include hotspots assigned to the user"
	SinceLeakPeriod bool   `url:"sinceLeakPeriod,omitempty"`
}

type Hotspot struct {
	Key                      string        `json:"key"`
	Component                string        `json:"component"`
	Project                  string        `json:"project"`
	SecurityCategory         string        `json:"securityCategory"`
	VulnerabilityProbability string        `json:"vulnerabilityProbability"`
	Status                   string        `json:"status"`
	Line                     int           `json:"line"`
	Message                  string        `json:"message"`
	Author                   string        `json:"author"`
	CreationDate             string        `json:"creationDate"`
	UpdateDate               string        `json:"updateDate"`
	TextRange                TextRange     `json:"textRange"`
	Flows                    []interface{} `json:"flows"`
}

type HotspotsSearchObject struct {
	Component []*Component `json:"components,omitempty"`
	Hotspots  []*Hotspot   `json:"hotspots,omitempty"`
	Paging    *Paging      `json:"paging,omitempty"`
}

// Search Search for issues.<br>At most one of the following parameters can be provided at the same time: componentKeys, componentUuids, components, componentRootUuids, componentRoots.<br>Requires the 'Browse' permission on the specified project(s).
func (s *HotspotsService) Search(opt *HotspotsSearchOption) (v *HotspotsSearchObject, resp *http.Response, err error) {
	err = s.ValidateSearchOpt(opt)
	if err != nil {
		return
	}
	req, err := s.client.NewRequest("GET", "hotspots/search", opt)
	if err != nil {
		return
	}
	v = new(HotspotsSearchObject)
	resp, err = s.client.Do(req, v)
	if err != nil {
		return nil, resp, err
	}
	return
}
