
function onDirectiveChange (text) {
  var directives = 0
  var lastChar = ''
  for (var i = 0; i < text.length; i++) {
    if (text.charAt(i) == '~') {
      if (lastChar == '~') {
        directives--
      } else {
        directives++
      }
    }
    lastChar = text.charAt(i)
  }

  var paramElem = document.getElementById('params')
  var params = paramElem.children
  var diffCnt = directives - (params.length / 2)
  console.log('diffCnt', diffCnt)
  if (diffCnt < 0) {
    // remove
    for (var i = diffCnt; i < 0; i++) {
      params.item(params.length - 1).remove()
      params.item(params.length - 1).remove()
    }
  } else if (diffCnt > 0) {
    // add
    for (var i = diffCnt; i > 0; i--) {
      paramElem.insertAdjacentHTML('beforeend', '<span>,</span>')
      paramElem.insertAdjacentHTML('beforeend', '<input type="text" style="display:inline-block">')
    }
  }
}

function format () {
  // TODO collect parameters and convert them to the correct types. Supply them to the function.

  var directive = document.getElementById('directive').value
  console.log('directive = ', directive)
  var formatted = Sformat(directive)
  var resultNode = document.createElement('div')
  resultNode.insertAdjacentText('beforeend', `format.Sformat("${directive}") => ${formatted}`)
  document.getElementById('output').prepend(resultNode)
}
