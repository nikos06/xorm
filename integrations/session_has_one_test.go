package integrations

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetHasOne(t *testing.T) {
	err := PrepareEngine()
	assert.NoError(t, err)

	type Dog struct {
		Id       int64
		MasterId int64
	}

	type Cat struct {
		Id       int64
		MasterId int64
	}

	type Master struct {
		Id    int64
		Name  string
		Dog *Dog `xorm:"has_one"`
		Cat Cat `xorm:"has_one"`
	}

	err = testEngine.Sync2(new(Master), new(Dog), new(Cat))
	assert.NoError(t, err)

	//test get
	master := Master{Id: 1, Name: `master1`}
	_, err = testEngine.InsertOne(&master)
	assert.NoError(t, err)

	dog := Dog{Id: 1, MasterId: master.Id}
	_, err = testEngine.InsertOne(dog)
	assert.NoError(t, err)

	cat := Cat{Id: 1, MasterId: master.Id}
	_, err = testEngine.InsertOne(cat)
	assert.NoError(t, err)

	assertMaster := Master{}
	_, err = testEngine.ID(master.Id).Get(&assertMaster)
	assert.NoError(t, err)
	assert.NotEmpty(t, assertMaster.Dog)
	assert.NotEmpty(t, assertMaster.Cat)
}


func testGetHasMany(t *testing.T) {

}