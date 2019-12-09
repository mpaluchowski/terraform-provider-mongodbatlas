package mongodbatlas

import (
	"context"
	"fmt"
	"net/http"
)

const globalClustersBasePath = "groups/%s/clusters/%s/globalWrites/%s"

//GlobalClustersService is an interface for interfacing with the Global Clusters
// endpoints of the MongoDB Atlas API.
//See more: https://docs.atlas.mongodb.com/reference/api/global-clusters/
type GlobalClustersService interface {
	Get(context.Context, string, string) (*GlobalCluster, *Response, error)
	AddManagedNamespace(context.Context, string, string, *ManagedNamespace) (*GlobalCluster, *Response, error)
	DeleteManagedNamespace(context.Context, string, string, *ManagedNamespace) (*GlobalCluster, *Response, error)
	AddCustomZoneMappings(context.Context, string, string, *CustomZoneMappingsRequest) (*GlobalCluster, *Response, error)
	DeleteCustomZoneMappings(context.Context, string, string) (*GlobalCluster, *Response, error)
}

//GlobalClustersServiceOp handles communication with the GlobalClusters related methos of the
//MongoDB Atlas API
type GlobalClustersServiceOp struct {
	client *Client
}

var _ GlobalClustersService = &GlobalClustersServiceOp{}

// GlobalCluster represents MongoDB Global Cluster Configuration in your Global Cluster.
type GlobalCluster struct {
	CustomZoneMapping map[string]string  `json:"customZoneMapping"`
	ManagedNamespaces []ManagedNamespace `json:"managedNamespaces"`
}

// ManagedNamespace represents the information about managed namespace configuration.
type ManagedNamespace struct {
	Db             string `json:"db"`
	Collection     string `json:"collection"`
	CustomShardKey string `json:"customShardKey,omitempty"`
}

type CustomZoneMappingsRequest struct {
	CustomZoneMappings []CustomZoneMapping `json:"customZoneMappings"`
}

type CustomZoneMapping struct {
	Location string `json:"location"`
	Zone     string `json:"zone"`
}

// //Get retrieves all managed namespaces and custom zone mappings associated with the specified Global Cluster.
//See more: https://docs.atlas.mongodb.com/reference/api/global-clusters-retrieve-namespaces/
func (s *GlobalClustersServiceOp) Get(ctx context.Context, groupID string, clusterName string) (*GlobalCluster, *Response, error) {
	if clusterName == "" {
		return nil, nil, NewArgError("username", "must be set")
	}

	path := fmt.Sprintf("groups/%s/clusters/%s/globalWrites", groupID, clusterName)

	req, err := s.client.NewRequest(ctx, http.MethodGet, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GlobalCluster)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

//AddManagedNamespace adds a managed namespace to the specified Global Cluster.
//See more: https://docs.atlas.mongodb.com/reference/api/database-users-create-a-user/
func (s *GlobalClustersServiceOp) AddManagedNamespace(ctx context.Context, groupID string, clusterName string, createRequest *ManagedNamespace) (*GlobalCluster, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := fmt.Sprintf(globalClustersBasePath, groupID, clusterName, "managedNamespaces")

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(GlobalCluster)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

//DeleteManagedNamespace deletes the managed namespace configuration of the global cluster given.
//See more: https://docs.atlas.mongodb.com/reference/api/global-clusters-delete-namespace/
func (s *GlobalClustersServiceOp) DeleteManagedNamespace(ctx context.Context, groupID string, clusterName string, deleteRequest *ManagedNamespace) (*GlobalCluster, *Response, error) {
	if deleteRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := fmt.Sprintf(globalClustersBasePath, groupID, clusterName, "managedNamespaces")

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, nil, err
	}

	q := req.URL.Query()
	q.Add("collection", deleteRequest.Collection)
	q.Add("db", deleteRequest.Db)
	req.URL.RawQuery = q.Encode()

	root := new(GlobalCluster)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

//AddCustomZoneMappings adds an entry to the list of custom zone mappings for the specified Global Cluster.
//See more: https://docs.atlas.mongodb.com/reference/api/global-clusters-add-customzonemapping/
func (s *GlobalClustersServiceOp) AddCustomZoneMappings(ctx context.Context, groupID string, clusterName string, createRequest *CustomZoneMappingsRequest) (*GlobalCluster, *Response, error) {
	if createRequest == nil {
		return nil, nil, NewArgError("createRequest", "cannot be nil")
	}

	path := fmt.Sprintf(globalClustersBasePath, groupID, clusterName, "customZoneMapping")

	req, err := s.client.NewRequest(ctx, http.MethodPost, path, createRequest)
	if err != nil {
		return nil, nil, err
	}

	root := new(GlobalCluster)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}

	return root, resp, err
}

//DeleteManagedNamespace removes all custom zone mappings from the specified Global Cluster.
//See more: https://docs.atlas.mongodb.com/reference/api/global-clusters-delete-namespace/
func (s *GlobalClustersServiceOp) DeleteCustomZoneMappings(ctx context.Context, groupID string, clusterName string) (*GlobalCluster, *Response, error) {
	path := fmt.Sprintf(globalClustersBasePath, groupID, clusterName, "customZoneMapping")

	req, err := s.client.NewRequest(ctx, http.MethodDelete, path, nil)
	if err != nil {
		return nil, nil, err
	}

	root := new(GlobalCluster)
	resp, err := s.client.Do(ctx, req, root)
	if err != nil {
		return nil, resp, err
	}
	return root, resp, err
}