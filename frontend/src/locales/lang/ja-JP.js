import antd from 'ant-design-vue/es/locale-provider/zh_CN'
import momentCN from 'moment/locale/zh-cn'
import global from './zh-CN/global'

import menu from './ja-JP/menu'
import setting from './ja-JP/setting'
import user from './ja-JP/user'
import dashboard from './ja-JP/dashboard'
import form from './ja-JP/form'
import result from './ja-JP/result'
import account from './ja-JP/account'

const components = {
  antLocale: antd,
  momentName: 'jp',
  momentLocale: momentCN
}

export default {
  message: '-',

  'layouts.usermenu.dialog.title': '信息',
  'layouts.usermenu.dialog.content': '您确定要注销吗？',
  'layouts.userLayout.title': 'Ant Design 是西湖区最具影响力的 Web 设计规范',
  ...components,
  ...global,
  ...menu,
  ...setting,
  ...user,
  ...dashboard,
  ...form,
  ...result,
  ...account
}
