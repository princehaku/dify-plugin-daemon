package definitions

import (
	"github.com/langgenius/dify-plugin-daemon/internal/core/plugin_daemon/access_types"
	"github.com/langgenius/dify-plugin-daemon/pkg/entities/dynamic_select_entities"
	"github.com/langgenius/dify-plugin-daemon/pkg/entities/model_entities"
	"github.com/langgenius/dify-plugin-daemon/pkg/entities/oauth_entities"
	"github.com/langgenius/dify-plugin-daemon/pkg/entities/requests"
	"github.com/langgenius/dify-plugin-daemon/pkg/entities/tool_entities"
)

// PluginDispatcher defines a plugin dispatcher
type PluginDispatcher struct {
	Name               string
	RequestType        interface{} // e.g. requests.RequestInvokeLLM
	ResponseType       interface{} // e.g. requests.ResponseInvokeLLM
	RequestTypeString  string
	ResponseTypeString string
	AccessType         access_types.PluginAccessType
	AccessAction       access_types.PluginAccessAction
	AccessTypeString   string
	AccessActionString string
	BufferSize         int
	Path               string // e.g. "/tool/invoke"
}

// Define all plugin dispatchers
var PluginDispatchers = []PluginDispatcher{
	// { // No need to implement this for now, it has its special implementation in the agent service
	// 	Name:               "InvokeTool",
	// 	RequestType:        requests.RequestInvokeTool{},
	// 	ResponseType:       tool_entities.ToolResponseChunk{},
	// 	AccessType:         access_types.PLUGIN_ACCESS_TYPE_TOOL,
	// 	AccessAction:       access_types.PLUGIN_ACCESS_ACTION_INVOKE_TOOL,
	// 	AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_TOOL",
	// 	AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_INVOKE_TOOL",
	// 	BufferSize:         1024,
	// 	Path:               "/tool/invoke",
	// },
	{
		Name:               "ValidateToolCredentials",
		RequestType:        requests.RequestValidateToolCredentials{},
		ResponseType:       tool_entities.ValidateCredentialsResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_TOOL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_VALIDATE_TOOL_CREDENTIALS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_TOOL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_VALIDATE_TOOL_CREDENTIALS",
		BufferSize:         1,
		Path:               "/tool/validate_credentials",
	},
	{
		Name:               "GetToolRuntimeParameters",
		RequestType:        requests.RequestGetToolRuntimeParameters{},
		ResponseType:       tool_entities.GetToolRuntimeParametersResponse{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_TOOL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_GET_TOOL_RUNTIME_PARAMETERS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_TOOL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_GET_TOOL_RUNTIME_PARAMETERS",
		BufferSize:         1,
		Path:               "/tool/get_runtime_parameters",
	},
	{
		Name:               "InvokeLLM",
		RequestType:        requests.RequestInvokeLLM{},
		ResponseType:       model_entities.LLMResultChunk{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_INVOKE_LLM,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_INVOKE_LLM",
		BufferSize:         512,
		Path:               "/llm/invoke",
	},
	{
		Name:               "GetLLMNumTokens",
		RequestType:        requests.RequestGetLLMNumTokens{},
		ResponseType:       model_entities.LLMGetNumTokensResponse{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_GET_LLM_NUM_TOKENS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_GET_LLM_NUM_TOKENS",
		BufferSize:         1,
		Path:               "/llm/num_tokens",
	},
	{
		Name:               "InvokeTextEmbedding",
		RequestType:        requests.RequestInvokeTextEmbedding{},
		ResponseType:       model_entities.TextEmbeddingResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_INVOKE_TEXT_EMBEDDING,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_INVOKE_TEXT_EMBEDDING",
		BufferSize:         1,
		Path:               "/text_embedding/invoke",
	},
	{
		Name:               "GetTextEmbeddingNumTokens",
		RequestType:        requests.RequestGetTextEmbeddingNumTokens{},
		ResponseType:       model_entities.GetTextEmbeddingNumTokensResponse{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_GET_TEXT_EMBEDDING_NUM_TOKENS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_GET_TEXT_EMBEDDING_NUM_TOKENS",
		BufferSize:         1,
		Path:               "/text_embedding/num_tokens",
	},
	{
		Name:               "InvokeRerank",
		RequestType:        requests.RequestInvokeRerank{},
		ResponseType:       model_entities.RerankResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_INVOKE_RERANK,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_INVOKE_RERANK",
		BufferSize:         1,
		Path:               "/rerank/invoke",
	},
	{
		Name:               "InvokeTTS",
		RequestType:        requests.RequestInvokeTTS{},
		ResponseType:       model_entities.TTSResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_INVOKE_TTS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_INVOKE_TTS",
		BufferSize:         512,
		Path:               "/tts/invoke",
	},
	{
		Name:               "GetTTSModelVoices",
		RequestType:        requests.RequestGetTTSModelVoices{},
		ResponseType:       model_entities.GetTTSVoicesResponse{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_GET_TTS_MODEL_VOICES,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_GET_TTS_MODEL_VOICES",
		BufferSize:         1,
		Path:               "/tts/model/voices",
	},
	{
		Name:               "InvokeSpeech2Text",
		RequestType:        requests.RequestInvokeSpeech2Text{},
		ResponseType:       model_entities.Speech2TextResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_INVOKE_SPEECH2TEXT,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_INVOKE_SPEECH2TEXT",
		BufferSize:         1,
		Path:               "/speech2text/invoke",
	},
	{
		Name:               "InvokeModeration",
		RequestType:        requests.RequestInvokeModeration{},
		ResponseType:       model_entities.ModerationResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_INVOKE_MODERATION,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_INVOKE_MODERATION",
		BufferSize:         1,
		Path:               "/moderation/invoke",
	},
	{
		Name:               "ValidateProviderCredentials",
		RequestType:        requests.RequestValidateProviderCredentials{},
		ResponseType:       model_entities.ValidateCredentialsResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_VALIDATE_PROVIDER_CREDENTIALS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_VALIDATE_PROVIDER_CREDENTIALS",
		BufferSize:         1,
		Path:               "/model/validate_provider_credentials",
	},
	{
		Name:               "ValidateModelCredentials",
		RequestType:        requests.RequestValidateModelCredentials{},
		ResponseType:       model_entities.ValidateCredentialsResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_VALIDATE_MODEL_CREDENTIALS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_VALIDATE_MODEL_CREDENTIALS",
		BufferSize:         1,
		Path:               "/model/validate_model_credentials",
	},
	{
		Name:               "GetAIModelSchema",
		RequestType:        requests.RequestGetAIModelSchema{},
		ResponseType:       model_entities.GetModelSchemasResponse{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_MODEL,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_GET_AI_MODEL_SCHEMAS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_MODEL",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_GET_AI_MODEL_SCHEMAS",
		BufferSize:         1,
		Path:               "/model/schema",
	},
	// { // No need to implement this for now, it has its special implementation in the agent service
	// 	Name:               "InvokeAgentStrategy",
	// 	RequestType:        requests.RequestInvokeAgentStrategy{},
	// 	ResponseType:       agent_entities.AgentStrategyResponseChunk{},
	// 	AccessType:         access_types.PLUGIN_ACCESS_TYPE_AGENT_STRATEGY,
	// 	AccessAction:       access_types.PLUGIN_ACCESS_ACTION_INVOKE_AGENT_STRATEGY,
	// 	AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_AGENT_STRATEGY",
	// 	AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_INVOKE_AGENT_STRATEGY",
	// 	BufferSize:         512,
	// 	Path:               "/agent_strategy/invoke",
	// },
	{
		Name:               "GetAuthorizationURL",
		RequestType:        requests.RequestOAuthGetAuthorizationURL{},
		ResponseType:       oauth_entities.OAuthGetAuthorizationURLResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_OAUTH,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_GET_AUTHORIZATION_URL,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_OAUTH",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_GET_AUTHORIZATION_URL",
		BufferSize:         1,
		Path:               "/oauth/get_authorization_url",
	},
	{
		Name:               "GetCredentials",
		RequestType:        requests.RequestOAuthGetCredentials{},
		ResponseType:       oauth_entities.OAuthGetCredentialsResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_OAUTH,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_GET_CREDENTIALS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_OAUTH",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_GET_CREDENTIALS",
		BufferSize:         1,
		Path:               "/oauth/get_credentials",
	},
	{
		Name:               "RefreshCredentials",
		RequestType:        requests.RequestOAuthRefreshCredentials{},
		ResponseType:       oauth_entities.OAuthRefreshCredentialsResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_OAUTH,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_REFRESH_CREDENTIALS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_OAUTH",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_REFRESH_CREDENTIALS",
		BufferSize:         1,
		Path:               "/oauth/refresh_credentials",
	},
	{
		Name:               "FetchDynamicParameterOptions",
		RequestType:        requests.RequestDynamicParameterSelect{},
		ResponseType:       dynamic_select_entities.DynamicSelectResult{},
		AccessType:         access_types.PLUGIN_ACCESS_TYPE_DYNAMIC_PARAMETER,
		AccessAction:       access_types.PLUGIN_ACCESS_ACTION_DYNAMIC_PARAMETER_FETCH_OPTIONS,
		AccessTypeString:   "access_types.PLUGIN_ACCESS_TYPE_DYNAMIC_SELECT",
		AccessActionString: "access_types.PLUGIN_ACCESS_ACTION_DYNAMIC_PARAMETER_FETCH_OPTIONS",
		BufferSize:         1,
		Path:               "/dynamic_select/fetch_parameter_options",
	},
}
