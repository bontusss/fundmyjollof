function CLEAR_ERROR(inputs, input) {
   inputs[input].formHasError = false
   inputs[input].info.type = 'info'
   inputs[input].info.status = inputs[input].info.default.status
   inputs[input].info.mssg = inputs[input].info.default.mssg
}

export {
   CLEAR_ERROR
}