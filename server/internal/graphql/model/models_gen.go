// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
)

type AddonList struct {
	Name  string `json:"name"`
	Owner string `json:"owner"`
}

type AddonStatusInput struct {
	Selector     *MeshType `json:"selector"`
	K8scontextID string    `json:"k8scontextID"`
	TargetStatus Status    `json:"targetStatus"`
}

type ApplicationPage struct {
	Page         int                  `json:"page"`
	PageSize     int                  `json:"page_size"`
	TotalCount   int                  `json:"total_count"`
	Applications []*ApplicationResult `json:"applications"`
}

type ApplicationResult struct {
	ID              string      `json:"id"`
	Name            string      `json:"name"`
	ApplicationFile string      `json:"application_file"`
	Type            *NullString `json:"type"`
	UserID          string      `json:"user_id"`
	Location        *Location   `json:"location"`
	CreatedAt       *string     `json:"created_at"`
	UpdatedAt       *string     `json:"updated_at"`
}

type CatalogFilter struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	FilterFile  string                 `json:"filter_file"`
	UserID      string                 `json:"user_id"`
	Location    *Location              `json:"location"`
	Visibility  string                 `json:"visibility"`
	CatalogData map[string]interface{} `json:"catalog_data"`
	CreatedAt   *string                `json:"created_at"`
	UpdatedAt   *string                `json:"updated_at"`
}

type CatalogPattern struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	UserID      string                 `json:"user_id"`
	PatternFile string                 `json:"pattern_file"`
	Location    *Location              `json:"location"`
	Visibility  string                 `json:"visibility"`
	CatalogData map[string]interface{} `json:"catalog_data"`
	CreatedAt   *string                `json:"created_at"`
	UpdatedAt   *string                `json:"updated_at"`
}

type CatalogSelector struct {
	Search string `json:"search"`
	Order  string `json:"order"`
}

type ClusterResources struct {
	Resources []*Resource `json:"resources"`
}

type ConfigurationPage struct {
	Applications *ApplicationPage   `json:"applications"`
	Patterns     *PatternPageResult `json:"patterns"`
	Filters      *FilterPage        `json:"filters"`
}

type Container struct {
	ControlPlaneMemberName string           `json:"controlPlaneMemberName"`
	ContainerName          string           `json:"containerName"`
	Image                  string           `json:"image"`
	Status                 *ContainerStatus `json:"status"`
	Ports                  []*ContainerPort `json:"ports"`
	Resources              interface{}      `json:"resources"`
}

type ContainerPort struct {
	Name          *string `json:"name"`
	ContainerPort int     `json:"containerPort"`
	Protocol      string  `json:"protocol"`
}

type ContainerStatus struct {
	ContainerStatusName string      `json:"containerStatusName"`
	Image               string      `json:"image"`
	State               interface{} `json:"state"`
	LastState           interface{} `json:"lastState"`
	Ready               bool        `json:"ready"`
	RestartCount        interface{} `json:"restartCount"`
	Started             bool        `json:"started"`
	ImageID             interface{} `json:"imageID"`
	ContainerID         interface{} `json:"containerID"`
}

type ControlPlane struct {
	Name    string                `json:"name"`
	Members []*ControlPlaneMember `json:"members"`
}

type ControlPlaneMember struct {
	Name       string       `json:"name"`
	Component  string       `json:"component"`
	Version    string       `json:"version"`
	Namespace  string       `json:"namespace"`
	DataPlanes []*Container `json:"data_planes"`
}

type DataPlane struct {
	Name    string       `json:"name"`
	Proxies []*Container `json:"proxies"`
}

type Error struct {
	Code        string `json:"code"`
	Description string `json:"description"`
}

type FilterPage struct {
	Page       int             `json:"page"`
	PageSize   int             `json:"page_size"`
	TotalCount int             `json:"total_count"`
	Filters    []*FilterResult `json:"filters"`
}

type FilterResult struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	FilterFile  string                 `json:"filter_file"`
	UserID      string                 `json:"user_id"`
	Location    *Location              `json:"location"`
	Visibility  string                 `json:"visibility"`
	CatalogData map[string]interface{} `json:"catalog_data"`
	CreatedAt   *string                `json:"created_at"`
	UpdatedAt   *string                `json:"updated_at"`
}

type K8sContext struct {
	ID                 string                 `json:"id"`
	Name               string                 `json:"name"`
	Auth               map[string]interface{} `json:"auth"`
	Cluster            map[string]interface{} `json:"cluster"`
	Server             string                 `json:"server"`
	Owner              string                 `json:"owner"`
	CreatedBy          string                 `json:"created_by"`
	MesheryInstanceID  string                 `json:"meshery_instance_id"`
	KubernetesServerID string                 `json:"kubernetes_server_id"`
	DeploymentType     string                 `json:"deployment_type"`
	UpdatedAt          string                 `json:"updated_at"`
	CreatedAt          string                 `json:"created_at"`
}

type K8sContextsPage struct {
	TotalCount int           `json:"total_count"`
	Contexts   []*K8sContext `json:"contexts"`
}

type KctlDescribeDetails struct {
	Describe *string `json:"describe"`
	Ctxid    *string `json:"ctxid"`
}

type Location struct {
	Branch *string `json:"branch"`
	Host   *string `json:"host"`
	Path   *string `json:"path"`
	Type   *string `json:"type"`
}

type MeshModelComponent struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type MeshModelRelationship struct {
	Name  string `json:"name"`
	Count int    `json:"count"`
}

type MeshModelSummary struct {
	Components    []*MeshModelComponent    `json:"components"`
	Relationships []*MeshModelRelationship `json:"relationships"`
}

type MeshModelSummarySelector struct {
	Type string `json:"type"`
}

type MeshSyncEvent struct {
	Type      string      `json:"type"`
	Object    interface{} `json:"object"`
	ContextID string      `json:"contextId"`
}

type MesheryControllersStatusListItem struct {
	ContextID  string                  `json:"contextId"`
	Controller MesheryController       `json:"controller"`
	Status     MesheryControllerStatus `json:"status"`
}

type MesheryResult struct {
	MesheryID          *string                `json:"meshery_id"`
	Name               *string                `json:"name"`
	Mesh               *string                `json:"mesh"`
	PerformanceProfile *string                `json:"performance_profile"`
	TestID             *string                `json:"test_id"`
	RunnerResults      map[string]interface{} `json:"runner_results"`
	ServerMetrics      *string                `json:"server_metrics"`
	ServerBoardConfig  *string                `json:"server_board_config"`
	TestStartTime      *string                `json:"test_start_time"`
	UserID             *string                `json:"user_id"`
	UpdatedAt          *string                `json:"updated_at"`
	CreatedAt          *string                `json:"created_at"`
}

type NameSpace struct {
	Namespace string `json:"namespace"`
}

type NullString struct {
	String string `json:"String"`
	Valid  bool   `json:"Valid"`
}

type OAMCapability struct {
	OamDefinition interface{} `json:"oam_definition"`
	ID            *string     `json:"id"`
	OamRefSchema  *string     `json:"oam_ref_schema"`
	Host          *string     `json:"host"`
	Restricted    *bool       `json:"restricted"`
	Metadata      interface{} `json:"metadata"`
}

type OperatorControllerStatus struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Status    Status `json:"status"`
	Error     *Error `json:"error"`
	ContextID string `json:"contextID"`
}

type OperatorControllerStatusPerK8sContext struct {
	ContextID                string                    `json:"contextID"`
	OperatorControllerStatus *OperatorControllerStatus `json:"OperatorControllerStatus"`
}

type OperatorStatus struct {
	Status      Status                      `json:"status"`
	Version     string                      `json:"version"`
	Controllers []*OperatorControllerStatus `json:"controllers"`
	Error       *Error                      `json:"error"`
	ContextID   string                      `json:"contextID"`
}

type OperatorStatusInput struct {
	TargetStatus Status `json:"targetStatus"`
	ContextID    string `json:"contextID"`
}

type OperatorStatusPerK8sContext struct {
	ContextID      string          `json:"contextID"`
	OperatorStatus *OperatorStatus `json:"operatorStatus"`
}

type PageFilter struct {
	Page     string  `json:"page"`
	PageSize string  `json:"pageSize"`
	Order    *string `json:"order"`
	Search   *string `json:"search"`
	From     *string `json:"from"`
	To       *string `json:"to"`
}

type PatternPageResult struct {
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalCount int              `json:"total_count"`
	Patterns   []*PatternResult `json:"patterns"`
}

type PatternResult struct {
	ID          string                 `json:"id"`
	Name        string                 `json:"name"`
	UserID      string                 `json:"user_id"`
	Location    *Location              `json:"location"`
	PatternFile string                 `json:"pattern_file"`
	Visibility  string                 `json:"visibility"`
	CatalogData map[string]interface{} `json:"catalog_data"`
	CanSupport  bool                   `json:"canSupport"`
	Errmsg      *string                `json:"errmsg"`
	CreatedAt   *string                `json:"created_at"`
	UpdatedAt   *string                `json:"updated_at"`
}

type PerfPageProfiles struct {
	Page       int            `json:"page"`
	PageSize   int            `json:"page_size"`
	TotalCount int            `json:"total_count"`
	Profiles   []*PerfProfile `json:"profiles"`
}

type PerfPageResult struct {
	Page       int              `json:"page"`
	PageSize   int              `json:"page_size"`
	TotalCount int              `json:"total_count"`
	Results    []*MesheryResult `json:"results"`
}

type PerfProfile struct {
	ConcurrentRequest int       `json:"concurrent_request"`
	CreatedAt         *string   `json:"created_at"`
	Duration          string    `json:"duration"`
	Endpoints         []*string `json:"endpoints"`
	ID                string    `json:"id"`
	LastRun           *string   `json:"last_run"`
	LoadGenerators    []*string `json:"load_generators"`
	Name              *string   `json:"name"`
	QPS               *int      `json:"qps"`
	TotalResults      *int      `json:"total_results"`
	UpdatedAt         *string   `json:"updated_at"`
	UserID            string    `json:"user_id"`
	RequestHeaders    *string   `json:"request_headers"`
	RequestCookies    *string   `json:"request_cookies"`
	RequestBody       *string   `json:"request_body"`
	ContentType       *string   `json:"content_type"`
	ServiceMesh       *string   `json:"service_mesh"`
}

type ReSyncActions struct {
	ClearDb   string `json:"clearDB"`
	ReSync    string `json:"ReSync"`
	HardReset string `json:"hardReset"`
}

type Resource struct {
	Kind  string `json:"kind"`
	Count int    `json:"count"`
}

type ServiceMeshFilter struct {
	Type          *MeshType `json:"type"`
	K8sClusterIDs []string  `json:"k8sClusterIDs"`
}

type MeshType string

const (
	MeshTypeAllMesh            MeshType = "ALL_MESH"
	MeshTypeInvalidMesh        MeshType = "INVALID_MESH"
	MeshTypeAppMesh            MeshType = "APP_MESH"
	MeshTypeCitrixServiceMesh  MeshType = "CITRIX_SERVICE_MESH"
	MeshTypeConsul             MeshType = "CONSUL"
	MeshTypeIstio              MeshType = "ISTIO"
	MeshTypeKuma               MeshType = "KUMA"
	MeshTypeLinkerd            MeshType = "LINKERD"
	MeshTypeTraefikMesh        MeshType = "TRAEFIK_MESH"
	MeshTypeOctarine           MeshType = "OCTARINE"
	MeshTypeNetworkServiceMesh MeshType = "NETWORK_SERVICE_MESH"
	MeshTypeTanzu              MeshType = "TANZU"
	MeshTypeOpenServiceMesh    MeshType = "OPEN_SERVICE_MESH"
	MeshTypeNginxServiceMesh   MeshType = "NGINX_SERVICE_MESH"
	MeshTypeCiliumServiceMesh  MeshType = "CILIUM_SERVICE_MESH"
)

var AllMeshType = []MeshType{
	MeshTypeAllMesh,
	MeshTypeInvalidMesh,
	MeshTypeAppMesh,
	MeshTypeCitrixServiceMesh,
	MeshTypeConsul,
	MeshTypeIstio,
	MeshTypeKuma,
	MeshTypeLinkerd,
	MeshTypeTraefikMesh,
	MeshTypeOctarine,
	MeshTypeNetworkServiceMesh,
	MeshTypeTanzu,
	MeshTypeOpenServiceMesh,
	MeshTypeNginxServiceMesh,
	MeshTypeCiliumServiceMesh,
}

func (e MeshType) IsValid() bool {
	switch e {
	case MeshTypeAllMesh, MeshTypeInvalidMesh, MeshTypeAppMesh, MeshTypeCitrixServiceMesh, MeshTypeConsul, MeshTypeIstio, MeshTypeKuma, MeshTypeLinkerd, MeshTypeTraefikMesh, MeshTypeOctarine, MeshTypeNetworkServiceMesh, MeshTypeTanzu, MeshTypeOpenServiceMesh, MeshTypeNginxServiceMesh, MeshTypeCiliumServiceMesh:
		return true
	}
	return false
}

func (e MeshType) String() string {
	return string(e)
}

func (e *MeshType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MeshType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MeshType", str)
	}
	return nil
}

func (e MeshType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MesheryController string

const (
	MesheryControllerBroker   MesheryController = "BROKER"
	MesheryControllerOperator MesheryController = "OPERATOR"
	MesheryControllerMeshsync MesheryController = "MESHSYNC"
)

var AllMesheryController = []MesheryController{
	MesheryControllerBroker,
	MesheryControllerOperator,
	MesheryControllerMeshsync,
}

func (e MesheryController) IsValid() bool {
	switch e {
	case MesheryControllerBroker, MesheryControllerOperator, MesheryControllerMeshsync:
		return true
	}
	return false
}

func (e MesheryController) String() string {
	return string(e)
}

func (e *MesheryController) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MesheryController(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MesheryController", str)
	}
	return nil
}

func (e MesheryController) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type MesheryControllerStatus string

const (
	MesheryControllerStatusDeployed    MesheryControllerStatus = "DEPLOYED"
	MesheryControllerStatusNotdeployed MesheryControllerStatus = "NOTDEPLOYED"
	MesheryControllerStatusDeploying   MesheryControllerStatus = "DEPLOYING"
	MesheryControllerStatusUnkown      MesheryControllerStatus = "UNKOWN"
)

var AllMesheryControllerStatus = []MesheryControllerStatus{
	MesheryControllerStatusDeployed,
	MesheryControllerStatusNotdeployed,
	MesheryControllerStatusDeploying,
	MesheryControllerStatusUnkown,
}

func (e MesheryControllerStatus) IsValid() bool {
	switch e {
	case MesheryControllerStatusDeployed, MesheryControllerStatusNotdeployed, MesheryControllerStatusDeploying, MesheryControllerStatusUnkown:
		return true
	}
	return false
}

func (e MesheryControllerStatus) String() string {
	return string(e)
}

func (e *MesheryControllerStatus) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = MesheryControllerStatus(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid MesheryControllerStatus", str)
	}
	return nil
}

func (e MesheryControllerStatus) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type Status string

const (
	StatusEnabled    Status = "ENABLED"
	StatusConnected  Status = "CONNECTED"
	StatusDisabled   Status = "DISABLED"
	StatusProcessing Status = "PROCESSING"
	StatusUnknown    Status = "UNKNOWN"
)

var AllStatus = []Status{
	StatusEnabled,
	StatusConnected,
	StatusDisabled,
	StatusProcessing,
	StatusUnknown,
}

func (e Status) IsValid() bool {
	switch e {
	case StatusEnabled, StatusConnected, StatusDisabled, StatusProcessing, StatusUnknown:
		return true
	}
	return false
}

func (e Status) String() string {
	return string(e)
}

func (e *Status) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = Status(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid Status", str)
	}
	return nil
}

func (e Status) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
