/*
 * @Author: 27
 * @LastEditors: 27
 * @Date: 2024-06-04 14:25:19
 * @LastEditTime: 2024-06-04 14:29:56
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
	"fmt"
	"log"

	cloudevents "github.com/cloudevents/sdk-go/v2"
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

func display(event cloudevents.Event) {
	fmt.Printf("☁️  cloudevents.Event\n --------- start -------- \n%s", event.String())
	fmt.Printf("=========== end ===========")
}

func main() {
	c, err := cloudevents.NewClientHTTP()
	if err != nil {
		log.Fatal("Failed to create client, ", err)
	}

	log.Fatal(c.StartReceiver(context.Background(), display))
}
