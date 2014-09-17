package docs

import (
	"encoding/json"
	"strings"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/swagger"
)

var rootinfo string = `{"apiVersion":"1.0.0","swaggerVersion":"1.2","apis":[{"path":"/object","description":"Operations about object\n"},{"path":"/user","description":"Operations about Users\n"}],"info":{"title":"beego Test API","description":"beego has a very cool tools to autogenerate documents for your API","contact":"astaxie@gmail.com","termsOfServiceUrl":"http://beego.me/","license":"Url http://www.apache.org/licenses/LICENSE-2.0.html"}}`
var subapi string = `{"/object":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/object","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"create","type":"","summary":"create object","parameters":[{"paramType":"body","name":"body","description":"\"The object content\"","dataType":"Object","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Object.Id","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/:objectId","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"find object by objectid","parameters":[{"paramType":"path","name":"objectId","description":"\"the objectid you want to get\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Object","responseModel":"Object"},{"code":403,"message":":objectId is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"GetAll","type":"","summary":"get all objects","responseMessages":[{"code":200,"message":"models.Object","responseModel":"Object"},{"code":403,"message":":objectId is empty","responseModel":""}]}]},{"path":"/:objectId","description":"","operations":[{"httpMethod":"PUT","nickname":"update","type":"","summary":"update the object","parameters":[{"paramType":"path","name":"objectId","description":"\"The objectid you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"The body\"","dataType":"Object","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.Object","responseModel":"Object"},{"code":403,"message":":objectId is empty","responseModel":""}]}]},{"path":"/:objectId","description":"","operations":[{"httpMethod":"DELETE","nickname":"delete","type":"","summary":"delete the object","parameters":[{"paramType":"path","name":"objectId","description":"\"The objectId you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success!","responseModel":""},{"code":403,"message":"objectId is empty","responseModel":""}]}]}],"models":{"Object":{"id":"Object","properties":{"ObjectId":{"type":"string","description":"","format":""},"PlayerName":{"type":"string","description":"","format":""},"Score":{"type":"int64","description":"","format":""}}}}},"/user":{"apiVersion":"1.0.0","swaggerVersion":"1.2","basePath":"","resourcePath":"/user","produces":["application/json","application/xml","text/plain","text/html"],"apis":[{"path":"/","description":"","operations":[{"httpMethod":"POST","nickname":"createUser","type":"","summary":"create users","parameters":[{"paramType":"body","name":"body","description":"\"body for user content\"","dataType":"User","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.User.Id","responseModel":""},{"code":403,"message":"body is empty","responseModel":""}]}]},{"path":"/","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get all Users","responseMessages":[{"code":200,"message":"models.User","responseModel":"User"}]}]},{"path":"/:uid","description":"","operations":[{"httpMethod":"GET","nickname":"Get","type":"","summary":"get user by uid","parameters":[{"paramType":"path","name":"uid","description":"\"The key for staticblock\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.User","responseModel":"User"},{"code":403,"message":":uid is empty","responseModel":""}]}]},{"path":"/:uid","description":"","operations":[{"httpMethod":"PUT","nickname":"update","type":"","summary":"update the user","parameters":[{"paramType":"path","name":"uid","description":"\"The uid you want to update\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"body","name":"body","description":"\"body for user content\"","dataType":"User","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"models.User","responseModel":"User"},{"code":403,"message":":uid is not int","responseModel":""}]}]},{"path":"/:uid","description":"","operations":[{"httpMethod":"DELETE","nickname":"delete","type":"","summary":"delete the user","parameters":[{"paramType":"path","name":"uid","description":"\"The uid you want to delete\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success!","responseModel":""},{"code":403,"message":"uid is empty","responseModel":""}]}]},{"path":"/login","description":"","operations":[{"httpMethod":"GET","nickname":"login","type":"","summary":"Logs user into the system","parameters":[{"paramType":"query","name":"username","description":"\"The username for login\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0},{"paramType":"query","name":"password","description":"\"The password for login\"","dataType":"string","type":"","format":"","allowMultiple":false,"required":true,"minimum":0,"maximum":0}],"responseMessages":[{"code":200,"message":"success","responseModel":""},{"code":403,"message":"user not exist","responseModel":""}]}]},{"path":"/logout","description":"","operations":[{"httpMethod":"GET","nickname":"logout","type":"","summary":"Logs out current logged in user session","responseMessages":[{"code":200,"message":"success","responseModel":""}]}]}],"models":{"Profile":{"id":"Profile","properties":{"Address":{"type":"string","description":"","format":""},"Age":{"type":"int","description":"","format":""},"Email":{"type":"string","description":"","format":""},"Gender":{"type":"string","description":"","format":""}}},"User":{"id":"User","properties":{"Id":{"type":"string","description":"","format":""},"Password":{"type":"string","description":"","format":""},"Profile":{"type":"Profile","description":"","format":""},"Username":{"type":"string","description":"","format":""}}}}}}`
var rootapi swagger.ResourceListing

var apilist map[string]*swagger.ApiDeclaration

func init() {
	basepath := "/v1"
	err := json.Unmarshal([]byte(rootinfo), &rootapi)
	if err != nil {
		beego.Error(err)
	}
	err = json.Unmarshal([]byte(subapi), &apilist)
	if err != nil {
		beego.Error(err)
	}
	beego.GlobalDocApi["Root"] = rootapi
	for k, v := range apilist {
		for i, a := range v.Apis {
			a.Path = urlReplace(k + a.Path)
			v.Apis[i] = a
		}
		v.BasePath = basepath
		beego.GlobalDocApi[strings.Trim(k, "/")] = v
	}
}


func urlReplace(src string) string {
	pt := strings.Split(src, "/")
	for i, p := range pt {
		if len(p) > 0 {
			if p[0] == ':' {
				pt[i] = "{" + p[1:] + "}"
			} else if p[0] == '?' && p[1] == ':' {
				pt[i] = "{" + p[2:] + "}"
			}
		}
	}
	return strings.Join(pt, "/")
}
