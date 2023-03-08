package db

import (
	"context"
	"github.com/philip-edekobi/bank/util"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func createRandomTransfer(t *testing.T, sender, recv Account) Transfer {
	arg := CreateTransferParams{
		FromAccountID: sender.ID,
		ToAccountID:   recv.ID,
		Amount:        util.RandomMoney(),
	}

	transfer, err := testQueries.CreateTransfer(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, transfer)

	require.Equal(t, transfer.FromAccountID, sender.ID)
	require.Equal(t, transfer.ToAccountID, recv.ID)
	require.Equal(t, arg.Amount, transfer.Amount)

	require.NotZero(t, transfer.ID)
	require.NotZero(t, transfer.CreatedAt)

	return transfer
}

func TestCreateTransfer(t *testing.T) {
	sender := createRandomAccount(t)
	receiver := createRandomAccount(t)
	createRandomTransfer(t, sender, receiver)
}

func TestGetTransfer(t *testing.T) {
	sender := createRandomAccount(t)
	receiver := createRandomAccount(t)

	trans1 := createRandomTransfer(t, sender, receiver)
	trans2, err := testQueries.GetTransfer(context.Background(), trans1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, trans2)

	require.Equal(t, trans1.ID, trans2.ID)
	require.Equal(t, trans1.Amount, trans2.Amount)
	require.Equal(t, trans1.FromAccountID, trans2.FromAccountID)
	require.Equal(t, trans1.ToAccountID, trans2.ToAccountID)
	require.WithinDuration(t, trans1.CreatedAt, trans2.CreatedAt, time.Second)
}

func TestListTransfers(t *testing.T) {
	sender := createRandomAccount(t)
	receiver := createRandomAccount(t)

	for i := 0; i < 10; i++ {
		createRandomTransfer(t, sender, receiver)
	}

	arg := ListTransfersParams{
		FromAccountID: sender.ID,
		ToAccountID:   receiver.ID,
		Limit:         5,
		Offset:        5,
	}

	transfers, err := testQueries.ListTransfers(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, transfers, 5)

	for _, transfer := range transfers {
		require.NotEmpty(t, transfer)
	}
}
