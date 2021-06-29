// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/piperenv"
	"github.com/SAP/jenkins-library/pkg/splunk"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type integrationArtifactGetServiceEndpointOptions struct {
	APIServiceKey     string `json:"apiServiceKey,omitempty"`
	IntegrationFlowID string `json:"integrationFlowId,omitempty"`
	CpiPlatform       string `json:"cpiPlatform,omitempty"`
}

type integrationArtifactGetServiceEndpointCommonPipelineEnvironment struct {
	custom struct {
		iFlowServiceEndpoint string
	}
}

func (p *integrationArtifactGetServiceEndpointCommonPipelineEnvironment) persist(path, resourceName string) {
	content := []struct {
		category string
		name     string
		value    interface{}
	}{
		{category: "custom", name: "iFlowServiceEndpoint", value: p.custom.iFlowServiceEndpoint},
	}

	errCount := 0
	for _, param := range content {
		err := piperenv.SetResourceParameter(path, resourceName, filepath.Join(param.category, param.name), param.value)
		if err != nil {
			log.Entry().WithError(err).Error("Error persisting piper environment.")
			errCount++
		}
	}
	if errCount > 0 {
		log.Entry().Fatal("failed to persist Piper environment")
	}
}

// IntegrationArtifactGetServiceEndpointCommand Get an deployed CPI intgeration flow service endpoint
func IntegrationArtifactGetServiceEndpointCommand() *cobra.Command {
	const STEP_NAME = "integrationArtifactGetServiceEndpoint"

	metadata := integrationArtifactGetServiceEndpointMetadata()
	var stepConfig integrationArtifactGetServiceEndpointOptions
	var startTime time.Time
	var commonPipelineEnvironment integrationArtifactGetServiceEndpointCommonPipelineEnvironment
	var logCollector *log.CollectorHook

	var createIntegrationArtifactGetServiceEndpointCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Get an deployed CPI intgeration flow service endpoint",
		Long:  `With this step you can obtain information about the service endpoints exposed by SAP Cloud Platform Integration on a tenant using OData API. Learn more about the SAP Cloud Integration remote API for getting service endpoint of deployed integration artifact [here](https://help.sap.com/viewer/368c481cd6954bdfa5d0435479fd4eaf/Cloud/en-US/d1679a80543f46509a7329243b595bdb.html).`,
		PreRunE: func(cmd *cobra.Command, _ []string) error {
			startTime = time.Now()
			log.SetStepName(STEP_NAME)
			log.SetVerbose(GeneralConfig.Verbose)

			path, _ := os.Getwd()
			fatalHook := &log.FatalHook{CorrelationID: GeneralConfig.CorrelationID, Path: path}
			log.RegisterHook(fatalHook)

			err := PrepareConfig(cmd, &metadata, STEP_NAME, &stepConfig, config.OpenPiperFile)
			if err != nil {
				log.SetErrorCategory(log.ErrorConfiguration)
				return err
			}
			log.RegisterSecret(stepConfig.APIServiceKey)

			if len(GeneralConfig.HookConfig.SentryConfig.Dsn) > 0 {
				sentryHook := log.NewSentryHook(GeneralConfig.HookConfig.SentryConfig.Dsn, GeneralConfig.CorrelationID)
				log.RegisterHook(&sentryHook)
			}

			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				logCollector = &log.CollectorHook{CorrelationID: GeneralConfig.CorrelationID}
				log.RegisterHook(logCollector)
			}

			return nil
		},
		Run: func(_ *cobra.Command, _ []string) {
			telemetryData := telemetry.CustomData{}
			telemetryData.ErrorCode = "1"
			handler := func() {
				config.RemoveVaultSecretFiles()
				commonPipelineEnvironment.persist(GeneralConfig.EnvRootPath, "commonPipelineEnvironment")
				telemetryData.Duration = fmt.Sprintf("%v", time.Since(startTime).Milliseconds())
				telemetryData.ErrorCategory = log.GetErrorCategory().String()
				telemetry.Send(&telemetryData)
				if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
					splunk.Send(&telemetryData, logCollector)
				}
			}
			log.DeferExitHandler(handler)
			defer handler()
			telemetry.Initialize(GeneralConfig.NoTelemetry, STEP_NAME)
			if len(GeneralConfig.HookConfig.SplunkConfig.Dsn) > 0 {
				splunk.Initialize(GeneralConfig.CorrelationID,
					GeneralConfig.HookConfig.SplunkConfig.Dsn,
					GeneralConfig.HookConfig.SplunkConfig.Token,
					GeneralConfig.HookConfig.SplunkConfig.Index,
					GeneralConfig.HookConfig.SplunkConfig.SendLogs)
			}
			integrationArtifactGetServiceEndpoint(stepConfig, &telemetryData, &commonPipelineEnvironment)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addIntegrationArtifactGetServiceEndpointFlags(createIntegrationArtifactGetServiceEndpointCmd, &stepConfig)
	return createIntegrationArtifactGetServiceEndpointCmd
}

func addIntegrationArtifactGetServiceEndpointFlags(cmd *cobra.Command, stepConfig *integrationArtifactGetServiceEndpointOptions) {
	cmd.Flags().StringVar(&stepConfig.APIServiceKey, "apiServiceKey", os.Getenv("PIPER_apiServiceKey"), "Service key JSON string to access the Cloud Integration API")
	cmd.Flags().StringVar(&stepConfig.IntegrationFlowID, "integrationFlowId", os.Getenv("PIPER_integrationFlowId"), "Specifies the ID of the Integration Flow artifact")
	cmd.Flags().StringVar(&stepConfig.CpiPlatform, "cpiPlatform", `cf`, "Specifies the running platform of the SAP Cloud platform integraion service")

	cmd.MarkFlagRequired("apiServiceKey")
	cmd.MarkFlagRequired("integrationFlowId")
}

// retrieve step metadata
func integrationArtifactGetServiceEndpointMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "integrationArtifactGetServiceEndpoint",
			Aliases:     []config.Alias{},
			Description: "Get an deployed CPI intgeration flow service endpoint",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "cpiApiServiceKeyCredentialsId", Description: "Jenkins credential ID for secret text containing the service key to the SAP Cloud Integration API", Type: "jenkins"},
				},
				Parameters: []config.StepParameters{
					{
						Name: "apiServiceKey",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "cpiApiServiceKeyCredentialsId",
								Param: "apiServiceKey",
								Type:  "secret",
							},
						},
						Scope:     []string{"PARAMETERS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_apiServiceKey"),
					},
					{
						Name:        "integrationFlowId",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "GENERAL", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   true,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_integrationFlowId"),
					},
					{
						Name:        "cpiPlatform",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"GLOBAL", "PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     `cf`,
					},
				},
			},
			Outputs: config.StepOutputs{
				Resources: []config.StepResources{
					{
						Name: "commonPipelineEnvironment",
						Type: "piperEnvironment",
						Parameters: []map[string]interface{}{
							{"Name": "custom/iFlowServiceEndpoint"},
						},
					},
				},
			},
		},
	}
	return theMetaData
}
