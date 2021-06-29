// Code generated by piper's step-generator. DO NOT EDIT.

package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/SAP/jenkins-library/pkg/config"
	"github.com/SAP/jenkins-library/pkg/log"
	"github.com/SAP/jenkins-library/pkg/splunk"
	"github.com/SAP/jenkins-library/pkg/telemetry"
	"github.com/spf13/cobra"
)

type integrationArtifactTriggerIntegrationTestOptions struct {
	IFlowServiceKey         string `json:"iFlowServiceKey,omitempty"`
	IntegrationFlowID       string `json:"integrationFlowId,omitempty"`
	CpiPlatform             string `json:"cpiPlatform,omitempty"`
	IFlowServiceEndpointURL string `json:"iFlowServiceEndpointUrl,omitempty"`
	ContentType             string `json:"contentType,omitempty"`
	MessageBodyPath         string `json:"messageBodyPath,omitempty"`
}

// IntegrationArtifactTriggerIntegrationTestCommand Test the service endpoint of your iFlow
func IntegrationArtifactTriggerIntegrationTestCommand() *cobra.Command {
	const STEP_NAME = "integrationArtifactTriggerIntegrationTest"

	metadata := integrationArtifactTriggerIntegrationTestMetadata()
	var stepConfig integrationArtifactTriggerIntegrationTestOptions
	var startTime time.Time
	var logCollector *log.CollectorHook

	var createIntegrationArtifactTriggerIntegrationTestCmd = &cobra.Command{
		Use:   STEP_NAME,
		Short: "Test the service endpoint of your iFlow",
		Long:  `With this step you can test your intergration flow  exposed by SAP Cloud Platform Integration on a tenant using OData API.Learn more about the SAP Cloud Integration remote API for getting service endpoint of deployed integration artifact [here](https://help.sap.com/viewer/368c481cd6954bdfa5d0435479fd4eaf/Cloud/en-US/d1679a80543f46509a7329243b595bdb.html).`,
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
			log.RegisterSecret(stepConfig.IFlowServiceKey)

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
			integrationArtifactTriggerIntegrationTest(stepConfig, &telemetryData)
			telemetryData.ErrorCode = "0"
			log.Entry().Info("SUCCESS")
		},
	}

	addIntegrationArtifactTriggerIntegrationTestFlags(createIntegrationArtifactTriggerIntegrationTestCmd, &stepConfig)
	return createIntegrationArtifactTriggerIntegrationTestCmd
}

func addIntegrationArtifactTriggerIntegrationTestFlags(cmd *cobra.Command, stepConfig *integrationArtifactTriggerIntegrationTestOptions) {
	cmd.Flags().StringVar(&stepConfig.IFlowServiceKey, "iFlowServiceKey", os.Getenv("PIPER_iFlowServiceKey"), "User to authenticate to the SAP Cloud Platform Integration Service")
	cmd.Flags().StringVar(&stepConfig.IntegrationFlowID, "integrationFlowId", os.Getenv("PIPER_integrationFlowId"), "Specifies the ID of the Integration Flow artifact")
	cmd.Flags().StringVar(&stepConfig.CpiPlatform, "cpiPlatform", `cf`, "Specifies the running platform of the SAP Cloud platform integraion service")
	cmd.Flags().StringVar(&stepConfig.IFlowServiceEndpointURL, "iFlowServiceEndpointUrl", os.Getenv("PIPER_iFlowServiceEndpointUrl"), "Specifies the URL endpoint of the iFlow. Please provide in the format `<protocol>://<host>:<port>`. Supported protocols are `http` and `https`.")
	cmd.Flags().StringVar(&stepConfig.ContentType, "contentType", os.Getenv("PIPER_contentType"), "Specifies the content type of the file defined in messageBodyPath e.g. application/json")
	cmd.Flags().StringVar(&stepConfig.MessageBodyPath, "messageBodyPath", os.Getenv("PIPER_messageBodyPath"), "Speficfies the relative file path to the message body.")

	cmd.MarkFlagRequired("iFlowServiceKey")
	cmd.MarkFlagRequired("integrationFlowId")
	cmd.MarkFlagRequired("iFlowServiceEndpointUrl")
}

// retrieve step metadata
func integrationArtifactTriggerIntegrationTestMetadata() config.StepData {
	var theMetaData = config.StepData{
		Metadata: config.StepMetadata{
			Name:        "integrationArtifactTriggerIntegrationTest",
			Aliases:     []config.Alias{},
			Description: "Test the service endpoint of your iFlow",
		},
		Spec: config.StepSpec{
			Inputs: config.StepInputs{
				Secrets: []config.StepSecrets{
					{Name: "iFlowCredentialsId", Description: "Jenkins credentials ID containing username and password for authentication to the SAP Cloud Platform Integration API's", Type: "jenkins"},
				},
				Parameters: []config.StepParameters{
					{
						Name: "iFlowServiceKey",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "iFlowCredentialsId",
								Param: "iFlowServiceKey",
								Type:  "secret",
							},
						},
						Scope:     []string{"PARAMETERS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_iFlowServiceKey"),
					},
					{
						Name:        "integrationFlowId",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS", "GENERAL"},
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
					{
						Name: "iFlowServiceEndpointUrl",
						ResourceRef: []config.ResourceReference{
							{
								Name:  "commonPipelineEnvironment",
								Param: "custom/iFlowServiceEndpoint",
							},
						},
						Scope:     []string{"PARAMETERS"},
						Type:      "string",
						Mandatory: true,
						Aliases:   []config.Alias{},
						Default:   os.Getenv("PIPER_iFlowServiceEndpointUrl"),
					},
					{
						Name:        "contentType",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_contentType"),
					},
					{
						Name:        "messageBodyPath",
						ResourceRef: []config.ResourceReference{},
						Scope:       []string{"PARAMETERS", "STAGES", "STEPS"},
						Type:        "string",
						Mandatory:   false,
						Aliases:     []config.Alias{},
						Default:     os.Getenv("PIPER_messageBodyPath"),
					},
				},
			},
		},
	}
	return theMetaData
}
