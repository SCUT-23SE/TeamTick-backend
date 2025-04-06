package handlers

import (
	"TeamTickBackend/gen"

	"github.com/gin-gonic/gin"
)

type GroupHandler struct {
}

func NewGroupHandler() *GroupHandler {
	return &GroupHandler{}
}

// DeleteGroupsGroupId 删除群组
// (DELETE /groups/{groupId})
func (h *GroupHandler) DeleteGroupsGroupId(c *gin.Context, groupId int64) {
}

// DeleteGroupsGroupIdMembersUserId 删除群组成员
// (DELETE /groups/{groupId}/members/{userId})
func (h *GroupHandler) DeleteGroupsGroupIdMembersUserId(c *gin.Context, groupId int64, userId int64) {
}

// GetGroups 获取群组列表
// (GET /groups)
func (h *GroupHandler) GetGroups(c *gin.Context, params gen.GetGroupsParams) {
}

// GetGroupsGroupId 获取群组详情
// (GET /groups/{groupId})
func (h *GroupHandler) GetGroupsGroupId(c *gin.Context, groupId int64) {
}

// GetGroupsGroupIdJoinRequests 获取群组加入请求列表
// (GET /groups/{groupId}/join-requests)
func (h *GroupHandler) GetGroupsGroupIdJoinRequests(c *gin.Context, groupId int64, params gen.GetGroupsGroupIdJoinRequestsParams) {
}

// GetGroupsGroupIdMembers 获取群组成员列表
// (GET /groups/{groupId}/members)
func (h *GroupHandler) GetGroupsGroupIdMembers(c *gin.Context, groupId int64) {
}

// PostGroups 创建群组
// (POST /groups)
func (h *GroupHandler) PostGroups(c *gin.Context) {
}

// PostGroupsGroupIdJoinRequests 申请加入群组
// (POST /groups/{groupId}/join-requests)
func (h *GroupHandler) PostGroupsGroupIdJoinRequests(c *gin.Context, groupId int64) {
}

// PutGroupsGroupId 更新群组
// (PUT /groups/{groupId})
func (h *GroupHandler) PutGroupsGroupId(c *gin.Context, groupId int64) {
}

// PutGroupsGroupIdJoinRequestsRequestId 处理加入请求
// (PUT /groups/{groupId}/join-requests/{requestId})
func (h *GroupHandler) PutGroupsGroupIdJoinRequestsRequestId(c *gin.Context, groupId int64, requestId int64) {
}
