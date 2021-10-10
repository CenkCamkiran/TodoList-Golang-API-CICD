package mongoDBConnect

import (
	"testing"
)

func TestConnectMongoDB(t *testing.T) {

	t.Run("can connect to the mongodb", func(t *testing.T) {
		_, err := ConnectMongoDB()

		assertMongoDBConnection(t, err)

	})

}

func assertMongoDBConnection(t *testing.T, err error) {
	t.Helper()

	if err != nil {
		t.Errorf("Mongodb connection failed")
	} else {
		t.Logf("Mongodb connection successfull")
	}
}
