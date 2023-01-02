package create

import (
	"bytes"
	"context"
	"testing"

	"github.com/MakeNowJust/heredoc"
	"github.com/google/go-cmp/cmp"
	"github.com/stretchr/testify/mock"

	ispb "github.com/input-stream/cli/build/stack/inputstream/v1beta1"
	"github.com/input-stream/cli/mocks"
	"github.com/input-stream/cli/pkg/test"
)

func TestRunGet(t *testing.T) {
	for name, tc := range map[string]struct {
		input *ispb.Input
		want  string
	}{
		"degenerate": {
			want: `{}`,
		},
		"success": {
			input: &ispb.Input{
				Id:     "id1",
				Login:  "octocat",
				Title:  "Cat",
				Status: ispb.Input_STATUS_DRAFT,
			},
			want: heredoc.Doc(`{
			  "login":  "octocat",
			  "id":  "id1",
			  "title":  "Cat",
			  "status":  "STATUS_DRAFT"
			}`),
		},
	} {
		t.Run(name, func(t *testing.T) {
			ctx := context.Background()
			client := mocks.NewInputsClient(t)
			client.
				On("CreateInput", mock.Anything, mock.Anything).
				Return(tc.input, nil)
			cmd := test.GetRootCmdWithSubCommands(createInputCmd())
			err := runCreate(ctx, client, cmd)
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
