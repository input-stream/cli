package list

import (
	"bytes"
	"context"
	"strings"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"

	ispb "github.com/input-stream/cli/build/stack/inputstream/v1beta1"
	"github.com/input-stream/cli/mocks"
	"github.com/input-stream/cli/pkg/test"
)

func TestRunList(t *testing.T) {
	for name, tc := range map[string]struct {
		inputs []*ispb.Input
		want   string
	}{
		"degenerate": {
			want: "ID   TITLE   STATUS\n",
		},
		"success": {
			inputs: []*ispb.Input{
				{
					Id:     "id1",
					Login:  "octocat",
					Title:  "Cat",
					Status: ispb.Input_STATUS_DRAFT,
				},
				{
					Id:     "id2",
					Login:  "octodog",
					Title:  "Dog",
					Status: ispb.Input_STATUS_DRAFT,
				},
			},
			want: strings.Join([]string{
				"ID    TITLE   STATUS",
				"id1   Cat     STATUS_DRAFT",
				"id2   Dog     STATUS_DRAFT",
				"",
			}, "\n"),
		},
	} {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			client := mocks.NewInputsClient(t)
			client.
				On("ListInputs", mock.Anything, mock.Anything).
				Return(&ispb.ListInputsResponse{
					Input: tc.inputs,
				}, nil)
			cmd := test.GetRootCmdWithSubCommands(listInputsCmd())
			err := runList(ctx, client, cmd)
			if err != nil {
				t.Fatal(err)
			}
			stdOut := cmd.OutOrStdout().(*bytes.Buffer).String()
			if diff := cmp.Diff(tc.want, stdOut); diff != "" {
				t.Errorf("(-want,+got): %s", diff)
			}
		})
	}
}
