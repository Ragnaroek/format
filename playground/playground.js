
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
  var directive = document.getElementById('directive').value
  var paramsHtml = document.getElementById('params').children

  var params = []
  params.push(directive)

  for (var i = 0; i < paramsHtml.length; i++) {
    var param = paramsHtml.item(i)
    if (param.tagName == 'INPUT') {
      params.push(param.value)
    }
  }

  var formatted = Sformat.apply(null, params)
  var resultNode = document.createElement('div')
  resultNode.insertAdjacentText('beforeend', `format.Sformat("${directive}") => ${formatted}`)
  document.getElementById('output').prepend(resultNode)
}
