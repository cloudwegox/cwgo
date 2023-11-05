/*
*
*  * Copyright 2022 CloudWeGo Authors
*  *
*  * Licensed under the Apache License, Version 2.0 (the "License");
*  * you may not use this file except in compliance with the License.
*  * You may obtain a copy of the License at
*  *
*  *     http://www.apache.org/licenses/LICENSE-2.0
*  *
*  * Unless required by applicable law or agreed to in writing, software
*  * distributed under the License is distributed on an "AS IS" BASIS,
*  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
*  * See the License for the specific language governing permissions and
*  * limitations under the License.
*
 */

package service

import (
	"context"
	"github.com/cloudwego/cwgo/platform/server/cmd/agent/internal/svc"
	agent "github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/agent"
	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/model"
	"net/http"
)

type UpdateRepositoryService struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
} // NewUpdateRepositoryService new UpdateRepositoryService
func NewUpdateRepositoryService(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateRepositoryService {
	return &UpdateRepositoryService{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Run create note info
func (s *UpdateRepositoryService) Run(req *agent.UpdateRepositoryReq) (resp *agent.UpdateRepositoryRes, err error) {
	err = s.svcCtx.DaoManager.Repository.UpdateRepository(s.ctx, model.Repository{
		Id:     req.Id,
		Token:  req.Token,
		Status: req.Status,
	})
	if err != nil {
		return &agent.UpdateRepositoryRes{
			Code: http.StatusBadRequest,
			Msg:  "internal error",
		}, err
	}

	return &agent.UpdateRepositoryRes{
		Code: 0,
		Msg:  "update repository successfully",
	}, nil
}
