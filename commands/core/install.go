// This file is part of arduino-cli.
//
// Copyright 2020 ARDUINO SA (http://www.arduino.cc/)
//
// This software is released under the GNU General Public License version 3,
// which covers the main part of arduino-cli.
// The terms of this license can be found at:
// https://www.gnu.org/licenses/gpl-3.0.en.html
//
// You can be released from the requirements of the above licenses by purchasing
// a commercial license. Buying such a license is mandatory if you want to
// modify or otherwise use the software for commercial activities involving the
// Arduino software without disclosing the source code of your own applications.
// To purchase a commercial license, send an email to license@arduino.cc.

package core

import (
	"context"
	"fmt"

	"github.com/arduino/arduino-cli/arduino"
	"github.com/arduino/arduino-cli/arduino/cores/packagemanager"
	"github.com/arduino/arduino-cli/commands"
	rpc "github.com/arduino/arduino-cli/rpc/cc/arduino/cli/commands/v1"
)

// PlatformInstall FIXMEDOC
func PlatformInstall(ctx context.Context, req *rpc.PlatformInstallRequest,
	downloadCB rpc.DownloadProgressCB, taskCB rpc.TaskProgressCB) (*rpc.PlatformInstallResponse, error) {

	pm := commands.GetPackageManager(req.GetInstance().GetId())
	if pm == nil {
		return nil, &arduino.InvalidInstanceError{}
	}

	version, err := commands.ParseVersion(req)
	if err != nil {
		return nil, &arduino.InvalidVersionError{Cause: err}
	}

	ref := &packagemanager.PlatformReference{
		Package:              req.PlatformPackage,
		PlatformArchitecture: req.Architecture,
		PlatformVersion:      version,
	}
	platformRelease, tools, err := pm.FindPlatformReleaseDependencies(ref)
	if err != nil {
		return nil, &arduino.PlatformNotFoundError{Platform: ref.String(), Cause: err}
	}

	// Prerequisite checks before install
	if platformRelease.IsInstalled() {
		taskCB(&rpc.TaskProgress{Name: tr("Platform %s already installed", platformRelease), Completed: true})
		return &rpc.PlatformInstallResponse{}, nil
	}

	if req.GetNoOverwrite() {
		if installed := pm.GetInstalledPlatformRelease(platformRelease.Platform); installed != nil {
			return nil, fmt.Errorf("%s: %s",
				tr("Platform %s already installed", installed),
				tr("could not overwrite"))
		}
	}

	if err := pm.DownloadAndInstallPlatformAndTools(platformRelease, tools, downloadCB, taskCB, req.GetSkipPostInstall()); err != nil {
		return nil, err
	}

	if err := commands.Init(&rpc.InitRequest{Instance: req.Instance}, nil); err != nil {
		return nil, err
	}

	return &rpc.PlatformInstallResponse{}, nil
}
