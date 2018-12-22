
import client from '../utils/fetch'

export function loginByAccount(account, password) {
  const data = {
    account,
    password
  }
  return client({
    url: '/login',
    method: 'post',
    data
  })
}
