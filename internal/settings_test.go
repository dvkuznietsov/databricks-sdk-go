package internal

import (
	"errors"
	"testing"

	"github.com/databricks/databricks-sdk-go/apierr"
	"github.com/databricks/databricks-sdk-go/service/settings"
	"github.com/stretchr/testify/require"
)

func TestAccDefaultNamespaceSetting(t *testing.T) {
	ctx, w := workspaceTest(t)
	_, err := w.Settings.DefaultNamespace().UpdateDefaultNamespaceSetting(ctx, settings.UpdateDefaultNamespaceSettingRequest{})
	require.Error(t, err)
	var apiError *apierr.APIError
	require.True(t, errors.As(err, &apiError))
	etag := apiError.Details[0].Metadata["etag"]
	_, err = w.Settings.DefaultNamespace().UpdateDefaultNamespaceSetting(ctx, settings.UpdateDefaultNamespaceSettingRequest{
		AllowMissing: true,
		FieldMask:    "namespace.value",
		Setting: settings.DefaultNamespaceSetting{
			Etag: etag,
			Namespace: settings.StringMessage{
				Value: "test",
			},
		},
	})
	require.NoError(t, err)
	res, err := w.Settings.DefaultNamespace().GetDefaultNamespaceSetting(ctx, settings.GetDefaultNamespaceSettingRequest{
		Etag: etag,
	})
	require.NoError(t, err)
	require.Equal(t, "test", res.Namespace.Value)
}
