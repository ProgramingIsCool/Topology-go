package topology

import (
	"encoding/xml"
)

// xmlCI is an internal helper struct for XML encoding/decoding.
type xmlCI struct {
	XMLName xml.Name `xml:"CI"` // so the kids are CI not xmlCI
	Name    string   `xml:"name,attr"`
	Child   []*xmlCI `xml:"CI,omitempty"`
}

// toXMLCI recursively converts CI → xmlCI.
func (ci *CI) toXMLCI() *xmlCI {
	if ci == nil {
		return nil
	}
	children := make([]*xmlCI, 0, len(ci.Child)) // create slice with capacity equla to te number of kids
	for _, c := range ci.Child {
		children = append(children, c.toXMLCI())
	}
	return &xmlCI{
		Name:  ci.Name,
		Child: children,
	}
}

// toCI recursively converts xmlCI → CI.
/*
func (x *xmlCI) toCI() *CI {
	if x == nil {
		return nil
	}
	children := make([]*CI, 0, len(x.Child))
	for _, c := range x.Child {
		children = append(children, c.toCI())
	}
	return &CI{
		Name:  x.Name,
		Child: children,
	}
}
*/
// MarshalToXML converts Topology → XML ([]byte).
func (t *Topology) MarshalToXML() ([]byte, error) {
	xmlRoots := make([]*xmlCI, 0, len(t.Roots))
	for _, r := range t.Roots {
		xmlRoots = append(xmlRoots, r.toXMLCI())
	}
	return xml.MarshalIndent(struct {
		XMLName xml.Name `xml:"Topology"`
		Roots   []*xmlCI `xml:"CI"`
	}{
		Roots: xmlRoots,
	}, "", "  ")
}

// UnmarshalFromXML loads a Topology from XML data ([]byte).
/*
func UnmarshalFromXML(data []byte) (*Topology, error) {
	var wrapper struct {
		Roots []*xmlCI `xml:"CI"`
	}
	if err := xml.Unmarshal(data, &wrapper); err != nil {
		return nil, err
	}

	t := &Topology{}
	for _, r := range wrapper.Roots {
		t.Roots = append(t.Roots, r.toCI())
	}
	return t, nil
}
*/
