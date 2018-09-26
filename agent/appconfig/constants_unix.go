// Copyright 2016 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may not
// use this file except in compliance with the License. A copy of the
// License is located at
//
// http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
// either express or implied. See the License for the specific language governing
// permissions and limitations under the License.

// +build freebsd linux netbsd openbsd

// Package appconfig manages the configuration of the agent.
package appconfig

const (

	// PackageRoot specifies the directory under which packages will be downloaded and installed
	PackageRoot = "/data/qvagent/lib/packages"

	// PackageLockRoot specifies the directory under which package lock files will reside
	PackageLockRoot = "/data/qvagent/lib/locks/packages"

	// PackagePlatform is the platform name to use when looking for packages
	PackagePlatform = "linux"

	// DaemonRoot specifies the directory where daemon registration information is stored
	DaemonRoot = "/data/qvagent/lib/daemons"

	// LocalCommandRoot specifies the directory where users can submit command documents offline
	LocalCommandRoot = "/data/qvagent/lib/localcommands"

	// LocalCommandRootSubmitted is the directory where locally submitted command documents
	// are moved when they have been picked up
	LocalCommandRootSubmitted = "/data/qvagent/lib/localcommands/submitted"
	LocalCommandRootCompleted = "/data/qvagent/lib/localcommands/completed"

	// LocalCommandRootInvalid is the directory where locally submitted command documents
	// are moved if the service cannot validate the document (generally impossible via cli)
	LocalCommandRootInvalid = "/data/qvagent/lib/localcommands/invalid"

	// DownloadRoot specifies the directory under which files will be downloaded
	DownloadRoot = "/data/qvagent/log/download/"

	// DefaultDataStorePath represents the directory for storing system data
	DefaultDataStorePath = "/data/qvagent/lib/"

	// EC2ConfigDataStorePath represents the directory for storing ec2 config data
	// EC2ConfigDataStorePath = "/var/lib/amazon/ec2config/"

	// EC2ConfigSettingPath represents the directory for storing ec2 config settings
	// EC2ConfigSettingPath = "/var/lib/amazon/ec2configservice/"

	// UpdaterArtifactsRoot represents the directory for storing update related information
	UpdaterArtifactsRoot = "/data/qvagent/lib/update/"

	// DefaultPluginPath represents the directory for storing plugins in SSM
	DefaultPluginPath = "/data/qvagent/lib/plugins"

	// ManifestCacheDirectory represents the directory for storing all downloaded manifest files
	ManifestCacheDirectory = "/data/qvagent/lib/manifests"

	// List all plugin names, unfortunately golang doesn't support const arrays of strings

	// RebootExitCode that would trigger a Soft Reboot
	RebootExitCode = 194

	// Default Custom Inventory Inventory Folder
	DefaultCustomInventoryFolder = DefaultDataStorePath + "inventory/custom"

	// Used to capture and return exit code for windows powershell script execution - empty for unix shell script case
	ExitCodeTrap = ""

	// PowerShellPluginCommandArgs is the arguments of powershell.exe to be used by the runPowerShellScript plugin
	// PowerShellPluginCommandArgs = ""

	// Exit Code for a command that exits before completion (generally due to timeout or cancel)
	CommandStoppedPreemptivelyExitCode = 137 // Fatal error (128) + signal for SIGKILL (9) = 137

	// RunCommandScriptName is the script name where all downloaded or provided commands will be stored
	RunCommandScriptName = "_qvc.sh"
)

// PowerShellPluginCommandName is the path of the powershell.exe to be used by the runPowerShellScript plugin
var PowerShellPluginCommandName string

// DefaultProgramFolder is the default folder for SSM
var DefaultProgramFolder = "/data/qvagent/etc/"
var DefaultDocumentWorker = "/data/qvagent/bin/qvagent-worker"

// var DefaultSessionWorker = "/data/qvagent/bin/ssm-session-worker"
// var DefaultSessionLogger = "/data/qvagent/bin/ssm-session-logger"

// AppConfigPath is the path of the AppConfig
var AppConfigPath = DefaultProgramFolder + AppConfigFileName

func init() {
	/*
	   Powershell command used to be poweshell in alpha versions, now it's pwsh in prod versions
	*/
	// PowerShellPluginCommandName = "/usr/bin/powershell"
	// if _, err := os.Stat(PowerShellPluginCommandName); err != nil {
	// 	PowerShellPluginCommandName = "/usr/bin/pwsh"
	// }

	// if document-worker is not in the default location, try using the snap installed location
	// if _, err := os.Stat(DefaultDocumentWorker); err != nil {
	// 	if _, err := os.Stat("/snap/amazon-ssm-agent/current/ssm-document-worker"); err == nil {
	// 		DefaultProgramFolder = "/snap/amazon-ssm-agent/current"
	// 		DefaultDocumentWorker = "/snap/amazon-ssm-agent/current/ssm-document-worker"
	// 		DefaultSessionWorker = "/snap/amazon-ssm-agent/current/ssm-session-worker"
	// 		DefaultSessionLogger = "/snap/amazon-ssm-agent/current/ssm-session-logger"
	// 	}
	// }
}
