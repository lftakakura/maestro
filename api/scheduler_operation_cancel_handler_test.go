// maestro
// +build unit
// https://github.com/topfreegames/maestro
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2018 Top Free Games <backend@tfgco.com>

package api_test

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"

	"github.com/golang/mock/gomock"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/topfreegames/maestro/login"
	"github.com/topfreegames/maestro/models"
	. "github.com/topfreegames/maestro/testing"
)

var _ = Describe("SchedulerOperationCancelHandler", func() {
	BeforeEach(func() {
		mockDb.EXPECT().Query(gomock.Any(), `SELECT access_token, refresh_token, expiry, token_type
						FROM users
						WHERE key_access_token = ?`, gomock.Any()).
			Do(func(destToken *login.DestinationToken, query string, modifier string) {
				destToken.RefreshToken = "refresh-token"
			}).AnyTimes()
		mockLogin.EXPECT().Authenticate(gomock.Any(), app.DB).Return("user@example.com", http.StatusOK, nil).AnyTimes()
	})

	Describe("PUT /scheduler/{schedulerName}/operations/{operationKey}/cancel", func() {
		var request *http.Request
		var recorder *httptest.ResponseRecorder
		var opManager *models.OperationManager
		var name = "scheduler-name"

		BeforeEach(func() {
			recorder = httptest.NewRecorder()
			opManager = models.NewOperationManager(name, mockRedisClient, logger)

			url := fmt.Sprintf("http://%s/scheduler/%s/operations/%s/cancel",
				app.Address, name, opManager.GetOperationKey())

			request, _ = http.NewRequest("PUT", url, nil)
		})

		It("should cancel operation", func() {
			MockDeleteRedisKey(opManager, mockRedisClient, mockPipeline, nil)

			app.Router.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(http.StatusOK))

			var response map[string]interface{}
			json.Unmarshal(recorder.Body.Bytes(), &response)
			Expect(response).To(HaveKeyWithValue("success", true))
		})

		It("should return error if redis fails", func() {
			MockDeleteRedisKey(opManager, mockRedisClient, mockPipeline, errors.New("redis error"))

			app.Router.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(http.StatusInternalServerError))

			var response map[string]interface{}
			json.Unmarshal(recorder.Body.Bytes(), &response)
			Expect(response).To(HaveKeyWithValue("success", false))
			Expect(response).To(HaveKeyWithValue("error", "error deleting operation key on redis"))
			Expect(response).To(HaveKeyWithValue("description", "redis error"))
			Expect(response).To(HaveKeyWithValue("code", "MAE-000"))
		})

		It("should return error if operation key is invalid", func() {
			key := "invalid key"
			url := fmt.Sprintf("http://%s/scheduler/%s/operations/%s/cancel",
				app.Address, name, key)

			request, _ = http.NewRequest("PUT", url, nil)
			app.Router.ServeHTTP(recorder, request)
			Expect(recorder.Code).To(Equal(http.StatusInternalServerError))

			var response map[string]interface{}
			json.Unmarshal(recorder.Body.Bytes(), &response)
			Expect(response).To(HaveKeyWithValue("success", false))
			Expect(response).To(HaveKeyWithValue("error", "error deleting operation key on redis"))
			Expect(response).To(HaveKeyWithValue("description", "operationKey is not valid: invalid key"))
			Expect(response).To(HaveKeyWithValue("code", "MAE-000"))
		})
	})
})
