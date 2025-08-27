package service

import (
	"github.com/gin-gonic/gin"
	"github.com/langgenius/dify-plugin-daemon/internal/core/plugin_daemon"
	"github.com/langgenius/dify-plugin-daemon/internal/core/plugin_daemon/access_types"
	"github.com/langgenius/dify-plugin-daemon/internal/core/session_manager"
	"github.com/langgenius/dify-plugin-daemon/internal/utils/stream"
	"github.com/langgenius/dify-plugin-daemon/pkg/entities/plugin_entities"
	"github.com/langgenius/dify-plugin-daemon/pkg/entities/requests"
	"github.com/langgenius/dify-plugin-daemon/pkg/entities/tool_entities"
)

func InvokeTool(
	r *plugin_entities.InvokePluginRequest[requests.RequestInvokeTool],
	ctx *gin.Context,
	max_timeout_seconds int,
) {
	baseSSEWithSession(
		func(session *session_manager.Session) (*stream.Stream[tool_entities.ToolResponseChunk], error) {
			return plugin_daemon.InvokeTool(session, &r.Data)
		},
		access_types.PLUGIN_ACCESS_TYPE_TOOL,
		access_types.PLUGIN_ACCESS_ACTION_INVOKE_TOOL,
		r,
		ctx,
		max_timeout_seconds,
	)
}
