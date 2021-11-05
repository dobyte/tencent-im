## 腾讯云IM官方API文档

点击查看 [官方文档](https://cloud.tencent.com/product/im/developer)

## 如何使用

```shell script
go get github.com/dobyte/tencent-im
```

## 调用方法

```go
package main

import (
    "fmt"
    
    "github.com/dobyte/tencent-im"
    "github.com/dobyte/tencent-im/account"
)

func main() {
    tim := im.NewIM(&im.Options{
        AppId:     1400579830,                                                         // 无效的AppId,请勿直接使用
        AppSecret: "0d2a321b087fdb8fd5ed5ea14fe0489139086eb1b03541283fc9feeab8f2bfd3", // 无效的AppSecret,请勿直接使用
        UserId:    "administrator",                                                    // 管理员用户账号，请在腾讯云IM后台设置管理账号
    })
    
    // 导入账号
    if err := tim.Account().ImportAccount(&account.Info{
        Account:  "test1",
        Nickname: "测试账号1",
        FaceUrl:  "https://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
    }); err != nil {
        if e, ok := err.(im.Error); ok {
            fmt.Println(fmt.Sprintf("import accout failed, code:%d, message:%s.", e.Code(), e.Message()))
        } else {
            fmt.Println(fmt.Sprintf("import accout failed:%s.", err.Error()))
        }
    }
    
    fmt.Println("import account success.")
}
```

## SDK列表

<table>
    <tr>
        <td width="100">模块</td>
        <td width="100">名称</td>
        <td>方法</td>
        <td>说明</td>
        <td>master</td>
    </tr>
    <tr>
        <td rowspan="9">账号管理</td>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1608">导入单个帐号</a>
        </td>
        <td>Account.ImportAccount</td>
        <td>本接口用于将 App 自有帐号导入即时通信 IM 帐号系统，为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/4919">导入多个帐号</a>
        </td>
        <td>Account.ImportAccounts</td>
        <td>本接口用于批量将 App 自有帐号导入即时通信 IM 帐号系统，为该帐号创建一个对应的内部 ID，使该帐号能够使用即时通信 IM 服务。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/36443">删除单个帐号</a>
        </td>
        <td>Account.DeleteAccount</td>
        <td>
            <ul>
                <li>本方法拓展于“删除多个帐号（DeleteAccounts）”方法。</li>
                <li>仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/36443">删除多个帐号</a>
        </td>
        <td>Account.DeleteAccounts</td>
        <td>仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/38417">查询单个帐号导入状态</a>
        </td>
        <td>Account.CheckAccount</td>
        <td>
            <ul>
                <li>本方法拓展于“查询多个帐号导入状态（CheckAccounts）”方法。</li>
                <li>用于查询自有帐号是否已导入即时通信 IM。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/38417">查询多个帐号导入状态</a>
        </td>
        <td>Account.CheckAccounts</td>
        <td>用于查询自有帐号是否已导入即时通信 IM，支持批量查询。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/3853">失效帐号登录状态</a>
        </td>
        <td>Account.KickAccount</td>
        <td>本接口适用于将 App 用户帐号的登录状态（例如 UserSig）失效。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/2566">查询单个帐号在线状态</a>
        </td>
        <td>Account.GetAccountOnlineState</td>
        <td>
            <ul>
                <li>本方法拓展于“查询多个帐号在线状态（GetAccountsOnlineState）”方法。</li>
                <li>获取用户当前的登录状态。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/2566">查询多个帐号在线状态</a>
        </td>
        <td>Account.GetAccountsOnlineState</td>
        <td>获取用户当前的登录状态。</td>
        <td>√</td>
    </tr>
    <tr>
        <td rowspan="2">资料管理</td>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1640">设置资料</a>
        </td>
        <td>Profile.SetProfile</td>
        <td>支持 <a href="https://cloud.tencent.com/document/product/269/1500#.E6.A0.87.E9.85.8D.E8.B5.84.E6.96.99.E5.AD.97.E6.AE.B5">标配资料字段</a> 和 <a href="https://cloud.tencent.com/document/product/269/1500#.E8.87.AA.E5.AE.9A.E4.B9.89.E8.B5.84.E6.96.99.E5.AD.97.E6.AE.B5">自定义资料字段</a> 的设置。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1639">拉取资料</a>
        </td>
        <td>Profile.GetProfiles</td>
        <td>
            <ul>
                <li>支持拉取好友和非好友的资料字段。</li>
                <li>支持拉取 <a href="https://cloud.tencent.com/document/product/269/1500#.E6.A0.87.E9.85.8D.E8.B5.84.E6.96.99.E5.AD.97.E6.AE.B5">标配资料字段</a> 和 <a href="https://cloud.tencent.com/document/product/269/1500#.E8.87.AA.E5.AE.9A.E4.B9.89.E8.B5.84.E6.96.99.E5.AD.97.E6.AE.B5">自定义资料字段</a>。</li>
                <li>建议每次拉取的用户数不超过100，避免因回包数据量太大导致回包失败。</li>
                <li>请确保请求中的所有帐号都已导入即时通信 IM，如果请求中含有未导入即时通信 IM 的帐号，即时通信 IM 后台将会提示错误。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td rowspan="23">关系链管理</td>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1643">添加单个好友</a>
        </td>
        <td>SNS.AddFriend</td>
        <td>
            <ul>
                <li>本方法拓展于“添加多个好友（AddFriends）”方法。</li>
                <li>添加好友，仅支持添加单个好友</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1643">添加多个好友</a>
        </td>
        <td>SNS.AddFriends</td>
        <td>添加好友，支持批量添加好友。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/8301">导入单个好友</a>
        </td>
        <td>SNS.ImportFriend</td>
        <td>本方法拓展于“添加多个好友（ImportFriends）”方法。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/8301">导入多个好友</a>
        </td>
        <td>SNS.ImportFriends</td>
        <td>
            <ul>
                <li>支持批量导入单向好友。</li>
                <li>往同一个用户导入好友时建议采用批量导入的方式，避免并发写导致的写冲突。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/12525">更新单个好友</a>
        </td>
        <td>SNS.UpdateFriend</td>
        <td>
            <ul>
                <li>本方法拓展于“更新多个好友（UpdateFriends）”方法。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/12525">更新多个好友</a>
        </td>
        <td>SNS.UpdateFriends</td>
        <td>
            <ul>
                <li>支持批量更新同一用户的多个好友的关系链数据。</li>
                <li>更新一个用户多个好友时，建议采用批量方式，避免并发写导致的写冲突。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1644">删除单个好友</a>
        </td>
        <td>SNS.DeleteFriend</td>
        <td>本方法拓展于“删除多个好友（DeleteFriends）”方法。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1644">删除多个好友</a>
        </td>
        <td>SNS.DeleteFriends</td>
        <td>删除好友，支持单向删除好友和双向删除好友。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1645">删除所有好友</a>
        </td>
        <td>SNS.DeleteAllFriends</td>
        <td>清除指定用户的标配好友数据和自定义好友数据。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1646">校验单个好友</a>
        </td>
        <td>SNS.CheckFriend</td>
        <td>本方法拓展于“校验多个好友（CheckFriends）”方法。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1646">校验多个好友</a>
        </td>
        <td>SNS.CheckFriends</td>
        <td>支持批量校验好友关系。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1647">拉取好友</a>
        </td>
        <td>SNS.FetchFriends</td>
        <td>
            <ul>
                <li>分页拉取全量好友数据。</li>
                <li>不支持资料数据的拉取。</li>
                <li>不需要指定请求拉取的字段，默认返回全量的标配好友数据和自定义好友数据。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1647">拉取好友</a>
        </td>
        <td>SNS.PullFriends</td>
        <td>
            <ul>
                <li>本API是借助"拉取好友"API进行扩展实现</li>
                <li>分页拉取全量好友数据。</li>
                <li>不支持资料数据的拉取。</li>
                <li>不需要指定请求拉取的字段，默认返回全量的标配好友数据和自定义好友数据。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/8609">拉取单个指定好友</a>
        </td>
        <td>SNS.GetFriend</td>
        <td>
            <ul>
                <li>本方法拓展于“拉取多个指定好友（GetFriends）”方法。</li>
                <li>支持拉取指定好友的好友数据和资料数据。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/8609">拉取多个指定好友</a>
        </td>
        <td>SNS.GetFriends</td>
        <td>
            <ul>
                <li>支持拉取指定好友的好友数据和资料数据。</li>
                <li>建议每次拉取的好友数不超过100，避免因数据量太大导致回包失败。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/3718">添加黑名单</a>
        </td>
        <td>SNS.AddBlacklist</td>
        <td>添加黑名单，支持批量添加黑名单。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/3719">删除黑名单</a>
        </td>
        <td>SNS.DeleteBlacklist</td>
        <td>删除指定黑名单。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/3722">拉取黑名单</a>
        </td>
        <td>SNS.FetchBlacklist</td>
        <td>支持分页拉取所有黑名单。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/3722">拉取黑名单</a>
        </td>
        <td>SNS.PullBlacklist</td>
        <td>
            <ul>
                <li>本API是借助"拉取黑名单"API进行扩展实现</li>
                <li>支持分页拉取所有黑名单。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/3725">校验黑名单</a>
        </td>
        <td>SNS.CheckBlacklist</td>
        <td>支持批量校验黑名单。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/10107">添加分组</a>
        </td>
        <td>SNS.AddGroups</td>
        <td>添加分组，支持批量添加分组，并将指定好友加入到新增分组中。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/10108">删除分组</a>
        </td>
        <td>SNS.DeleteGroups</td>
        <td>删除指定分组。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/54763">拉取分组</a>
        </td>
        <td>SNS.GetGroups</td>
        <td>拉取分组，支持指定分组以及拉取分组下的好友列表。</td>
        <td>√</td>
    </tr>
    <tr>
        <td rowspan="8">私聊消息</td>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/2282">单发单聊消息</a>
        </td>
        <td>Private.SendMessage</td>
        <td>
            <ul>
                <li>管理员向帐号发消息，接收方看到消息发送者是管理员。</li>
                <li>管理员指定某一帐号向其他帐号发消息，接收方看到发送者不是管理员，而是管理员指定的帐号。</li>
                <li>该接口不会检查发送者和接收者的好友关系（包括黑名单），同时不会检查接收者是否被禁言。</li>
                <li>单聊消息 MsgSeq 字段的作用及说明：该字段在发送消息时由用户自行指定，该值可以重复，非后台生成，非全局唯一。与群聊消息的 MsgSeq 字段不同，群聊消息的 MsgSeq 由后台生成，每个群都维护一个 MsgSeq，从1开始严格递增。单聊消息历史记录对同一个会话的消息先以时间戳排序，同秒内的消息再以 MsgSeq 排序。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1612">批量发单聊消息</a>
        </td>
        <td>Private.SendMessages</td>
        <td>
            <ul>
                <li>支持一次对最多500个用户进行单发消息。</li>
                <li>与单发消息相比，该接口更适用于营销类消息、系统通知 tips 等时效性较强的消息。</li>
                <li>管理员指定某一帐号向目标帐号批量发消息，接收方看到发送者不是管理员，而是管理员指定的帐号。</li>
                <li>该接口不触发回调请求。</li>
                <li>该接口不会检查发送者和接收者的好友关系（包括黑名单），同时不会检查接收者是否被禁言。</li>
                <li>单聊消息 MsgSeq 字段的作用及说明：该字段在发送消息时由用户自行指定，该值可以重复，非后台生成，非全局唯一。与群聊消息的 MsgSeq 字段不同，群聊消息的 MsgSeq 由后台生成，每个群都维护一个 MsgSeq，从1开始严格递增。单聊消息历史记录对同一个会话的消息先以时间戳排序，同秒内的消息再以 MsgSeq 排序。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/2568">导入单聊消息</a>
        </td>
        <td>Private.ImportMessage</td>
        <td>
            <ul>
                <li>导入历史单聊消息到即时通信 IM。</li>
                <li>平滑过渡期间，将原有即时通信实时单聊消息导入到即时通信 IM。</li>
                <li>该接口不会触发回调。</li>
                <li>该接口会根据 From_Account ， To_Account ，MsgSeq ， MsgRandom ， MsgTimeStamp 字段的值对导入的消息进行去重。仅当这五个字段的值都对应相同时，才判定消息是重复的，消息是否重复与消息内容本身无关。</li>
                <li>重复导入的消息不会覆盖之前已导入的消息（即消息内容以首次导入的为准）。</li>
                <li>单聊消息 MsgSeq 字段的作用及说明：该字段在发送消息时由用户自行指定，该值可以重复，非后台生成，非全局唯一。与群聊消息的 MsgSeq 字段不同，群聊消息的 MsgSeq 由后台生成，每个群都维护一个 MsgSeq，从1开始严格递增。单聊消息历史记录对同一个会话的消息先以时间戳排序，同秒内的消息再以 MsgSeq 排序。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/42794">查询单聊消息</a>
        </td>
        <td>Private.FetchMessages</td>
        <td>
            <ul>
                <li>管理员按照时间范围查询某单聊会话的消息记录。</li>
                <li>查询的单聊会话由请求中的 From_Account 和 To_Account 指定。查询结果包含会话双方互相发送的消息，具体每条消息的发送方和接收方由每条消息里的 From_Account 和 To_Account 指定。</li>
                <li>一般情况下，请求中的 From_Account 和 To_Account 字段值互换，查询结果不变。但通过 单发单聊消息 或 批量发单聊消息 接口发送的消息，如果指定 SyncOtherMachine 值为2，则需要指定正确的 From_Account 和 To_Account 字段值才能查询到该消息。</li>
                <li>查询结果包含被撤回的消息，由消息里的 MsgFlagBits 字段标识。</li>
                <li>若想通过 REST API 撤回单聊消息 接口撤回某条消息，可先用本接口查询出该消息的 MsgKey，然后再调用撤回接口进行撤回。</li>
                <li>可查询的消息记录的时间范围取决于漫游消息存储时长，默认是7天。支持在控制台修改消息漫游时长，延长消息漫游时长是增值服务。具体请参考 漫游消息存储。</li>
                <li>若请求时间段内的消息总大小超过应答包体大小限制（目前为13K）时，则需要续拉。您可以通过应答中的 Complete 字段查看是否已拉取请求的全部消息。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/42794">续拉取单聊消息</a>
        </td>
        <td>Private.PullMessages</td>
        <td>
            <ul>
                <li>本API是借助"查询单聊消息"API进行扩展实现。</li>
                <li>管理员按照时间范围查询某单聊会话的消息记录。</li>
                <li>查询的单聊会话由请求中的 From_Account 和 To_Account 指定。查询结果包含会话双方互相发送的消息，具体每条消息的发送方和接收方由每条消息里的 From_Account 和 To_Account 指定。</li>
                <li>一般情况下，请求中的 From_Account 和 To_Account 字段值互换，查询结果不变。但通过 单发单聊消息 或 批量发单聊消息 接口发送的消息，如果指定 SyncOtherMachine 值为2，则需要指定正确的 From_Account 和 To_Account 字段值才能查询到该消息。</li>
                <li>查询结果包含被撤回的消息，由消息里的 MsgFlagBits 字段标识。</li>
                <li>若想通过 REST API 撤回单聊消息 接口撤回某条消息，可先用本接口查询出该消息的 MsgKey，然后再调用撤回接口进行撤回。</li>
                <li>可查询的消息记录的时间范围取决于漫游消息存储时长，默认是7天。支持在控制台修改消息漫游时长，延长消息漫游时长是增值服务。具体请参考 漫游消息存储。</li>
                <li>若请求时间段内的消息总大小超过应答包体大小限制（目前为13K）时，则需要续拉。您可以通过应答中的 Complete 字段查看是否已拉取请求的全部消息。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/38980">撤回单聊消息</a>
        </td>
        <td>Private.RevokeMessage</td>
        <td>
            <ul>
                <li>管理员撤回单聊消息。</li>
                <li>该接口可以撤回所有单聊消息，包括客户端发出的单聊消息，由 REST API 单发 和 批量发 接口发出的单聊消息。</li>
                <li>若需要撤回由客户端发出的单聊消息，您可以开通 发单聊消息之前回调 或 发单聊消息之后回调 ，通过该回调接口记录每条单聊消息的 MsgKey ，然后填在本接口的 MsgKey 字段进行撤回。您也可以通过 查询单聊消息 查询出待撤回的单聊消息的 MsgKey 后，填在本接口的 MsgKey 字段进行撤回。</li>
                <li>若需要撤回由 REST API 单发 和 批量发 接口发出的单聊消息，需要记录这些接口回包里的 MsgKey 字段以进行撤回。</li>
                <li>调用该接口撤回消息后，该条消息的离线、漫游存储，以及消息发送方和接收方的客户端的本地缓存都会被撤回。</li>
                <li>该接口可撤回的单聊消息没有时间限制，即可以撤回任何时间的单聊消息。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/50349">设置单聊消息已读</a>
        </td>
        <td>Private.SetMessageRead</td>
        <td>设置用户的某个单聊会话的消息全部已读。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/56043">查询单聊未读消息计数</a>
        </td>
        <td>Private.GetUnreadMessageNum</td>
        <td>App 后台可以通过该接口查询特定账号的单聊总未读数（包含所有的单聊会话）或者单个单聊会话的未读数。</td>
        <td>√</td>
    </tr>
    <tr>
            <td rowspan="10">全员推送</td>
            <td>
                <a href="https://cloud.tencent.com/document/product/269/45934">设置应用属性名称</a>
            </td>
            <td>Push.PushMessage</td>
            <td>
                <ul>
                    <li>支持全员推送。</li>
                    <li>支持按用户属性推送。</li>
                    <li>支持按用户标签推送。</li>
                    <li>管理员推送消息，接收方看到消息发送者是管理员。</li>
                    <li>管理员指定某一帐号向其他帐号推送消息，接收方看到发送者不是管理员，而是管理员指定的帐号。</li>
                    <li>支持消息离线存储，不支持漫游。</li>
                    <li>由于全员推送需要下发的帐号数量巨大，下发完全部帐号需要一定时间（根据帐号总数而定，一般在一分钟内）。</li>
                </ul>
            </td>
            <td>√</td>
        </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/45935">设置应用属性名称</a>
        </td>
        <td>Push.SetAttrNames</td>
        <td>每个应用可以设置自定义的用户属性，最多可以有10个。通过本接口可以设置每个属性的名称，设置完成后，即可用于按用户属性推送等。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/45936">获取应用属性名称</a>
        </td>
        <td>Push.GetAttrNames</td>
        <td>管理员获取应用属性名称。使用前请先 设置应用属性名称 。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/45937">获取用户属性</a>
        </td>
        <td>Push.GetUserAttrs</td>
        <td>获取用户属性（必须以管理员帐号调用）；每次最多只能获取100个用户的属性。使用前请先 设置应用属性名称 。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/45938">设置用户属性</a>
        </td>
        <td>Push.SetUserAttrs</td>
        <td>管理员给用户设置属性。每次最多只能给100个用户设置属性。使用前请先 设置应用属性名称 。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/45939">删除用户属性</a>
        </td>
        <td>Push.DeleteUserAttrs</td>
        <td>管理员给用户删除属性。注意每次最多只能给100个用户删除属性。使用前请先 设置应用属性名称。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/45940">获取用户标签</a>
        </td>
        <td>Push.GetUserTags</td>
        <td>获取用户标签（必须以管理员帐号调用）。每次最多只能获取100个用户的标签。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/45941">添加用户标签</a>
        </td>
        <td>Push.AddUserTags</td>
        <td>
            <ul>
                <li>管理员给用户添加标签。</li>
                <li>每次请求最多只能给100个用户添加标签，请求体中单个用户添加标签数最多为10个。</li>
                <li>单个用户可设置最大标签数为100个，若用户当前标签超过100，则添加新标签之前请先删除旧标签。</li>
                <li>单个标签最大长度为50字节。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/45942">删除用户标签</a>
        </td>
        <td>Push.DeleteUserTags</td>
        <td>管理员给用户删除标签。注意每次最多只能给100个用户删除标签。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/45943">删除用户所有标签</a>
        </td>
        <td>Push.DeleteUserAllTags</td>
        <td>管理员给用户删除所有标签。注意每次最多只能给100个用户删除所有标签。</td>
        <td>√</td>
    </tr>
    <tr>
        <td rowspan="2">全局禁言管理</td>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/4230">设置全局禁言</a>
        </td>
        <td>Mute.SetNoSpeaking</td>
        <td>
            <ul>
                <li>设置帐号的单聊消息全局禁言。</li>
                <li>设置帐号的群组消息全局禁言。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/4229">查询全局禁言</a>
        </td>
        <td>Mute.GetNoSpeaking</td>
        <td>
            <ul>
                <li>查询帐号的单聊消息全局禁言。</li>
                <li>查询帐号的群组消息全局禁言。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td rowspan="3">运营管理</td>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/4193">拉取运营数据</a>
        </td>
        <td>Operation.GetOperationData</td>
        <td>App 管理员可以通过该接口拉取最近30天的运营数据，可拉取的字段见下文可拉取的运营字段。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1650">下载最近消息记录</a>
        </td>
        <td>Operation.GetHistoryData</td>
        <td>App 管理员可以通过该接口获取 App 中最近7天中某天某小时的所有单发或群组消息记录的下载地址。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/45438">获取服务器IP地址</a>
        </td>
        <td>Operation.GetIPList</td>
        <td>基于安全等考虑，您可能需要获知服务器的 IP 地址列表，以便进行相关限制。App 管理员可以通过该接口获得 SDK、第三方回调所使用到的服务器 IP 地址列表或 IP 网段信息。</td>
        <td>√</td>
    </tr>
    <tr>
        <td rowspan="32">群组管理</td>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1614">拉取App中的所有群组ID</a>
        </td>
        <td>Group.FetchGroupIds</td>
        <td>App 管理员可以通过该接口获取App中所有群组的ID。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1614">拉取App中的所有群组</a>
        </td>
        <td>Group.FetchGroups</td>
        <td>本方法由“拉取App中的所有群组ID（FetchGroupIds）”拓展而来</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1614">续拉取App中的所有群组</a>
        </td>
        <td>Group.PullGroups</td>
        <td>本方法由“拉取App中的所有群组（FetchGroups）”拓展而来</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1615">创建群组</a>
        </td>
        <td>Group.CreateGroup</td>
        <td>App 管理员可以通过该接口创建群组。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1616">获取单个群详细资料</a>
        </td>
        <td>Group.GetGroup</td>
        <td>本方法由“获取多个群详细资料（GetGroups）”拓展而来</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1616">获取多个群详细资料</a>
        </td>
        <td>Group.GetGroups</td>
        <td>App 管理员可以根据群组 ID 获取群组的详细信息。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1617">拉取群成员详细资料</a>
        </td>
        <td>Group.FetchMembers</td>
        <td>App管理员可以根据群组ID获取群组成员的资料。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1617">拉取群成员详细资料</a>
        </td>
        <td>Group.PullMembers</td>
        <td>本方法由“拉取群成员详细资料（FetchMembers）”拓展而来</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1620">修改群基础资料</a>
        </td>
        <td>Group.UpdateGroup</td>
        <td>App管理员可以通过该接口修改指定群组的基础信息。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1621">增加群成员</a>
        </td>
        <td>Group.AddMembers</td>
        <td>App管理员可以通过该接口向指定的群中添加新成员。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1622">删除群成员</a>
        </td>
        <td>Group.DeleteMembers</td>
        <td>App管理员可以通过该接口删除群成员。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1623">修改群成员资料</a>
        </td>
        <td>Group.UpdateMember</td>
        <td>App管理员可以通过该接口修改群成员资料。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1624">解散群组</a>
        </td>
        <td>Group.DestroyGroup</td>
        <td>App管理员通过该接口解散群。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1625">拉取用户所加入的群组</a>
        </td>
        <td>Group.FetchMemberGroups</td>
        <td>App管理员可以通过本接口获取某一用户加入的群信息。默认不获取用户已加入但未激活好友工作群（Work）以及直播群（AVChatRoom）群信息。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1625">拉取用户所加入的群组</a>
        </td>
        <td>Group.PullMemberGroups</td>
        <td>本方法由“拉取用户所加入的群组（FetchMemberGroups）”拓展而来</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1626">查询用户在群组中的身份</a>
        </td>
        <td>Group.GetRolesInGroup</td>
        <td>App管理员可以通过该接口获取一批用户在群内的身份，即“成员角色”。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1627">批量禁言</a>
        </td>
        <td>Group.ForbidSendMessage</td>
        <td>
            <ul>
                <li>App 管理员禁止指定群组中某些用户在一段时间内发言。</li>
                <li>App 管理员取消对某些用户的禁言。</li>
                <li>被禁言用户退出群组之后再进入同一群组，禁言仍然有效。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1627">取消禁言</a>
        </td>
        <td>Group.AllowSendMessage</td>
        <td>本方法由“批量禁言（ForbidSendMessage）”拓展而来</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/2925">获取被禁言群成员列表</a>
        </td>
        <td>Group.GetShuttedUpMembers</td>
        <td>App管理员可以根据群组ID获取群组中被禁言的用户列表。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1629">在群组中发送普通消息</a>
        </td>
        <td>Group.SendMessage</td>
        <td>App管理员可以通过该接口在群组中发送普通消息。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1630">在群组中发送系统通知</a>
        </td>
        <td>Group.SendNotification</td>
        <td>App 管理员可以通过该接口在群组中发送系统通知。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1633">转让群主</a>
        </td>
        <td>Group.ChangeGroupOwner</td>
        <td>
            <ul>
                <li>App 管理员可以通过该接口将群主身份转移给他人。</li>
                <li>没有群主的群，App 管理员可以通过此接口指定他人作为群主。</li>
                <li>新群主必须为群内成员。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/12341">撤回单条群消息</a>
        </td>
        <td>Group.RevokeMessage</td>
        <td>本方法由“撤回多条群消息（RevokeMessages）”拓展而来</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/12341">撤回多条群消息</a>
        </td>
        <td>Group.RevokeMessages</td>
        <td>App 管理员通过该接口撤回指定群组的消息，消息需要在漫游有效期以内。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1634">导入群基础资料</a>
        </td>
        <td>Group.ImportGroup</td>
        <td>App 管理员可以通过该接口导入群组，不会触发回调、不会下发通知；当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群组数据。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1635">导入群消息</a>
        </td>
        <td>Group.ImportMembers</td>
        <td>
            <ul>
                <li>该 API 接口的作用是导入群组的消息，不会触发回调、不会下发通知。</li>
                <li>当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群消息数据。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1636">导入多个群成员</a>
        </td>
        <td>Group.ImportMembers</td>
        <td>
            <ul>
                <li>该 API 接口的作用是导入群组成员，不会触发回调、不会下发通知。</li>
                <li>当 App 需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议导入存量群成员数据。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1637">设置成员未读消息计数</a>
        </td>
        <td>Group.SetMemberUnreadMsgNum</td>
        <td>
            <ul>
                <li>App管理员使用该接口设置群组成员未读消息数，不会触发回调、不会下发通知。</li>
                <li>当App需要从其他即时通信系统迁移到即时通信 IM 时，使用该协议设置群成员的未读消息计数。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/2359">撤回指定用户发送的消息</a>
        </td>
        <td>Group.RevokeMemberMessages</td>
        <td>该API接口的作用是撤回最近1000条消息中指定用户发送的消息。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/2738">拉取群历史消息</a>
        </td>
        <td>Group.FetchMessages</td>
        <td>
            <ul>
                <li>即时通信 IM 的群消息是按 Seq 排序的，按照 server 收到群消息的顺序分配 Seq，先发的群消息 Seq 小，后发的 Seq 大。</li>
                <li>如果用户想拉取一个群的全量消息，首次拉取时不用填拉取 Seq，Server 会自动返回最新的消息，以后拉取时拉取 Seq 填上次返回的最小 Seq 减1。</li>
                <li>如果返回消息的 IsPlaceMsg 为1，表示这个 Seq 的消息或者过期、或者存储失败、或者被删除了。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/2738">续拉取群历史消息</a>
        </td>
        <td>Group.PullMessages</td>
        <td>本方法由“拉取群历史消息（FetchMessages）”拓展而来</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/49180">获取直播群在线人数</a>
        </td>
        <td>Group.GetOnlineMemberNum</td>
        <td>App 管理员可以根据群组 ID 获取直播群在线人数。</td>
        <td>√</td>
    </tr>
    <tr>
        <td rowspan="3">最近联系人</td>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/62118">拉取会话列表</a>
        </td>
        <td>RecentContact.FetchSessions</td>
        <td>支持分页拉取会话列表。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/62118">拉取会话列表</a>
        </td>
        <td>RecentContact.PullSessions</td>
        <td>
            <ul>
                <li>本API是借助"拉取会话列表"API进行扩展实现</li>
                <li>支持分页拉取会话列表。</li>
            </ul>
        </td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/62119">删除单个会话</a>
        </td>
        <td>RecentContact.DeleteSession</td>
        <td>删除指定会话，支持同步清理漫游消息。</td>
        <td>√</td>
    </tr>
</table>