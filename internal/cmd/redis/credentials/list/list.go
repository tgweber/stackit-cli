package list

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/stackitcloud/stackit-cli/internal/pkg/args"
	"github.com/stackitcloud/stackit-cli/internal/pkg/errors"
	"github.com/stackitcloud/stackit-cli/internal/pkg/examples"
	"github.com/stackitcloud/stackit-cli/internal/pkg/flags"
	"github.com/stackitcloud/stackit-cli/internal/pkg/globalflags"
	"github.com/stackitcloud/stackit-cli/internal/pkg/print"
	"github.com/stackitcloud/stackit-cli/internal/pkg/services/redis/client"
	redisUtils "github.com/stackitcloud/stackit-cli/internal/pkg/services/redis/utils"
	"github.com/stackitcloud/stackit-cli/internal/pkg/tables"

	"github.com/spf13/cobra"
	"github.com/stackitcloud/stackit-sdk-go/services/redis"
)

const (
	instanceIdFlag = "instance-id"
	limitFlag      = "limit"
)

type inputModel struct {
	*globalflags.GlobalFlagModel
	InstanceId string
	Limit      *int64
}

func NewCmd(p *print.Printer) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "Lists all credentials' IDs for a Redis instance",
		Long:  "Lists all credentials' IDs for a Redis instance.",
		Args:  args.NoArgs,
		Example: examples.Build(
			examples.NewExample(
				`List all credentials' IDs for a Redis instance`,
				"$ stackit redis credentials list --instance-id xxx"),
			examples.NewExample(
				`List all credentials' IDs for a Redis instance in JSON format`,
				"$ stackit redis credentials list --instance-id xxx --output-format json"),
			examples.NewExample(
				`List up to 10 credentials' IDs for a Redis instance`,
				"$ stackit redis credentials list --instance-id xxx --limit 10"),
		),
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := context.Background()
			model, err := parseInput(cmd)
			if err != nil {
				return err
			}

			// Configure API client
			apiClient, err := client.ConfigureClient(p)
			if err != nil {
				return err
			}

			// Call API
			req := buildRequest(ctx, model, apiClient)
			resp, err := req.Execute()
			if err != nil {
				return fmt.Errorf("list Redis credentials: %w", err)
			}
			credentials := *resp.CredentialsList
			if len(credentials) == 0 {
				instanceLabel, err := redisUtils.GetInstanceName(ctx, apiClient, model.ProjectId, model.InstanceId)
				if err != nil {
					instanceLabel = model.InstanceId
				}
				p.Info("No credentials found for instance %q\n", instanceLabel)
				return nil
			}

			// Truncate output
			if model.Limit != nil && len(credentials) > int(*model.Limit) {
				credentials = credentials[:*model.Limit]
			}
			return outputResult(p, model.OutputFormat, credentials)
		},
	}
	configureFlags(cmd)
	return cmd
}

func configureFlags(cmd *cobra.Command) {
	cmd.Flags().Var(flags.UUIDFlag(), instanceIdFlag, "Instance ID")
	cmd.Flags().Int64(limitFlag, 0, "Maximum number of entries to list")

	err := flags.MarkFlagsRequired(cmd, instanceIdFlag)
	cobra.CheckErr(err)
}

func parseInput(cmd *cobra.Command) (*inputModel, error) {
	globalFlags := globalflags.Parse(cmd)
	if globalFlags.ProjectId == "" {
		return nil, &errors.ProjectIdError{}
	}

	limit := flags.FlagToInt64Pointer(cmd, limitFlag)
	if limit != nil && *limit < 1 {
		return nil, &errors.FlagValidationError{
			Flag:    limitFlag,
			Details: "must be greater than 0",
		}
	}

	return &inputModel{
		GlobalFlagModel: globalFlags,
		InstanceId:      flags.FlagToStringValue(cmd, instanceIdFlag),
		Limit:           limit,
	}, nil
}

func buildRequest(ctx context.Context, model *inputModel, apiClient *redis.APIClient) redis.ApiListCredentialsRequest {
	req := apiClient.ListCredentials(ctx, model.ProjectId, model.InstanceId)
	return req
}

func outputResult(p *print.Printer, outputFormat string, credentials []redis.CredentialsListItem) error {
	switch outputFormat {
	case globalflags.JSONOutputFormat:
		details, err := json.MarshalIndent(credentials, "", "  ")
		if err != nil {
			return fmt.Errorf("marshal Redis credentials list: %w", err)
		}
		p.Outputln(string(details))

		return nil
	default:
		table := tables.NewTable()
		table.SetHeader("ID")
		for i := range credentials {
			c := credentials[i]
			table.AddRow(*c.Id)
		}
		err := table.Display(p)
		if err != nil {
			return fmt.Errorf("render table: %w", err)
		}

		return nil
	}
}
