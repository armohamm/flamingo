package templatefunctions

import (
	"testing"

	"flamingo.me/flamingo/v3/core/pugtemplate/pugjs"
	"flamingo.me/flamingo/v3/framework/template"

	"github.com/stretchr/testify/assert"
)

func TestJsObject(t *testing.T) {
	var jsObject template.Func = new(JsObject)

	object := jsObject.Func().(func() Object)()

	m := &pugjs.Map{
		Items: make(map[pugjs.Object]pugjs.Object),
	}
	m2 := &pugjs.Map{
		Items: map[pugjs.Object]pugjs.Object{
			pugjs.String("foo"): pugjs.String("bar"),
			pugjs.String("asd"): pugjs.String("dsa"),
		},
	}
	m3 := &pugjs.Map{
		Items: map[pugjs.Object]pugjs.Object{
			pugjs.String("foo"): pugjs.String("bbb"),
		},
	}

	mx := &pugjs.Map{
		Items: map[pugjs.Object]pugjs.Object{
			pugjs.String("foo"): pugjs.String("bbb"),
			pugjs.String("asd"): pugjs.String("dsa"),
		},
	}

	object.Assign(m, m2, m3)
	assert.Equal(t, mx, m)

	arr := object.Keys(mx)
	assert.Equal(t, "asd, foo", arr.Join(", ").String())
	assert.Equal(t, "", object.Keys(nil).Join(", ").String())
}
