package influxdbclient

import "github.com/stretchr/testify/assert"
import "testing"
import "time"


var fields = map[string]interface{}{
            "rsc": 3711,
            "r":   2138,
            "gri": 1908,
            "adg": 912,
}

var fields2 = map[string]interface{}{
            "rst": 3711,
            "r":   2138,
            "gri": 1908,
            "adg": 912,
}

func Test_BadConnect(t *testing.T) {
    testDB := NewInfluxDB("localhost", "8087", "testdb", "admin", "admin")
    err := testDB.Connect()
    assert.NotNil(t, err, "We are expecting error and got one")
}

func Test_GoodConnect(t *testing.T) {
    testDB := NewInfluxDB("localhost", "8086", "testdb", "admin", "admin")
    err := testDB.Connect()
    assert.Nil(t, err, "We are expecting no errors and got one")
}

func Test_CreateDB(t *testing.T) {
    testDB := NewInfluxDB("localhost", "8086", "testdb", "admin", "admin")
    testDB.Connect()
    _, err := testDB.CreateDB("testdb")

    assert.Nil(t, err, "We are expecting no errors and got one")
}

func Test_AddPoint(t *testing.T) {
    testDB := InitSession()
    testDB.AddPoint("test", time.Now(), fields)
    assert.Equal(t, len(testDB.points), 1)
}

func Test_WritePoints(t *testing.T) {
    testDB := InitSession()
    testDB.AddPoint("test", time.Now(), fields)
    testDB.AddPoint("test2", time.Now(), fields)

    err := testDB.WritePoints()
    assert.Nil(t, err, "We are expecting no errors and got one")
}

func Test_ReadPoints(t *testing.T) {
    testDB := InitSession()
    testDB.AddPrecisePoint("test", time.Now(), fields2, "n")
    time.Sleep(100 * time.Millisecond)
    testDB.AddPrecisePoint("test", time.Now(), fields, "n")
    time.Sleep(100 * time.Millisecond)
    testDB.WritePoints()

    time.Sleep(1000 * time.Millisecond)
    _, err := testDB.ReadPoints("test", "*")
    assert.Nil(t, err, "We are expecting error and got one")
}

func Test_DropDB(t *testing.T) {
    testDB := NewInfluxDB("localhost", "8086", "testdb", "admin", "admin")
    testDB.Connect()
    _, err := testDB.DropDB("testdb")

    assert.Nil(t, err, "We are expecting error and got one")
}

func InitSession() (db *InfluxDB) {
   db = NewInfluxDB("localhost", "8086", "testdb", "admin", "admin")
   db.Connect()
   return
}
