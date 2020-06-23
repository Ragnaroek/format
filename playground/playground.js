
function onDirectiveChange (text) {
  let directives = 0
  let lastChar = ''
  for (let i = 0; i < text.length; i++) {
    if (text.charAt(i) == '~') {
      if (lastChar == '~') {
        directives--
      } else {
        directives++
      }
    }
    lastChar = text.charAt(i)
  }

  let paramElem = document.getElementById('params')
  let params = paramElem.children
  let diffCnt = directives - (params.length / 2)
  if (diffCnt < 0) {
    // remove
    for (let i = diffCnt; i < 0; i++) {
      params.item(params.length - 1).remove()
      params.item(params.length - 1).remove()
    }
  } else if (diffCnt > 0) {
    // add
    for (let i = diffCnt; i > 0; i--) {
      paramElem.insertAdjacentHTML('beforeend', '<span>,</span>')
      paramElem.insertAdjacentHTML('beforeend', `<label style="display:inline-block">
        <span style="position:absolute; top:6.7em; color: silver"><em class="paramType">string</em></span><input type="text" style="display:inline-block" oninput="onParamChange(this)">
      </label>`)
    }
  }
}

function onParamChange (inputElem) {
  let num = Number(inputElem.value)
  let typeElem = inputElem.previousSibling.children.item(0)
  if (isNaN(num)) {
    typeElem.innerText = 'string'
  } else {
    typeElem.innerText = 'number'
  }
}

function format () {
  let directive = document.getElementById('directive').value
  let paramsHtml = document.getElementById('params').children

  let params = []
  params.push(directive)

  for (let i = 0; i < paramsHtml.length; i++) {
    let param = paramsHtml.item(i)
    if (param.tagName == 'LABEL') {
      for (let z = 0; z < param.children.length; z++) {
        let subParam = param.children.item(z)
        if (subParam.tagName == 'INPUT') {
          params.push(normaliseValue(subParam.value))
        }
      }
    }
  }

  let formatted = Sformat.apply(null, params)
  let resultNode = document.createElement('div')
  resultNode.insertAdjacentText('beforeend', `format.Sformat("${directive}") => ${formatted}`)
  document.getElementById('output').prepend(resultNode)
}

// Removes "" that forces a string type and converts to numbers (if possible)
function normaliseValue (val) {
  let num = Number(val)
  if (isNaN(num)) {
    if (val.length >= 2) {
      if (val.charAt(0) == '"' && val.charAt(val.length - 1) == '"') {
        return val.substring(1, val.length - 1)
      }
      return val
    }
  } else {
    return num
  }
}
