const SA_CITIZEN                        = 0;
const PERMANENT_RESIDENT                = 1;

const ATTRIBUTE_DATA_ERROR              = 'data-error'

const ERROR_MESSAGE                     = 'error'
const ERROR_MESSAGE_CLIENT_SIDE         = 'ID Number Invalid'
const ERROR_MESSAGE_FROM_SERVER_SIDE    = 'said-message-error'

const SUCCESS_MESSAGE_FROM_SERVER_SIDE  = 'said-card-success'
const SUCCESS_MESSAGE                   = 'success'

const idNumber                          = document.getElementById('southAfricaNumberId')
const form                              = document.getElementById('form')
const subButton                         = document.getElementById('sub');
const searchElement                     = document.getElementById('said-card__search')

subButton.disabled=true;

function addMessageError(el){
    const searchInput =  el.parentElement
    const small = searchInput.querySelector('small');
    small.innerText = ERROR_MESSAGE_CLIENT_SIDE
    
    el.classList.add(ERROR_MESSAGE)
}

function removeMessageError(el){
    el.classList.remove(ERROR_MESSAGE)
}

function addMessageSuccess(el){
    el.classList.add(SUCCESS_MESSAGE)
}

function removeMessageSuccess(el){
    el.classList.remove(SUCCESS_MESSAGE)
}

idNumber.addEventListener('input', ()=>{
    idNumber.parentElement.removeAttribute(ATTRIBUTE_DATA_ERROR)
    searchElement.classList.remove(ERROR_MESSAGE)
})

idNumber.addEventListener('focus', ()=>{
    if (document.getElementById(ERROR_MESSAGE_FROM_SERVER_SIDE) != null)
        document.getElementById(ERROR_MESSAGE_FROM_SERVER_SIDE).remove()
})

form.addEventListener('submit', ()=>{
    subButton.disabled = true;
    subButton.classList.remove(SUCCESS_MESSAGE_FROM_SERVER_SIDE);
    removeMessageSuccess(searchElement)
})

//TODO continue refactoring from here

function setInputFilter(textbox, inputFilter) {
    ["input", "keydown", "keyup", "mousedown", "mouseup", "select", "contextmenu", "drop"].forEach(function(event) {
      textbox.addEventListener(event, function() {
        if (inputFilter(this.value)) {
          this.oldValue = this.value;
          this.oldSelectionStart = this.selectionStart;
          this.oldSelectionEnd = this.selectionEnd;
        } else if (this.hasOwnProperty("oldValue")) {
          this.value = this.oldValue;
          this.setSelectionRange(this.oldSelectionStart, this.oldSelectionEnd);
        } else {
          this.value = "";
        }

        if(this.value.length === 13){
            if(!checksumDig(this.value)){
                subButton.classList.remove(SUCCESS_MESSAGE_FROM_SERVER_SIDE);
                removeMessageSuccess(searchElement)
                subButton.disabled = true;
                idNumber.parentElement.setAttribute(ATTRIBUTE_DATA_ERROR,ERROR_MESSAGE_CLIENT_SIDE);
                addMessageError(searchElement)
                return;
            }
            subButton.classList.add(SUCCESS_MESSAGE_FROM_SERVER_SIDE);
            searchElement.classList.add(SUCCESS_MESSAGE)
            return subButton.disabled = false;
        }
        subButton.classList.remove(SUCCESS_MESSAGE_FROM_SERVER_SIDE);
        removeMessageSuccess(searchElement)
        subButton.disabled = true;
        
      });
    });
}
  
setInputFilter(idNumber, function(value) {return /^-?\d*$/.test(value); });

function checksumDig(southAfricaId){
    let dateOfBirth = southAfricaId.substring(0,6)
    if(invalidDate(dateOfBirth)) return false
    
    let citizenshipCode = southAfricaId.substring(10,11)
    if(!(citizenshipCode == SA_CITIZEN || citizenshipCode == PERMANENT_RESIDENT)) return false
    
    let ACode = parseInt(southAfricaId.substring(11,12))

    let original = [];
    for (let i = 0; i < southAfricaId.length; i++) {
        original.push(parseInt(southAfricaId.charAt(i)));
    }
    
    const reversed = original.reverse();
    
    let sum = 0;

    for(let index = 0; index < reversed.length; index++){
        if(index % 2 === 0){ 
            sum += reversed[index];
        }else{  // is odd index
            let result = reversed[index] * 2; // multiplay by 2
            if(result > ACode){ // result is greather then A code
                let resultString = result.toString();
                let sumCharByChar = 0;
                for(let i = 0; i < resultString.length; i++){
                    sumCharByChar += parseInt(resultString.charAt(i))
                }
                sum += sumCharByChar;
            }else{
                sum += result;
            }
        }
    }

    return sum % 10 === 0;
}

function invalidDate(yymmdd){
    let yy = yymmdd.substring(0,2)
    let mm = yymmdd.substring(2,4)
    let dd = yymmdd.substring(4,6)

    let date19yy = new Date(parseInt(`19${yy}`), parseInt(mm)-1, parseInt(dd))
    let date20yy = new Date(parseInt(`20${yy}`), parseInt(mm)-1, parseInt(dd))

    let month19yy = (date19yy.getMonth()+1) < 10 ?  `0${date19yy.getMonth()+1}` : `${date19yy.getMonth()+1}`
    let day19yy = date19yy.getDate() < 10 ? `0${date19yy.getDate()}`: `${date19yy.getDate()}`
    let full19yyFromDate = `${date19yy.getFullYear()}${month19yy}${day19yy}`
    let full19yy = `19${yymmdd}`
    if(full19yyFromDate == full19yy) return false;

    let month20yy = (date20yy.getMonth()+1) < 10 ?  `0${date20yy.getMonth()+1}` : `${date20yy.getMonth()+1}`
    let day20yy = date20yy.getDate() < 10 ? `0${date20yy.getDate()}`: `${date20yy.getDate()}`
    let full20yyFromDate = `${date20yy.getFullYear()}${month20yy}${day20yy}`
    let full20yy = `20${yymmdd}`
    if(full20yyFromDate == full20yy) return false;

    return true
}