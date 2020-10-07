package packngo

import "fmt"

const emailBasePath = "emails"

// EmailRequest type used to add an email address to the current user
type EmailRequest struct {
	Address string `json:"address,omitempty"`
	Default *bool  `json:"default,omitempty"`
}

// EmailService interface defines available email methods
type EmailService interface {
	Get(string, *GetOptions) (*Email, *Response, error)
	Create(*EmailRequest) (*Email, *Response, error)
	Update(string, *EmailRequest) (*Email, *Response, error)
	Delete(string) (*Response, error)
}

// Email represents a user's email address
type Email struct {
	ID      string `json:"id"`
	Address string `json:"address"`
	Default bool   `json:"default,omitempty"`
	URL     string `json:"href,omitempty"`
}

func (e Email) String() string {
	return Stringify(e)
}

// EmailServiceOp implements EmailService
type EmailServiceOp struct {
	client *Client
}

// Get retrieves an email by id
func (s *EmailServiceOp) Get(emailID string, getOpt *GetOptions) (*Email, *Response, error) {
	params := urlQuery(getOpt)
	path := fmt.Sprintf("%s/%s?%s", emailBasePath, emailID, params)
	email := new(Email)

	resp, err := s.client.DoRequest("GET", path, nil, email)
	if err != nil {
		return nil, resp, err
	}

	return email, resp, err
}

// Create adds a new email address to the current user.
func (s *EmailServiceOp) Create(request *EmailRequest) (*Email, *Response, error) {
	email := new(Email)

	resp, err := s.client.DoRequest("POST", emailBasePath, request, email)
	if err != nil {
		return nil, resp, err
	}

	return email, resp, err
}

// Delete removes the email addres from the current user account
func (s *EmailServiceOp) Delete(emailID string) (*Response, error) {
	path := fmt.Sprintf("%s/%s", emailBasePath, emailID)

	resp, err := s.client.DoRequest("DELETE", path, nil, nil)
	if err != nil {
		return resp, err
	}

	return resp, err
}

// Update email parameters
func (s *EmailServiceOp) Update(emailID string, request *EmailRequest) (*Email, *Response, error) {
	email := new(Email)
	path := fmt.Sprintf("%s/%s", emailBasePath, emailID)

	resp, err := s.client.DoRequest("PUT", path, request, email)
	if err != nil {
		return nil, resp, err
	}

	return email, resp, err
}
