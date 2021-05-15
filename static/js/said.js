const saNumberId = document.getElementById('southAfricaNumberId')
const form = document.getElementById('form')
const subButton = document.getElementById('sub');
subButton.disabled=true;

const SA_CITIZEN = 0;
const PERMANENT_RESIDENT = 1;
const ERROR_MESSAGE = 'ID Number Invalid'

saNumberId.addEventListener('input', ()=>{
    saNumberId.parentElement.removeAttribute('data-error')
})

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
                subButton.classList.remove("success");
                subButton.disabled = true;
                saNumberId.parentElement.setAttribute('data-error',ERROR_MESSAGE);
                return;
            }
            subButton.classList.add("success");
            return subButton.disabled = false;
        }
        subButton.classList.remove("success");
        subButton.disabled = true;
        
      });
    });
}
  
setInputFilter(saNumberId, function(value) {return /^-?\d*$/.test(value); });

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