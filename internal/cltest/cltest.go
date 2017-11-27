package cltest

import (
	"github.com/asdine/storm"
	"github.com/gin-gonic/gin"
	"github.com/smartcontractkit/chainlink-go/orm"
	"github.com/smartcontractkit/chainlink-go/web"
	"io/ioutil"
	"log"
	"net/http/httptest"
)

var server *httptest.Server

func SetUpDB() *storm.DB {
	orm.InitTest()
	return orm.GetDB()
}

func TearDownDB() {
	orm.Close()
}

func SetUpWeb() *httptest.Server {
	gin.SetMode(gin.TestMode)
	server = httptest.NewServer(web.Router())
	return server
}

func TearDownWeb() {
	gin.SetMode(gin.DebugMode)
	server.Close()
}

func LoadJSON(file string) []byte {
	content, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return content
}