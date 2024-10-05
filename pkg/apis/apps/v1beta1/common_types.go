package v1beta1

// MinerAddressType describes a valid MinerAddress type.
type MinerAddressType string

// MinerAddress contains information for the miner's address.
type MinerAddress struct {
	// Miner address type, one of Hostname, ExternalIP or InternalIP.
	Type MinerAddressType `json:"type" protobuf:"bytes,1,opt,name=type,casttype=MinerAddressType"`

	// The machine address.
	Address string `json:"address" protobuf:"bytes,2,opt,name=address"`
}

// MinerAddresses is a slice of MinerAddress items to be used by infrastructure providers.
type MinerAddresses []MinerAddress

type ObjectMeta struct {
	// Map of string keys and values that can be used to organize and categorize
	// (scope and select) objects. May match selectors of replication controllers
	// and services.
	// More info: http://kubernetes.io/docs/user-guide/labels
	// +optional
	Labels map[string]string `json:"labels,omitempty" protobuf:"bytes,1,rep,name=labels"`

	// Annotations is an unstructured key value map stored with a resource that may be
	// set by external tools to store and retrieve arbitrary metadata. They are not
	// queryable and should be preserved when modifying objects.
	// More info: http://kubernetes.io/docs/user-guide/annotations
	// +optional
	Annotations map[string]string `json:"annotations,omitempty" protobuf:"bytes,2,rep,name=annotations"`
}
