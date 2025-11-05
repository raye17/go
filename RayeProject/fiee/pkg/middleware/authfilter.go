/*
 * Licensed to the Apache Software Foundation (ASF) under one or more
 * contributor license agreements.  See the NOTICE file distributed with
 * this work for additional information regarding copyright ownership.
 * The ASF licenses this file to You under the Apache License, Version 2.0
 * (the "License"); you may not use this file except in compliance with
 * the License.  You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package middleware

import (
	"context"
	"fmt"

	"dubbo.apache.org/dubbo-go/v3/common/extension"
	"dubbo.apache.org/dubbo-go/v3/filter"
	"dubbo.apache.org/dubbo-go/v3/protocol"
)

func init() {
	extension.SetFilter("authFilter", NewAuthFilter)
}

func NewAuthFilter() filter.Filter {
	return &AuthFilter{}
}

type AuthFilter struct {
}

func (f *AuthFilter) Invoke(ctx context.Context, invoker protocol.Invoker, invocation protocol.Invocation) protocol.Result {

	//xxx 接收凭证
	invocation.SetAttachment("usertoken", "this is jwt token")
	return invoker.Invoke(ctx, invocation)
}

func (f *AuthFilter) OnResponse(ctx context.Context, result protocol.Result, invoker protocol.Invoker, protocol protocol.Invocation) protocol.Result {
	fmt.Println("AuthFilter OnResponse is called")
	myAttachmentMap := make(map[string]interface{})
	myAttachmentMap["key1"] = "value1"
	myAttachmentMap["key2"] = []string{"value1", "value2"}
	result.SetAttachments(myAttachmentMap)
	fmt.Printf("result %v--", result)
	return result
}
