// maestro
// https://github.com/topfreegames/maestro
//
// Licensed under the MIT license:
// http://www.opensource.org/licenses/mit-license
// Copyright © 2017 Top Free Games <backend@tfgco.com>

package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Sirupsen/logrus"
	"github.com/topfreegames/extensions/clock"
	"github.com/topfreegames/extensions/redis"
	"github.com/topfreegames/maestro/controller"
	maestroErrors "github.com/topfreegames/maestro/errors"
	"github.com/topfreegames/maestro/models"
)

// SchedulerCreateHandler handler
type SchedulerCreateHandler struct {
	App *App
}

// NewSchedulerCreateHandler creates a new scheduler create handler
func NewSchedulerCreateHandler(a *App) *SchedulerCreateHandler {
	m := &SchedulerCreateHandler{App: a}
	return m
}

// ServeHTTP method
func (g *SchedulerCreateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l := loggerFromContext(r.Context())
	mr := metricsReporterFromCtx(r.Context())
	payload := configYamlFromCtx(r.Context())

	logger := l.WithFields(logrus.Fields{
		"source":    "schedulerHandler",
		"operation": "create",
	})

	logger.Debug("Creating scheduler...")

	timeoutSec := g.App.Config.GetInt("scaleUpTimeoutSeconds")
	err := mr.WithSegment(models.SegmentController, func() error {
		return controller.CreateScheduler(l, mr, g.App.DB, g.App.RedisClient, g.App.KubernetesClient, payload, timeoutSec)
	})

	if err != nil {
		status := http.StatusInternalServerError
		if strings.Contains(err.Error(), "already exists") {
			status = http.StatusConflict
		} else if err.Error() == "node without label error" {
			status = http.StatusUnprocessableEntity
		}
		logger.WithError(err).Error("Create scheduler failed.")
		g.App.HandleError(w, status, "Create scheduler failed", err)
		return
	}

	mr.WithSegment(models.SegmentSerialization, func() error {
		Write(w, http.StatusCreated, `{"success": true}`)
		return nil
	})
	logger.Debug("Create scheduler succeeded.")
}

// SchedulerDeleteHandler handler
type SchedulerDeleteHandler struct {
	App *App
}

// NewSchedulerDeleteHandler creates a new scheduler delete handler
func NewSchedulerDeleteHandler(a *App) *SchedulerDeleteHandler {
	m := &SchedulerDeleteHandler{App: a}
	return m
}

// ServeHTTP method
func (g *SchedulerDeleteHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l := loggerFromContext(r.Context())
	mr := metricsReporterFromCtx(r.Context())
	params := schedulerParamsFromContext(r.Context())
	logger := l.WithFields(logrus.Fields{
		"source":    "schedulerHandler",
		"operation": "delete",
	})

	logger.Debug("Deleting scheduler...")

	timeoutSec := g.App.Config.GetInt("deleteTimeoutSeconds")
	err := mr.WithSegment(models.SegmentController, func() error {
		return controller.DeleteScheduler(l, mr, g.App.DB, g.App.KubernetesClient, params.SchedulerName, timeoutSec)
	})

	if err != nil {
		logger.WithError(err).Error("Delete scheduler failed.")
		g.App.HandleError(w, http.StatusInternalServerError, "Delete scheduler failed", err)
		return
	}

	mr.WithSegment(models.SegmentSerialization, func() error {
		Write(w, http.StatusOK, `{"success": true}`)
		return nil
	})
	logger.Debug("Delete scheduler succeeded.")
}

// SchedulerUpdateHandler handler
type SchedulerUpdateHandler struct {
	App *App
}

// NewSchedulerUpdateHandler creates a new scheduler delete handler
func NewSchedulerUpdateHandler(a *App) *SchedulerUpdateHandler {
	m := &SchedulerUpdateHandler{App: a}
	return m
}

// ServeHTTP method
func (g *SchedulerUpdateHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l := loggerFromContext(r.Context())
	mr := metricsReporterFromCtx(r.Context())
	params := schedulerParamsFromContext(r.Context())
	payload := configYamlFromCtx(r.Context())
	logger := l.WithFields(logrus.Fields{
		"source":    "schedulerHandler",
		"operation": "update",
	})

	logger.Debugf("Updating scheduler %s", params.SchedulerName)

	if params.SchedulerName != payload.Name {
		msg := fmt.Sprintf("url name %s doesn't match payload name %s", params.SchedulerName, payload.Name)
		err := maestroErrors.NewValidationFailedError(errors.New(msg))
		g.App.HandleError(w, http.StatusBadRequest, "Update scheduler failed", err)
		return
	}

	redisClient, err := redis.NewClient("extensions.redis", g.App.Config, g.App.RedisClient)
	if err != nil {
		logger.WithError(err).Error("error getting redisClient")
		g.App.HandleError(w, http.StatusInternalServerError, "Update scheduler failed", err)
		return
	}

	timeoutSec := g.App.Config.GetInt("updateTimeoutSeconds")
	logger.WithField("time", time.Now()).Info("Starting update")
	err = mr.WithSegment(models.SegmentController, func() error {
		return controller.UpdateSchedulerConfig(
			l,
			mr,
			g.App.DB,
			redisClient,
			g.App.KubernetesClient,
			payload,
			timeoutSec, g.App.Config.GetInt("watcher.lockTimeoutMs"),
			g.App.Config.GetString("watcher.lockKey"),
			&clock.Clock{},
		)
	})
	logger.WithField("time", time.Now()).Info("Finished update")

	if err != nil {
		status := http.StatusInternalServerError
		if _, ok := err.(*maestroErrors.ValidationFailedError); ok {
			status = http.StatusNotFound
		} else if err.Error() == "node without label error" {
			status = http.StatusUnprocessableEntity
		}
		logger.WithError(err).Error("Update scheduler failed.")
		g.App.HandleError(w, status, "Update scheduler failed", err)
		return
	}

	mr.WithSegment(models.SegmentSerialization, func() error {
		Write(w, http.StatusOK, `{"success": true}`)
		return nil
	})
	logger.Debug("Update scheduler succeeded.")
}

// SchedulerListHandler handler
type SchedulerListHandler struct {
	App *App
}

// NewSchedulerListHandler lists schedulers
func NewSchedulerListHandler(a *App) *SchedulerListHandler {
	m := &SchedulerListHandler{App: a}
	return m
}

// ServeHTTP method
func (g *SchedulerListHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l := loggerFromContext(r.Context())
	mr := metricsReporterFromCtx(r.Context())
	logger := l.WithFields(logrus.Fields{
		"source":    "schedulerHandler",
		"operation": "list",
	})
	names, err := controller.ListSchedulersNames(l, mr, g.App.DB)
	if err != nil {
		status := http.StatusInternalServerError
		logger.WithError(err).Error("List scheduler failed.")
		g.App.HandleError(w, status, "List scheduler failed", err)
		return
	}
	resp := map[string]interface{}{
		"schedulers": names,
	}
	bts, err := json.Marshal(resp)
	if err != nil {
		logger.WithError(err).Error("List scheduler failed.")
		g.App.HandleError(w, http.StatusInternalServerError, "List scheduler failed", err)
		return
	}

	mr.WithSegment(models.SegmentSerialization, func() error {
		WriteBytes(w, http.StatusOK, bts)
		return nil
	})
	logger.Debug("List scheduler succeeded.")
}

// SchedulerStatusHandler handler
type SchedulerStatusHandler struct {
	App *App
}

// NewSchedulerStatusHandler get scheduler and its rooms status
func NewSchedulerStatusHandler(a *App) *SchedulerStatusHandler {
	m := &SchedulerStatusHandler{App: a}
	return m
}

// ServeHTTP method
func (g *SchedulerStatusHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	l := loggerFromContext(r.Context())
	mr := metricsReporterFromCtx(r.Context())
	params := schedulerParamsFromContext(r.Context())

	_, getConfig := r.URL.Query()["config"]
	if getConfig {
		g.returnConfig(w, r, l, mr, params.SchedulerName)
		return
	}

	logger := l.WithFields(logrus.Fields{
		"source":    "schedulerHandler",
		"operation": "status",
	})

	logger.Debugf("Getting scheduler %s status", params.SchedulerName)

	scheduler, _, roomCountByStatus, err := controller.GetSchedulerScalingInfo(
		l,
		mr,
		g.App.DB,
		g.App.RedisClient,
		params.SchedulerName,
	)
	if err != nil {
		status := http.StatusInternalServerError
		if _, ok := err.(*maestroErrors.ValidationFailedError); ok {
			status = http.StatusNotFound
		}
		logger.WithError(err).Error("Status scheduler failed.")
		g.App.HandleError(w, status, "Status scheduler failed", err)
		return
	}

	resp := map[string]interface{}{
		"game":               scheduler.Game,
		"state":              scheduler.State,
		"stateLastChangedAt": scheduler.StateLastChangedAt,
		"lastScaleOpAt":      scheduler.LastScaleOpAt,
		"roomsAtCreating":    roomCountByStatus.Creating,
		"roomsAtOccupied":    roomCountByStatus.Occupied,
		"roomsAtReady":       roomCountByStatus.Ready,
		"roomsAtTerminating": roomCountByStatus.Terminating,
	}
	bts, err := json.Marshal(resp)
	if err != nil {
		logger.WithError(err).Error("Status scheduler failed.")
		g.App.HandleError(w, http.StatusInternalServerError, "Status scheduler failed", err)
		return
	}

	mr.WithSegment(models.SegmentSerialization, func() error {
		WriteBytes(w, http.StatusOK, bts)
		return nil
	})
	logger.Debug("Status scheduler succeeded.")
}

func (g *SchedulerStatusHandler) returnConfig(
	w http.ResponseWriter,
	r *http.Request,
	l logrus.FieldLogger,
	mr *models.MixedMetricsReporter,
	schedulerName string,
) {
	logger := l.WithFields(logrus.Fields{
		"source":    "schedulerHandler",
		"operation": "get config",
	})

	logger.Debugf("Getting scheduler %s config", schedulerName)

	var yamlStr string
	err := mr.WithSegment(models.SegmentSelect, func() error {
		var err error
		yamlStr, err = models.LoadConfig(g.App.DB, schedulerName)
		return err
	})
	if err != nil {
		logger.WithError(err).Error("config scheduler failed.")
		g.App.HandleError(w, http.StatusInternalServerError, "config scheduler failed", err)
		return
	}
	if len(yamlStr) == 0 {
		logger.Error("config scheduler not found.")
		g.App.HandleError(w, http.StatusNotFound, "get config error", errors.New("config scheduler not found"))
		return
	}
	resp := map[string]interface{}{
		"yaml": yamlStr,
	}

	bts, err := json.Marshal(resp)
	if err != nil {
		logger.WithError(err).Error("config scheduler failed.")
		g.App.HandleError(w, http.StatusInternalServerError, "config scheduler failed", err)
		return
	}

	mr.WithSegment(models.SegmentSerialization, func() error {
		WriteBytes(w, http.StatusOK, bts)
		return nil
	})
	logger.Debug("config scheduler succeeded.")
}
