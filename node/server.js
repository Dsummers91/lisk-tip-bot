var http = require("http");
var app = require("express")();
var Mnemonic = require("bitcore-mnemonic");

// LSK.setNode();
// console.log(LSK);
app.get('/account/new', (err, res) => {
  let code = new Mnemonic(Mnemonic.Words.ENGLISH).toString();
  let response = {};
  err !== null ? response.secret = code : response.error = "error generating account";
  res.send(response);
});

app.listen("7200", (err, res) => {
  console.log("listening on poer 7200");
})