const uuid = require('uuid'),
      fs = require('fs');

const testJSON = fs.readFileSync('test.json').toString();
let testObj = JSON.parse(testJSON, (key, val) => {
    if (typeof val === 'string' && val.indexOf('{uuid}') !== -1) {
        const valParts = val.split('{uuid}');
        return valParts[0] + uuid.v4().toString() + valParts[1];
    } 
    return val;
});
console.log(testObj);
