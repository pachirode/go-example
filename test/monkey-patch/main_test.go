package main

import (
	"gorm.io/gorm"
	"net/http/httptest"
	"reflect"
	"strings"
	"testing"

	"github.com/agiledragon/gomonkey/v2"
	"github.com/stretchr/testify/assert"
	"github.com/xhd2015/xgo/runtime/mock"
)

func TestUserHandler_CreateUser_by_gomonkey(t *testing.T) {
	mysqlDB := &gorm.DB{}
	handler := NewUserHandler(mysqlDB)
	router := setupRouter(handler)

	patches := gomonkey.ApplyMethod(reflect.TypeOf(mysqlDB), "Create", func(in *gorm.DB, value interface{}) (tx *gorm.DB) {
		expected := &User{Name: "user"}
		actual := value.(*User)
		assert.Equal(t, expected, actual)
		return in
	})
	defer patches.Reset() // 测试完成之后还原

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/users", strings.NewReader(`{"name":"user"}`))
	router.ServeHTTP(w, req)

	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, "", w.Body.String())
}

func TestUserHandler_CreateUser_by_xgo(t *testing.T) {
	mysqlDB := &gorm.DB{}
	handler := NewUserHandler(mysqlDB)
	router := setupRouter(handler)

	// 为 mysqlDB 打上猴子补丁，替换其 Create 方法
	mock.Patch(mysqlDB.Create, func(value interface{}) (tx *gorm.DB) {
		expected := &User{
			Name: "user1",
		}
		actual := value.(*User)
		assert.Equal(t, expected, actual)
		return mysqlDB
	})

	w := httptest.NewRecorder()
	req := httptest.NewRequest("POST", "/users", strings.NewReader(`{"name": "user1"}`))
	router.ServeHTTP(w, req)

	// 断言成功响应
	assert.Equal(t, 201, w.Code)
	assert.Equal(t, "application/json", w.Header().Get("Content-Type"))
	assert.Equal(t, "", w.Body.String())
}
