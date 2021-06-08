/*
 * Copyright © 2018-2019 A Bunch Tell LLC.
 *
 * This file is part of WriteFreely.
 *
 * WriteFreely is free software: you can redistribute it and/or modify
 * it under the terms of the GNU Affero General Public License, included
 * in the LICENSE file in this source code package.
 */

package writefreely

import (
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
	"time"
)

// updatesCacheTime is the default interval between cache updates for new
// software versions
const defaultUpdatesCacheTime = 12 * time.Hour

// updatesCache holds data about current and new releases of the writefreely
// software
type updatesCache struct {
	mu             sync.Mutex
	frequency      time.Duration
	lastCheck      time.Time
	latestVersion  string
	currentVersion string
}

// CheckNow asks for the latest released version of writefreely and updates
// the cache last checked time. If the version postdates the current 'latest'
// the version value is replaced.
func (uc *updatesCache) CheckNow() error {
	uc.mu.Lock()
	defer uc.mu.Unlock()
	latestRemote, err := newVersionCheck()
	if err != nil {
		return err
	}
	uc.lastCheck = time.Now()
	if CompareSemver(latestRemote, uc.latestVersion) == 1 {
		uc.latestVersion = latestRemote
	}
	return nil
}

// AreAvailable updates the cache if the frequency duration has passed
// then returns if the latest release is newer than the current running version.
func (uc updatesCache) AreAvailable() bool {
	if time.Since(uc.lastCheck) > uc.frequency {
		uc.CheckNow()
	}
	return CompareSemver(uc.latestVersion, uc.currentVersion) == 1
}

// LatestVersion returns the latest stored version available.
func (uc updatesCache) LatestVersion() string {
	return uc.latestVersion
}

// ReleaseURL returns the full URL to the blog.writefreely.org release notes
// for the latest version as stored in the cache.
func (uc updatesCache) ReleaseURL() string {
	ver := strings.TrimPrefix(uc.latestVersion, "v")
	ver = strings.TrimSuffix(ver, ".0")
	// hack until go 1.12 in build/travis
	seg := strings.Split(ver, ".")
	return "https://blog.writefreely.org/version-" + strings.Join(seg, "-")
}

// newUpdatesCache returns an initialized updates cache
func newUpdatesCache(expiry time.Duration) *updatesCache {
	cache := updatesCache{
		frequency:      expiry,
		currentVersion: "v" + softwareVer,
	}
	cache.CheckNow()
	return &cache
}

// InitUpdates initializes the updates cache, if the config value is set
// It uses the defaultUpdatesCacheTime for the cache expiry
func (app *App) InitUpdates() {
	if app.cfg.App.UpdateChecks {
		app.updates = newUpdatesCache(defaultUpdatesCacheTime)
	}
}

func newVersionCheck() (string, error) {
	res, err := http.Get("https://version.writefreely.org")
	if err == nil && res.StatusCode == http.StatusOK {
		defer res.Body.Close()

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return "", err
		}
		return string(body), nil
	}
	return "", err
}
