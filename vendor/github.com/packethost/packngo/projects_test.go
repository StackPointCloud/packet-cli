package packngo

import (
	"testing"
)

func TestAccProject(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)

	c := setup(t)
	defer projectTeardown(c)

	rs := testProjectPrefix + randString8()
	pcr := ProjectCreateRequest{Name: rs}
	p, _, err := c.Projects.Create(&pcr)
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != rs {
		t.Fatalf("Expected new project name to be %s, not %s", rs, p.Name)
	}
	rs = testProjectPrefix + randString8()
	pur := ProjectUpdateRequest{Name: &rs}
	p, _, err = c.Projects.Update(p.ID, &pur)
	if err != nil {
		t.Fatal(err)
	}
	if p.Name != rs {
		t.Fatalf("Expected the name of the updated project to be %s, not %s", rs, p.Name)
	}
	gotProject, _, err := c.Projects.Get(p.ID)
	if err != nil {
		t.Fatal(err)
	}
	if gotProject.Name != rs {
		t.Fatalf("Expected the name of the GOT project to be %s, not %s", rs, gotProject.Name)
	}

	if gotProject.PaymentMethod.URL == "" {
		t.Fatalf("Empty payment_method: %v", gotProject)
	}
	_, err = c.Projects.Delete(p.ID)
	if err != nil {
		t.Fatal(err)
	}

}

func TestAccCreateOrgProject(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)

	c := setup(t)
	defer projectTeardown(c)

	u, _, err := c.Users.Current()
	if err != nil {
		t.Fatal(err)
	}

	rs := testProjectPrefix + randString8()

	orgPath := "/organizations/" + u.DefaultOrganizationID
	pcr := ProjectCreateRequest{Name: rs}
	p, _, err := c.Projects.Create(&pcr)
	if err != nil {
		t.Fatal(err)
	}
	if p.Organization.URL != orgPath {
		t.Fatalf("Expected new project to be part of org %s, not %v", orgPath, p.Organization)
	}
}

func TestAccCreateNonDefaultOrgProject(t *testing.T) {
	skipUnlessAcceptanceTestsAllowed(t)

	c := setup(t)
	defer organizationTeardown(c)
	defer projectTeardown(c)

	u, _, err := c.Users.Current()
	if err != nil {
		t.Fatal(err)
	}

	orgName := testProjectPrefix + randString8()
	ocr := OrganizationCreateRequest{
		Name:        orgName,
		Description: "Managed by Packngo.",
		Website:     "http://example.com",
		Twitter:     "foo",
	}
	org, _, err := c.Organizations.Create(&ocr)
	if err != nil {
		t.Fatal(err)
	}

	rs := testProjectPrefix + randString8()

	if org.ID == u.DefaultOrganizationID {
		t.Fatalf("Expected new organization %s to not have same ID as Default org %s", org.ID, u.DefaultOrganizationID)
	}

	orgPath := "/organizations/" + org.ID
	pcr := ProjectCreateRequest{Name: rs, OrganizationID: org.ID}
	p, _, err := c.Projects.Create(&pcr)
	if err != nil {
		t.Fatal(err)
	}

	if p.Organization.URL != orgPath {
		t.Fatalf("Expected new project to be part of org %s, not %v", orgPath, p.Organization)
	}

	defaultOrgPath := "/organizations/" + u.DefaultOrganizationID
	if p.Organization.URL == defaultOrgPath {
		t.Fatalf("Expected new project to not be part of org %s", orgPath)
	}
}
