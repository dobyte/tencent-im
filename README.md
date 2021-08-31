#### Tencent-IM SDK

<table>
    <tr>
        <td>模块</td>
        <td>名称</td>
        <td>方法</td>
        <td width="40%">说明</td>
        <td>master</td>
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
        <td rowspan="3">关系链管理</td>
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
</table>