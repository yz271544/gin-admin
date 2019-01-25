package test

import (
	"net/http/httptest"
	"testing"

	"github.com/LyricTian/gin-admin/src/schema"
	"github.com/LyricTian/gin-admin/src/util"
	"github.com/stretchr/testify/assert"
)

func TestDemo(t *testing.T) {
	const router = "v1/demos"
	var err error

	w := httptest.NewRecorder()

	// post /demos
	addItem := &schema.Demo{
		Code:   util.MustUUID(),
		Name:   "测试用例",
		Status: 1,
	}
	engine.ServeHTTP(w, newPostRequest(router, addItem))
	assert.Equal(t, 200, w.Code)

	var addNewItem schema.Demo
	err = parseReader(w.Body, &addNewItem)
	assert.Nil(t, err)
	assert.NotEmpty(t, addNewItem.RecordID)

	// get /demos/:id
	engine.ServeHTTP(w, newGetRequest("%s/%s", nil, router, addNewItem.RecordID))
	assert.Equal(t, 200, w.Code)

	var addGetItem schema.Demo
	err = parseReader(w.Body, &addGetItem)

	assert.Nil(t, err)
	assert.Equal(t, addGetItem.Code, addItem.Code)
	assert.Equal(t, addGetItem.Name, addItem.Name)
	assert.Equal(t, addGetItem.Status, addItem.Status)

	// get /demos?q=page
	engine.ServeHTTP(w, newGetRequest(router,
		newPageParam(map[string]string{"q": "page"})))
	assert.Equal(t, 200, w.Code)
	var pageItems []*schema.Demo
	err = parsePageReader(w.Body, &pageItems)
	assert.Nil(t, err)
	assert.Equal(t, len(pageItems), 1)
	assert.Equal(t, pageItems[0].RecordID, addNewItem.RecordID)

	// get /demos/:id
	engine.ServeHTTP(w, newGetRequest("%s/%s", nil, router, addNewItem.RecordID))
	assert.Equal(t, 200, w.Code)
	var putItem schema.Demo
	err = parseReader(w.Body, &putItem)
	// put /demos/:id
	putItem.Code = util.MustUUID()
	putItem.Name = "测试用例2"
	engine.ServeHTTP(w, newPutRequest("%s/%s", putItem, router, addNewItem.RecordID))
	assert.Equal(t, 200, w.Code)
	releaseReader(w.Body)

	// get /demos/:id
	engine.ServeHTTP(w, newGetRequest("%s/%s", nil, router, addNewItem.RecordID))
	assert.Equal(t, 200, w.Code)

	var getItem schema.Demo
	err = parseReader(w.Body, &getItem)

	assert.Nil(t, err)
	assert.Equal(t, getItem.RecordID, addNewItem.RecordID)
	assert.Equal(t, getItem.Code, putItem.Code)
	assert.Equal(t, getItem.Name, putItem.Name)

	// delete /demos/:id
	engine.ServeHTTP(w, newDeleteRequest("%s/%s", router, addNewItem.RecordID))
	assert.Equal(t, 200, w.Code)
	releaseReader(w.Body)
}
