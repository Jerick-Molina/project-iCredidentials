package mongodb

import (
	"context"
	"fmt"
	"project/iCredidentials/util"
	"testing"

	"github.com/stretchr/testify/require"
)

var letters_test = "ABCDEFGHIJKLMNOQRZTUVWXYZ"
var numbers_test = "1234567890"
var special_test = "!?$&%#^"

func createRandomAccount(t *testing.T) AccountCreateAccountParams {
	arg := AccountCreateAccountParams{

		FirstName: util.RandomName(),
		LastName:  "testSubject",
		Password:  util.RandomChars(10),
	}
	fullName := fmt.Sprintf("%s%s", arg.FirstName, arg.LastName)
	arg.Email = fmt.Sprintf("%s%d@testSubject.com", fullName, util.RandomNumber(200))
	token, err := testCollections.CreateAccount(context.Background(), arg)

	require.NoError(t, err)

	require.NotEmpty(t, token)

	return arg
}

func TestCreateAccount(t *testing.T) {
	createRandomAccount(t)
}

func TestAccountSignIn(t *testing.T) {
	acc := createRandomAccount(t)
	results, err := testCollections.SignIn(context.Background(), acc.Email, acc.Password)

	require.NoError(t, err)
	require.Equal(t, acc.FirstName, results.FirstName)
	require.Equal(t, results.LastName, results.LastName)
	require.Equal(t, acc.Email, results.Email)
}

func TestDuplicateEmailError(t *testing.T) {
	acc := createRandomAccount(t)
	err := testCollections.EmailDuplicateValidation(context.Background(), acc.Email)
	fmt.Println(err)
	require.Error(t, err)
}
