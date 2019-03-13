package packngo

import "fmt"

const (
	connectBasePath = "/packet-connect/connections"
)

type ConnectService interface {
	List(string, *ListOptions) ([]Connect, *Response, error)
	Get(string, *GetOptions) (*Connect, *Response, error)
	Delete(string) (*Response, error)
	Create(*ConnectCreateRequest) (*Connect, *Response, error)
	//Update(string, *VolumeUpdateRequest) (*Volume, *Response, error)
	//Provision(string) (*Response, error)
	//Deprovision(string) (*Response, error)
}

type ConnectCreateRequest struct {
	Name            string   `json:"name"`
	ProjectID       string   `json:"project_id"`
	ProviderID      string   `json:"provider_id"`
	ProviderPayload string   `json:"provider_payload"`
	Facility        string   `json:"facility"`
	PortSpeed       int      `json:"port_speed"`
	VLAN            int      `json:"vlan"`
	Tags            []string `json:"tags,omitempty"`
	Description     string   `json:"description,omitempty"`
}

type Connect struct {
	ID              string `json:"id"`
	Status          string `json:"status"`
	Name            string `json:"name"`
	ProjectID       string `json:"project_id"`
	ProviderID      string `json:"provider_id"`
	ProviderPayload string `json:"provider_payload"`
	Facility        string `json:"facility"`
	PortSpeed       int    `json:"port_speed"`
	VLAN            int    `json:"vlan"`
	Description     string `json:"description,omitempty"`
}

type ConnectServiceOp struct {
	client *Client
}

type connectsRoot struct {
	Connects []Connect `json:"connections"`
	Meta     meta      `json:"meta"`
}

func (c *ConnectServiceOp) List(projectID string, listOpt *ListOptions) (connects []Connect, resp *Response, err error) {
	params := createListOptionsURL(listOpt)

	project_param := fmt.Sprintf("project_id=%s", projectID)
	if params == "" {
		params = project_param
	} else {
		params = fmt.Sprintf("%s&%s", params, project_param)
	}
	path := fmt.Sprintf("%s/?%s", connectBasePath, params)

	for {
		subset := new(connectsRoot)

		resp, err = c.client.DoRequest("GET", path, nil, subset)
		if err != nil {
			return nil, resp, err
		}

		connects = append(connects, subset.Connects...)

		if subset.Meta.Next != nil && (listOpt == nil || listOpt.Page == 0) {
			path = subset.Meta.Next.Href
			if params != "" {
				path = fmt.Sprintf("%s&%s", path, params)
			}
			continue
		}

		return
	}
}

func (c *ConnectServiceOp) Create(createRequest *ConnectCreateRequest) (*Connect, *Response, error) {
	url := fmt.Sprintf("%s", connectBasePath)
	connect := new(Connect)

	resp, err := c.client.DoRequest("POST", url, createRequest, connect)
	if err != nil {
		return nil, resp, err
	}

	return connect, resp, err
}

func (c *ConnectServiceOp) Get(connectID string, getOpt *GetOptions) (*Connect, *Response, error) {
	params := createGetOptionsURL(getOpt)
	path := fmt.Sprintf("%s/%s?%s", connectBasePath, connectID, params)
	connect := new(Connect)

	resp, err := c.client.DoRequest("GET", path, nil, connect)
	if err != nil {
		return nil, resp, err
	}

	return connect, resp, err
}

func (c *ConnectServiceOp) Delete(connectID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", connectBasePath, connectID)

	return c.client.DoRequest("DELETE", path, nil, nil)
}
