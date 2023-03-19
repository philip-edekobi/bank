package api

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	mockdb "github.com/philip-edekobi/bank/db/mock"
	db "github.com/philip-edekobi/bank/db/sqlc"
	"github.com/philip-edekobi/bank/util"
	"github.com/stretchr/testify/require"
)

func TestCreateUserAPI(t *testing.T) {
	user := randomUser(t, "balablu")

	testCases := []struct {
		name          string
		buildStubs    func(store *mockdb.MockStore)
		checkResponse func(t *testing.T, recorder *httptest.ResponseRecorder)
	}{
		{
			name: "OK",
			buildStubs: func(store *mockdb.MockStore) {
				hash, _ := util.HashPassword("balablu")
				store.EXPECT().
					CreateUser(gomock.Any(), db.CreateUserParams{}).
					Times(1).
					Return(user, nil)
			},
			checkResponse: func(t *testing.T, recorder *httptest.ResponseRecorder) {
				require.Equal(t, http.StatusOK, recorder.Code)
				requireBodyMatchUser(t, recorder.Body, user)
			},
		},
	}

	for i := range testCases {
		tc := testCases[i]
		t.Run(tc.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			store := mockdb.NewMockStore(ctrl)

			tc.buildStubs(store)

			//start test server and send requests
			server := NewServer(store)
			recorder := httptest.NewRecorder()

			args := bytes.NewReader([]byte(`{"username": "ororo", "full_name": "Mil Keith", "password": "thebigbank", "email": "abbael@email.com"}`))

			request, err := http.NewRequest(http.MethodPost, "/users", args)

			require.NoError(t, err)

			server.router.ServeHTTP(recorder, request)

			tc.checkResponse(t, recorder)
		})
	}
}

func randomUser(t *testing.T, password string) db.User {
	hash, err := util.HashPassword(password)
	if err != nil {
		log.Fatal(err)
	}

	return db.User{
		FullName:       util.RandomOwner(),
		Email:          util.RandomEmail(),
		HashedPassword: hash,
	}
}

func requireBodyMatchUser(t *testing.T, body *bytes.Buffer, user db.User) {
	data, err := io.ReadAll(body)
	require.NoError(t, err)

	var gotUser db.User
	err = json.Unmarshal(data, &gotUser)
	require.NoError(t, err)
	require.Equal(t, gotUser, user)
}
