package config_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/abemedia/config"
)

type example struct {
	Foo string
	Bar struct {
		Greeting string
		Planet   string
	}
}

func TestNew(t *testing.T) {
	_, err := config.New(example{})
	assert.NotNil(t, err, "should fail on non-pointer")
	c := "Invalid config var"
	_, err = config.New(&c)
	assert.NotNil(t, err, "should fail on non-struct")
}

func TestReadFile(t *testing.T) {
	assert := assert.New(t)

	for _, format := range []string{"json", "yml"} {
		c, err := config.New(&example{})
		assert.Nil(err)

		err = c.ReadFile("./testdata/config." + format)
		assert.Nil(err)

		val, err := c.Get("Bar.Greeting")
		assert.Nil(err)
		assert.Equal(val, "Hello")
	}
}

func TestSetGet(t *testing.T) {
	assert := assert.New(t)

	c, err := config.New(&example{})
	assert.Nil(err)

	// set value
	err = c.Set("Bar.Greeting", "Howdy")
	assert.Nil(err)

	// get existing value
	val, err := c.Get("Bar.Greeting")
	assert.Nil(err)
	assert.Equal(val, "Howdy")

	// get non-existant field
	_, err = c.Get("Bar.NotHereReally")
	assert.NotNil(err, "should throw error: field not found")

	// get child of non-struct
	_, err = c.Get("Foo.Bar")
	assert.NotNil(err, "should throw error: field not found")
}
