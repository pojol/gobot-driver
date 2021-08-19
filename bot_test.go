package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"reflect"
	"testing"

	"github.com/pojol/apibot/assert"
	"github.com/pojol/apibot/plugins"
)

var srv *httptest.Server

func TestMain(m *testing.M) {

	srv = httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		parm := GetAccountInfoParam{}

		body, _ := ioutil.ReadAll(req.Body)
		json.Unmarshal(body, &parm)

		w.Write(body)
	}))
	defer srv.Close()

	os.Exit(m.Run())
}

type Metadata struct {
	Val string
}

type GetAccountInfoParam struct {
	Token string
}
type GetAccountPack struct {
}

func (p *GetAccountPack) Marshal(meta interface{}, param interface{}) []byte {

	byt, err := json.Marshal(&param)
	if err != nil {
		fmt.Println(err.Error())
	}

	return byt
}

func (p *GetAccountPack) Unmarshal(meta interface{}, body []byte, header http.Header) {
	mp := meta.(*Metadata)
	mp.Val = string(body)
}

type GetAccountAssert struct {
}

func (p *GetAccountAssert) Do(meta interface{}) error {
	mp := meta.(*Metadata)
	return assert.Equal(mp.Val, "aabbcc", reflect.TypeOf(*p).Name())
}

var compose = `
{
    "id":"e41d33fa-4658-49ef-ab00-db2919de2dbc",
    "ty":"RootNode",
    "pos":{
        "x":0,
        "y":0
    },
    "children":[
        {
            "id":"65c0cb5c-8cf3-427d-a6a2-5ee636627fef",
            "ty":"SelectorNode",
            "pos":{
                "x":-15,
                "y":67
            },
            "children":[
                {
                    "id":"f758e386-c521-4e50-8be7-7300975b354c",
                    "ty":"ConditionNode",
                    "pos":{
                        "x":-20,
                        "y":130
                    },
                    "children":[
                        {
                            "id":"a7609b7b-064d-4a0e-b166-d89f20fcabcd",
                            "ty":"HTTPActionNode",
                            "pos":{
                                "x":-35,
                                "y":192
                            },
                            "children":[

                            ],
                            "api":"/login/guest",
                            "parm":{

                            }
                        }
                    ],
                    "script":{
                        "$eq":{
                            "meta.Token":""
                        }
                    }
                },
                {
                    "id":"24f27096-453f-4035-819b-d5dbfca1471b",
                    "ty":"ConditionNode",
                    "pos":{
                        "x":50,
                        "y":130
                    },
                    "children":[
                        {
                            "id":"ea54ba5d-366c-4248-a12e-2fda71122514",
                            "ty":"HTTPActionNode",
                            "pos":{
                                "x":35,
                                "y":192
                            },
                            "children":[

                            ],
                            "api":"/base/acc.info",
                            "parm":{
                                "Token":"meta.Token"
                            }
                        }
                    ],
                    "script":{
                        "$ne":{
                            "meta.Token":""
                        }
                    }
                }
            ]
        }
    ]
}
`

var structmap map[string]interface{}

type composeInfo struct {
	Behavior string      `json:"behavior"`
	Url      string      `json:"url"`
	Name     string      `json:"name"`
	Script   string      `json:"script"`
	Param    interface{} `json:"param"`
}

func TestBot(t *testing.T) {
	/*
		md := Metadata{}
		b := New("", &md)

		structmap = make(map[string]interface{})
		structmap["GetAccountPack"] = &GetAccountPack{}

		info := &composeInfo{}

		l := lua.NewState()
		defer l.Close()

		l.DoString(` print("hello,lua!") `)

		err := json.Unmarshal([]byte(compose), info)
		if err != nil {
			fmt.Println(err.Error())
			t.FailNow()
		}

		if info.Behavior == "post" {

			pack := reflect.New(reflect.TypeOf(structmap[info.Script]).Elem())

			b.Post(&behavior.HTTPPost{
				URL:   srv.URL,
				Meta:  b.metadata,
				Param: info.Param,
				Api:   pack.Interface(),
			})
		}

		b.Run()
	*/
}

func TestPlugin(t *testing.T) {

	plugins.Load("plugins/jsonmarshal/json_marshal.so")
	m := plugins.Get("jsonmarshal")
	if m != nil {
		m.Marshal("hello")
	}

}
