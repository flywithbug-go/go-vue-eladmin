
export function formatDate(date, fmt) {
  if (/(y+)/.test(fmt)) {
    fmt = fmt.replace(RegExp.$1, (date.getFullYear() + '').substr(4 - RegExp.$1.length))
  }
  const o = {
    'M+': date.getMonth() + 1,
    'd+': date.getDate(),
    'h+': date.getHours(),
    'm+': date.getMinutes(),
    's+': date.getSeconds()
  }
  for (const k in o) {
    if (new RegExp(`(${k})`).test(fmt)) {
      const str = o[k] + ''
      fmt = fmt.replace(RegExp.$1, (RegExp.$1.length === 1) ? str : padLeftZero(str))
    }
  }
  return fmt
}

function padLeftZero(str) {
  return ('00' + str).substr(str.length)
}

// 入参是纳秒ns
export function formatTimeDuration(duration) {
  duration = duration / 1000 // µs
  if (duration < 1000) {
    return duration.toFixed(3) + 'µs'
  }
  duration = duration / 1000
  if (duration < 1000) {
    return duration.toFixed(3) + 'ms'
  }
  duration = duration / 1000
  if (duration < 60) {
    return duration.toFixed(3) + 's'
  }
  duration = duration / 60
  return duration.toFixed(3) + 'min'
}
