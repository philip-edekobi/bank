package db

import (
	"context"
	"testing"
	"time"

	"github.com/philip-edekobi/bank/util"

	"github.com/stretchr/testify/require"
)

func createRandomEntry(t *testing.T, acc Account) Entry {
	arg := CreateEntryParams{
		AccountID: acc.ID,
		Amount:    util.RandomMoney(),
	}

	entry, err := testQueries.CreateEntry(context.Background(), arg)

	require.NoError(t, err)
	require.NotEmpty(t, entry)

	require.Equal(t, entry.AccountID, acc.ID)
	require.Equal(t, entry.AccountID, arg.AccountID)
	require.Equal(t, entry.Amount, arg.Amount)

	require.NotZero(t, entry.ID)
	require.NotZero(t, entry.CreatedAt)

	return entry
}

func TestCreateEntry(t *testing.T) {
	acc := createRandomAccount(t)
	createRandomEntry(t, acc)
}

func TestGetEntry(t *testing.T) {
	acc := createRandomAccount(t)
	ent1 := createRandomEntry(t, acc)

	ent2, err := testQueries.GetEntry(context.Background(), ent1.ID)

	require.NoError(t, err)
	require.NotEmpty(t, ent2)

	require.Equal(t, ent1.ID, ent2.ID)
	require.Equal(t, ent1.AccountID, ent2.AccountID)
	require.Equal(t, ent1.Amount, ent2.Amount)
	require.WithinDuration(t, ent1.CreatedAt, ent2.CreatedAt, time.Second)
}

func TestListEntries(t *testing.T) {
	acc := createRandomAccount(t)
	for i := 1; i <= 10; i++ {
		createRandomEntry(t, acc)
	}

	arg := ListEntriesParams{
		Limit:     5,
		AccountID: acc.ID,
		Offset:    5,
	}

	entries, err := testQueries.ListEntries(context.Background(), arg)

	require.NoError(t, err)
	require.Len(t, entries, 5)

	for _, entry := range entries {
		require.NotEmpty(t, entry)
	}
}
