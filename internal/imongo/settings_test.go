package imongo

import (
	"context"
	"projects/iCredidentials/util"
	"testing"

	"github.com/stretchr/testify/require"
)

func createRandomSettings(t *testing.T) SetSettingsParams {
	setBool := util.RandomBool()
	args := SetSettingsParams{
		SettingCode:     util.RandomChars(10),
		OwnerId:         "Default_Testing",
		PassSpecialChar: util.RandomBool(),
		PassNumbers:     util.RandomBool(),
		PassMinLength:   int(util.RandomNumber(7)),
		UserMinLength:   int(util.RandomNumber(7)),
		FindUsername:    setBool,
		TwoFactor:       util.RandomBool(),
		EmailConfirm:    setBool,
	}

	err := testCollections.SetUniqueSettings(context.Background(), args)

	require.NoError(t, err)
	return args
}

func TestSetUniqueSettings(t *testing.T) {
	createRandomSettings(t)
}
func TestDefaultSettings(t *testing.T) {

	results, err := testCollections.GetSettingsDefault(context.Background())

	require.NoError(t, err)
	require.NotEmpty(t, results)
}

func TestGetUniqueSettings(t *testing.T) {
	set := createRandomSettings(t)

	results, err := testCollections.GetSettingsUnique(context.Background(), set.SettingCode)

	require.NoError(t, err)

	require.Equal(t, set.OwnerId, results.OwnerId)
	require.Equal(t, set.FindUsername, results.FindUsername)
	require.Equal(t, set.PassMinLength, results.PassMinLength)
}
