export function clean(phonenumber: string): string {
  if (phonenumber.match(/[@:!]/)) { 
    throw new Error("Punctuations not permitted");
  }

  if (phonenumber.match(/[a-z]/i)) {
    throw new Error("Letters not permitted");
  } 

  phonenumber = phonenumber.replace(/\D/g, '');
  if (phonenumber.length === 11) {
    if (phonenumber[0] === '1') {
      phonenumber = phonenumber.slice(1);
    } else {
      throw new Error("11 digits must start with 1");
    }
  } else if (phonenumber.length < 10) {
    throw new Error("Incorrect number of digits");
  } else if (phonenumber.length > 11) {
    throw new Error("More than 11 digits");
  }


    // split the phone number into 3 parts
  const areaCode = phonenumber.slice(0, 3);
  const exchangeCode = phonenumber.slice(3, 6);
  const subscriberNumber = phonenumber.slice(6);

  // check if the area code starts with 0 or 1
  if (areaCode[0] == '0'){
    throw new Error("Area code cannot start with zero");
  }
  if (areaCode[0] == '1'){
    throw new Error("Area code cannot start with one");
  }

  if (exchangeCode[0] === '0') {
    throw new Error("Exchange code cannot start with zero");
  }

  if (exchangeCode[0] === '1') {
    throw new Error("Exchange code cannot start with one");
  }



  if (areaCode.length !== 3 || exchangeCode.length !== 3 || subscriberNumber.length !== 4) {
    throw new Error("Incorrect number of digits");
  }

  phonenumber = areaCode + exchangeCode + subscriberNumber;
   return phonenumber;
}
