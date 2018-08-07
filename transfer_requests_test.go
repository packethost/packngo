package packngo

import (
	"testing"
)

func TestAccTransferRequests(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)
	c := setup(t)

	rs := testProjectPrefix + randString8()
	ocr := OrganizationCreateRequest{
		Name:        rs,
		Description: "Managed by Packngo.",
		Website:     "http://example.com",
		Twitter:     "foo",
	}
	org, _, err := c.Organizations.Create(&ocr)
	if err != nil {
		t.Fatal(err)
	}
	defer organizationTeardown(c)

	rs = testProjectPrefix + randString8()

	pcr := ProjectCreateRequest{Name: rs}
	p, _, err := c.Projects.Create(&pcr)
	if err != nil {
		t.Fatal(err)
	}
	defer projectTeardown(c)

	_, err = c.TransferRequests.TransferProject(p.ID, org.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.TransferRequests.Accept(p.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.TransferRequests.Decline(p.ID)
	if err != nil {
		t.Fatal(err)
	}

	_, err = c.TransferRequests.TransferProject(p.ID, org.ID)
	if err != nil {
		t.Fatal(err)
	}
}