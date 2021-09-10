/**
 * @Author: fuxiao
 * @Author: 576101059@qq.com
 * @Date: 2021/9/7 10:22
 * @Desc: TODO
 */

package private

import (
    "github.com/dobyte/tencent-im/internal/enum"
)

const (
    // 同步至其他设备
    SyncOtherMachineYes = enum.SyncOtherMachineYes // 把消息同步到From_Account在线终端和漫游上
    SyncOtherMachineNo  = enum.SyncOtherMachineNo  // 消息不同步至From_Account
    
    // 推送标识
    PushFlagYes = enum.PushFlagYes // 正常推送
    PushFlagNo  = enum.PushFlagYes // 不离线推送
    
    // 华为推送通知消息分类
    HuaWeiImportanceLow    = enum.HuaWeiImportanceLow    // LOW类消息
    HuaWeiImportanceNormal = enum.HuaWeiImportanceNormal // NORMAL类消息
    
    // 华为推送为“打开应用内指定页面”的前提下透传参数行为
    HuaweiIntentParamAction = enum.HuaweiIntentParamAction // 将透传内容Ext作为Action参数
    HuaweiIntentParamIntent = enum.HuaweiIntentParamIntent // 将透传内容Ext作为Intent参数
    
    // VIVO手机推送消息分类
    VivoClassificationOperation = enum.VivoClassificationOperation // 运营类消息
    VivoClassificationSystem    = enum.VivoClassificationSystem    // 系统类消息
    
    // IOS徽章计数模式
    BadgeModeNormal = enum.BadgeModeNormal // 本条消息需要计数
    BadgeModeIgnore = enum.BadgeModeIgnore // 本条消息不需要计数
    
    // IOS10的推送扩展开关
    MutableContentNormal = enum.MutableContentNormal // 关闭iOS10的推送扩展
    MutableContentEnable = enum.MutableContentEnable // 开启iOS10的推送扩展
)
