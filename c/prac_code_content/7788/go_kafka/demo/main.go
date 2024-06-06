/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-06-04 14:25:19
 * @LastEditTime: 2024-06-06 08:42:46
 * @FilePath: /let-sGo/c/prac_code_content/7788/go_kafka/demo/main.go
 * @description: type some description
 */
/*
Copyright 2019 The Knative Authors

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	cloudEvents "github.com/cloudevents/sdk-go/v2"
	"github.com/google/uuid"
	"google.golang.org/grpc/metadata"
)

/*
Example Output:

☁️  cloudevents.Event
Validation: valid
Context Attributes,
  specversion: 1.0
  type: dev.knative.eventing.samples.heartbeat
  source: https://knative.dev/eventing-contrib/cmd/heartbeats/#event-test/mypod
  id: 2b72d7bf-c38f-4a98-a433-608fbcdd2596
  time: 2019-10-18T15:23:20.809775386Z
  contenttype: application/json
Extensions,
  beats: true
  heart: yes
  the: 42
Data,
  {
    "id": 2,
    "label": ""
  }
*/

func NewRandomRequestId() string {
	u, _ := uuid.NewRandom()
	return u.String()
}

type GroupUser struct {
	GroupUserId uint64 `protobuf:"varint,1,opt,name=group_user_id,json=groupUserId,proto3" json:"group_user_id,omitempty" gorm:"primaryKey"`
	GroupId     uint64 `protobuf:"varint,10,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty" gorm:"INDEX:group_id"`
	Name        string `protobuf:"bytes,20,opt,name=name,proto3" json:"name,omitempty"` // 用户名字，Group 内部唯一
}

type Group struct {
	GroupId      uint64             `protobuf:"varint,1,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty" gorm:"primaryKey"`
	BusinessType Group_BusinessType `protobuf:"varint,19,opt,name=business_type,json=businessType,proto3,enum=ceres.enterprise.Group_BusinessType" json:"business_type,omitempty" gorm:"default:1"` //判断企业业务类型
}

type UserInfo struct {
	ClientId  uint64     `protobuf:"varint,1,opt,name=client_id,json=clientId,proto3" json:"client_id,omitempty" gorm:"primaryKey"` // 查看文件顶部的注释获取更多关于 client_id 的信息
	AccountId uint64     `protobuf:"varint,2,opt,name=account_id,json=accountId,proto3" json:"account_id,omitempty"`
	GroupId   uint64     `protobuf:"varint,3,opt,name=group_id,json=groupId,proto3" json:"group_id,omitempty"`
	StationId uint64     `protobuf:"varint,4,opt,name=station_id,json=stationId,proto3" json:"station_id,omitempty"` //当前station_id，station_id=0表示集团角色，前提是已企业集团化
	Group     *Group     `protobuf:"bytes,100,opt,name=group,proto3" json:"group,omitempty"`
	GroupUser *GroupUser `protobuf:"bytes,101,opt,name=group_user,json=groupUser,proto3" json:"group_user,omitempty"`
	EventType string     `json:"type"`
	Source    string     `json:"source"`
}

type Group_BusinessType int32

const (
	Group_UNSPECIFIED         Group_BusinessType = 0
	Group_STANDARD            Group_BusinessType = 1     // 标准版
	Group_LITE                Group_BusinessType = 2     // 轻巧版
	Group_PRO                 Group_BusinessType = 3     // 专业版
	Group_SIMPLE_LITE         Group_BusinessType = 4     // 简易轻巧版
	Group_CENTRAL_KITCHEN_STD Group_BusinessType = 5     // 央厨标准版
	Group_CENTRAL_KITCHEN_PRO Group_BusinessType = 6     // 央厨旗舰版
	Group_LITE_FREE           Group_BusinessType = 7     // 轻巧版免费版
	Group_LITE_PURCHASER      Group_BusinessType = 1000  // 菜小蜜采购小程序，已废弃
	Group_LITE_SUPPLIER       Group_BusinessType = 10001 // 菜小蜜收单小程序，已废弃
	Group_LITE_COLLABORATION  Group_BusinessType = 10002 // 供应商协同
)

func NewMDPairsWithUserInfo(md metadata.MD, userInfo *UserInfo) ([]string, error) {
	return newMDPairsWithUserInfo(md, userInfo)
}

func newMDPairsWithUserInfo(md metadata.MD, userInfo *UserInfo) ([]string, error) {
	if userInfo == nil {
		return nil, errors.New("grpcerrors.GRPCError(commonproto.Status_INTERNAL)")
	}
	data, err := json.Marshal(userInfo)
	if err != nil {
		return nil, err
	}
	accountId := ""
	if userInfo.AccountId > 0 {
		accountId = fmt.Sprint(userInfo.AccountId)
	}

	res := []string{
		"x-user-info", base64.StdEncoding.EncodeToString(data),
		"x-group-id", fmt.Sprintf("%d", userInfo.GroupId),
		"x-station-id", fmt.Sprintf("%d", userInfo.StationId),
		"x-account-id", accountId,
	}
	requestIds := md.Get("x-request-id")
	if len(requestIds) == 0 {
		res = append(res, []string{"x-request-id", NewRandomRequestId()}...)
	} else {
		for _, requestId := range requestIds {
			res = append(res, "x-request-id", requestId)
		}
	}
	// apiVersions := md.Get("x-api-version")
	// if len(apiVersions) == 0 {
	// 	_, version := GetPodVersion()
	// 	if version != "" {
	// 		res = append(res, "x-api-version", version)
	// 	}
	// } else {
	// 	res = append(res, "x-api-version", apiVersions[0])
	// }
	res = append(res, "x-api-version", "api-version-xxxx")

	return res, nil
}

// Enum value maps for Group_BusinessType.
var (
	Group_BusinessType_name = map[int32]string{
		0:     "UNSPECIFIED",
		1:     "STANDARD",
		2:     "LITE",
		3:     "PRO",
		4:     "SIMPLE_LITE",
		5:     "CENTRAL_KITCHEN_STD",
		6:     "CENTRAL_KITCHEN_PRO",
		7:     "LITE_FREE",
		1000:  "LITE_PURCHASER",
		10001: "LITE_SUPPLIER",
		10002: "LITE_COLLABORATION",
	}
	Group_BusinessType_value = map[string]int32{
		"UNSPECIFIED":         0,
		"STANDARD":            1,
		"LITE":                2,
		"PRO":                 3,
		"SIMPLE_LITE":         4,
		"CENTRAL_KITCHEN_STD": 5,
		"CENTRAL_KITCHEN_PRO": 6,
		"LITE_FREE":           7,
		"LITE_PURCHASER":      1000,
		"LITE_SUPPLIER":       10001,
		"LITE_COLLABORATION":  10002,
	}
)

// UpdateIncomingContextWithMap 从一个 map 中获取上下游需传递的 metadata 后更新 ctx
func UpdateIncomingContextWithMap(ctx context.Context, aMap map[string]any) (context.Context, error) {
	// var groupId, stationId, groupUserId, clientId uint64
	// var requestIds, groupUserName, podVersion, eventType, eventSource string
	var eventType, eventSource string

	if typeStr, ok := aMap["type"]; ok {
		eventType = typeStr.(string)
	}

	if sourceStr, ok := aMap["source"]; ok {
		eventSource = sourceStr.(string)
	}

	var res []string
	res = append(res, "source", eventSource)
	res = append(res, "type", eventType)
	ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(res...))

	// userInfo := &UserInfo{
	// 	EventType: eventType,
	// 	Source:    eventSource,
	// }
	// metaPairs, err := NewMDPairsWithUserInfo(metadata.Pairs(res...), userInfo)
	// if err != nil {
	// 	return nil, err
	// }
	// ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(metaPairs...))

	// if clientIdStr, ok := aMap["clientid"]; ok {
	// 	clientId, _ = strconv.ParseUint(clientIdStr.(string), 10, 64)
	// }
	// if groupIdStr, ok := aMap["groupid"]; ok {
	// 	groupId, _ = strconv.ParseUint(groupIdStr.(string), 10, 64)
	// }
	// if groupUserIdStr, ok := aMap["userid"]; ok {
	// 	groupUserId, _ = strconv.ParseUint(groupUserIdStr.(string), 10, 64)
	// }
	// if username, ok := aMap["username"]; ok {
	// 	if bytes, err := hex.DecodeString(username.(string)); err == nil {
	// 		groupUserName = string(bytes)
	// 	}
	// }
	// if podVersionStr, ok := aMap["apiversion"]; ok {
	// 	podVersion = podVersionStr.(string)
	// }

	// if stationIdStr, ok := aMap["stationid"]; ok {
	// 	stationId, _ = strconv.ParseUint(stationIdStr.(string), 10, 64)
	// }

	// if requestIdStr, ok := aMap["requestid"]; ok {
	// 	requestIds = requestIdStr.(string)
	// }

	// if businessTypeStr, ok := aMap["businesstype"]; ok {
	// 	businessTypeInt, _ := strconv.ParseInt(businessTypeStr.(string), 10, 32)
	// 	businessType = Group_BusinessType(businessTypeInt)
	// }

	// if semesterIdStr, ok := aMap["semesterid"]; ok {
	// 	semesterId, _ = strconv.ParseUint(semesterIdStr.(string), 10, 64)
	// }

	// if groupId > 0 {
	// 	if requestIds != "" { // 新事件一定有
	// 		var res []string
	// 		for _, requestId := range strings.Split(requestIds, ",") {
	// 			res = append(res, "x-request-id", requestId)
	// 		}
	// 		if podVersion != "" {
	// 			res = append(res, "x-api-version", podVersion)
	// 		}
	// 		userInfo := &UserInfo{
	// 			ClientId:  clientId,
	// 			GroupId:   groupId,
	// 			StationId: stationId,
	// 			GroupUser: &GroupUser{
	// 				GroupUserId: groupUserId,
	// 				GroupId:     groupId,
	// 				Name:        groupUserName,
	// 			},
	// 			Group: &Group{
	// 				GroupId:      groupId,
	// 				BusinessType: businessType,
	// 			},
	// 			EventType: eventType,
	// 			Source:    eventSource,
	// 		}
	// 		metaPairs, err := NewMDPairsWithUserInfo(metadata.Pairs(res...), userInfo)
	// 		if err != nil {
	// 			return nil, err
	// 		}
	// 		ctx = metadata.NewIncomingContext(ctx, metadata.Pairs(metaPairs...))
	// 	}
	// }
	// if semesterId > 0 {
	// 	ctx = UpdateIncomingContextWithSemesterIdForUpdate(ctx, semesterId)
	// }

	return ctx, nil
}

func UpdateIncomingContextWithEvent(ctx context.Context, event cloudEvents.Event) (context.Context, error) {
	ctx, err := UpdateIncomingContextWithMap(ctx, event.Extensions())
	if err != nil {
		return nil, err
	}
	return ctx, nil
}

func extractEventExtensions(ctx context.Context, event cloudEvents.Event) (context.Context, error) {
	var eventType, eventSource string

	typeStr, e := event.Context.GetExtension("type")
	if e != nil {
		return nil, e
	}
	eventType = typeStr.(string)

	sourceStr, e2 := event.Context.GetExtension("source")
	if e2 != nil {
		return nil, e2
	}
	eventSource = sourceStr.(string)

	var res []string
	res = append(res, "source", eventSource)
	res = append(res, "type", eventType)
	return metadata.NewIncomingContext(ctx, metadata.Pairs(res...)), nil
}

func display(event cloudEvents.Event) {

	fmt.Printf("☁️  cloudevents.Event =========== start ===========\n")

	// ctx, e1 := UpdateIncomingContextWithEvent(context.Background(), event)
	// if e1 != nil {
	// 	fmt.Println("err =========> ", e1.Error())
	// } else {
	// 	// 打印 source 和 type
	// 	a, b := metadata.FromIncomingContext(ctx)
	// 	if b {
	// 		type1 := a.Get("type")
	// 		source := a.Get("source")
	// 		fmt.Println("取出成功-type: ", type1)
	// 		fmt.Println("取出成功-source: ", source)
	// 	} else {
	// 		fmt.Println("取出失败")
	// 	}
	// }
	// ctx, e1 := extractEventExtensions(context.Background(), event)
	// if e1 != nil {
	// 	fmt.Println("err =========> ", e1.Error())
	// } else {
	// 	// 打印 source 和 type
	// 	a, b := metadata.FromIncomingContext(ctx)
	// 	if b {
	// 		type1 := a.Get("type")
	// 		source := a.Get("source")
	// 		fmt.Println("取出成功-type: ", type1)
	// 		fmt.Println("取出成功-source: ", source)
	// 	} else {
	// 		fmt.Println("取出失败")
	// 	}
	// }

	// fmt.Println(event.Extensions())
	fmt.Println(" >>>>>>>>>>>>>>>>>>>>>>> ", event.Context.GetType())
	fmt.Println(" >>>>>>>>>>>>>>>>>>>>>>> ", event.Context.GetSource())
	// fmt.Println(" >>>>>>>>>>>>>>>>>>>>>>> ", event.Extensions()["type"])
	var m MsgT
	e1 := json.Unmarshal(event.Data(), &m)
	if e1 != nil {
		fmt.Println("error->", e1)
	} else {
		fmt.Println(" msg data is ====>", m)
	}

	fmt.Printf("%s --------- ---- end ---- -------- \n", event.String())
}

type MsgT struct {
	Msg string `json:"msg"`
}

func main() {
	options := []cloudEvents.HTTPOption{
		cloudEvents.WithPort(8080),
		cloudEvents.WithPath("/msg/receive"),
	}
	c, err := cloudEvents.NewClientHTTP(options...)
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	log.Fatal(c.StartReceiver(context.Background(), display))
}
