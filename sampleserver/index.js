/*
Part of exercise file for go lang course at
https://web.learncodeonline.in
*/

const express = require('express')
const app = express()
const port = 3000

app.use(express.json()); 
app.use(express.urlencoded({extended: true}));

app.get('/', (req, res) => {
  res.status(200).send("Hello my name is Ananthapadmanabhan")
})

app.get('/get', (req, res) => {
    res.status(200).json({message: "My name is kuttuuu"})
  })


app.post('/post', (req, res) => {
 res.status(200).json({message: "Hello from post method", data: req.body
  })
})

app.post('/postform', (req, res) => {
    res.status(200).send(JSON.stringify(req.body));
})
  

app.listen(port, () => {
  console.log(`Example app listening at http://localhost:${port}`)
})