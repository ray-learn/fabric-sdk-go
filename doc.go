/*
Copyright SecureKey Technologies Inc. All Rights Reserved.

SPDX-License-Identifier: Apache-2.0
*/

// Package fabricsdk enables Go developers to build solutions that interact with Hyperledger Fabric.
// fabric包让Go开发者可以搭建与超级账本交互的解决方案
//
// Packages for end developer usage
// 终端开发者的用法
//
// pkg/fabsdk: The main package of the Fabric SDK. This package enables creation of contexts based on
// configuration. These contexts are used by the client packages listed below.
// Reference: https://godoc.org/github.com/hyperledger/fabric-sdk-go/pkg/fabsdk
// pkg/fabsdk: Fabric SDK的主要包，该包可以基于配置创建上下文，这些上下文可由下述客户包使用
//
// pkg/client/channel: Provides channel transaction capabilities.
// Reference: https://godoc.org/github.com/hyperledger/fabric-sdk-go/pkg/client/channel
// pkg/client/channel 提供通道交易能力
//
// pkg/client/event: Provides channel event capabilities.
// Reference: https://godoc.org/github.com/hyperledger/fabric-sdk-go/pkg/client/event
// pkg/client/event 提供通道事件能力
//
// pkg/client/ledger: Enables queries to a channel's underlying ledger.
// Reference: https://godoc.org/github.com/hyperledger/fabric-sdk-go/pkg/client/ledger
// pkg/client/ledger 提供通道底层账本的查询
//
// pkg/client/resmgmt: Provides resource management capabilities such as installing chaincode.
// Reference: https://godoc.org/github.com/hyperledger/fabric-sdk-go/pkg/client/resmgmt
// pkg/client/resmgmt 提供如链码安装等的资源管理
//
// pkg/client/msp: Enables identity management capability.
// Reference: https://godoc.org/github.com/hyperledger/fabric-sdk-go/pkg/client/msp
// pkg/client/msp 提供身份管理
//
// Basic workflow 基本工作流
//
//      1) Instantiate a fabsdk instance using a configuration.
//         Note: fabsdk maintains caches so you should minimize instances of fabsdk itself.
//      2) Create a context based on a user and organization, using your fabsdk instance.
//         Note: A channel context additionally requires the channel ID.
//      3) Create a client instance using its New func, passing the context.
//         Note: you create a new client instance for each context you need.
//      4) Use the funcs provided by each client to create your solution!
//      5) Call fabsdk.Close() to release resources and caches.
//
//		1) 使用配置，实例化fabsdk
// 		   注意：fabsdk保存有缓存，所以可以最小化实力
//		2) 使用fabsdk实例，创建用户和组织创建上下文
//		   注意：通道上下文额外要求通道ID
//		3) 传递上下文参数，使用New函数创建客户端实例
//		   注意：为每个需要的上下文创建客户端实例
//		4) 使用为每个客户端提供的函数创建解决方案
//		5) 调用fabsdk.Close()函数释放资源和缓存
package fabricsdk
