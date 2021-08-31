#### Tencent-IM SDK

<table>
    <tr>
        <td width="10%">模块</td>
        <td width="20%">名称</td>
        <td width="30%">方法</td>
        <td>说明</td>
        <td width="5%">master</td>
    </tr>
    <tr>
        <td rowspan="6">账号管理</td>
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
            <a href="https://cloud.tencent.com/document/product/269/36443">删除帐号</a>
        </td>
        <td>Account.DeleteAccounts</td>
        <td>仅支持删除套餐包类型为 IM 体验版的帐号，其他类型的账号（如：TRTC、白板、专业版、旗舰版）无法删除。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/38417">查询帐号</a>
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
            <a href="https://cloud.tencent.com/document/product/269/2566">查询帐号在线状态</a>
        </td>
        <td>Account.QueryAccountsOnlineStatus</td>
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
        <td rowspan="15">关系链管理</td>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/1643">添加好友</a>
        </td>
        <td>SNS.AddFriends</td>
        <td>添加好友，支持批量添加好友。</td>
        <td>√</td>
    </tr>
    <tr>
        <td>
            <a href="https://cloud.tencent.com/document/product/269/8301">导入好友</a>
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
            <a href="https://cloud.tencent.com/document/product/269/12525">更新好友</a>
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
            <a href="https://cloud.tencent.com/document/product/269/1644">删除好友</a>
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
            <a href="https://cloud.tencent.com/document/product/269/1646">校验好友</a>
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
            <a href="https://cloud.tencent.com/document/product/269/8609">拉取指定好友</a>
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
</table>