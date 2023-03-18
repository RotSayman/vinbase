function validate() {
    if( document.myform.vin.value.match("^(?=.*[0-9])(?=.*[A-z])[0-9A-z-]{17}$")  != null ) {
        return true
    }
    return false
}
