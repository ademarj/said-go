const saNumberId = document.getElementById('southAfricaNumberId')
const form = document.getElementById('form')
const errorElement = document.getElementById('error');
const subButton = document.getElementById('sub');
subButton.disabled=true;

const SA_CITIZEN = 0;
const PERMANENT_RESIDENT = 1;
const INVALID_MESSAGE = 'Invalid ID Number'


form.addEventListener('submit', (el)=>{
    let messages = []
    if (saNumberId.value === '' || saNumberId.value == null ){
        messages.push('ID Number is required')
    }

    if(messages.length > 0){
        el.preventDefault();
        setMessageError(messages.join(', '))
    }
})

function setMessageError(message){
    errorElement.innerText = message;
}

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
            subButton.disabled = !checksumDig(this.value);
            if(subButton.disabled){
                setMessageError(INVALID_MESSAGE)
            }else{
                setMessageError('')
            }
        }else{
            subButton.disabled = true;
            setMessageError('')
        }
      });
    });
}
  
setInputFilter(saNumberId, function(value) {return /^-?\d*$/.test(value); });

function checksumDig(southAfricaId){
    let dateOfBirth = southAfricaId.substring(0,6)
    let genderCode = southAfricaId.substring(6,7)
    let sssCode = southAfricaId.substring(7,10)
    let citizenshipCode = southAfricaId.substring(10,11)
    let ACode = parseInt(southAfricaId.substring(11,12))
    let checksumCode = southAfricaId.substring(12)

    if(!(citizenshipCode == SA_CITIZEN || citizenshipCode == PERMANENT_RESIDENT)) return false

    // https://www.youtube.com/watch?v=XJ7Z8dAPjxI
    // 12 11 10 9 8 7 6 5 4 3 2 1 0
    //  9  2  0 2 2 0 4 7 2 0 0 8 2
    //  9202204720082
    //  9501127062097

    let original = [];

    for (let i = 0; i < southAfricaId.length; i++) {
        original.push(
             parseInt(southAfricaId.charAt(i))
        );
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

    return sum % 10 == 0;
}