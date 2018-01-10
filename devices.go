package packngo

import (
	"fmt"
	"strings"
)

const deviceBasePath = "/devices"

// DeviceService interface defines available device methods
type DeviceService interface {
	List(ProjectID string) ([]Device, *Response, error)
	Get(string) (*Device, *Response, error)
	GetAndInclude(string, []string) (*Device, *Response, error)
	Create(*DeviceCreateRequest) (*Device, *Response, error)
	Update(string, *DeviceUpdateRequest) (*Device, *Response, error)
	Delete(string) (*Response, error)
	Reboot(string) (*Response, error)
	PowerOff(string) (*Response, error)
	PowerOn(string) (*Response, error)
	Lock(string) (*Response, error)
	Unlock(string) (*Response, error)
}

type devicesRoot struct {
	Devices []Device `json:"devices"`
}

// Device represents a Packet device
type Device struct {
	ID                  string                 `json:"id"`
	Href                string                 `json:"href,omitempty"`
	Hostname            string                 `json:"hostname,omitempty"`
	State               string                 `json:"state,omitempty"`
	Created             string                 `json:"created_at,omitempty"`
	Updated             string                 `json:"updated_at,omitempty"`
	Locked              bool                   `json:"locked,omitempty"`
	BillingCycle        string                 `json:"billing_cycle,omitempty"`
	Storage             map[string]interface{} `json:"storage,omitempty"`
	Tags                []string               `json:"tags,omitempty"`
	Network             []*IPAddressAssignment `json:"ip_addresses"`
	Volumes             []*Volume              `json:"volumes"`
	OS                  *OS                    `json:"operating_system,omitempty"`
	Plan                *Plan                  `json:"plan,omitempty"`
	Facility            *Facility              `json:"facility,omitempty"`
	Project             *Project               `json:"project,omitempty"`
	ProvisionEvents     []*ProvisionEvent      `json:"provisioning_events,omitempty"`
	ProvisionPer        float32                `json:"provisioning_percentage,omitempty"`
	UserData            string                 `json:"userdata,omitempty"`
	RootPassword        string                 `json:"root_password,omitempty"`
	IPXEScriptURL       string                 `json:"ipxe_script_url,omitempty"`
	AlwaysPXE           bool                   `json:"always_pxe,omitempty"`
	HardwareReservation Href                   `json:"hardware_reservation,omitempty"`
	SpotInstance        bool                   `json:"spot_instance,omitempty"`
	SpotPriceMax        float64                `json:"spot_price_max,omitempty"`
	TerminationTime     *Timestamp             `json:"termination_time,omitempty"`
	NetworkPorts        []Port                 `json:"network_ports"`
}

type ProvisionEvent struct {
	ID            string     `json:"id"`
	Body          string     `json:"body"`
	CreatedAt     *Timestamp `json:"created_at,omitempty"`
	Href          string     `json:"href"`
	Interpolated  string     `json:"interpolated"`
	Relationships []Href     `json:"relationships"`
	State         string     `json:"state"`
	Type          string     `json:"type"`
}

func (d Device) String() string {
	return Stringify(d)
}

// DeviceCreateRequest type used to create a Packet device
type DeviceCreateRequest struct {
	Hostname              string     `json:"hostname"`
	Plan                  string     `json:"plan"`
	Facility              string     `json:"facility"`
	OS                    string     `json:"operating_system"`
	BillingCycle          string     `json:"billing_cycle"`
	ProjectID             string     `json:"project_id"`
	UserData              string     `json:"userdata"`
	Storage               string     `json:"storage,omitempty"`
	Tags                  []string   `json:"tags"`
	IPXEScriptURL         string     `json:"ipxe_script_url,omitempty"`
	PublicIPv4SubnetSize  int        `json:"public_ipv4_subnet_size,omitempty"`
	AlwaysPXE             bool       `json:"always_pxe,omitempty"`
	HardwareReservationID string     `json:"hardware_reservation_id,omitempty"`
	SpotInstance          bool       `json:"spot_instance,omitempty"`
	SpotPriceMax          float64    `json:"spot_price_max,omitempty,string"`
	TerminationTime       *Timestamp `json:"termination_time,omitempty"`
}

// DeviceUpdateRequest type used to update a Packet device
type DeviceUpdateRequest struct {
	Hostname      string   `json:"hostname"`
	Description   string   `json:"description"`
	UserData      string   `json:"userdata"`
	Locked        bool     `json:"locked"`
	Tags          []string `json:"tags"`
	AlwaysPXE     bool     `json:"always_pxe,omitempty"`
	IPXEScriptURL string   `json:"ipxe_script_url,omitempty"`
}

func (d DeviceCreateRequest) String() string {
	return Stringify(d)
}

// DeviceActionRequest type used to execute actions on devices
type DeviceActionRequest struct {
	Type string `json:"type"`
}

func (d DeviceActionRequest) String() string {
	return Stringify(d)
}

// DeviceServiceOp implements DeviceService
type DeviceServiceOp struct {
	client *Client
}

// List returns devices on a project
func (s *DeviceServiceOp) List(projectID string) ([]Device, *Response, error) {
	path := fmt.Sprintf("%s/%s/devices?include=facility", projectBasePath, projectID)
	root := new(devicesRoot)

	resp, err := s.client.DoRequest("GET", path, nil, root)
	if err != nil {
		return nil, resp, err
	}

	return root.Devices, resp, err
}

// Get returns a device by id
func (s *DeviceServiceOp) Get(deviceID string) (*Device, *Response, error) {
	return s.GetAndInclude(deviceID, []string{"facility"})
}

// GetAndInclude returns a device by id with full information of the provided resources
// via 'includes' rather than href links
func (s *DeviceServiceOp) GetAndInclude(deviceID string, includes []string) (*Device, *Response, error) {
	path := fmt.Sprintf("%s/%s?include=%s", deviceBasePath, deviceID, strings.Join(includes, ","))
	device := new(Device)

	resp, err := s.client.DoRequest("GET", path, nil, device)
	if err != nil {
		return nil, resp, err
	}

	return device, resp, err
}

// Create creates a new device
func (s *DeviceServiceOp) Create(createRequest *DeviceCreateRequest) (*Device, *Response, error) {
	path := fmt.Sprintf("%s/%s/devices", projectBasePath, createRequest.ProjectID)
	device := new(Device)

	resp, err := s.client.DoRequest("POST", path, createRequest, device)
	if err != nil {
		return nil, resp, err
	}

	return device, resp, err
}

// Update updates an existing device
func (s *DeviceServiceOp) Update(deviceID string, updateRequest *DeviceUpdateRequest) (*Device, *Response, error) {
	path := fmt.Sprintf("%s/%s?include=facility", deviceBasePath, deviceID)
	device := new(Device)

	resp, err := s.client.DoRequest("PUT", path, updateRequest, device)
	if err != nil {
		return nil, resp, err
	}

	return device, resp, err
}

// Delete deletes a device
func (s *DeviceServiceOp) Delete(deviceID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", deviceBasePath, deviceID)

	return s.client.DoRequest("DELETE", path, nil, nil)
}

// Reboot reboots on a device
func (s *DeviceServiceOp) Reboot(deviceID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s/actions", deviceBasePath, deviceID)
	action := &DeviceActionRequest{Type: "reboot"}

	return s.client.DoRequest("POST", path, action, nil)
}

// PowerOff powers on a device
func (s *DeviceServiceOp) PowerOff(deviceID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s/actions", deviceBasePath, deviceID)
	action := &DeviceActionRequest{Type: "power_off"}

	return s.client.DoRequest("POST", path, action, nil)
}

// PowerOn powers on a device
func (s *DeviceServiceOp) PowerOn(deviceID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s/actions", deviceBasePath, deviceID)
	action := &DeviceActionRequest{Type: "power_on"}

	return s.client.DoRequest("POST", path, action, nil)
}

type lockDeviceType struct {
	Locked bool `json:"locked"`
}

// Lock sets a device to "locked"
func (s *DeviceServiceOp) Lock(deviceID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", deviceBasePath, deviceID)
	action := lockDeviceType{Locked: true}

	return s.client.DoRequest("PATCH", path, action, nil)
}

// Unlock sets a device to "locked"
func (s *DeviceServiceOp) Unlock(deviceID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", deviceBasePath, deviceID)
	action := lockDeviceType{Locked: false}

	return s.client.DoRequest("PATCH", path, action, nil)
}
