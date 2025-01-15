function clearMessage(input) {
   input.formHasError = false
   input.info.type = 'info'
   input.info.status = input.info.default.status
   input.info.mssg = input.info.default.mssg
}

function showMessage(input, info) {
   // input.formHasError = false

   input.info.type = info.type
   input.info.status = info.status
   input.info.mssg = info.mssg
}

export {
   clearMessage,
   showMessage
}