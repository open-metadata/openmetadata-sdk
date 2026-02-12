package ometa

import "time"

type Client struct {
	config  *ClientConfig
	backend Backend

	Tables *TableService
	Databases *DatabaseSvc
	DatabaseSchemas *DatabaseSchemaService
	Users *UserService
	Teams *TeamService
	Glossaries *GlossaryService
	GlossaryTerms *GlossaryTermService
	Dashboards *DashboardSvc
	Pipelines *PipelineSvc
	Topics *TopicService
	Charts *ChartService
	Bots *BotService
	Classifications *ClassificationService
	Containers *ContainerService
	DataProducts *DataProductService
	Domains *DomainService
	Tags *TagService
	Roles *RoleService
	Policies *PolicyService
	MlModels *MlModelSvc
	Metrics *MetricService
	StoredProcedures *StoredProcedureService
	SearchIndexes *SearchIndexService
	Queries *QueryService
	Personas *PersonaService
	DashboardDataModels *DashboardDataModelService
	APICollections *APICollectionService
	APIEndpoints *APIEndpointService
	TestCases *TestCaseService
	TestSuites *TestSuiteService
	TestDefinitions *TestDefinitionService
	EventSubscriptions *EventSubscriptionService
	WorkflowDefinitions *WorkflowDefinitionService
	DatabaseServices *DatabaseServiceService
	DashboardServices *DashboardServiceService
	MessagingServices *MessagingServiceService
	PipelineServices *PipelineServiceService
	MlModelServices *MlModelServiceService
	StorageServices *StorageServiceService
	SearchServices *SearchServiceService
	MetadataServices *MetadataServiceService
	ApiServices *ApiServiceService
	DriveServices *DriveServiceService
	LLMServices *LLMServiceService
	SecurityServices *SecurityServiceService
	AIApplications *AIApplicationService
	AIGovernancePolicies *AIGovernancePolicyService
	DataContracts *DataContractService
	Directories *DirectoryService
	Documents *DocumentService
	Files *FileService
	LLMModels *LLMModelService
	LearningResources *LearningResourceService
	NotificationTemplates *NotificationTemplateService
	PromptTemplates *PromptTemplateService
	Spreadsheets *SpreadsheetService
	Worksheets *WorksheetService
	IngestionPipelines *IngestionPipelineService
}

func NewClient(
	baseUrl string,
	opts ...ClientOption) *Client {
	config := defaultConfig(baseUrl)
	for _, opt := range opts {
		opt(config)
	}

	backend := NewHTTPBackend(config)

	c := &Client{
		config:  config,
		backend: backend,
	}

	c.Tables = &TableService{backend: backend}
	c.Databases = &DatabaseSvc{backend: backend}
	c.DatabaseSchemas = &DatabaseSchemaService{backend: backend}
	c.Users = &UserService{backend: backend}
	c.Teams = &TeamService{backend: backend}
	c.Glossaries = &GlossaryService{backend: backend}
	c.GlossaryTerms = &GlossaryTermService{backend: backend}
	c.Dashboards = &DashboardSvc{backend: backend}
	c.Pipelines = &PipelineSvc{backend: backend}
	c.Topics = &TopicService{backend: backend}
	c.Charts = &ChartService{backend: backend}
	c.Bots = &BotService{backend: backend}
	c.Classifications = &ClassificationService{backend: backend}
	c.Containers = &ContainerService{backend: backend}
	c.DataProducts = &DataProductService{backend: backend}
	c.Domains = &DomainService{backend: backend}
	c.Tags = &TagService{backend: backend}
	c.Roles = &RoleService{backend: backend}
	c.Policies = &PolicyService{backend: backend}
	c.MlModels = &MlModelSvc{backend: backend}
	c.Metrics = &MetricService{backend: backend}
	c.StoredProcedures = &StoredProcedureService{backend: backend}
	c.SearchIndexes = &SearchIndexService{backend: backend}
	c.Queries = &QueryService{backend: backend}
	c.Personas = &PersonaService{backend: backend}
	c.DashboardDataModels = &DashboardDataModelService{backend: backend}
	c.APICollections = &APICollectionService{backend: backend}
	c.APIEndpoints = &APIEndpointService{backend: backend}
	c.TestCases = &TestCaseService{backend: backend}
	c.TestSuites = &TestSuiteService{backend: backend}
	c.TestDefinitions = &TestDefinitionService{backend: backend}
	c.EventSubscriptions = &EventSubscriptionService{backend: backend}
	c.WorkflowDefinitions = &WorkflowDefinitionService{backend: backend}
	c.DatabaseServices = &DatabaseServiceService{backend: backend}
	c.DashboardServices = &DashboardServiceService{backend: backend}
	c.MessagingServices = &MessagingServiceService{backend: backend}
	c.PipelineServices = &PipelineServiceService{backend: backend}
	c.MlModelServices = &MlModelServiceService{backend: backend}
	c.StorageServices = &StorageServiceService{backend: backend}
	c.SearchServices = &SearchServiceService{backend: backend}
	c.MetadataServices = &MetadataServiceService{backend: backend}
	c.ApiServices = &ApiServiceService{backend: backend}
	c.DriveServices = &DriveServiceService{backend: backend}
	c.LLMServices = &LLMServiceService{backend: backend}
	c.SecurityServices = &SecurityServiceService{backend: backend}
	c.AIApplications = &AIApplicationService{backend: backend}
	c.AIGovernancePolicies = &AIGovernancePolicyService{backend: backend}
	c.DataContracts = &DataContractService{backend: backend}
	c.Directories = &DirectoryService{backend: backend}
	c.Documents = &DocumentService{backend: backend}
	c.Files = &FileService{backend: backend}
	c.LLMModels = &LLMModelService{backend: backend}
	c.LearningResources = &LearningResourceService{backend: backend}
	c.NotificationTemplates = &NotificationTemplateService{backend: backend}
	c.PromptTemplates = &PromptTemplateService{backend: backend}
	c.Spreadsheets = &SpreadsheetService{backend: backend}
	c.Worksheets = &WorksheetService{backend: backend}
	c.IngestionPipelines = &IngestionPipelineService{backend: backend}

	return c
}

func defaultConfig(baseURL string) *ClientConfig {
	return &ClientConfig{
		BaseURL:    baseURL,
		APIVersion: "v1",
		Retry:      3,
		RetryWait:  30 * time.Second,
		RetryCodes: []int{429, 504},
	}
}
